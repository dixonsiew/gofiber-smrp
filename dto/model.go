package dto

type LoginDto struct {
    Username string `json:"username" validate:"required" default:"admin"`
    Password string `json:"password" validate:"required" default:"admin123"`
}

type RefreshTokenDto struct {
    RefreshToken string `json:"refresh_token" validate:"required"`
}

type ChangePasswordDto struct {
    Password        string `json:"password" validate:"required,max=150"`
    ConfirmPassword string `json:"confirm_password" validate:"required,max=150"`
}

type KeywordDto struct {
    Keyword string `json:"keyword"`
}

type CommonSetupDto struct {
    Code string `json:"code" validate:"required,max=30"`
    Desc string `json:"desc" validate:"required,max=300"`
    Ref  string `json:"ref" validate:"required,max=200"`
}

type UserDto struct {
    Username  string `json:"username" validate:"required,max=150"`
    Password  string `json:"password" validate:"required,max=150"`
    Firstname string `json:"first_name" validate:"required,max=150"`
    Lastname  string `json:"last_name" validate:"max=150"`
    RoleId    int    `json:"role_id" validate:"required,number"`
}

type ReportQueryDto struct {
    Page     int    `json:"_page" validate:"number"`
    Limit    int    `json:"_limit" validate:"number"`
    Vt       int    `json:"vt" validate:"number"`
    DateFrom string `json:"datefrom" validate:"required"`
    DateTo   string `json:"dateto" validate:"required"`
}
