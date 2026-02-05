-- public.app_user definition

-- Drop table

-- DROP TABLE public.app_user;

CREATE TABLE public.app_user (
	id bigserial NOT NULL,
	active bool NOT NULL,
	first_name varchar(150) NOT NULL,
	last_name varchar(150) NULL,
	"password" varchar(150) NOT NULL,
	username varchar(150) NOT NULL,
	last_login timestamp(0) NULL,
	CONSTRAINT app_user_pkey PRIMARY KEY (id),
	CONSTRAINT uk_3k4cplvh82srueuttfkwnylq0 UNIQUE (username)
);

-- public."role" definition

-- Drop table

-- DROP TABLE public."role";

CREATE TABLE public."role" (
	id bigserial NOT NULL,
	"name" varchar(150) NOT NULL,
	CONSTRAINT role_pkey PRIMARY KEY (id)
);

-- public.app_user_roles definition

-- Drop table

-- DROP TABLE public.app_user_roles;

CREATE TABLE public.app_user_roles (
    id bigserial NOT NULL,
	app_user_id int8 NOT NULL,
	roles_id int8 NOT NULL,
	CONSTRAINT app_user_roles_pkey PRIMARY KEY (id)
);


-- public.app_user_roles foreign keys

ALTER TABLE public.app_user_roles ADD CONSTRAINT fk23e7b5jyl3ql41rk3566gywdd FOREIGN KEY (roles_id) REFERENCES role(id);
ALTER TABLE public.app_user_roles ADD CONSTRAINT fkkwxexnudtp5gmt82j0qtytnoe FOREIGN KEY (app_user_id) REFERENCES app_user(id);

-- public.adm_status definition

-- Drop table

-- DROP TABLE public.adm_status;

CREATE TABLE public.adm_status (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT adm_status_pkey PRIMARY KEY (id)
);

-- public.city definition

-- Drop table

-- DROP TABLE public.city;

CREATE TABLE public.city (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT city_pkey PRIMARY KEY (id)
);

-- public.country definition

-- Drop table

-- DROP TABLE public.country;

CREATE TABLE public.country (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT country_pkey PRIMARY KEY (id)
);

-- public.diag_item_type definition

-- Drop table

-- DROP TABLE public.diag_item_type;

CREATE TABLE public.diag_item_type (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT diag_item_type_pkey PRIMARY KEY (id)
);

-- public.discharge_officer definition

-- Drop table

-- DROP TABLE public.discharge_officer;

CREATE TABLE public.discharge_officer (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT discharge_officer_pkey PRIMARY KEY (id)
);

-- public.discharge_type definition

-- Drop table

-- DROP TABLE public.discharge_type;

CREATE TABLE public.discharge_type (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT discharge_type_pkey PRIMARY KEY (id)
);

-- public.education definition

-- Drop table

-- DROP TABLE public.education;

CREATE TABLE public.education (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT education_pkey PRIMARY KEY (id)
);

-- public.ethnic_group definition

-- Drop table

-- DROP TABLE public.ethnic_group;

CREATE TABLE public.ethnic_group (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT ethnic_group_pkey PRIMARY KEY (id)
);

-- public.gender definition

-- Drop table

-- DROP TABLE public.gender;

CREATE TABLE public.gender (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT gender_pkey PRIMARY KEY (id)
);

-- public.id_type definition

-- Drop table

-- DROP TABLE public.id_type;

CREATE TABLE public.id_type (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT id_type_pkey PRIMARY KEY (id)
);

-- public.income definition

-- Drop table

-- DROP TABLE public.income;

CREATE TABLE public.income (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT income_pkey PRIMARY KEY (id)
);

-- public.marital_status definition

-- Drop table

-- DROP TABLE public.marital_status;

CREATE TABLE public.marital_status (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT marital_status_pkey PRIMARY KEY (id)
);

-- public.occupation definition

-- Drop table

-- DROP TABLE public.occupation;

CREATE TABLE public.occupation (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT occupation_pkey PRIMARY KEY (id)
);

-- public.person_category_code definition

-- Drop table

-- DROP TABLE public.person_category_code;

CREATE TABLE public.person_category_code (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT person_category_code_pkey PRIMARY KEY (id)
);

-- public.referral definition

-- Drop table

-- DROP TABLE public.referral;

CREATE TABLE public.referral (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT referral_pkey PRIMARY KEY (id)
);

-- public.relationship definition

-- Drop table

-- DROP TABLE public.relationship;

CREATE TABLE public.relationship (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT relationship_pkey PRIMARY KEY (id)
);

-- public.religion definition

-- Drop table

-- DROP TABLE public.religion;

CREATE TABLE public.religion (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT religion_pkey PRIMARY KEY (id)
);

-- public.speciality definition

-- Drop table

-- DROP TABLE public.speciality;

CREATE TABLE public.speciality (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT speciality_pkey PRIMARY KEY (id)
);

-- public.state definition

-- Drop table

-- DROP TABLE public.state;

CREATE TABLE public.state (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT state_pkey PRIMARY KEY (id)
);

-- public.title definition

-- Drop table

-- DROP TABLE public.title;

CREATE TABLE public.title (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT title_pkey PRIMARY KEY (id)
);

-- public.visit_type definition

-- Drop table

-- DROP TABLE public.visit_type;

CREATE TABLE public.visit_type (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT visit_type_pkey PRIMARY KEY (id)
);

-- public.ward_class definition

-- Drop table

-- DROP TABLE public.ward_class;

CREATE TABLE public.ward_class (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT ward_class_pkey PRIMARY KEY (id)
);

-- public.delivery_type definition

-- Drop table

-- DROP TABLE public.delivery_type;

CREATE TABLE public.delivery_type (
	id bigserial NOT NULL,
	code varchar(30) NOT NULL,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL,
	deleted bool NULL,
	deleted_by int8 NULL,
	deleted_date timestamp NULL,
	"desc" varchar(300) NOT NULL,
	modified_by int8 NULL,
	modified_date timestamp NULL,
	"ref" varchar(200) NOT NULL,
	CONSTRAINT delivery_type_pkey PRIMARY KEY (id)
);
