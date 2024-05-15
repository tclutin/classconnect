-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS public.groups (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    code TEXT NOT NULL UNIQUE,
    is_exists_schedule BOOL NOT NULL DEFAULT FALSE,
    members_count INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS public.weeks (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT,
    is_even BOOLEAN NOt NULL,
    FOREIGN KEY (group_id) REFERENCES public.groups(id)
);

CREATE TABLE IF NOT EXISTS public.days (
    id BIGSERIAL PRIMARY KEY,
    week_id BIGINT,
    day_of_week INT NOT NULL,
    FOREIGN KEY (week_id) REFERENCES public.weeks(id)
);

CREATE TABLE IF NOT EXISTS public.subjects (
    id BIGSERIAL PRIMARY KEY,
    day_id BIGINT,
    teacher TEXT NOT NULL,
    name TEXT NOT NULL,
    cabinet TEXT NOT NULL,
    description TEXT NOT NULL,
    time_start TIME NOT NULL,
    time_end TIME NOT NULL,
    FOREIGN KEY (day_id) REFERENCES public.days(id)
);

CREATE TABLE IF NOT EXISTS public.users (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (group_id) REFERENCES public.groups(id)
);

CREATE TABLE IF NOT EXISTS public.subscribers (
    id BIGSERIAL PRIMARY KEY,
    group_id BIGINT,
    telegram_chat_id BIGINT UNIQUE,
    device_id BIGINT UNIQUE,
    notification_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (group_id) REFERENCES public.groups(id)
);

-- +goose Down
DROP TABLE IF EXISTS public.subscribers;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.subjects;
DROP TABLE IF EXISTS public.days;
DROP TABLE IF EXISTS public.weeks;
DROP TABLE IF EXISTS public.groups;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
