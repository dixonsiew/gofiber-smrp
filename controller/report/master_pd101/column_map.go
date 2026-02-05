package masterpd101

import (
    "smrp/controller/report"
)

var COLUMN_MAP = []report.ColumnMap{
    { Field: "ACCOUNT_NO", Text: "ACCOUNT NO" },
    { Field: "PRN", Text: "PRN" },
    { Field: "REGISTRATION_DATE", Text: "REG DATE" },
    { Field: "REGISTRATION_TIME", Text: "REG TIME" },
    { Field: "TITLE", Text: "TITLE" },

    { Field: "PATIENT_NAME", Text: "NAME" },
    { Field: "GENDER", Text: "GENDER" },
    { Field: "DOB", Text: "DOB" },
    { Field: "MARITAL_STATUS", Text: "MARITAL STATUS" },
    { Field: "RELIGION", Text: "RELIGION" },

    { Field: "NATIONALITY", Text: "NATIONALITY" },
    { Field: "ETHNIC_GROUP", Text: "ETHNIC GROUP" },
    { Field: "OCCUPATION", Text: "OCCUPATION" },
    { Field: "HEIGHT", Text: "HEIGHT" },
    { Field: "WEIGHT", Text: "WEIGHT" },

    { Field: "COUNTRY_OF_BIRTH", Text: "COUNTRY OF BIRTH" },
    { Field: "REFPERSONCATEGORYCODE", Text: "PATIENT CATEGORY" },
    { Field: "DOCUMENT_TYPE", Text: "DOC TYPE" },
    { Field: "DOCUMENT_NUMBER", Text: "DOC NO" },
    { Field: "STREET1", Text: "STREET1" },

    { Field: "STREET2", Text: "STREET2" },
    { Field: "CITYCODE", Text: "STREET3" },
    { Field: "POSTCODE", Text: "POSTCODE" },
    { Field: "OCITY", Text: "STATE" },
    { Field: "COUNTRY", Text: "COUNTRY" },

    { Field: "HOME_PHONE", Text: "HOME PHONE" },
    { Field: "NOK_TITLE", Text: "NOK TITLE" },
    { Field: "PATIENT_NOK_NAME", Text: "NOK NAME" },
    { Field: "RELATION_DESCRIPTION", Text: "RELATIONSHIP" },
    { Field: "NOK_ID_TYPE", Text: "NOK DOC TYPE" },

    { Field: "NOK_ID", Text: "NOK DOC NO" },
    { Field: "NOK_STREET1", Text: "STREET1" },
    { Field: "NOK_STREET2", Text: "STREET2" },
    { Field: "NOK_CITYCODE", Text: "STREET3" },
    { Field: "NOK_POSTCODE", Text: "POSTCODE" },

    { Field: "NOK_OCITY", Text: "STATE" },
    { Field: "COUNTRY", Text: "COUNTRY" },
    { Field: "NOK_MOBILE_PHONE", Text: "NOK MOBILE NO" },
    { Field: "ADMISSION_DATE", Text: "ADMISSION DATE" },
    { Field: "ADMISSION_TIME", Text: "ADMISSION TIME" },

    { Field: "WARD_NO", Text: "WARD NO" },
    { Field: "PRIMARY_SPECIALITY", Text: "PRIMARY SPECIALITY" },
    { Field: "PAYMENT_CLASS_CODE", Text: "PAYMENT CLASS" },
}
