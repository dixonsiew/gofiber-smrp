package masterpd301

import (
    "encoding/json"
    "fmt"
    "smrp/config"
    "smrp/database"
    "smrp/dto"
    "smrp/middleware"
    "smrp/model"
    rpt "smrp/service/report"
    "smrp/sql"
    u "smrp/utils"
    "strings"

    // "github.com/flosch/pongo2/v6"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

// JsonRH101
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        datefrom           query      string  false  "datefrom"
// @Param        dateto             query      string  false  "dateto"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/export/rpt2 [get]
func JsonRH101(c *fiber.Ctx) error {
    datefrom := c.Query("datefrom", "")
    dateto := c.Query("dateto", "")
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusNotFound, "User not found")
    }

    username := user.Username
    cli := database.GetMongoClient()
    col := getCollection(cli, username, "1")
    ls, err := u.GetCollectionList(database.GetMongoCtx(), col)
    if err != nil {
        return err
    }

    dt1 := strings.Split(datefrom, "-")
    dt2 := strings.Split(dateto, "-")
    ds1 := fmt.Sprintf("%s%s%s", dt1[2], dt1[1], dt1[0])
    ds2 := fmt.Sprintf("%s%s%s", dt2[2], dt2[1], dt2[0])

    forms := make([]fiber.Map, 0)
    for _, d := range ls {
        person := fiber.Map{
            "refPersonTitleCode":        rpt.RefPersonTitleCode(d),
            "fullName":                  u.GetStr(d["PATIENT_NAME"]),
            "refIdentificationTypeCode": rpt.RefIdentificationTypeCode(d),
            "identificationNo":          u.GetStr(d["DOCUMENT_NUMBER"]),
            "refAddressTypeCode":        "C",
            "street1":                   u.GetStr(d["STREET1"]),
            "street2":                   u.GetStr(d["STREET2"]),
            "refCityCode":               rpt.RefCityCode(d),
            "refPostCode":               u.GetStr(d["POSTCODE"]),
            "refStateCode":              rpt.RefStateCode(d),
            "refCountryCode":            rpt.RefCitizenshipCode(d),
            "refContactTypeCode":        "02",
            "contactInfo":               u.GetStr(d["HOME_PHONE"]),
        }

        nok := fiber.Map{
            "refPersonTitleCode":        rpt.RefPersonTitleCodeNOK(d),
            "fullName":                  u.GetStr(d["PATIENT_NOK_NAME"]),
            "refIdentificationTypeCode": rpt.RefIdentificationTypeCodeNOK(d),
            "identificationNo":          u.GetStr(d["NOK_ID"]),
            "refAddressTypeCode":        "C",
            "street1":                   u.GetStr(d["NOK_STREET1"]),
            "street2":                   u.GetStr(d["NOK_STREET2"]),
            "refCityCode":               rpt.RefCityCodeNOK(d),
            "refPostCode":               u.GetStr(d["NOK_POSTCODE"]),
            "refStateCode":              rpt.RefStateCodeNOK(d),
            "refCountryCode":            rpt.RefCitizenshipCodeNOK(d),
            "refContactTypeCode":        "02",
            "contactInfo":               u.GetStr(d["NOK_MOBILE_PHONE"]),
        }

        m := fiber.Map{
            "rn":                               u.GetStr(d["ACCOUNT_NO"]),
            "mrn":                              u.GetStr(d["PRN"]),
            "eventDate":                        fmt.Sprintf("%s %s:00", d["REGISTRATION_DATE"], d["REGISTRATION_TIME"]),
            "isPoliceCase":                     "02",
            "internalReferral":                 "false",
            "refReferralSourceCode":            rpt.RefReferralSourceCode(d),
            "refGenderCode":                    rpt.RefGenderCode(d),
            "dob":                              u.GetStr(d["DOB"]),
            "refMaritalStatusCode":             rpt.RefMaritalStatusCode(d),
            "refReligionCode":                  rpt.RefReligionCode(d),
            "refCitizenshipCode":               rpt.RefCitizenshipCode(d),
            "refEthnicCode":                    rpt.RefEthnicCode(d),
            "height":                           u.GetNum(u.GetStr(d["HEIGHT"])),
            "weight":                           u.GetNum(u.GetStr(d["WEIGHT"])),
            "refForeignerOriginCountryCode":    rpt.RefForeignerOriginCountryCode(d),
            "refForeignerResidenceCountryCode": rpt.RefForeignerResidenceCountryCode(d),
            "refPersonCategoryCode":            rpt.RefPersonCategoryCode(d),
            "refRelationshipCode":              rpt.RefRelationshipCode(d),
            "totalDurationDay":                 "0",
            "admissionDate":                    fmt.Sprintf("%s %s:00", d["ADMISSION_DATE"], d["ADMISSION_TIME"]),
            "refDischargeTypeCode":             rpt.RefDischargeTypeCode(d),
            "dischargeDateTime":                fmt.Sprintf("%s %s:00", d["DISCHARGE_DATE"], d["DISCHARGE_TIME"]),
            "refDischargeOfficerTypeCode":      "02",
            "mmc":                              "00",
            "refDiagnosisItemTypeCode":         rpt.RefDiagnosisItemTypeCode(d),
            "description":                      u.GetStr(d["ICD10_DESCRIPTION"]),
            "refIcd10Main":                     u.GetStr(d["ICD10_CODE"]),
            "person":                           person,
            "nextOfKins":                       nok,
        }

        forms = append(forms, m)
    }

    facilityCode := config.Config("facilityCode")
    filename := fmt.Sprintf("%s_%s_%s_RH301.json", facilityCode, ds1, ds2)

    c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filename))
    c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
    c.Set(fiber.HeaderPragma, "no-cache")
    c.Set(fiber.HeaderExpires, "0")
    c.Set("filename", filename)
    c.Set(fiber.HeaderContentType, "application/json")

    /* js := "views/master-rh301.django"
    tpl := pongo2.Must(pongo2.FromFile(js))
    s, _ := tpl.Execute(pongo2.Context{
        "filename":           filename,
        "dischargeFrom":      datefrom,
        "dischargeTo":        dateto,
        "refServiceTypeCode": "02",
        "facilityCode":       facilityCode,
        "forms":              forms,
    }) */
    data := fiber.Map{
        "filename":           filename,
        "dischargeFrom":      datefrom,
        "dischargeTo":        dateto,
        "refServiceTypeCode": "02",
        "facilityCode":       facilityCode,
        "forms":              forms,
    }
    jdata, _ := json.MarshalIndent(data, "", "    ")
    return c.SendString(string(jdata))
}

