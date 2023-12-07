package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/naviscom/dbSchemaReader"
)

func PrintHeaderInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File, projectFolderName string) {
	_, _ = file.WriteString("package api" + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("import (" + "\n")
	_, _ = file.WriteString(`	// "errors"` + "\n")
	_, _ = file.WriteString(`	"net/http"` + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(`	"github.com/gin-gonic/gin"` + "\n")
	_, _ = file.WriteString(`	db "github.com/naviscom/` + projectFolderName + `/db/sqlc` + `"` + "\n")
	_, _ = file.WriteString(`	// "github.com/naviscom/` + projectFolderName + `/tocken` + `"` + "\n")
	_, _ = file.WriteString(`	// "time"` + "\n")
	_, _ = file.WriteString(")" + "\n")
	_, _ = file.WriteString("\n")
}

func PrintCreateInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File, projectFolderName string) {
	tableName_str := strings.ToUpper(strings.TrimSpace(table[i].OutputFileName[0:1])) + strings.TrimSpace(table[i].OutputFileName[1:])
	_, _ = file.WriteString("type create" + tableName_str + "Request struct {" + "\n")
	for j := 1; j < len(table[i].Table_Columns); j++ {
		var columnType string
		if table[i].Table_Columns[j].ColumnType == "varchar" || table[i].Table_Columns[j].ColumnType == "varchar," {
			columnType = "string"
		}
		if table[i].Table_Columns[j].ColumnType == "bigint" || table[i].Table_Columns[j].ColumnType == "bigint," {
			columnType = "int64"
		}
		if table[i].Table_Columns[j].ColumnType == "real" || table[i].Table_Columns[j].ColumnType == "real," {
			columnType = "float32"
		}
		if table[i].Table_Columns[j].ColumnType == "timestamptz" || table[i].Table_Columns[j].ColumnType == "timestamptz," {
			columnType = "time.Time"
		}
		_, _ = file.WriteString("    " + strings.ToUpper(strings.TrimSpace(table[i].Table_Columns[j].Column_name[0:1])) + strings.TrimSpace(table[i].Table_Columns[j].Column_name[1:]) + `    ` + columnType + ` ` + "`json:" + `"` + table[i].Table_Columns[j].Column_name + `"` + ` binding:"required"` + "`" + "\n")
	}
	_, _ = file.WriteString("}" + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("  func (server *Server) create" + tableName_str + "(ctx *gin.Context) {" + "\n")
	_, _ = file.WriteString("	var req create" + tableName_str + "Request" + "\n")
	_, _ = file.WriteString("	if err := ctx.ShouldBindJSON(&req); err != nil {" + "\n")
	_, _ = file.WriteString("		ctx.JSON(http.StatusBadRequest, errorResponse(err))" + "\n")
	_, _ = file.WriteString("		return" + "\n")
	_, _ = file.WriteString("	}" + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("	arg := db.Create" + tableName_str + "Params{" + "\n")
	var column_name_str string
	for j := 1; j < len(table[i].Table_Columns); j++ {
		column_name_slice := strings.Split(table[i].Table_Columns[j].Column_name, "_")
		// fmt.Println(column_name_slice)
		for k := 0; k < len(column_name_slice); k++ {
			if column_name_slice[k] == "id" {
				column_name_slice[k] = strings.ToUpper(strings.TrimSpace(column_name_slice[k]))
			} else {
				column_name_slice[k] = strings.ToUpper(strings.TrimSpace(column_name_slice[k][0:1])) + strings.TrimSpace(column_name_slice[k][1:])

			}
			column_name_str = strings.Join(column_name_slice, "")
		}
		// fmt.Println(column_name_str)
		_, _ = file.WriteString("		" + column_name_str + ":" + "    req." + strings.ToUpper(strings.TrimSpace(table[i].Table_Columns[j].Column_name[0:1])) + strings.TrimSpace(table[i].Table_Columns[j].Column_name[1:]) + "," + "\n")
	}
	_, _ = file.WriteString("	}" + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("	" + table[i].Table_name + ", err := server.store.Create" + tableName_str + "(ctx, arg)" + "\n")
	_, _ = file.WriteString("	if err != nil {" + "\n")
	_, _ = file.WriteString("		ctx.JSON(http.StatusInternalServerError, errorResponse(err))" + "\n")
	_, _ = file.WriteString("		return" + "\n")
	_, _ = file.WriteString("	}" + "\n")
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString("	ctx.JSON(http.StatusOK, " + table[i].Table_name + ")" + "\n")
	_, _ = file.WriteString("}" + "\n")
}

func PrintInsertBlock(table []dbSchemaReader.Table_Struct, i int) {
	var firstLineInsert, secondLineInsert, footer1, footer2, footer3 string
	firstLineInsert = "-- name: Create" + table[i].FunctionSignature + " :one"
	secondLineInsert = "INSERT INTO " + table[i].Table_name + " ("
	footer1 = ") VALUES ("
	footer2 = ")"
	footer3 = "RETURNING *;"
	fmt.Println(firstLineInsert)
	fmt.Println(secondLineInsert)
	for j := 1; j < len(table[i].Table_Columns); j++ {
		if j >= 1 && j < len(table[i].Table_Columns)-1 {
			fmt.Println("    " + table[i].Table_Columns[j].Column_name + ",")
		}
		if j == len(table[i].Table_Columns)-1 {
			fmt.Println("    " + table[i].Table_Columns[j].Column_name)
		}
	}
	fmt.Println(footer1)
	fmt.Print(" ")
	for j := 1; j < len(table[i].Table_Columns); j++ {
		if j >= 1 && j < len(table[i].Table_Columns)-1 {
			fmt.Print("$" + strconv.Itoa(j) + ", ")
		}
		if j == len(table[i].Table_Columns)-1 {
			fmt.Print("$" + strconv.Itoa(j))
		}
	}
	fmt.Println()
	fmt.Println(footer2)
	fmt.Println(footer3)
}

func PrintInsertBlockInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File) {
	var firstLineInsert, secondLineInsert, footer1, footer2, footer3 string
	firstLineInsert = "-- name: Create" + table[i].FunctionSignature + " :one"
	secondLineInsert = "INSERT INTO " + table[i].Table_name + " ("
	footer1 = ") VALUES ("
	footer2 = ")"
	footer3 = "RETURNING *;"
	_, _ = file.WriteString(firstLineInsert + "\n")
	_, _ = file.WriteString(secondLineInsert + "\n")
	for j := 1; j < len(table[i].Table_Columns); j++ {
		if j >= 1 && j < len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString("    " + table[i].Table_Columns[j].Column_name + "," + "\n")
		}
		if j == len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString("    " + table[i].Table_Columns[j].Column_name + "\n")
		}
	}
	_, _ = file.WriteString(footer1 + "\n")
	_, _ = file.WriteString(" ")
	for j := 1; j < len(table[i].Table_Columns); j++ {
		if j >= 1 && j < len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString("$" + strconv.Itoa(j) + ", ")
		}
		if j == len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString("$" + strconv.Itoa(j))
		}
	}
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(footer2 + "\n")
	_, _ = file.WriteString(footer3 + "\n")
}

