package database

import (
	"database/sql"
	"smrp/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDataList(rows *sql.Rows) ([]any, []string) {
    cols, _ := rows.Columns()
    colTypes, _ := rows.ColumnTypes()
    la := make([]any, 0)

    values := make([]any, len(cols))
    for i, colType := range colTypes {
        var val any
        switch colType.DatabaseTypeName() {
        case "NCHAR", "VARCHAR", "VARCHAR2", "NVARCHAR2", "DATE", "TIMESTAMP", "TIMESTAMP WITH TIME ZONE", "TIMESTAMP WITH LOCAL TIME ZONE":
            val = new(sql.NullString)
        case "INTEGER", "LONG", "INT", "SMALLINT":
            val = new(sql.NullInt64)
        case "REAL", "DOUBLE PRECISION", "FLOAT", "DECIMAL", "NUMBER":
            val = new(sql.NullFloat64)
        default:
            val = new(interface{})
        }
        values[i] = val
    }

    for rows.Next() {
        err := rows.Scan(values...)
        if err != nil {
            utils.LogError(err)
        }

        mx := make(bson.M)
        for i, col := range cols {
            switch v := values[i].(type) {
            case *sql.NullString:
                mx[col] = v.String
            case *sql.NullInt64:
                mx[col] = v.Int64
            case *sql.NullFloat64:
                mx[col] = v.Float64
            }
        }
        la = append(la, mx)
    }

    return la, cols
}
