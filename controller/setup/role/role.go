package role

import (
    roleService "smrp/service/role"

    "github.com/gofiber/fiber/v2"
)

// LookupList
//
// @Tags Setup/Role
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Role
// @Router /api/lookup/groups [get]
func LookupList(c *fiber.Ctx) error {
    lx, err := roleService.FindAll("name", "asc")
    if err != nil {
        return err
    }

    return c.JSON(lx)
}
