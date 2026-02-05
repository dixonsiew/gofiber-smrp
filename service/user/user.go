package user

import (
	"fmt"
	"smrp/database"
	"smrp/model"
	"smrp/utils"

	"golang.org/x/crypto/bcrypt"
)

func FindById(id int) (*model.User, error) {
    o := model.DbUser{}
    k := model.User{}
    var x *model.User
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return x, nil
    }

    q := `select id, username, first_name, last_name, password, last_login from app_user where id = $1 limit 1`
    rows, err := db.Query(database.GetCtx(), q, id)
    if err != nil {
        utils.LogError(err)
        return x, err
    }

    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&o.Id, &o.Username, &o.Firstname, &o.Lastname, &o.Password, &o.LastLogin)
        if err != nil {
            utils.LogError(err)
            return x, err
        }

        k.FromDbModel(o, db)
        x = &k
    }

    return x, nil
}

func FindByUsername(username string) (*model.User, error) {
    o := model.DbUser{}
    k := model.User{}
    var x *model.User
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return x, nil
    }

    q := `select id, username, first_name, last_name, password, last_login from app_user where username = $1 limit 1`
    rows, err := db.Query(database.GetCtx(), q, username)
    if err != nil {
        utils.LogError(err)
        return x, err
    }

    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&o.Id, &o.Username, &o.Firstname, &o.Lastname, &o.Password, &o.LastLogin)
        if err != nil {
            utils.LogError(err)
            return x, err
        }

        k.FromDbModel(o, db)
        x = &k
    }

    return x, nil
}

func FindAll(offset int, limit int, sortBy string, sortDir string) ([]model.User, error) {
    lx := make([]model.User, 0)
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return lx, nil
    }

    q := `select t.id, t.username, t.first_name, t.last_name, t.password, t.last_login from app_user t order by %s %s offset $1 limit $2`
    q = fmt.Sprintf(q, sortBy, sortDir)
    rows, err := db.Query(database.GetCtx(), q, offset, limit)
    if err != nil {
        utils.LogError(err)
        return lx, err
    }

    defer rows.Close()

    for rows.Next() {
        o := model.DbUser{}
        err := rows.Scan(&o.Id, &o.Username, &o.Firstname, &o.Lastname, &o.Password, &o.LastLogin)
        if err != nil {
            utils.LogError(err)
            return lx, err
        }

        k := model.User{}
        k.FromDbModel(o, db)
        lx = append(lx, k)
    }

    return lx, nil
}

func Count() (int, error) {
    n := 0
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return n, nil
    }

    q := `select count(id) from app_user`
    err := db.QueryRow(database.GetCtx(), q).Scan(&n)
    if err != nil {
        utils.LogError(err)
        return n, err
    }

    return n, nil
}

func ExistsByOtherUsername(username string, id int) (bool, error) {
    b := false
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return b, nil
    }

    q := `select exists (select 1 from app_user t where t.username = $1 and t.id <> $2)`
    err := db.QueryRow(database.GetCtx(), q, username, id).Scan(&b)
    if err != nil {
        utils.LogError(err)
        return b, err
    }

    return b, nil
}

func ExistsByUsername(username string) (bool, error) {
    b := false
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return b, nil
    }

    q := `select exists (select 1 from app_user t where t.username = $1)`
    err := db.QueryRow(database.GetCtx(), q, username).Scan(&b)
    if err != nil {
        utils.LogError(err)
        return b, err
    }

    return b, nil
}

func FindByKeyword(keyword string, offset int, limit int, sortBy string, sortDir string) ([]model.User, error) {
    lx := make([]model.User, 0)
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return lx, nil
    }

    q := `select t.id, t.username, t.first_name, t.last_name, t.password, t.last_login from app_user t where (t.username ilike $1 or t.first_name ilike $2 or t.last_name ilike $3) order by %s %s offset $4 limit $5`
    q = fmt.Sprintf(q, sortBy, sortDir)
    rows, err := db.Query(database.GetCtx(), q, keyword, keyword, keyword, offset, limit)
    if err != nil {
        utils.LogError(err)
        return lx, err
    }

    defer rows.Close()

    for rows.Next() {
        o := model.DbUser{}
        err := rows.Scan(&o.Id, &o.Username, &o.Firstname, &o.Lastname, &o.Password, &o.LastLogin)
        if err != nil {
            utils.LogError(err)
            return lx, err
        }

        k := model.User{}
        k.FromDbModel(o, db)
        lx = append(lx, k)
    }

    return lx, nil
}

