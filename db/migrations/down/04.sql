SELECT 'db/migrations/down/4.sql' AS filepath;

DROP TRIGGER zip_codes_touch_trg ON zip_codes;

DROP TABLE zip_codes;

DELETE FROM migrations WHERE version = 4;