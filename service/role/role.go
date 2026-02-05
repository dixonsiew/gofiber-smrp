package role

import (
    "fmt"
    "smrp/database"
    "smrp/model"
    "smrp/utils"
)

func FindAll(sortBy string, sortDir string) ([]model.Role, error) {
    lx := make([]model.Role, 0)
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return lx, nil
    }

    q := `select id, name from role order by %s %s`
    q = fmt.Sprintf(q, sortBy, sortDir)
    rows, err := db.Query(database.GetCtx(), q)
    if err != nil {
        utils.LogError(err)
        return lx, err
    }

    defer rows.Close()

    for rows.Next() {
        o := model.DbRole{}
        err := rows.Scan(&o.Id, &o.Name)
        if err != nil {
            utils.LogError(err)
            return lx, err
        }

        k := model.Role{}
        k.FromDbModel(o)
        lx = append(lx, k)
    }

    return lx, nil
}

func FindById(id int) (*model.Role, error) {
    o := model.DbRole{}
    k := model.Role{}
    var x *model.Role
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return x, nil
    }

    q := `select id, name from role where id = $1 limit 1`
    rows, err := db.Query(database.GetCtx(), q, id)
    if err != nil {
        utils.LogError(err)
        return x, err
    }

    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&o.Id, &o.Name)
        if err != nil {
            utils.LogError(err)
            return x, err
        }

        k.FromDbModel(o)
        x = &k
    }

    return x, nil
}
