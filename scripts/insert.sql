INSERT INTO public."role"
("name")
VALUES('ADMIN');

INSERT INTO public."role"
("name")
VALUES('USER');


INSERT INTO public.app_user
(active, first_name, last_name, "password", username)
VALUES(true, 'sys admin', '', '$2a$10$gp9o1kUR6kYXv0qNDIhPpO8pTVEo/bsRdo/vbaEbJmFBqpt47n1fa', 'admin');


INSERT INTO public.app_user_roles
(app_user_id, roles_id)
VALUES(1, 1);
