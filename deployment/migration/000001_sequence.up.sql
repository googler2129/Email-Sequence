CREATE TABLE IF NOT EXISTS sequences (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN NOT NULL,
    click_tracking_enabled BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    created_by_id VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    updated_by_id VARCHAR(255),
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );
