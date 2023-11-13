package main

import (
	"bufio"
	"math/rand"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type table_struct struct {
  Table_name          string
  Table_Columns       []tableColumns
  IndexDetails        []index_name_details
  OutputFileName      string
  FunctionSignature   string
}

type index_name_details struct {
  IndexName       string
  IndexColumn     []string
}


type tableColumns struct {
  Column_name     string
  PrimaryFlag     bool
  UniqueFlag      bool
}

func PrintInsertBlock(table []table_struct, i int) {
  var firstLineInsert, secondLineInsert, footer1, footer2, footer3 string
  firstLineInsert = "-- name: Create"+table[i].FunctionSignature+" :one"
  secondLineInsert = "INSERT INTO "+table[i].Table_name+" ("
  footer1 = ") VALUES ("
  footer2 = ")"
  footer3 = "RETURNING *;"
  fmt.Println(firstLineInsert)
  fmt.Println(secondLineInsert)
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j >= 1 && j < len(table[i].Table_Columns)-1 {
      fmt.Println("    "+table[i].Table_Columns[j].Column_name+"," )
    }
    if j == len(table[i].Table_Columns)-1 {
      fmt.Println("    "+table[i].Table_Columns[j].Column_name)      
    }
  }
  fmt.Println(footer1)
  fmt.Print(" ")
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j >= 1 && j < len(table[i].Table_Columns)-1 {
      fmt.Print("$"+strconv.Itoa(j)+", ")
    }
    if j == len(table[i].Table_Columns)-1 {
      fmt.Print("$"+strconv.Itoa(j))         
    }
  }
  fmt.Println()
  fmt.Println(footer2)
  fmt.Println(footer3)
}

func PrintInsertBlockInFile(table []table_struct, i int, file *os.File) {
  var firstLineInsert, secondLineInsert, footer1, footer2, footer3 string
  firstLineInsert = "-- name: Create"+table[i].FunctionSignature+" :one"
  secondLineInsert = "INSERT INTO "+table[i].Table_name+" ("
  footer1 = ") VALUES ("
  footer2 = ")"
  footer3 = "RETURNING *;"
  _, _ = file.WriteString(firstLineInsert+"\n")
  _, _ = file.WriteString(secondLineInsert+"\n")
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j >= 1 && j < len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString("    "+table[i].Table_Columns[j].Column_name+","+"\n" )
    }
    if j == len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString("    "+table[i].Table_Columns[j].Column_name+"\n")      
    }
  }
  _, _ = file.WriteString(footer1+"\n")
  _, _ = file.WriteString(" ")
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j >= 1 && j < len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString("$"+strconv.Itoa(j)+", ")
    }
    if j == len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString("$"+strconv.Itoa(j))         
    }
  }
  _, _ = file.WriteString("\n")
  _, _ = file.WriteString(footer2+"\n")
  _, _ = file.WriteString(footer3+"\n")
}

func PrintGetBlock(table []table_struct, i int) {
  var firstLineGet, secondLineGet, thirdLineGet string
  for j := 0; j < len(table[i].Table_Columns); j++ {
    if table[i].Table_Columns[j].PrimaryFlag || table[i].Table_Columns[j].UniqueFlag {
      firstLineGet = "-- name: Get"+table[i].FunctionSignature+strconv.Itoa(j)+" :one"
      secondLineGet = "SELECT * FROM "+table[i].Table_name
      thirdLineGet = "WHERE "+table[i].Table_Columns[j].Column_name+" = $1 LIMIT 1;"
      fmt.Println()
      fmt.Println(firstLineGet)
      fmt.Println(secondLineGet)
      fmt.Println(thirdLineGet)
    }
  }
}

func PrintGetBlockInFile(table []table_struct, i int, file *os.File) {
  var firstLineGet, secondLineGet, thirdLineGet string
  for j := 0; j < len(table[i].Table_Columns); j++ {
    if table[i].Table_Columns[j].PrimaryFlag || table[i].Table_Columns[j].UniqueFlag {
      firstLineGet = "-- name: Get"+table[i].FunctionSignature+strconv.Itoa(j)+" :one"
      secondLineGet = "SELECT * FROM "+table[i].Table_name
      thirdLineGet = "WHERE "+table[i].Table_Columns[j].Column_name+" = $1 LIMIT 1;"
      _, _ = file.WriteString("\n")
      _, _ = file.WriteString(firstLineGet+"\n")
      _, _ = file.WriteString(secondLineGet+"\n")
      _, _ = file.WriteString(thirdLineGet+"\n")  
    }
  }
}

func PrintListBlock(table []table_struct, i int) {
  var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
  firstLineList = "-- name: List"+table[i].FunctionSignature+" :many"
  secondLineList = "SELECT * FROM "+table[i].Table_name
  thirdLineList = "ORDER BY id"
  fourthLineList = "LIMIT $1"
  fifthLineList = "OFFSET $2;"
  fmt.Println()
  fmt.Println(firstLineList)
  fmt.Println(secondLineList)
  fmt.Println(thirdLineList)
  fmt.Println(fourthLineList)
  fmt.Println(fifthLineList)
}

