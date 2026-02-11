package model

import (
    "database/sql"
    "math"
    "smrp/database"
    "smrp/utils"
    "strconv"

    "github.com/jackc/pgx/v5/pgxpool"
)

type Pager struct {
    Total    int
    PageNum  int
    PageSize int
}

func (o *Pager) SetPageSize(pageSize int) {
    if (o.Total < pageSize || pageSize < 1) && o.Total > 0 {
        o.PageSize = o.Total
    } else {
        o.PageSize = pageSize
    }

    if o.GetTotalPages() < o.PageNum {
        o.PageNum = o.GetTotalPages()
    }

    if o.PageNum < 1 {
        o.PageNum = 1
    }
}

func (o *Pager) GetLowerBound() int {
    return (o.PageNum - 1) * o.PageSize
}

func (o *Pager) GetUpperBound() int {
    x := o.PageNum * o.PageSize
    if o.Total < x {
        x = o.Total
    }

    return x
}

func (o *Pager) GetTotalPages() int {
    v := float64(o.Total) / float64(o.PageSize)
    x := math.Ceil(v)
    return int(x)
}

func GetPager(total int, page string, limit string) Pager {
    pageNum, _ := strconv.Atoi(page)
    pageSize, _ := strconv.Atoi(limit)
    pg := Pager{
        Total: total,
        PageNum: pageNum,
        PageSize: pageSize,
    }
    pg.SetPageSize(pageSize)
    return pg
}

type DbRole struct {
    Id   sql.NullInt32
    Name sql.NullString
}

type Role struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func (o *Role) FromDbModel(m DbRole) {
    o.Id = int(m.Id.Int32)
    o.Name = m.Name.String
}

type DbUser struct {
    Id        sql.NullInt32
    Username  sql.NullString
    Firstname sql.NullString
    Lastname  sql.NullString
    Password  sql.NullString
    LastLogin sql.NullString
}

type User struct {
    Id        int    `json:"id"`
    Username  string `json:"username"`
    Firstname string `json:"first_name"`
    Lastname  string `json:"last_name"`
    Password  string `json:"-"`
    LastLogin string `json:"last_login"`
    Roles     []Role `json:"roles"`
}

func (o *User) FromDbModel(m DbUser, db *pgxpool.Pool) {
    o.Id = int(m.Id.Int32)
    o.Username = m.Username.String
    o.Firstname = m.Firstname.String
    o.Lastname = m.Lastname.String
    o.Password = m.Password.String
    o.LastLogin = utils.GetDateTimeStr(m.LastLogin.String)
    o.SetRoles(db)
}

func (o *User) SetRoles(db *pgxpool.Pool) {
    lx := make([]Role, 0)
    rows, _ := db.Query(database.GetCtx(), `select aur.app_user_id, aur.roles_id, r.id, r.name from app_user_roles aur inner join role r on aur.roles_id = r.id where aur.app_user_id = $1`, o.Id)
    defer rows.Close()
    for rows.Next() {
        var (
            app_user_id sql.NullInt32
            roles_id    sql.NullInt32
            id          sql.NullInt32
            name        sql.NullString
        )
        _ = rows.Scan(&app_user_id, &roles_id, &id, &name)
        k := Role{
            Id:   int(id.Int32),
            Name: name.String,
        }
        lx = append(lx, k)
    }

    o.Roles = lx
}

type DbCommonSetup struct {
    Id           sql.NullInt32
    Code         sql.NullString
    Desc         sql.NullString
    Ref          sql.NullString
    CreatedBy    sql.NullInt32
    CreatedDate  sql.NullString
    ModifiedBy   sql.NullInt32
    ModifiedDate sql.NullString
    Deleted      sql.NullBool
    DeletedBy    sql.NullInt32
    DeletedDate  sql.NullString
}

type CommonSetup struct {
    Id           int    `json:"id"`
    Code         string `json:"code"`
    Desc         string `json:"desc"`
    Ref          string `json:"ref"`
    CreatedBy    int    `json:"created_by"`
    CreatedDate  string `json:"created_date"`
    ModifiedBy   int    `json:"modified_by"`
    ModifiedDate string `json:"modified_date"`
    Deleted      bool   `json:"deleted"`
    DeletedBy    int    `json:"deleted_by"`
    DeletedDate  string `json:"deleted_date"`
}

func (o *CommonSetup) FromDbModel(m DbCommonSetup) {
    o.Id = int(m.Id.Int32)
    o.Code = m.Code.String
    o.Desc = m.Desc.String
    o.Ref = m.Ref.String
    o.CreatedBy = int(m.CreatedBy.Int32)
    o.CreatedDate = m.CreatedDate.String
    o.ModifiedBy = int(m.ModifiedBy.Int32)
    o.ModifiedDate = m.ModifiedDate.String
    o.Deleted = m.Deleted.Bool
    o.DeletedBy = int(m.DeletedBy.Int32)
    o.DeletedDate = m.DeletedDate.String
}
