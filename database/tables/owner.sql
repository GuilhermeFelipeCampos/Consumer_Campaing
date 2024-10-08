CREATE TABLE IF NOT EXISTS owner (
    id UUID PRIMARY KEY NOT NULL,
    email VARCHAR(50) NOT NULL,
    status VARCHAR(10),
    created_by VARCHAR(60),
    updated_by VARCHAR(60),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);