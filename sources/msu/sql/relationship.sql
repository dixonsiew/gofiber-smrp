INSERT INTO public.relationship (code,created_by,created_date,"desc","ref") VALUES
	 ('0',1,'2021-05-04 00:00:00','N/A','No Information'),
	 ('1',1,'2021-05-04 00:00:00','','Wife'),
	 ('2',1,'2021-05-04 00:00:00','','Husband'),
	 ('3',1,'2021-05-04 00:00:00','MOTHER','Mother'),
	 ('4',1,'2021-05-04 00:00:00','FATHER','Father'),
	 ('5',1,'2021-05-04 00:00:00','','Biological Child'),
	 ('6',1,'2021-05-04 00:00:00','','Step Child'),
	 ('7',1,'2021-05-04 00:00:00','','Adopted Child'),
	 ('9',1,'2021-05-04 00:00:00','','Grandmother'),
	 ('10',1,'2021-05-04 00:00:00','','Grandfather');
INSERT INTO public.relationship (code,created_by,created_date,"desc","ref") VALUES
	 ('11',1,'2021-05-04 00:00:00','','Sibling'),
	 ('12',1,'2021-05-04 00:00:00','','Guardian'),
	 ('13',1,'2021-05-04 00:00:00','','Self'),
	 ('99',1,'2021-05-04 00:00:00','OTHER','Others'),
	 ('99',1,'2021-05-04 00:00:00','SPOUSE','Others'),
	 ('99',1,'2021-05-04 00:00:00','SON','Others'),
	 ('99',1,'2021-05-04 00:00:00','DAUGHTER','Others'),
	 ('99',1,'2021-05-04 00:00:00','MOTHER-IN-LAW','Others'),
	 ('99',1,'2021-05-04 00:00:00','FATHER-IN-LAW','Others'),
	 ('99',1,'2021-05-04 00:00:00','BROTHER-IN-LAW','Others');
INSERT INTO public.relationship (code,created_by,created_date,"desc","ref") VALUES
	 ('99',1,'2021-05-04 00:00:00','SISTER-IN-LAW','Others'),
	 ('11',1,'2021-05-04 00:00:00','SISTER','Sibling'),
	 ('11',1,'2021-05-04 00:00:00','BROTHER','Sibling'),
	 ('99',1,'2021-05-04 00:00:00','GRANDPARENT','Others');