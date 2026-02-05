package masterpd102

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
    { Field: "PERSON_HEIGHT", Text: "HEIGHT" },
    { Field: "PERSON_WEIGHT", Text: "WEIGHT" },

    { Field: "COUNTRY_OF_BIRTH", Text: "COUNTRY OF BIRTH" },
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
    { Field: "STREET1", Text: "STREET1" },
    { Field: "STREET2", Text: "STREET2" },
    { Field: "CITYCODE", Text: "STREET3" },
    { Field: "POSTCODE", Text: "POSTCODE" },

    { Field: "OCITY", Text: "STATE" },
    { Field: "COUNTRY", Text: "COUNTRY" },
    { Field: "NOK_MOBILE_PHONE", Text: "NOK MOBILE NO" },
    { Field: "ADMISSION_DATE", Text: "ADMISSION DATE" },
    { Field: "ADMISSION_TIME", Text: "ADMISSION TIME" },

    { Field: "WARD_NO", Text: "WARD NO" },
    { Field: "PAYMENT_CLASS_CODE", Text: "PAYMENT CLASS" },

    { Field: "GRAVIDA", Text: "GRAVIDA" },
    { Field: "PARITY", Text: "PARITY" },
    { Field: "GESTATION_PERIOD", Text: "GESTATION PERIOD" },
    { Field: "ISMOTHERALIVE", Text: "IS MOTHER ALIVE" },
    { Field: "REFANTENATALCARECODE", Text: "ANTENATAL CARE" },
    { Field: "LABOUR_METHOD", Text: "LABOUR METHOD" },

    { Field: "DELIVERY_DATE", Text: "DELIVERY DATE" },
    { Field: "RESULT_OF_BIRTH", Text: "RESULT OF BIRTH" },
    { Field: "DELIVERY_TYPE", Text: "DELIVERY TYPE" },
    { Field: "CHILD_SEX", Text: "CHILD SEX" },
    { Field: "WEIGHT", Text: "WEIGHT" },
    { Field: "LENGTH", Text: "LENGTH" },
}
