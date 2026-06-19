-- Aufruf: psql --dbname=kunde --username=postgres --file=/init/kunde/sql/copy-csv.sql

SET search_path TO kunde;

TRUNCATE TABLE bestellung, adresse, kunde RESTART IDENTITY CASCADE;

COPY kunde (id, nachname, email, username, version)
FROM '/init/kunde/csv/kunde.csv'
WITH (FORMAT csv, HEADER true);

COPY adresse (id, strasse, hausnummer, plz, ort, kunde_id)
FROM '/init/kunde/csv/adresse.csv'
WITH (FORMAT csv, HEADER true);

COPY bestellung (id, produktname, menge, kunde_id)
FROM '/init/kunde/csv/bestellung.csv'
WITH (FORMAT csv, HEADER true);

SELECT setval(pg_get_serial_sequence('kunde', 'id'), COALESCE((SELECT MAX(id) FROM kunde), 1), true);
SELECT setval(pg_get_serial_sequence('adresse', 'id'), COALESCE((SELECT MAX(id) FROM adresse), 1), true);
SELECT setval(pg_get_serial_sequence('bestellung', 'id'), COALESCE((SELECT MAX(id) FROM bestellung), 1), true);
