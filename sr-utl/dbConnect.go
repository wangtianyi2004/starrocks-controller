package  utl

import(
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func RunSQL(userName string, password string, ip string, port int, dbName string, sqlStat string) (rows *sql.Rows, err error){

    var infoMess string
    dbPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, password, ip, port, dbName)
    DB, err := sql.Open("mysql", dbPath)
    if err != nil{
        infoMess = fmt.Sprintf("Error in open db [dbPath = %s], error = %v", dbPath, err)
	Log("ERROR", infoMess)
	return nil, err
    }
    defer DB.Close()


    err = DB.Ping()
    if err != nil{
        infoMess = fmt.Sprintf("Error in ping db [dbPath = %s], error = %v", dbPath, err)
	Log("ERROR", infoMess)
	return nil, err
    }

    rows, err = DB.Query(sqlStat)
    if err != nil{
        infoMess = fmt.Sprintf("Error in run sql [dbPath = %s, SQL = %s], error = %v", dbPath, sqlStat, err)
	Log("ERROR", infoMess)
	return nil, err
    }

    return rows, err

}
