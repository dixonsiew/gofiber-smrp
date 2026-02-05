package report

import (
    "fmt"
    cs "smrp/service/common_setup"
    "smrp/utils"
    "strings"

    "go.mongodb.org/mongo-driver/v2/bson"
)

func RefReferralSourceCode(doc bson.M) string {
    return getCode("REFERRAL", doc, "referral")
}

func RefPersonTitleCode(doc bson.M) string {
    return getCode("TITLE", doc, "title")
}

func RefGenderCode(doc bson.M) string {
    return getCode("GENDER", doc, "gender")
}

func RefGenderCode1(doc bson.M) string {
    return getCode("CHILD_SEX", doc, "gender")
}

func RefMaritalStatusCode(doc bson.M) string {
    return getCode("MARITAL_STATUS", doc, "marital_status")
}

func RefReligionCode(doc bson.M) string {
    return getCode("RELIGION", doc, "religion")
}

func RefCitizenshipCode(doc bson.M) string {
    return getCode("NATIONALITY", doc, "country")
}

func RefCitizenshipCodeNOK(doc bson.M) string {
    return getCode("NOK_NATIONALITY", doc, "country")
}

func RefEthnicCode(doc bson.M) string {
    return getCode("ETHNIC_GROUP", doc, "ethnic_group")
}

func RefForeignerOriginCountryCode(doc bson.M) string {
    return getCode("COUNTRY_OF_BIRTH", doc, "country")
  }

func RefForeignerResidenceCountryCode(doc bson.M) string {
    return getCode("REFFOREIGNRCOUNTRYCODE", doc, "country")
}

func RefPersonCategoryCode(doc bson.M) string {
    return getCode("REFPERSONCATEGORYCODE", doc, "person_category_code")
}

func RefIdentificationTypeCode(doc bson.M) string {
    return getCode("DOCUMENT_TYPE", doc, "id_type")
}

func RefCityCode(doc bson.M) string {
    return getCode("CITYCODE", doc, "city")
}

func RefCityCodeNOK(doc bson.M) string {
    return getCode("NOK_CITYCODE", doc, "city")
}

func RefStateCode(doc bson.M) string {
    return getCode("OCITY", doc, "state")
}

func RefStateCodeNOK(doc bson.M) string {
    return getCode("NOK_OCITY", doc, "state")
}

func RefPersonTitleCodeNOK(doc bson.M) string {
    return getCode("NOK_TITLE", doc, "title")
}

func RefRelationshipCode(doc bson.M) string {
    return getCode("RELATION_DESCRIPTION", doc, "relationship")
}

func RefIdentificationTypeCodeNOK(doc bson.M) string {
    return getCode("NOK_ID_TYPE", doc, "id_type")
}

func RefDisciplineCode(doc bson.M) string {
    return getCode("PRIMARY_SPECIALITY", doc, "speciality")
}

func RefDisciplineCode1(doc bson.M) string {
    return getCode("PRIMARY_SPECIALTY", doc, "speciality")
}

func RefWardClassCode(doc bson.M) string {
    return getCode("PAYMENT_CLASS_CODE", doc, "ward_class")
}

func RefDischargeTypeCode(doc bson.M) string {
    return getCode("DISCHARGE_REASON", doc, "discharge_type")
}

func RefDiagnosisItemTypeCode(doc bson.M) string {
    return getCode("DIAGNOSIS_DESC", doc, "diag_item_type")
}

func RefLabourModeCode(doc bson.M) string {
    return getCode("DELIVERY_TYPE", doc, "delivery_type")
}

func getCode(key string, doc bson.M, table string) string {
    x := utils.NO_INFO
    s := get(key, doc)
    if s != "" {
        o, err := cs.FindByDesc(s, table)
        if err == nil && o != nil {
            x = o.Code
        }
    }

    return x
}

func get(key string, doc bson.M) string {
    s := ""
    v, ok := doc[key]
    if ok {
        s = fmt.Sprintf("%v", v)
        s = strings.TrimSpace(s)
    }

    return s
}
