-- export DB_ADMIN_PASSWORD="password"
-- psql -h your_host -U your_user -d your_database -v db_admin_password="'$DB_ADMIN_PASSWORD'" -f your_script.sql

-- Create the db_admin role
CREATE USER db_admin WITH LOGIN PASSWORD :'db_admin_password';

-- give schema_ownership to the db_admin
ALTER SCHEMA public OWNER TO db_admin;

-- Give db_admin control over the database
GRANT ALL PRIVILEGES ON DATABASE c2nofficialsitetestdb TO db_admin;

-- Allow db_admin control over all tables, sequences, functions in public schema
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO db_admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO db_admin;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO db_admin;

-- Make sure future tables/sequences/functions automatically get privileges too
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL PRIVILEGES ON TABLES TO db_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL PRIVILEGES ON SEQUENCES TO db_admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL PRIVILEGES ON FUNCTIONS TO db_admin;