-- Generate 10 dummy records
INSERT INTO plane_informations (
        code_plane,
        total_bagage_capacity,
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