// JsonPD101
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        datefrom           query      string  false  "datefrom"
// @Param        dateto             query      string  false  "dateto"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/export/rpt1 [get]
func JsonPD101(c *fiber.Ctx) error {
    datefrom := c.Query("datefrom", "")
    dateto := c.Query("dateto", "")
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusNotFound, "User not found")
    }

    username := user.Username
    cli := database.GetMongoClient()
    col := getCollection(cli, username, "0")
    ls, err := u.GetCollectionList(database.GetMongoCtx(), col)
    if err != nil {
        return err
    }

    dt1 := strings.Split(datefrom, "-")
    dt2 := strings.Split(dateto, "-")
    ds1 := fmt.Sprintf("%s%s%s", dt1[2], dt1[1], dt1[0])
    ds2 := fmt.Sprintf("%s%s%s", dt2[2], dt2[1], dt2[0])

    forms := make([]fiber.Map, 0)
    for _, d := range ls {
        person := fiber.Map{
            "refPersonTitleCode":        rpt.RefPersonTitleCode(d),
            "fullName":                  u.GetStr(d["PATIENT_NAME"]),
            "refIdentificationTypeCode": rpt.RefIdentificationTypeCode(d),
            "identificationNo":          u.GetStr(d["DOCUMENT_NUMBER"]),
            "refAddressTypeCode":        "C",
            "street1":                   u.GetStr(d["STREET1"]),
            "street2":                   u.GetStr(d["STREET2"]),
            "refCityCode":               rpt.RefCityCode(d),
            "refPostCode":               u.GetStr(d["POSTCODE"]),
            "refStateCode":              rpt.RefStateCode(d),
            "refCountryCode":            rpt.RefCitizenshipCode(d),
            "refContactTypeCode":        "02",
            "contactInfo":               u.GetStr(d["HOME_PHONE"]),
        }

        nok := fiber.Map{
            "refPersonTitleCode":        rpt.RefPersonTitleCodeNOK(d),
            "fullName":                  u.GetStr(d["PATIENT_NOK_NAME"]),
            "refIdentificationTypeCode": rpt.RefIdentificationTypeCodeNOK(d),
            "identificationNo":          u.GetStr(d["NOK_ID"]),
            "refAddressTypeCode":        "C",
            "street1":                   fmt.Sprintf("%v", d["NOK_STREET1"]),
            "street2":                   u.GetStr(d["NOK_STREET2"]),
            "refCityCode":               rpt.RefCityCodeNOK(d),
            "refPostCode":               u.GetStr(d["NOK_POSTCODE"]),
            "refStateCode":              rpt.RefStateCodeNOK(d),
            "refCountryCode":            rpt.RefCitizenshipCodeNOK(d),
            "refContactTypeCode":        "02",
            "contactInfo":               u.GetStr(d["NOK_MOBILE_PHONE"]),
        }

        m := fiber.Map{
            "rn":                               u.GetStr(d["ACCOUNT_NO"]),
            "mrn":                              u.GetStr(d["PRN"]),
            "eventDate":                        fmt.Sprintf("%s %s:00", d["REGISTRATION_DATE"], d["REGISTRATION_TIME"]),
            "isPoliceCase":                     "02",
            "internalReferral":                 "false",
            "refReferralSourceCode":            rpt.RefReferralSourceCode(d),
            "refGenderCode":                    rpt.RefGenderCode(d),
            "dob":                              fmt.Sprintf("%v", d["DOB"]),
            "refMaritalStatusCode":             rpt.RefMaritalStatusCode(d),
            "refReligionCode":                  rpt.RefReligionCode(d),
            "refCitizenshipCode":               rpt.RefCitizenshipCode(d),
            "refEthnicCode":                    rpt.RefEthnicCode(d),
            "height":                           u.GetNum(fmt.Sprintf("%v", d["HEIGHT"])),
            "weight":                           u.GetNum(fmt.Sprintf("%v", d["WEIGHT"])),
            "refForeignerOriginCountryCode":    rpt.RefForeignerOriginCountryCode(d),
            "refForeignerResidenceCountryCode": rpt.RefForeignerResidenceCountryCode(d),
            "refPersonCategoryCode":            rpt.RefPersonCategoryCode(d),
            "refRelationshipCode":              rpt.RefRelationshipCode(d),
            "totalDurationDay":                 "0",
            "refWardTransitionTypeCode":        "A",
            "wardDateTime":                     fmt.Sprintf("%s %s:00", d["ADMISSION_DATE"], d["ADMISSION_TIME"]),
            "wardCode":                         u.GetStr(d["WARD_NO"]),
            "refDisciplineCode":                rpt.RefDisciplineCode(d),
            "refSpecialityCode":                rpt.RefDisciplineCode(d),
            "refSubSpecialityCode":             rpt.RefDisciplineCode(d),
            "refWardClassCode":                 rpt.RefWardClassCode(d),
            "refWardCategoryCode":              "00",
            "refDischargeTypeCode":             rpt.RefDischargeTypeCode(d),
            "dischargeDateTime":                fmt.Sprintf("%s %s:00", d["DISCHARGE_DATE"], d["DISCHARGE_TIME"]),
            "refDischargeOfficerTypeCode":      "02",
            "mmc":                              "00",
            "refDiagnosisItemTypeCode":         rpt.RefDiagnosisItemTypeCode(d),
            "description":                      u.GetStr(d["ICD10_DESCRIPTION"]),
            "refIcd10Main":                     u.GetStr(d["ICD10_CODE"]),
            "person":                           person,
            "nextOfKins":                       nok,
        }

        forms = append(forms, m)
    }

    facilityCode := config.Config("facilityCode")
    filename := fmt.Sprintf("%s_%s_PD301.json", ds1, ds2)

    c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filename))
    c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
    c.Set(fiber.HeaderPragma, "no-cache")
    c.Set(fiber.HeaderExpires, "0")
    c.Set("filename", filename)
    c.Set(fiber.HeaderContentType, "application/json")

    /* js := "views/master-pd301.django"
    tpl := pongo2.Must(pongo2.FromFile(js))
    s, _ := tpl.Execute(pongo2.Context{
        "filename":           filename,
        "dischargeFrom":      datefrom,
        "dischargeTo":        dateto,
        "refServiceTypeCode": "01",
        "facilityCode":       facilityCode,
        "forms":              forms,
    }) */
    data := fiber.Map{
        "filename":           filename,
        "dischargeFrom":      datefrom,
        "dischargeTo":        dateto,
        "refServiceTypeCode": "01",
        "facilityCode":       facilityCode,
        "forms":              forms,
    }
    jdata, _ := json.MarshalIndent(data, "", "    ")
    return c.SendString(string(jdata))
}