func PrintGetBlock(table []dbSchemaReader.Table_Struct, i int) {
	var firstLineGet, secondLineGet, thirdLineGet string
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag || table[i].Table_Columns[j].UniqueFlag {
			firstLineGet = "-- name: Get" + table[i].FunctionSignature + strconv.Itoa(j) + " :one"
			secondLineGet = "SELECT * FROM " + table[i].Table_name
			thirdLineGet = "WHERE " + table[i].Table_Columns[j].Column_name + " = $1 LIMIT 1;"
			fmt.Println()
			fmt.Println(firstLineGet)
			fmt.Println(secondLineGet)
			fmt.Println(thirdLineGet)

		}
	}
}

func PrintGetBlockInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File) {
	var firstLineGet, secondLineGet, thirdLineGet string
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag || table[i].Table_Columns[j].UniqueFlag {
			firstLineGet = "-- name: Get" + table[i].FunctionSignature + strconv.Itoa(j) + " :one"
			secondLineGet = "SELECT * FROM " + table[i].Table_name
			thirdLineGet = "WHERE " + table[i].Table_Columns[j].Column_name + " = $1 LIMIT 1;"
			_, _ = file.WriteString("\n")
			_, _ = file.WriteString(firstLineGet + "\n")
			_, _ = file.WriteString(secondLineGet + "\n")
			_, _ = file.WriteString(thirdLineGet + "\n")
		}
	}
}

func PrintListBlock(table []dbSchemaReader.Table_Struct, i int) {
	var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
	firstLineList = "-- name: List" + table[i].FunctionSignature2 + " :many"
	secondLineList = "SELECT * FROM " + table[i].Table_name
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

func PrintListBlockInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File) {
	var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
	firstLineList = "-- name: List" + table[i].FunctionSignature2 + " :many"
	secondLineList = "SELECT * FROM " + table[i].Table_name
	thirdLineList = "ORDER BY id"
	fourthLineList = "LIMIT $1"
	fifthLineList = "OFFSET $2;"
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineList + "\n")
	_, _ = file.WriteString(secondLineList + "\n")
	_, _ = file.WriteString(thirdLineList + "\n")
	_, _ = file.WriteString(fourthLineList + "\n")
	_, _ = file.WriteString(fifthLineList + "\n")
}

