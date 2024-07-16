CREATE TABLE IF NOT EXISTS ledger (
    id  UUID PRIMARY KEY NOT NULL,
    campaing_id UUID,
    spent_id UUID,
    slug_id UUID,
    user_id UUID,
    event_type VARCHAR(50) NOT NULL,
    latitude FLOAT,
    longitude FLOAT,
    cost DECIMAL(4,3)NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE
 
);