func CountByKeyword(keyword string) (int, error) {
    n := 0
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return n, nil
    }

    q := `select count(id) from app_user t where (t.username ilike $1 or t.first_name ilike $2 or t.last_name ilike $3)`
    err := db.QueryRow(database.GetCtx(), q, keyword, keyword, keyword).Scan(&n)
    if err != nil {
        utils.LogError(err)
        return n, err
    }

    return n, nil
}

func Save(o model.User) error {
    pw := []byte(o.Password)
    pwd, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
    if err != nil {
        utils.LogError(err)
        return err
    }

    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    ctx := database.GetCtx()
    tx, err := db.Begin(ctx)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer tx.Rollback(ctx)

    q := `insert into app_user (id, username, password, first_name, last_name, active) values(nextval('app_user_id_seq'),$1,$2,$3,$4,$5) returning id as app_user_id`
    var id int
    err = tx.QueryRow(ctx, q, o.Username, pwd, o.Firstname, o.Lastname, true).Scan(&id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    for _, r := range o.Roles {
        qr := `insert into app_user_roles (app_user_id, roles_id) values($1, $2)`
        _, err := tx.Exec(ctx, qr, id, r.Id)
        if err != nil {
            utils.LogError(err)
            return err
        }
    }
    
    if err = tx.Commit(ctx); err != nil {
        return err
    }

    return nil
}

func Update(o model.User) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    ctx := database.GetCtx()
    tx, err := db.Begin(ctx)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer tx.Rollback(ctx)

    if o.Password != "" {
        pw := []byte(o.Password)
        pwd, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
        if err != nil {
            utils.LogError(err)
            return err
        }

        q := `update app_user set password = $1, first_name = $2, last_name = $3 where id = $4`
        _, err = tx.Exec(ctx, q, pwd, o.Firstname, o.Lastname, o.Id)
        if err != nil {
            utils.LogError(err)
            return err
        }
    } else {
        q := `update app_user set first_name = $1, last_name = $2 where id = $3`
        _, err = tx.Exec(ctx, q, o.Firstname, o.Lastname, o.Id)
        if err != nil {
            utils.LogError(err)
            return err
        }
    }
    
    q := `delete from app_user_roles where app_user_id = $1`
    _, err = tx.Exec(ctx, q, o.Id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    for _, r := range o.Roles {
        qr := `insert into app_user_roles (app_user_id, roles_id) values($1,$2)`
        _, err := tx.Exec(ctx, qr, o.Id, r.Id)
        if err != nil {
            utils.LogError(err)
            return err
        }
    }

    if err = tx.Commit(ctx); err != nil {
        return err
    }

    return nil
}

func DeleteById(id int) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    ctx := database.GetCtx()
    tx, err := db.Begin(ctx)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer tx.Rollback(ctx)

    q := `delete from app_user_roles where app_user_id = $1`
    _, err = tx.Exec(ctx, q, id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    q = `delete from app_user where id = $1`
    _, err = tx.Exec(ctx, q, id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    if err = tx.Commit(ctx); err != nil {
        return err
    }

    return nil
}

func UpdateLastLogin(id int) error {
    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    q := `update app_user set last_login = now() where id = $1`
    _, err := db.Exec(database.GetCtx(), q, id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    return nil
}

func UpdatePassword(o model.User) error {
    pw := []byte(o.Password)
    pwd, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
    if err != nil {
        utils.LogError(err)
        return err
    }

    db := database.GetDb()
    if db == nil {
        utils.LogInfo("db is nil")
        return nil
    }

    q := `update app_user set password = $1 where id = $2`
    _, err = db.Exec(database.GetCtx(), q, pwd, o.Id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    return nil
}

func ValidateCredentials(user model.User, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}
