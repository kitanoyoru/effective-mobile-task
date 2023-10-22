-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE "Person"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "name"       TEXT NOT NULL,
    "surname"    TEXT NOT NULL,
    "patronymic" TEXT,
    "age"        BIGINT

);

CREATE TABLE "person_gender"
(
    "id"          SMALLINT PRIMARY KEY,
    "gender"      TEXT    NOT NULL,
    "probability" DECIMAL NOT NULL,
    "personid"    BIGINT,

    CONSTRAINT "FK_Person_gender_Person" FOREIGN KEY ("personid") REFERENCES "Person" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE "person_country"
(
    "id"          SMALLINT PRIMARY KEY,
    "probability" DECIMAL NOT NULL,
    "personid"    BIGINT,

    CONSTRAINT "FK_Person_Country_Person" FOREIGN KEY ("personid") REFERENCES "Person" ("id") ON DELETE SET NULL ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
