CREATE TABLE IF NOT EXISTS campaing (
    id  UUID PRIMARY KEY NOT NULL,
    merchant_id UUID,
    status VARCHAR(10),
    latitude FLOAT,
    longitude FLOAT,
    budget DECIMAL(4,3)NOT NULL,
    created_by VARCHAR(60),
    updated_by VARCHAR(60),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);