func PrintListBlockInFile(table []table_struct, i int, file *os.File) {
  var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
  firstLineList = "-- name: List"+table[i].FunctionSignature+" :many"
  secondLineList = "SELECT * FROM "+table[i].Table_name
  thirdLineList = "ORDER BY id"
  fourthLineList = "LIMIT $1"
  fifthLineList = "OFFSET $2;"
  _, _ = file.WriteString("\n")
  _, _ = file.WriteString(firstLineList+"\n")
  _, _ = file.WriteString(secondLineList+"\n")
  _, _ = file.WriteString(thirdLineList+"\n")
  _, _ = file.WriteString(fourthLineList+"\n")
  _, _ = file.WriteString(fifthLineList+"\n")
}

func PrintUpdateBlock(table []table_struct, i int) {
  var firstLineUpdate, secondLineUpdate, footer1, footer2, footer3 string
  firstLineUpdate = "-- name: Update"+table[i].FunctionSignature+" :one"
  secondLineUpdate = "UPDATE "+table[i].Table_name
  footer1 = "SET "
  footer2 = "WHERE id = $1"
  footer3 = "RETURNING *;"
  fmt.Println()
  fmt.Println(firstLineUpdate)
  fmt.Println(secondLineUpdate)
  fmt.Print(footer1)
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j == 1 {
      fmt.Println(table[i].Table_Columns[j].Column_name, " = $"+strconv.Itoa(j+1)+",")
    }
    if j >= 2 && j < len(table[i].Table_Columns)-1 {
      fmt.Println(table[i].Table_Columns[j].Column_name, " = $"+strconv.Itoa(j+1)+",")
    }
    if j == len(table[i].Table_Columns)-1 {
      fmt.Println(table[i].Table_Columns[j].Column_name, " = $"+strconv.Itoa(j+1))         
    }
  }
  fmt.Println(footer2)
  fmt.Println(footer3)
}

func PrintUpdateBlockInFile(table []table_struct, i int, file *os.File) {
  var firstLineUpdate, secondLineUpdate, footer1, footer2, footer3 string
  firstLineUpdate = "-- name: Update"+table[i].FunctionSignature+" :one"
  secondLineUpdate = "UPDATE "+table[i].Table_name
  footer1 = "SET "
  footer2 = "WHERE id = $1"
  footer3 = "RETURNING *;"
  _, _ = file.WriteString("\n")
  _, _ = file.WriteString(firstLineUpdate+"\n")
  _, _ = file.WriteString(secondLineUpdate+"\n")
  _, _ = file.WriteString(footer1)
  for j := 1; j < len(table[i].Table_Columns); j++ {
    if j == 1 {
      _, _ = file.WriteString(table[i].Table_Columns[j].Column_name+" = $"+strconv.Itoa(j+1)+","+"\n")
    }
    if j >= 2 && j < len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString(table[i].Table_Columns[j].Column_name+" = $"+strconv.Itoa(j+1)+","+"\n")
    }
    if j == len(table[i].Table_Columns)-1 {
      _, _ = file.WriteString(table[i].Table_Columns[j].Column_name+" = $"+strconv.Itoa(j+1)+"\n")         
    }
  }
  _, _ = file.WriteString(footer2+"\n")
  _, _ = file.WriteString(footer3+"\n")
}

func PrintDeleteBlock(table []table_struct, i int) {
  var firstLineDelete, secondLineDelete, thirdLineDelete string
  firstLineDelete = "-- name: Delete"+table[i].FunctionSignature+" :exec"
  secondLineDelete = "DELETE FROM "+table[i].Table_name
  thirdLineDelete = "WHERE id = $1"
  fmt.Println()
  fmt.Println(firstLineDelete)
  fmt.Println(secondLineDelete)
  fmt.Println(thirdLineDelete+";")
}

func PrintDeleteBlockInFile(table []table_struct, i int, file *os.File) {
  var firstLineDelete, secondLineDelete, thirdLineDelete string
  firstLineDelete = "-- name: Delete"+table[i].FunctionSignature+" :exec"
  secondLineDelete = "DELETE FROM "+table[i].Table_name
  thirdLineDelete = "WHERE id = $1"
  _, _ = file.WriteString("\n")
  _, _ = file.WriteString(firstLineDelete+"\n")
  _, _ = file.WriteString(secondLineDelete+"\n")
  _, _ = file.WriteString(thirdLineDelete+";"+"\n")
}

