INSERT INTO payment_methods (
        "name",
        is_active,
        created_at,
        updated_at,
        deleted_at
    )
VALUES ('Credit Card', true, NOW(), null, null),
    ('PayPal', true, NOW(), null, null),
    ('Bank Transfer', true, NOW(), null, null),
    ('Cash on Delivery', true, NOW(), null, null),
    ('Google Pay', true, NOW(), null, null);