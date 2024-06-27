CREATE TABLE IF NOT EXISTS region (
    id  UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    status VARCHAR(10),
    latitude FLOAT,
    longitude FLOAT,
    cost DECIMAL(4,3)NOT NULL,
    created_by VARCHAR(60),
    updated_by VARCHAR(60),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);