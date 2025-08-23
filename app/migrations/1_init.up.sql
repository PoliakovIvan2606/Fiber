BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;


-- TABLES --

CREATE TABLE public.book
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(50) NOT NULL,
    number_pages INT CHECK (number_pages > 0)
);

-- name
-- description
-- author
-- number_pages


COMMIT;