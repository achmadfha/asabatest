CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO items (items_id, code, name, amount, description, statusActive, isDeleted, created_at, updated_at)
VALUES
    (uuid_generate_v4(), 'ABC1234567', 'Item 1', 10, 'Description for Item 1', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'DEF4567890', 'Item 2', 20, 'Description for Item 2', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'GHI7890123', 'Item 3', 30, 'Description for Item 3', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'JKL0123456', 'Item 4', 15, 'Description for Item 4', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'MNO1234567', 'Item 5', 25, 'Description for Item 5', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'PQR8901234', 'Item 6', 12, 'Description for Item 6', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'STU2345678', 'Item 7', 18, 'Description for Item 7', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'VWX3456789', 'Item 8', 22, 'Description for Item 8', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'YZA4567890', 'Item 9', 28, 'Description for Item 9', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'BCD5678901', 'Item 10', 35, 'Description for Item 10', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'EFG6789012', 'Item 11', 19, 'Description for Item 11', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'HIJ7890123', 'Item 12', 27, 'Description for Item 12', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'KLM8901234', 'Item 13', 14, 'Description for Item 13', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'NOP9012345', 'Item 14', 32, 'Description for Item 14', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'QRS0123456', 'Item 15', 21, 'Description for Item 15', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'TUV1234567', 'Item 16', 17, 'Description for Item 16', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'WXYZ234567', 'Item 17', 23, 'Description for Item 17', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'ZAB3456789', 'Item 18', 29, 'Description for Item 18', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'CDE4567890', 'Item 19', 16, 'Description for Item 19', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'FGH5678901', 'Item 20', 24, 'Description for Item 20', TRUE, FALSE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert transaction items
INSERT INTO transaction_items (transaction_id, items_code, transaction_type, quantity, created_at, updated_at)
VALUES
    (uuid_generate_v4(), 'ABC1234567', 'IN', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'DEF4567890', 'OUT', 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'GHI7890123', 'IN', 10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'JKL0123456', 'OUT', 8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'MNO1234567', 'IN', 12, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'PQR8901234', 'OUT', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'STU2345678', 'IN', 6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'VWX3456789', 'OUT', 7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'YZA4567890', 'IN', 9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'BCD5678901', 'OUT', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'EFG6789012', 'IN', 15, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'HIJ7890123', 'OUT', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'KLM8901234', 'IN', 18, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'NOP9012345', 'OUT', 10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'QRS0123456', 'IN', 7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'TUV1234567', 'OUT', 6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'WXYZ234567', 'IN', 11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'ZAB3456789', 'OUT', 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'CDE4567890', 'IN', 9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'FGH5678901', 'OUT', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
