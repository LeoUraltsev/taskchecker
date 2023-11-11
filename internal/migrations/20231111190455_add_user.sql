-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"
(
    id         SERIAL primary key,
    username   varchar(128)                not null,
    email      varchar(256),
    password   varchar(64)                 not null,
    created_at timestamp without time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table if exists "user";
-- +goose StatementEnd
