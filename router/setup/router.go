package setup

import (
    "smrp/router/setup/adm_status"
	"smrp/router/setup/city"
	"smrp/router/setup/country"
	"smrp/router/setup/delivery_type"
	"smrp/router/setup/diag_item_type"
	"smrp/router/setup/discharge_officer"
	"smrp/router/setup/discharge_type"
	"smrp/router/setup/education"
	"smrp/router/setup/ethnic_group"
	"smrp/router/setup/gender"
    "smrp/router/setup/id_type"
    "smrp/router/setup/income"
    "smrp/router/setup/marital_status"
    "smrp/router/setup/occupation"
    "smrp/router/setup/person_category_code"
    "smrp/router/setup/referral"
    "smrp/router/setup/relationship"
    "smrp/router/setup/religion"
    "smrp/router/setup/role"
    "smrp/router/setup/speciality"
    "smrp/router/setup/state"
    "smrp/router/setup/title"
    "smrp/router/setup/user"
    "smrp/router/setup/visit_type"
    "smrp/router/setup/ward_class"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
    admstatus.SetupRoutes(router)
    city.SetupRoutes(router)
    country.SetupRoutes(router)
    deliverytype.SetupRoutes(router)
    diagitemtype.SetupRoutes(router)
    dischargeofficer.SetupRoutes(router)
    dischargetype.SetupRoutes(router)
    education.SetupRoutes(router)
    ethnicgroup.SetupRoutes(router)
    gender.SetupRoutes(router)
    idtype.SetupRoutes(router)
    income.SetupRoutes(router)
    maritalstatus.SetupRoutes(router)
    occupation.SetupRoutes(router)
    personcategorycode.SetupRoutes(router)
    referral.SetupRoutes(router)
    relationship.SetupRoutes(router)
    religion.SetupRoutes(router)
    role.SetupRoutes(router)
    speciality.SetupRoutes(router)
    state.SetupRoutes(router)
    title.SetupRoutes(router)
    user.SetupRoutes(router)
    visittype.SetupRoutes(router)
    wardclass.SetupRoutes(router)
}