// Xlsx
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        vt                 query      string  false  "vt"
// @Param        datefrom           query      string  false  "datefrom"
// @Param        dateto             query      string  false  "dateto"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/export/rpt1/xlsx [get]
func Xlsx(c *fiber.Ctx) error {
    vt := c.Query("vt", "0")
    datefrom := c.Query("datefrom", "")
    dateto := c.Query("dateto", "")
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusNotFound, "User not found")
    }

    username := user.Username
    cli := database.GetMongoClient()
    col := getCollection(cli, username, vt)
    lx, err := u.GetCollectionList(database.GetMongoCtx(), col)
    if err != nil {
        u.LogError(err)
        return err
    }

    dt1 := strings.Split(datefrom, "-")
    dt2 := strings.Split(dateto, "-")
    ds1 := fmt.Sprintf("%s%s%s", dt1[2], dt1[1], dt1[0])
    ds2 := fmt.Sprintf("%s%s%s", dt2[2], dt2[1], dt2[0])
    pf := "RH301"
    if vt == "0" {
        pf = "PD301"
    }

    facilityCode := config.Config("facilityCode")
    filename := fmt.Sprintf("%s_%s_%s_%s.xlsx",facilityCode, ds1, ds2, pf)
    bx, err := u.GetXlsx(COLUMN_MAP, lx)
    if err != nil {
        u.LogError(err)
        return err
    }
   
    c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filename))
    c.Set(fiber.HeaderCacheControl, "no-cache, no-store, must-revalidate")
    c.Set(fiber.HeaderPragma, "no-cache")
    c.Set(fiber.HeaderExpires, "0")
    c.Set("filename", filename)
    c.Set(fiber.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    return c.Send(bx.Bytes())
}

