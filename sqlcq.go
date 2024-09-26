package sqlcq

import (
    "fmt"
    "os"
    "strconv"
    // "strings"

    "github.com/naviscom/dbschemareader"
)

// func PrintHeaderInFile(table []dbschemareader.Table_Struct, i int, file *os.File, projectFolderName string) {
//     _, _ = file.WriteString("package api" + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString("import (" + "\n")
//     _, _ = file.WriteString(`   // "errors"` + "\n")
//     _, _ = file.WriteString(`   "net/http"` + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString(`   "github.com/gin-gonic/gin"` + "\n")
//     _, _ = file.WriteString(`   db "github.com/naviscom/` + projectFolderName + `/db/sqlc` + `"` + "\n")
//     _, _ = file.WriteString(`   // "github.com/naviscom/` + projectFolderName + `/tocken` + `"` + "\n")
//     _, _ = file.WriteString(`   // "time"` + "\n")
//     _, _ = file.WriteString(")" + "\n")
//     _, _ = file.WriteString("\n")
// }

// func PrintCreateInFile(table []dbschemareader.Table_Struct, i int, file *os.File, projectFolderName string) {
//     tableName_str := strings.ToUpper(strings.TrimSpace(table[i].OutputFileName[0:1])) + strings.TrimSpace(table[i].OutputFileName[1:])
//     _, _ = file.WriteString("type create" + tableName_str + "Request struct {" + "\n")
//     for j := 1; j < len(table[i].Table_Columns); j++ {
//         var columnType string
//         if table[i].Table_Columns[j].ColumnType == "varchar" || table[i].Table_Columns[j].ColumnType == "varchar," {
//             columnType = "string"
//         }
//         if table[i].Table_Columns[j].ColumnType == "bigint" || table[i].Table_Columns[j].ColumnType == "bigint," {
//             columnType = "int64"
//         }
//         if table[i].Table_Columns[j].ColumnType == "real" || table[i].Table_Columns[j].ColumnType == "real," {
//             columnType = "float32"
//         }
//         if table[i].Table_Columns[j].ColumnType == "timestamptz" || table[i].Table_Columns[j].ColumnType == "timestamptz," {
//             columnType = "time.Time"
//         }
//         _, _ = file.WriteString("    " + strings.ToUpper(strings.TrimSpace(table[i].Table_Columns[j].Column_name[0:1])) + strings.TrimSpace(table[i].Table_Columns[j].Column_name[1:]) + `    ` + columnType + ` ` + "`json:" + `"` + table[i].Table_Columns[j].Column_name + `"` + ` binding:"required"` + "`" + "\n")
//     }
//     _, _ = file.WriteString("}" + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString("  func (server *Server) create" + tableName_str + "(ctx *gin.Context) {" + "\n")
//     _, _ = file.WriteString("   var req create" + tableName_str + "Request" + "\n")
//     _, _ = file.WriteString("   if err := ctx.ShouldBindJSON(&req); err != nil {" + "\n")
//     _, _ = file.WriteString("       ctx.JSON(http.StatusBadRequest, errorResponse(err))" + "\n")
//     _, _ = file.WriteString("       return" + "\n")
//     _, _ = file.WriteString("   }" + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString("   arg := db.Create" + tableName_str + "Params{" + "\n")
//     var column_name_str string
//     for j := 1; j < len(table[i].Table_Columns); j++ {
//         column_name_slice := strings.Split(table[i].Table_Columns[j].Column_name, "_")
//         // fmt.Println(column_name_slice)
//         for k := 0; k < len(column_name_slice); k++ {
//             if column_name_slice[k] == "id" {
//                 column_name_slice[k] = strings.ToUpper(strings.TrimSpace(column_name_slice[k]))
//             } else {
//                 column_name_slice[k] = strings.ToUpper(strings.TrimSpace(column_name_slice[k][0:1])) + strings.TrimSpace(column_name_slice[k][1:])

//             }
//             column_name_str = strings.Join(column_name_slice, "")
//         }
//         // fmt.Println(column_name_str)
//         _, _ = file.WriteString("       " + column_name_str + ":" + "    req." + strings.ToUpper(strings.TrimSpace(table[i].Table_Columns[j].Column_name[0:1])) + strings.TrimSpace(table[i].Table_Columns[j].Column_name[1:]) + "," + "\n")
//     }
//     _, _ = file.WriteString("   }" + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString("   " + table[i].Table_name + ", err := server.store.Create" + tableName_str + "(ctx, arg)" + "\n")
//     _, _ = file.WriteString("   if err != nil {" + "\n")
//     _, _ = file.WriteString("       ctx.JSON(http.StatusInternalServerError, errorResponse(err))" + "\n")
//     _, _ = file.WriteString("       return" + "\n")
//     _, _ = file.WriteString("   }" + "\n")
//     _, _ = file.WriteString("\n")
//     _, _ = file.WriteString("   ctx.JSON(http.StatusOK, " + table[i].Table_name + ")" + "\n")
//     _, _ = file.WriteString("}" + "\n")
// }

func PrintInsertBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
    var firstLineInsert, secondLineInsert, footer1, footer2, footer3 string
    firstLineInsert = "-- name: Create" + table[i].FunctionSignature + " :one"
    secondLineInsert = "INSERT INTO " + table[i].Table_name + " ("
    footer1 = ") VALUES ("
    footer2 = ")"
    footer3 = "RETURNING *;"
    _, _ = file.WriteString(firstLineInsert + "\n")
    _, _ = file.WriteString(secondLineInsert + "\n")
    var z int
    if (table[i].Table_Columns[0].ColumnType == "bigserial" && table[i].Table_Columns[0].PrimaryFlag) || (table[i].Table_Columns[0].ColumnType == "uuid" && table[i].Table_Columns[0].PrimaryFlag) {
        z = 1
    }
    if (table[i].Table_Columns[0].ColumnType != "bigserial"  && table[i].Table_Columns[0].PrimaryFlag) ||  (table[i].Table_Columns[0].ColumnType != "uuid"  && table[i].Table_Columns[0].PrimaryFlag){
        z = 0
    }
    for j := z; j < len(table[i].Table_Columns); j++ {
        if j >= z && j <= len(table[i].Table_Columns)-1 {
            if table[i].Table_name == "users" && (table[i].Table_Columns[j].Column_name == "password_changed_at" || table[i].Table_Columns[j].Column_name == "password_created_at") {
                continue
            }
            if table[i].Table_name == "sessions" && (table[i].Table_Columns[j].Column_name == "created_at") {
                continue
            }
            if table[i].Table_name == "userpaymenttokens" && (table[i].Table_Columns[j].Column_name == "created_at" || table[i].Table_Columns[j].Column_name == "updated_at") {
                continue
            }
            if table[i].Table_name == "subusers" && (table[i].Table_Columns[j].Column_name == "password_changed_at" || table[i].Table_Columns[j].Column_name == "password_created_at") {
                continue
            }
            if table[i].Table_name == "activities" && (table[i].Table_Columns[j].Column_name == "service_used_at") {
                continue
            }
            if j > z {
                _, _ = file.WriteString("," + "\n")
            }
            _, _ = file.WriteString("    " + table[i].Table_Columns[j].Column_name)
        }
    }
    _, _ = file.WriteString(footer1 + "\n")
    _, _ = file.WriteString(" ")
    if z == 1 {
        for j := z; j <= len(table[i].Table_Columns); j++ {
            if j >= z && j <= len(table[i].Table_Columns)-1 {
                if j > z {
                    _, _ = file.WriteString(",")
                }
                _, _ = file.WriteString("$" + strconv.Itoa(j))
            }
        }
    }
    if z == 0 {
        u := 0
        for j := z; j < len(table[i].Table_Columns); j++ {
            if j >= z && j <= len(table[i].Table_Columns)-1 {
                if table[i].Table_name == "users" && (table[i].Table_Columns[j].Column_name == "password_changed_at" || table[i].Table_Columns[j].Column_name == "password_created_at") {
                    continue
                }
                if table[i].Table_name == "sessions" && (table[i].Table_Columns[j].Column_name == "created_at") {
                    continue
                }
                if table[i].Table_name == "userpaymenttokens" && (table[i].Table_Columns[j].Column_name == "created_at" || table[i].Table_Columns[j].Column_name == "updated_at") {
                    continue
                }
                if table[i].Table_name == "subusers" && (table[i].Table_Columns[j].Column_name == "password_changed_at" || table[i].Table_Columns[j].Column_name == "password_created_at") {
                    continue
                }
                if table[i].Table_name == "activities" && (table[i].Table_Columns[j].Column_name == "service_used_at") {
                    continue
                }    
                if j > z {
                    _, _ = file.WriteString(",")
                }
                _, _ = file.WriteString("$" + strconv.Itoa(u+1))
                u++
            }
        }
    }
    _, _ = file.WriteString("\n")
    _, _ = file.WriteString(footer2 + "\n")
    _, _ = file.WriteString(footer3 + "\n")
}

func PrintGetBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
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

