-- Create the db_admin role with a login password injected via env variable
CREATE USER db_admin WITH LOGIN PASSWORD :'db_admin_password' CREATEROLE;

-- Give db ownership to db_admin
ALTER DATABASE c2nofficialdb OWNER TO db_admin;

-- db_admin now can manage the db
SET ROLE db_admin;