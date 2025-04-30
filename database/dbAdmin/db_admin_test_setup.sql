-- export DB_ADMIN_PASSWORD="password"
-- psql -h your_host -U your_user -d your_database -v db_admin_password="'$DB_ADMIN_PASSWORD'" -f your_script.sql

-- Create the db_admin role with a login password injected via env variable
CREATE USER db_admin WITH LOGIN PASSWORD :'db_admin_password' CREATEROLE;

-- give schema_ownership to the db_admin
ALTER DATABASE c2nofficialtestdb OWNER TO db_admin;

SET ROLE db_admin;