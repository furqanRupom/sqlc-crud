CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE opinions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    subject_type TEXT NOT NULL,     
    subject_id BIGINT NOT NULL,     
    opinion TEXT NOT NULL,         
    rating INT CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_opinions_subject ON opinions(subject_type, subject_id);
CREATE INDEX idx_opinions_user ON opinions(user_id);
