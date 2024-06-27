CREATE SCHEMA consumer_campaing_pg;

GRANT ALL PRIVILEGES ON DATABASE "consumer-campaing-pg" TO "postgres";

GRANT USAGE ON SCHEMA consumer_campaing_pg TO "postgres";
ALTER USER "postgres" SET search_path = 'consumer_campaing_pg';

SET SCHEMA 'consumer-campaing-pg';
ALTER DEFAULT PRIVILEGES
    IN SCHEMA consumer_campaing_pg
GRANT SELECT, UPDATE, INSERT, DELETE ON TABLES
    TO "postgres";

ALTER DEFAULT PRIVILEGES
    IN SCHEMA consumer_campaing_pg
GRANT USAGE ON SEQUENCES
    TO "postgres";