// List
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        _page              query      string  false  "_page"
// @Param        _limit             query      string  false  "_limit"
// @Param        vt                 query      string  false  "vt"
// @Param        datefrom           query      string  false  "datefrom"
// @Param        dateto             query      string  false  "dateto"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/rpt1 [get]
func List(c *fiber.Ctx) error {
    page := c.Query("_page", "1")
    limit := c.Query("_limit", "20")
    vt := c.Query("vt", "0")
    datefrom := c.Query("datefrom", "")
    dateto := c.Query("dateto", "")
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusNotFound, "User not found")
    }

    username := user.Username
    cli := database.GetMongoClient()
    db := getDb(cli, vt)
    col := db.Collection(fmt.Sprintf("__%s__", username))
    col2 := db.Collection(fmt.Sprintf("__%s-q__", username))
    total, err := col.CountDocuments(database.GetMongoCtx(), bson.D{})
    if err != nil {
        u.LogError(err)
        return err
    }

    dateFrom := datefrom
    dateTo := dateto
    t2, err := col2.CountDocuments(database.GetMongoCtx(), bson.D{})
    if err != nil {
        u.LogError(err)
        return err
    }

    if t2 > 0 {
        ld, err := u.FindList(database.GetMongoCtx(), col2)
        if err != nil {
            return err
        }

        dateFrom = fmt.Sprintf("%v", ld[0]["datefrom"])
        dateTo = fmt.Sprintf("%v", ld[0]["dateto"])
    }

    pg := model.GetPager(int(total), page, limit)
    findOptions := options.Find()
    findOptions.SetSkip(int64(pg.GetLowerBound()))
    findOptions.SetLimit(int64(pg.PageSize))
    ls, err := u.GetCollectionList(database.GetMongoCtx(), col, findOptions)
    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "columnmaps":  COLUMN_MAP,
        "total_count": total,
        "total_page":  pg.GetTotalPages(),
        "page":        pg.PageNum,
        "data":        ls,
        "datefrom":    dateFrom,
        "dateto":      dateTo,
    })
}

// SearchList
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        request  body       dto.ReportQueryDto true   "Search Request"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/rpt1 [post]
func SearchList(c *fiber.Ctx) error {
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    data := dto.ReportQueryDto{}
    if err := c.BodyParser(&data); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    errs := u.ValidatePayload(data, c)
    if errs != nil {
        return fiber.NewError(fiber.StatusBadRequest, errs.Error())
    }

    username := user.Username
    md, err := queryAndSave(data, username)
    if err != nil {
        return err
    }

    return c.JSON(md)
}