func main() {
  filePath := os.Args[1]
  destPath := os.Args[2]
  readFile, err := os.Open(filePath)
  fmt.Println(destPath)
  if err != nil {
      fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  var tableX []table_struct
  var table table_struct
  var tabColumns tableColumns
  for fileScanner.Scan() {
    res1 := strings.Split(fileScanner.Text(), " ")
    if len(res1) > 1 {
      if res1[0] == "CREATE" && res1[1] == "TABLE" {
        table.Table_name = strings.TrimSpace(res1[2][1:len(res1[2])-1])
        if strings.TrimSpace(table.Table_name[len(table.Table_name)-3:]) == `ies` {
          table.OutputFileName = strings.TrimSpace(table.Table_name[:len(table.Table_name)-3])+"y"
        }else if strings.TrimSpace(table.Table_name[len(table.Table_name)-1:]) == `s` {
          table.OutputFileName = strings.TrimSpace(table.Table_name[:len(table.Table_name)-1])
        }else {
          table.OutputFileName = table.Table_name
        }
        if strings.TrimSpace(table.Table_name[len(table.Table_name)-3:]) == `ies` {
          table.FunctionSignature = strings.ToUpper(strings.TrimSpace(table.Table_name[0:1]))+strings.TrimSpace(table.Table_name[1:len(table.Table_name)-3]+"y")
          } else if strings.TrimSpace(table.Table_name[len(table.Table_name)-1:]) == `s` {
            table.FunctionSignature = strings.ToUpper(strings.TrimSpace(table.Table_name[0:1]))+strings.TrimSpace(table.Table_name[1:len(table.Table_name)-1])
          } else {
            table.FunctionSignature = strings.ToUpper(strings.TrimSpace(table.Table_name[0:1]))+strings.TrimSpace(table.Table_name[1:])
        }
        table.Table_Columns = nil
      }
      if res1[0] == "" && res1[1] == "" && strings.TrimSpace(res1[2][0:1]) == `"` {
        tabColumns.Column_name = strings.TrimSpace(res1[2][1:len(res1[2])-1])
        if len(res1) > 4 {
          if res1[4] == `PRIMARY` {
            tabColumns.PrimaryFlag = true
          } else{
            tabColumns.PrimaryFlag = false
          }
          if res1[4] == `UNIQUE` {
            tabColumns.UniqueFlag = true
          } else{
            tabColumns.UniqueFlag = false            
          }
        } else {
          tabColumns.PrimaryFlag = false
          tabColumns.UniqueFlag = false
        }
        table.Table_Columns = append(table.Table_Columns, tabColumns)
      }
      if res1[0] == "CREATE" && res1[1] == "INDEX" {
        for i:=0; i<len(tableX); i++{
          if tableX[i].Table_name == strings.TrimSpace(res1[3][1:len(res1[3])-1]) { 
            var index index_name_details
            index.IndexName = strings.TrimSpace(res1[3][1:len(res1[3])-1]) + strconv.Itoa(rand.Intn(90000))
            for m :=4; m<len(res1); m++ {            
              indexColumnName := res1[m]
              if strings.TrimSpace(indexColumnName[0:1]) == `(` {
                indexColumnName = strings.TrimSpace(indexColumnName[2:len(indexColumnName)-1])
              } else if strings.TrimSpace(indexColumnName[0:1]) == `"` {
                indexColumnName = strings.TrimSpace(indexColumnName[1:len(indexColumnName)-1])
              }
              if strings.TrimSpace(indexColumnName[len(indexColumnName)-1:]) == `)` {
                indexColumnName = strings.TrimSpace(indexColumnName[0:len(indexColumnName)-2])
              } else if strings.TrimSpace(indexColumnName[len(indexColumnName)-1:]) == `"` {
                indexColumnName = strings.TrimSpace(indexColumnName[0:len(indexColumnName)-1])
              }
              // fmt.Println(indexColumnName)
              index.IndexColumn = append(index.IndexColumn,   indexColumnName)
            }
            tableX[i].IndexDetails = append(tableX[i].IndexDetails, index)
          }
        }
      }
    }
    if len(res1) == 1 {
      if res1[0] == ");" {
        tableX = append(tableX, table)
      }
    }
  }
  for i:=0; i<len(tableX); i++{
    fmt.Println("table Name: ", tableX[i].Table_name)
    for j:=0; j<len(tableX[i].Table_Columns); j++{
      fmt.Println("    column name: ", tableX[i].Table_Columns[j].Column_name, tableX[i].Table_Columns[j].PrimaryFlag, tableX[i].Table_Columns[j].UniqueFlag)
    }
    for j:=0; j<len(tableX[i].IndexDetails); j++{
      fmt.Println("    index name: ", tableX[i].IndexDetails[j].IndexName)
      for k:=0; k<len(tableX[i].IndexDetails[j].IndexColumn); k++{
        fmt.Println("    index column name: ", tableX[i].IndexDetails[j].IndexColumn[k])
      }
    }
  }
  for i:=0; i<len(tableX); i++{
    file, errs := os.Create(destPath+"/"+tableX[i].OutputFileName+".sql")
    if errs != nil {
      fmt.Println("Failed to create file:", errs)
      return
    }
    defer file.Close()
    PrintInsertBlock(tableX[:], i)
    PrintInsertBlockInFile(tableX[:], i, file)
    PrintGetBlock(tableX[:], i)
    PrintGetBlockInFile(tableX[:], i, file)
    PrintListBlock(tableX[:], i)
    PrintListBlockInFile(tableX[:], i, file)
    PrintUpdateBlock(tableX[:], i)
    PrintUpdateBlockInFile(tableX[:], i, file)
    PrintDeleteBlock(tableX[:], i)
    PrintDeleteBlockInFile(tableX[:], i, file)
    file.Close()
  }
}
