package sqlcq

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strings"

	"github.com/naviscom/dbschemareader"
)

func PrintInsertBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
	firstLineInsert := "-- name: Create" + table[i].FunctionSignature + " :one"
	secondLineInsert := "INSERT INTO " + table[i].Table_name + " ("
	footer1 := ") VALUES ("
	footer2 := ")"
	footer3 := "RETURNING *;"
	_, _ = file.WriteString(firstLineInsert + "\n")
	_, _ = file.WriteString(secondLineInsert + "\n")
	
	var columns []string
	
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag &&
		 (strings.TrimSpace(table[i].Table_Columns[j].ColumnType) == "uuid" &&
		  strings.TrimSpace(table[i].Table_Columns[j].DefaultValue) == "gen_random_uuid()") {
			continue
		}
		if (table[i].Table_Columns[j].ColumnType == "timestamptz" &&
		 table[i].Table_Columns[j].DefaultValue == "now()") {
			continue
		}
		if (table[i].Table_Columns[j].ColumnType == "date" &&
		 table[i].Table_Columns[j].DefaultValue == "CURRENT_DATE") {
			continue
		}
		
		columns = append(columns, "    "+table[i].Table_Columns[j].Column_name)
	}
	
	// Join with commas and write
	_, _ = file.WriteString(strings.Join(columns, ",\n") + "\n")
	_, _ = file.WriteString(footer1 + "\n")
	var params []string
	u := 1
	
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag &&
		 (strings.TrimSpace(table[i].Table_Columns[j].ColumnType) == "uuid" &&
		  strings.TrimSpace(table[i].Table_Columns[j].DefaultValue) == "gen_random_uuid()") {
			continue
		}
		if (table[i].Table_Columns[j].ColumnType == "timestamptz" &&
		table[i].Table_Columns[j].DefaultValue == "now()") {
			continue
		}
		if (table[i].Table_Columns[j].ColumnType == "date" &&
		table[i].Table_Columns[j].DefaultValue == "CURRENT_DATE") {
			continue
		}
		
		params = append(params, "$"+strconv.Itoa(u))
		u++
	}
	
	// Write all parameters with proper comma separation
	_, _ = file.WriteString(strings.Join(params, ","))
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(footer2 + "\n")
	_, _ = file.WriteString(footer3 + "\n")
}

func PrintGetBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
	var firstLineGet, secondLineGet, thirdLineGet string
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag || table[i].Table_Columns[j].UniqueFlag {
			// fmt.Println("from PrintGetBlockInFile------>", table[i].Table_Columns[j].Column_name, table[i].Table_Columns[j].PrimaryFlag, table[i].Table_Columns[j].UniqueFlag)
			firstLineGet = "-- name: Get" + table[i].FunctionSignature + strconv.Itoa(j) + " :one"
			secondLineGet = "SELECT * FROM " + table[i].Table_name
			thirdLineGet = "WHERE " + table[i].Table_Columns[j].Column_name + " = $1 LIMIT 1;"
			_, _ = file.WriteString("\n")
			_, _ = file.WriteString(firstLineGet + "\n")
			_, _ = file.WriteString(secondLineGet + "\n")
			_, _ = file.WriteString(thirdLineGet + "\n")
		}
	}

	for j, element := range table[i].CompositeUniqueConstraints {
		uniqueColumns := strings.Split(element.ConstraintColumns, ",")
		firstLineGet = "-- name: Get" + table[i].FunctionSignature + strconv.Itoa(j)+strconv.Itoa(j) + " :one"
		secondLineGet = "SELECT * FROM " + table[i].Table_name
		thirdLineGet = "WHERE "
		for x, column := range uniqueColumns {
			thirdLineGet = thirdLineGet + strings.TrimSpace(column) +" = $"+ strconv.Itoa(x+1)
			if x < len(uniqueColumns)-1 {
				thirdLineGet = thirdLineGet + " AND "
			}else if x == len(uniqueColumns)-1 {
				thirdLineGet = thirdLineGet + " LIMIT 1;"
			}
		}		
		_, _ = file.WriteString("\n")
		_, _ = file.WriteString(firstLineGet + "\n")
		_, _ = file.WriteString(secondLineGet + "\n")
		_, _ = file.WriteString(thirdLineGet + "\n")
	}
}

func PrintListBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
	var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
	firstLineList = "-- name: List" + table[i].FunctionSignature2 + " :many"
	secondLineList = "SELECT * FROM " + table[i].Table_name
	var newLine string
	// var fkFlag, firstFKFlag bool
	// fkFlag = false
	// firstFKFlag = false
	// var w int = 2
	
	// // Only process nullable foreign keys for WHERE clause
	// for g := 0; g < len(table[i].Table_Columns); g++ {
	// 	if table[i].Table_Columns[g].ForeignFlag && !table[i].Table_Columns[g].Not_Null {
	// 		w++
	// 		if !firstFKFlag {
	// 			newLine = "WHERE "
	// 			firstFKFlag = true
	// 		}
	// 		if fkFlag {
	// 			newLine = newLine + " OR "
	// 		}
	// 		newLine = newLine + table[i].Table_Columns[g].Column_name + " = $" + strconv.Itoa(w)
	// 		fkFlag = true
	// 	}
	// }
	
	thirdLineList = "ORDER BY "
	for g := 0; g < len(table[i].Table_Columns); g++ {
		if table[i].Table_Columns[g].PrimaryFlag {
			thirdLineList = thirdLineList + table[i].Table_Columns[g].Column_name
			break
		}
	}
	fourthLineList = "LIMIT $1"
	fifthLineList = "OFFSET $2;"
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineList + "\n")
	_, _ = file.WriteString(secondLineList + "\n")
	if len(newLine) > 0 {
		_, _ = file.WriteString(newLine + "\n")
	}
	_, _ = file.WriteString(thirdLineList + "\n")
	_, _ = file.WriteString(fourthLineList + "\n")
	_, _ = file.WriteString(fifthLineList + "\n")
}

func PrintUpdateBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
	var firstLineUpdate, secondLineUpdate, footer1, footer2, footer3 string
	firstLineUpdate = "-- name: Update" + table[i].FunctionSignature + " :one"
	secondLineUpdate = "UPDATE " + table[i].Table_name
	footer1 = "SET "
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag {
			footer2 = "WHERE " + table[i].Table_Columns[j].Column_name + " = $1"
		}
	}
	footer3 = "RETURNING *;"
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineUpdate + "\n")
	_, _ = file.WriteString(secondLineUpdate + "\n")
	_, _ = file.WriteString(footer1)

	// First, collect all the columns that will be in the update statement
	var columnsToUpdate []string
	u := 2 // Start parameter counter from 2 (assuming $1 is used in the WHERE clause)

	for j := 0; j < len(table[i].Table_Columns); j++ {
		if (table[i].Table_Columns[j].PrimaryFlag ||
		table[i].Table_Columns[j].UniqueFlag ) {
			continue
		}

		// Add this column to our update list
		columnStr := table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(u)
		columnsToUpdate = append(columnsToUpdate, columnStr)
		u++
	}

	// Now write all columns with commas in between
	for idx, columnStr := range columnsToUpdate {
		if idx == 0 {
			// First column - no indentation needed for first line after SET
			_, _ = file.WriteString(columnStr)
		} else {
			// Subsequent columns - add comma after previous line and indent this line
			_, _ = file.WriteString(",\n    " + columnStr)
		}
	}

	// Add a newline before the WHERE clause
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(footer2 + "\n")
	_, _ = file.WriteString(footer3 + "\n")
}

func PrintDeleteBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
	var firstLineDelete, secondLineDelete, thirdLineDelete string
	firstLineDelete = "-- name: Delete" + table[i].FunctionSignature + " :exec"
	secondLineDelete = "DELETE FROM " + table[i].Table_name
	for j := 0; j < len(table[i].Table_Columns); j++ {
		if table[i].Table_Columns[j].PrimaryFlag {
			thirdLineDelete = "WHERE " + table[i].Table_Columns[j].Column_name + " = $1"
		}
	}
	_, _ = file.WriteString("\n")
	_, _ = file.WriteString(firstLineDelete + "\n")
	_, _ = file.WriteString(secondLineDelete + "\n")
	_, _ = file.WriteString(thirdLineDelete + ";" + "\n")
}

// func WriteQuery(upSqlFile string, dest string) {
func WriteQuery(tableX []dbschemareader.Table_Struct, dest string) {
	// filePath := upSqlFile
	destPath := dest
	for i := 0; i < len(tableX); i++ {
		file, errs := os.Create(destPath + "/" + tableX[i].OutputFileName + ".sql")
		if errs != nil {
			fmt.Println("Failed to create file:", errs)
			return
		}
		defer file.Close()
		PrintInsertBlockInFile(tableX[:], i, file)
		PrintGetBlockInFile(tableX[:], i, file)
		PrintListBlockInFile(tableX[:], i, file)
		PrintUpdateBlockInFile(tableX[:], i, file)
		PrintDeleteBlockInFile(tableX[:], i, file)
		file.Close()
	}
}
