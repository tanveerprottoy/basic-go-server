-- DROP CREATE DB
DROP DATABASE IF EXISTS basic_server_db;
CREATE DATABASE basic_server_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS resources;
CREATE TABLE resources (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    `name` VARCHAR NOT NULL,
    created_at BIGINT,
    updated_at BIGINT
);