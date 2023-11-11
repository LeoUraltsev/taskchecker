-- +goose Up
-- +goose StatementBegin
alter table "user"
    add updated_at timestamp without time zone default now() not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table if exists "user";
-- +goose StatementEnd
