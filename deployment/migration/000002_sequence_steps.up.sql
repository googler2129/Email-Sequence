CREATE TABLE IF NOT EXISTS sequence_steps (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    subject TEXT NOT NULL,
    content TEXT NOT NULL,
    waiting_days BIGINT NOT NULL,
    sequence_id BIGINT NOT NULL,
    serial_order BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    created_by_id VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255),
    updated_by_id VARCHAR(255),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
