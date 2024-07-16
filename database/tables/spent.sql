CREATE TABLE IF NOT EXISTS spent (
    id  UUID PRIMARY KEY NOT NULL,
    campaing_id UUID,
    bucket TEXT,
    total_spent DECIMAL(4,3)NOT NULL,
    total_clicks INT NOT NULL,
    total_impressions INT NOT NULL,
    );