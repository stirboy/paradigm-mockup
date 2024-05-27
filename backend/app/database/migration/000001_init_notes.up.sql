CREATE TABLE notes (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    icon TEXT,
    is_archived BOOLEAN NOT NULL DEFAULT FALSE,
    parent_id UUID,
    cover_image TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);