INSERT INTO flights (
        code_flight,
        departure_time,
        arrival_time,
        departure_airport_name,
        arrival_airport_name,
        departure_airport_id,
        arrival_airport_id,
        status,
        plane_id
    )
VALUES (
        'FL001',
        '2022-01-01 08:00:00',
        '2022-01-01 10:00:00',
        'Soekarno-Hatta International Airport',
        'Ngurah Rai International Airport',
        1,
        2,
        'On Time',
        1
    ),
    (
        'FL002',
        '2022-01-02 08:00:00',
        '2022-01-02 10:00:00',
        'Ngurah Rai International Airport',
        'Soekarno-Hatta International Airport',
        2,
        1,
        'On Time',
        2
    ),
    (
        'FL003',
        '2022-01-03 08:00:00',
        '2022-01-03 10:00:00',
        'Juanda International Airport',
        'Adisutjipto International Airport',
        3,
        4,
        'On Time',
        3
    ),
    (
        'FL004',
        '2022-01-04 08:00:00',
        '2022-01-04 10:00:00',
        'Adisutjipto International Airport',
        'Juanda International Airport',
        4,
        3,
        'On Time',
        4
    ),
    (
        'FL005',
        '2022-01-05 08:00:00',
        '2022-01-05 10:00:00',
        'Sultan Hasanuddin International Airport',
        'Kualanamu International Airport',
        5,
        6,
        'On Time',
        5
    ),
    (
        'FL006',
        '2022-01-06 08:00:00',
        '2022-01-06 10:00:00',
        'Kualanamu International Airport',
        'Sultan Hasanuddin International Airport',
        6,
        5,
        'On Time',
        6
    ),
    (
        'FL007',
        '2022-01-07 08:00:00',
        '2022-01-07 10:00:00',
        'Husein Sastranegara International Airport',
        'Sultan Mahmud Badaruddin II International Airport',
        7,
        8,
        'On Time',
        7
    ),
    (
        'FL008',
        '2022-01-08 08:00:00',
        '2022-01-08 10:00:00',
        'Sultan Mahmud Badaruddin II International Airport',
        'Husein Sastranegara International Airport',
        8,
        7,
        'On Time',
        8
    ),
    (
        'FL009',
        '2022-01-09 08:00:00',
        '2022-01-09 10:00:00',
        'Sam Ratulangi International Airport',
        'Sultan Aji Muhammad Sulaiman Airport',
        9,
        10,
        'On Time',
        9
    ),
    (
        'FL010',
        '2022-01-10 08:00:00',
        '2022-01-10 10:00:00',
        'Sultan Aji Muhammad Sulaiman Airport',
        'Sam Ratulangi International Airport',
        10,
        9,
        'On Time',
        10
    );
INSERT INTO flight_prices (price, class, flight_id)
VALUES (100, 'Economy', 1),
    (150, 'Economy', 2),
    (200, 'Economy', 3),
    (250, 'Economy', 4),
    (300, 'Economy', 5),
    (350, 'Economy', 6),
    (400, 'Economy', 7),
    (450, 'Economy', 8),
    (500, 'Economy', 9),
    (550, 'Economy', 10);
INSERT INTO flight_reservations (class, remining_seat, total_seat, flight_id)
VALUES ('Economy' , 172, 172, 1),
    ('Economy', 172, 172, 2),
    ('Economy', 172, 172, 3),
    ('Economy', 172, 172, 4),
    ('Economy', 172, 172, 5),
    ('Economy', 172, 172, 6),
    ('Economy', 172, 172, 7),
    ('Economy', 172, 172, 8),
    ('Economy', 172, 172, 9),
    ('Economy', 172, 172, 10);
