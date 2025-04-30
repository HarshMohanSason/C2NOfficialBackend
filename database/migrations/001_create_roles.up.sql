CREATE ROLE c2n_admin;
CREATE ROLE c2n_user;

GRANT c2n_admin TO db_admin;
GRANT c2n_user TO db_admin;