func PrintUpdateBlock(table []dbSchemaReader.Table_Struct, i int) {
	var firstLineUpdate, secondLineUpdate, footer1, footer2, footer3 string
	firstLineUpdate = "-- name: Update" + table[i].FunctionSignature + " :one"
	secondLineUpdate = "UPDATE " + table[i].Table_name
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

func PrintUpdateBlockInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File) {
	var firstLineUpdate, secondLineUpdate, footer1, footer2, footer3 string
	firstLineUpdate = "-- name: Update" + table[i].FunctionSignature + " :one"
	secondLineUpdate = "UPDATE " + table[i].Table_name
	footer1 = "SET "
	footer2 = "WHERE id = $1"
	footer3 = "RETURNING *;"
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineUpdate + "\n")
	_, _ = file.WriteString(secondLineUpdate + "\n")
	_, _ = file.WriteString(footer1)
	for j := 1; j < len(table[i].Table_Columns); j++ {
		if j == 1 {
			_, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "," + "\n")
		}
		if j >= 2 && j < len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "," + "\n")
		}
		if j == len(table[i].Table_Columns)-1 {
			_, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "\n")
		}
	}
	_, _ = file.WriteString(footer2 + "\n")
	_, _ = file.WriteString(footer3 + "\n")
}

func PrintDeleteBlock(table []dbSchemaReader.Table_Struct, i int) {
	var firstLineDelete, secondLineDelete, thirdLineDelete string
	firstLineDelete = "-- name: Delete" + table[i].FunctionSignature + " :exec"
	secondLineDelete = "DELETE FROM " + table[i].Table_name
	thirdLineDelete = "WHERE id = $1"
	fmt.Println()
	fmt.Println(firstLineDelete)
	fmt.Println(secondLineDelete)
	fmt.Println(thirdLineDelete + ";")
}

func PrintDeleteBlockInFile(table []dbSchemaReader.Table_Struct, i int, file *os.File) {
	var firstLineDelete, secondLineDelete, thirdLineDelete string
	firstLineDelete = "-- name: Delete" + table[i].FunctionSignature + " :exec"
	secondLineDelete = "DELETE FROM " + table[i].Table_name
	thirdLineDelete = "WHERE id = $1"
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineDelete + "\n")
	_, _ = file.WriteString(secondLineDelete + "\n")
	_, _ = file.WriteString(thirdLineDelete + ";" + "\n")
}

func main() {
	filePath := os.Args[1]
	destPath := os.Args[2]
	tableX, _ := dbSchemaReader.ReadSchema(filePath)
	for i := 0; i < len(tableX); i++ {
		fmt.Println("table Name: ", tableX[i].Table_name, "OutputFileName: ", tableX[i].OutputFileName, "FunctionSignature: ", tableX[i].FunctionSignature, "FunctionSignature2: ", tableX[i].FunctionSignature2)
		for j := 0; j < len(tableX[i].Table_Columns); j++ {
			fmt.Println("    column name: ", tableX[i].Table_Columns[j].Column_name, tableX[i].Table_Columns[j].ColumnType, tableX[i].Table_Columns[j].PrimaryFlag, tableX[i].Table_Columns[j].UniqueFlag, tableX[i].Table_Columns[j].ColumnNameParams)
		}
		for j := 0; j < len(tableX[i].IndexDetails); j++ {
			fmt.Println("    index name: ", tableX[i].IndexDetails[j].IndexName)
			for k := 0; k < len(tableX[i].IndexDetails[j].IndexColumn); k++ {
				fmt.Println("    index column name: ", tableX[i].IndexDetails[j].IndexColumn[k])
			}
		}
	}
	for i := 0; i < len(tableX); i++ {
		file, errs := os.Create(destPath + "/" + tableX[i].OutputFileName + ".sql")
		if errs != nil {
			fmt.Println("Failed to create file:", errs)
			return
		}
		defer file.Close()
		// PrintInsertBlock(tableX[:], i)
		PrintInsertBlockInFile(tableX[:], i, file)
		// PrintGetBlock(tableX[:], i)
		PrintGetBlockInFile(tableX[:], i, file)
		// PrintListBlock(tableX[:], i)
		PrintListBlockInFile(tableX[:], i, file)
		// PrintUpdateBlock(tableX[:], i)
		PrintUpdateBlockInFile(tableX[:], i, file)
		// PrintDeleteBlock(tableX[:], i)
		PrintDeleteBlockInFile(tableX[:], i, file)
		file.Close()
	}
}
