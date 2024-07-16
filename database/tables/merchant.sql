CREATE TABLE IF NOT EXISTS merchant (
    id  UUID PRIMARY KEY NOT NULL,
    owner_id UUID,
    region_id UUID,
    slug TEXT[],
    name VARCHAR(50) NOT NULL,
    status VARCHAR(10),
    created_by VARCHAR(60),
    updated_by VARCHAR(60),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);