// Edit
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        id              path      int     true  "id"
// @Param        vt              query     string  false  "vt"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/rpt1/{id} [get]
func Edit(c *fiber.Ctx) error {
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    ids := c.Params("id")
    vt := c.Query("vt", "0")
    username := user.Username
    cli := database.GetMongoClient()
    col := getCollection(cli, username, vt)
    id, err := bson.ObjectIDFromHex(ids)
    if err != nil {
        u.LogError(err)
        return err
    }

    filter := bson.D{{"_id", id}}
    ls, err := u.GetCollectionListWithFilter(database.GetMongoCtx(), filter, col)
    if err != nil {
        return err
    }

    if len(ls) > 0 {
        o := ls[0]
        return c.JSON(o)
    }
    
    return fiber.NewError(fiber.StatusNotFound, "Record not found")
}

// Update
//
// @Tags Report/MasterPD301
// @Produce json
// @Param        id              path      int            true  "id"
// @Param        vt              query     string         false  "vt"
// @Param        request         body      map[string]any true  "Update Request"
// @Security BearerAuth
// @Success 200
// @Router /api/master-pd301/rpt1/{id} [put]
func Update(c *fiber.Ctx) error {
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    data := bson.M{}
    if err := c.BodyParser(&data); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    ids := c.Params("id")
    vt := c.Query("vt", "0")
    username := user.Username
    cli := database.GetMongoClient()
    col := getCollection(cli, username, vt)
    id, err := bson.ObjectIDFromHex(ids)
    if err != nil {
        u.LogError(err)
        return err
    }

    filter := bson.D{{"_id", id}}
    update := bson.M{
        "$set": data,
    }
    col.FindOneAndUpdate(database.GetMongoCtx(), filter, update)
    return c.JSON(fiber.Map{
        "success": 1,
    })
}

func getCollection(cli *mongo.Client, username string, vt string) *mongo.Collection {
    db := getDb(cli, vt)
    s := fmt.Sprintf("__%s__", username)
    col := db.Collection(s)
    return col
}

func getDb(cli *mongo.Client, vt string) *mongo.Database {
    suffix := ""
    var db *mongo.Database
    if config.Config("mongodb.prefix") == "prod" {
        suffix = "_prod"
    }

    if vt == "0" {
        s := fmt.Sprintf("master_pd301%s", suffix)
        db = cli.Database(s)
    } else {
        s := fmt.Sprintf("master_rh301%s", suffix)
        db = cli.Database(s)
    }

    return db
}

func queryAndSave(data dto.ReportQueryDto, username string) (fiber.Map, error) {
    md := make(fiber.Map)
    page := data.Page
    limit := data.Limit
    vt := fmt.Sprintf("%d", data.Vt)
    datefrom := data.DateFrom
    dateto := data.DateTo
    vs := "('INPATIENT')"
    if vt == "1" {
        vs = "('DAY-SURGERY')"
    }

    q := sql.GetMasterPD101(vs)
    db := database.GetDbrs()
    rows, err := db.Query(q, datefrom, dateto)
    if err != nil {
        u.LogError(err)
        return md, err
    }

    defer rows.Close()

    lx, colnames := database.GetDataList(rows)
    ld := make([]bson.M, 0)
    total := len(lx)
    pg := model.Pager{
        Total:    total,
        PageNum:  page,
        PageSize: limit,
    }
    pg.SetPageSize(pg.PageSize)

    if total > 0 {
        cli := database.GetMongoClient()
        dm := getDb(cli, vt)
        col := dm.Collection(fmt.Sprintf("__%s__", username))
        col.Drop(database.GetMongoCtx())
        _, err := col.InsertMany(database.GetMongoCtx(), lx)
        if err != nil {
            u.LogError(err)
            return md, err
        }

        col1 := dm.Collection(fmt.Sprintf("__%s-c__", username))
        col1.Drop(database.GetMongoCtx())
        _, err = col1.InsertOne(database.GetMongoCtx(), bson.M{"columns": colnames})
        if err != nil {
            u.LogError(err)
            return md, err
        }

        col2 := dm.Collection(fmt.Sprintf("__%s-q__", username))
        col2.Drop(database.GetMongoCtx())
        _, err = col2.InsertOne(database.GetMongoCtx(), bson.M{"datefrom": datefrom, "dateto": dateto})
        if err != nil {
            u.LogError(err)
            return md, err
        }

        findOptions := options.Find()
        findOptions.SetSkip(int64(pg.GetLowerBound()))
        findOptions.SetLimit(int64(pg.PageSize))
        ld, err = u.GetCollectionList(database.GetMongoCtx(), col, findOptions)
        if err != nil {
            return md, err
        }
    }

    return fiber.Map{
        "columnmaps":  COLUMN_MAP,
        "total_count": total,
        "total_page":  pg.GetTotalPages(),
        "page":        pg.PageNum,
        "data":        ld,
    }, nil
}
