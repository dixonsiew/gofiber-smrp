package admstatus

import (
	"fmt"
	"smrp/dto"
	"smrp/middleware"
	"smrp/model"
	cs "smrp/service/common_setup"
	"smrp/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var table string = "adm_status"

// LookupList
//
// @Tags Setup/AdmStatus
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.CommonSetup
// @Router /api/lookup/adm-statuses [get]
func LookupList(c *fiber.Ctx) error {
    _, _, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    lx, err := cs.FindAll(table, 0, 0, "", "")
    if err != nil {
        return err
    }

    return c.JSON(lx)
}

// List
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        _page              query      string  false  "_page"
// @Param        _limit             query      string  false  "_limit"
// @Param        sort               query      string  false  "sort"
// @Security BearerAuth
// @Success 200 {array} model.CommonSetup
// @Router /api/adm-statuses [get]
func List(c *fiber.Ctx) error {
    page := c.Query("_page", "1")
    limit := c.Query("_limit", "20")
    sorts := c.Query("sort", "")

    sortBy := "code"
    sortDir := "asc"

    if sorts != "" {
        lis := strings.Split(sorts, "$")
        s := lis[0]
        arr := strings.Split(s, ":")
        sortBy = arr[0]
        sortDir = arr[1]
    }

    total, err := cs.Count(table)
    if err != nil {
        return err
    }

    pg := model.GetPager(total, page, limit)
    lx, err := cs.FindAll(table, pg.GetLowerBound(), pg.PageSize, sortBy, sortDir)
    if err != nil {
        return err
    }

    c.Set(utils.X_TOTAL_COUNT, fmt.Sprintf("%d", total))
    c.Set(utils.X_TOTAL_PAGE, fmt.Sprintf("%d", pg.GetTotalPages()))
    return c.JSON(lx)
}

// SearchList
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        _page    query      string         false  "_page"
// @Param        _limit   query      string         false  "_limit"
// @Param        sort     query      string         false  "sort"
// @Param        request  body       dto.KeywordDto true   "Search Request"
// @Security BearerAuth
// @Success 200 {array} model.CommonSetup
// @Router /api/adm-statuses [post]
func SearchList(c *fiber.Ctx) error {
    page := c.Query("_page", "1")
    limit := c.Query("_limit", "20")
    sorts := c.Query("sort", "")

    data := dto.KeywordDto{}
    if err := c.BodyParser(&data); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
    }

    key := fmt.Sprintf("%%%s%%", data.Keyword)
    sortBy := "code"
    sortDir := "asc"

    if sorts != "" {
        lis := strings.Split(sorts, "$")
        s := lis[0]
        arr := strings.Split(s, ":")
        sortBy = arr[0]
        sortDir = arr[1]
    }

    total, err := cs.CountByKeyword(key, table)
    if err != nil {
        return err
    }

    pg := model.GetPager(total, page, limit)
    lx, err := cs.FindByKeyword(key, pg.GetLowerBound(), pg.PageSize, sortBy, sortDir, table)
    if err != nil {
        return err
    }

    c.Set(utils.X_TOTAL_COUNT, fmt.Sprintf("%d", total))
    c.Set(utils.X_TOTAL_PAGE, fmt.Sprintf("%d", pg.GetTotalPages()))
    return c.JSON(lx)
}

// Create
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        request  body       dto.CommonSetupDto true   "Create City Request"
// @Security BearerAuth
// @Success 200
// @Router /api/adm-status [post]
func Create(c *fiber.Ctx) error {
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    data := dto.CommonSetupDto{}
    if err := c.BodyParser(&data); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    errs := utils.ValidatePayload(data, c)
    if errs != nil {
        return fiber.NewError(fiber.StatusBadRequest, errs.Error())
    }

    o := model.CommonSetup{
        Code: data.Code,
        Desc: data.Desc,
        Ref: data.Ref,
        CreatedBy: user.Id,
    }
    cs.Save(o, table)
    return c.JSON(fiber.Map{
        "success": 1,
    })
}

// Edit
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        id              path      int  true  "id"
// @Security BearerAuth
// @Success 200 {object} model.CommonSetup
// @Router /api/adm-status/{id} [get]
func Edit(c *fiber.Ctx) error {
    ids := c.Params("id")
    id, _ := strconv.Atoi(ids)
    o, err := cs.FindById(id, table)
    if err != nil {
        return err
    }

    if o == nil {
        return fiber.NewError(fiber.StatusNotFound, "Record not found")
    }

    return c.JSON(o)
}

// Update
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        id              path      int                true  "id"
// @Param        request         body      dto.CommonSetupDto true  "Update City Request"
// @Security BearerAuth
// @Success 200
// @Router /api/adm-status/{id} [put]
func Update(c *fiber.Ctx) error {
    ids := c.Params("id")
    id, _ := strconv.Atoi(ids)
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    data := dto.CommonSetupDto{}
    if err := c.BodyParser(&data); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    errs := utils.ValidatePayload(data, c)
    if errs != nil {
        return fiber.NewError(fiber.StatusBadRequest, errs.Error())
    }

    o, err := cs.FindById(id, table)
    if err != nil {
        return err
    }

    if o == nil {
        return fiber.NewError(fiber.StatusNotFound, "Record not found")
    }

    o.Code = data.Code
    o.Desc = data.Desc
    o.Ref = data.Ref
    o.ModifiedBy = user.Id
    cs.Update(*o, table)
    return c.JSON(fiber.Map{
        "success": 1,
    })
}

// Delete
//
// @Tags Setup/AdmStatus
// @Produce json
// @Param        id              path      int                true  "id"
// @Security BearerAuth
// @Success 200
// @Router /api/adm-status/{id} [delete]
func Delete(c *fiber.Ctx) error {
    ids := c.Params("id")
    id, _ := strconv.Atoi(ids)
    _, user, err := middleware.ValidateToken(c)
    if err != nil {
        return err
    }

    if user == nil {
        return fiber.NewError(fiber.StatusUnauthorized, "User not found")
    }

    err = cs.DeleteById(id, user.Id, table)
    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "success": 1,
    })
}
