-- Copyright (C) 2022 - present Juergen Zimmermann, Hochschule Karlsruhe
--
-- This program is free software: you can redistribute it and/or modify
-- it under the terms of the GNU General Public License as published by
-- the Free Software Foundation, either version 3 of the License, or
-- (at your option) any later version.
--
-- This program is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
-- GNU General Public License for more details.
--
-- You should have received a copy of the GNU General Public License
-- along with this program.  If not, see <https://www.gnu.org/licenses/>.

-- Aufruf:   psql --dbname=kunde --username=kunde --file=/init/kunde/sql/create-table.sql

-- Schema "kunde" wurde bereits durch create-schema.sql angelegt; der
-- search_path der Rolle "kunde" ist bereits darauf gesetzt, daher reicht
-- hier ein unqualifizierter Tabellenname.

-- Spalten passend zu csv/kunde.csv (id, nachname, email, username, version).
-- "version" ist die klassische Spalte für optimistisches Sperren (Optimistic
-- Locking) bei nebenläufigen Änderungen.
CREATE TABLE IF NOT EXISTS kunde
(
    id       BIGINT GENERATED ALWAYS AS IDENTITY (START WITH 1) PRIMARY KEY,
    nachname VARCHAR(40)    NOT NULL,
    email    VARCHAR(40)    NOT NULL UNIQUE,
    username VARCHAR(40)    NOT NULL UNIQUE,
    version  INTEGER        NOT NULL DEFAULT 0
) TABLESPACE kundespace;
