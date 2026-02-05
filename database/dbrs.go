package database

import (
    "database/sql"
    "fmt"
    "smrp/config"
    "smrp/utils"

    _ "github.com/sijms/go-ora/v2"
)

var dbrsVar *sql.DB

func SetDbrs(db *sql.DB) {
    dbrsVar = db
}

func GetDbrs() *sql.DB {
    if dbrsVar == nil {
        ConnectDBRs()
    }

    return dbrsVar
}

func ConnectDBRs() {
    username := config.Config("report.db.username")
    pwd := config.Config("report.db.pwd")
    url := config.Config("report.db.url")
    connStr := fmt.Sprintf("oracle://%s:%s@%s", username, pwd, url)
    db, err := sql.Open("oracle", connStr)
    if err != nil {
        utils.LogError(err)
    } else {
        db.SetMaxOpenConns(10)
        db.SetMaxIdleConns(5)
        SetDbrs(db)
        utils.LogInfo("Connection Opened to Database")
    }
}

func CloseDBRs() {
    dbrsVar.Close()
}
