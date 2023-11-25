-- Generate 10 dummy records
INSERT INTO plane_informations (
        code_plane,
        total_baggage_capacity,
        "type",
        variant,
        created_at
    )
VALUES ('PL001', 100, 'Airbus', 'A320', NOW()),
    ('PL002', 100, 'Airbus', 'A320', NOW()),
    ('PL003', 100, 'Airbus', 'A320', NOW()),
    ('PL004', 100, 'Airbus', 'A320', NOW()),
    ('PL005', 100, 'Airbus', 'A320', NOW()),
    ('PL006', 100, 'Airbus', 'A320', NOW()),
    ('PL007', 100, 'Airbus', 'A320', NOW()),
    ('PL008', 100, 'Airbus', 'A320', NOW()),
    ('PL009', 100, 'Airbus', 'A320', NOW()),
    ('PL010', 100, 'Airbus', 'A320', NOW());
INSERT INTO plane_information_details (
        class,
        total_seat_capacity,
        created_at,
        plane_id
    )
VALUES ('Economy', 172, NOW(), 1),
    ('Economy', 172, NOW(), 2),
    ('Economy', 172, NOW(), 3),
    ('Economy', 172, NOW(), 4),
    ('Economy', 172, NOW(), 5),
    ('Economy', 172, NOW(), 6),
    ('Economy', 172, NOW(), 7),
    ('Economy', 172, NOW(), 8),
    ('Economy', 172, NOW(), 9),
    ('Economy', 172, NOW(), 10);