package commonsetup

import (
	"fmt"
	"smrp/database"
	"smrp/model"
	"smrp/utils"
)

func FindById(id int, table string) (*model.CommonSetup, error) {
    o := model.DbCommonSetup{}
    k := model.CommonSetup{}
    var x *model.CommonSetup
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return x, nil
    }

    q := `select id, code, "desc", ref, created_by, created_date, modified_by, modified_date, deleted, deleted_by, deleted_date from %s where id = $1 limit 1`
    q = fmt.Sprintf(q, table)
    rows, err := db.Query(database.GetCtx(), q, id)
    if err != nil {
        utils.LogError(err)
        return x, err
    }

    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&o.Id, &o.Code, &o.Desc, &o.Ref, &o.CreatedBy, &o.CreatedDate, &o.ModifiedBy, &o.ModifiedDate, &o.Deleted, &o.DeletedBy, &o.DeletedDate)
        if err != nil {
            utils.LogError(err)
            return x, err
        }

        k.FromDbModel(o)
        x = &k
    }

    return x, nil
}

func FindByDesc(desc string, table string) (*model.CommonSetup, error) {
    o := model.DbCommonSetup{}
    k := model.CommonSetup{}
    var x *model.CommonSetup
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return x, nil
    }

    q := `select t.id, t.code, t.desc, t.ref, t.created_by, t.created_date, t.modified_by, t.modified_date, t.deleted, t.deleted_by, t.deleted_date from %s t where lower(t.desc) = lower($1) limit 1`
    q = fmt.Sprintf(q, table)
    rows, err := db.Query(database.GetCtx(), q, desc)
    if err != nil {
        utils.LogError(err)
        return x, err
    }

    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&o.Id, &o.Code, &o.Desc, &o.Ref, &o.CreatedBy, &o.CreatedDate, &o.ModifiedBy, &o.ModifiedDate, &o.Deleted, &o.DeletedBy, &o.DeletedDate)
        if err != nil {
            utils.LogError(err)
            return x, err
        }

        k.FromDbModel(o)
        x = &k
    }

    return x, nil
}

func FindAll(table string, offset int, limit int, sortBy string, sortDir string) ([]model.CommonSetup, error) {
    lx := make([]model.CommonSetup, 0)
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return lx, nil
    }

    q := ""
    lp := make([]any, 0)
    if limit > 0 {
        q = `select t.id, t.code, t.desc, t.ref, t.created_by, t.created_date, t.modified_by, t.modified_date, t.deleted, t.deleted_by, t.deleted_date from %s t where t.deleted is not true order by "%s" %s offset $1 limit $2`
        q = fmt.Sprintf(q, table, sortBy, sortDir)
        lp = []any{ offset, limit }
    } else {
        q = `select t.id, t.code, t.desc, t.ref, t.created_by, t.created_date, t.modified_by, t.modified_date, t.deleted, t.deleted_by, t.deleted_date from %s t where t.desc <> '' and t.deleted is not true order by t.code`
        q = fmt.Sprintf(q, table)
    }

    rows, err := db.Query(database.GetCtx(), q, lp...)
    if err != nil {
        utils.LogError(err)
        return lx, err
    }

    defer rows.Close()

    for rows.Next() {
        o := model.DbCommonSetup{}
        err := rows.Scan(&o.Id, &o.Code, &o.Desc, &o.Ref, &o.CreatedBy, &o.CreatedDate, &o.ModifiedBy, &o.ModifiedDate, &o.Deleted, &o.DeletedBy, &o.DeletedDate)
        if err != nil {
            utils.LogError(err)
            return lx, err
        }

        k := model.CommonSetup{}
        k.FromDbModel(o)
        lx = append(lx, k)
    }

    return lx, nil
}

func Count(table string) (int, error) {
    n := 0
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return n, nil
    }

    q := `select count(id) from %s t where t.deleted is not true`
    q = fmt.Sprintf(q, table)
    err := db.QueryRow(database.GetCtx(), q).Scan(&n)
    if err != nil {
        utils.LogError(err)
        return n, err
    }

    return n, nil
}

func FindByKeyword(keyword string, offset int, limit int, sortBy string, sortDir string, table string) ([]model.CommonSetup, error) {
    lx := make([]model.CommonSetup, 0)
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return lx, nil
    }

    q := `select t.id, t.code, t.desc, t.ref, t.created_by, t.created_date, t.modified_by, t.modified_date, t.deleted, t.deleted_by, t.deleted_date from %s t where (t.code ilike $1 or t.desc ilike $2 or t.ref ilike $3) and t.deleted is not true order by "%s" %s offset $4 limit $5`
    q = fmt.Sprintf(q, table, sortBy, sortDir)
    rows, err := db.Query(database.GetCtx(), q, keyword, keyword, keyword, offset, limit)
    if err != nil {
        utils.LogError(err)
        return lx, err
    }

    defer rows.Close()

    for rows.Next() {
        o := model.DbCommonSetup{}
        err := rows.Scan(&o.Id, &o.Code, &o.Desc, &o.Ref, &o.CreatedBy, &o.CreatedDate, &o.ModifiedBy, &o.ModifiedDate, &o.Deleted, &o.DeletedBy, &o.DeletedDate)
        if err != nil {
            utils.LogError(err)
            return lx, err
        }

        k := model.CommonSetup{}
        k.FromDbModel(o)
        lx = append(lx, k)
    }

    return lx, nil
}

func CountByKeyword(keyword string, table string) (int, error) {
    n := 0
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return n, nil
    }

    q := `select count(id) from %s t where (t.code ilike $1 or t.desc ilike $2 or t.ref ilike $3) and t.deleted is not true`
    q = fmt.Sprintf(q, table)
    err := db.QueryRow(database.GetCtx(), q, keyword, keyword, keyword).Scan(&n)
    if err != nil {
        utils.LogError(err)
        return n, err
    }

    return n, nil
}

func Save(o model.CommonSetup, table string) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    q := `insert into %s (id, code, "desc", ref, created_by, created_date) values(nextval('%s_id_seq'),$1,$2,$3,$4,now())`
    q = fmt.Sprintf(q, table, table)
    _, err := db.Exec(database.GetCtx(), q, o.Code, o.Desc, o.Ref, o.CreatedBy)
    if err != nil {
        utils.LogError(err)
        return err
    }

    return nil
}

func Update(o model.CommonSetup, table string) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    q := `update %s set code = $1, "desc" = $2, ref = $3, modified_by = $4, modified_date = now() where id = $5`
    q = fmt.Sprintf(q, table)
    _, err := db.Exec(database.GetCtx(), q, o.Code, o.Desc, o.Ref, o.ModifiedBy, o.Id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    return nil
}

func DeleteById(id int, user_id int, table string) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    q := `update %s set deleted = true, deleted_by = $1, deleted_date = now() where id = $2`
    q = fmt.Sprintf(q, table)
    _, err := db.Exec(database.GetCtx(), q, user_id, id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    return nil
}