func PrintListBlockInFile(table []dbschemareader.Table_Struct, i int, file *os.File) {
    var firstLineList, secondLineList, thirdLineList, fourthLineList, fifthLineList string
    firstLineList = "-- name: List" + table[i].FunctionSignature2 + " :many"
    secondLineList = "SELECT * FROM " + table[i].Table_name
    var newLine string
    var fkFlag, firstFKFlag bool
    fkFlag = false
    firstFKFlag = false
    var w int = 2
    for g := 0; g < len(table[i].Table_Columns); g++ {
        if table[i].Table_Columns[g].ForeignFlag {
            w++
            if !firstFKFlag {
                newLine = "WHERE "
                firstFKFlag = true
            }
            if fkFlag {
                newLine = newLine + " OR "
            }
            newLine = newLine + table[i].Table_Columns[g].Column_name + " = $" + strconv.Itoa(w)
            fkFlag = true
        }
    }
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

    // len = 7
    // 0    username  varchar [pk]
    // 1    role varchar [not null, default: 'level_1_user']
    // 2    hashed_password varchar [not null]
    // 3    full_name varchar [not null]
    // 4    email varchar [unique, not null]
    // 5    password_changed_at timestamptz [not null, default: '0001-01-01 00:00:00Z']
    // 6    password_created_at timestamptz [not null, default: `now()`]


    // len = 7
    // 0    username  varchar [pk]
    // 1    hashed_password varchar [not null]
    // 2    full_name varchar [not null]
    // 3    email varchar [unique, not null]
    // 4    password_changed_at timestamptz [not null, default: '0001-01-01 00:00:00Z']
    // 5    password_created_at timestamptz [not null, default: `now()`]
    // 6    role varchar [not null, default: 'level_1_user']

    // j = 1, hashed_password, j+1 = 2 = full_name , len(table[i].Table_Columns)-1 = 6, 2<6
    // j = 2, full_name, j+1 = 3 = email , len(table[i].Table_Columns)-1 = 6, 3<6 
    // j = 3, email, j+1 = 4 = password_changed_at , len(table[i].Table_Columns)-1 = 6, 4<6
    // j = 4, password_changed_at, j+1 = 5 = password_created_at , len(table[i].Table_Columns)-1 = 6, 5<6
    // j = 5, password_created_at, j+1 = 6 = role , len(table[i].Table_Columns)-1 = 6, 6=6
    // j = 6, role, j+1 = 6 = role , len(table[i].Table_Columns)-1 = 6, 6=6


    if table[i].Table_name == "users" || table[i].Table_name == "subusers" {
        for j := 1; j < len(table[i].Table_Columns); j++ {
            if j > 0 && j < len(table[i].Table_Columns)-1 {
                if table[i].Table_Columns[j].Column_name == "password_created_at" {
                    continue
                }
                if table[i].Table_Columns[j+1].Column_name == "password_created_at" && j+1 < len(table[i].Table_Columns)-1 {
                    _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1)+ "," + "\n")
                    continue
                }
                if table[i].Table_Columns[j+1].Column_name == "password_created_at" && j+1 == len(table[i].Table_Columns)-1 {
                    _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "\n")
                    continue
                }
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "," + "\n")
                continue    
            }
            if j == len(table[i].Table_Columns)-1 {
                if table[i].Table_Columns[j].Column_name == "password_created_at"{
                    continue
                }
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j) + "\n")
            }
        }   
    } else if table[i].Table_name == "userpaymenttokens"{
        for j := 1; j < len(table[i].Table_Columns); j++ {
            if j > 0 && j < len(table[i].Table_Columns)-1 {
                if table[i].Table_Columns[j].Column_name == "password_created_at" {
                    continue
                }
                if table[i].Table_Columns[j+1].Column_name == "password_created_at" && j+1 < len(table[i].Table_Columns)-1 {
                    _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1)+ "," + "\n")
                    continue
                }
                if table[i].Table_Columns[j+1].Column_name == "password_created_at" && j+1 == len(table[i].Table_Columns)-1 {
                    _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "\n")
                    continue
                }
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "," + "\n")
                continue    
            }
            if j == len(table[i].Table_Columns)-1 {
                if table[i].Table_Columns[j].Column_name == "password_created_at"{
                    continue
                }
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j) + "\n")
            }
        }
    } else{
        for j := 1; j < len(table[i].Table_Columns); j++ {
            if j > 0 && j < len(table[i].Table_Columns)-1 {
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "," + "\n")
                continue
            }
            if j == len(table[i].Table_Columns)-1 {
                _, _ = file.WriteString(table[i].Table_Columns[j].Column_name + " = $" + strconv.Itoa(j+1) + "\n")
            }
        }   
    }
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
