-- Author: Christian Mahardhika
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS passengers (
    id BIGINT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female') NOT NULL,
    date_of_birth TIMESTAMP NOT NULL,
    id_number VARCHAR(255) NOT NULL,
    id_type ENUM('passport', 'ktp', 'driver_license') NOT NULL,
    covid_vaccine_status ENUM(
        'Vaccine I',
        'Vaccine II',
        'Vaccine III',
        'Not Vaccinated'
    ) NOT NULL,
    is_id_verified BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS plane_informations (
    id BIGINT PRIMARY KEY,
    code_plane VARCHAR(255) NOT NULL,
    total_bagage_capacity INT NOT NULL,
    type ENUM('airbus', 'boeing') NOT NULL,
    variant VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS plane_information_details (
    id BIGINT PRIMARY KEY,
    class ENUM('economy', 'business', 'first') NOT NULL,
    total_seat_capacity INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    plane_id BIGINT NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (plane_id) REFERENCES plane_informations(id)
);
CREATE TABLE IF NOT EXISTS flights (
    id BIGINT PRIMARY KEY,
    code_flight VARCHAR(255) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    departure_airport_id BIGINT NOT NULL,
    arrival_airport_id BIGINT NOT NULL,
    status ENUM('on_time', 'delayed', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    plane_id BIGINT NOT NULL,
    FOREIGN KEY (departure_airport_id) REFERENCES airport(id),
    FOREIGN KEY (arrival_airport_id) REFERENCES airport(id),
    FOREIGN KEY (plane_id) REFERENCES plane_informations(id)
);
CREATE TABLE IF NOT EXISTS flight_prices (
    id BIGINT PRIMARY KEY,
    price FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS flight_reservations (
    id BIGINT PRIMARY KEY,
    class ENUM('business', 'economy') NOT NULL,
    reserved_seat INT NOT NULL,
    total_seat INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS airports (
    id BIGINT PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS bookings (
    id BIGINT PRIMARY KEY,
    code_booking VARCHAR(255) NOT NULL,
    booking_date TIMESTAMP NOT NULL,
    booking_status ENUM('pending', 'paid', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    user_id BIGINT NOT NULL,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS booking_details (
    id BIGINT PRIMARY KEY,
    passenger_id BIGINT NOT NULL,
    seat_number VARCHAR(255) NOT NULL,
    baggage_capacity INT NOT NULL,
    class ENUM('economy', 'business', 'first') NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (passenger_id) REFERENCES passengers(id),
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
CREATE TABLE IF NOT EXISTS payments (
    id BIGINT PRIMARY KEY,
    invoice_number VARCHAR(255) NOT NULL,
    total_payment FLOAT NOT NULL,
    payment_method ENUM('credit_card', 'debit_card', 'bank_transfer') NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    payment_status ENUM('pending', 'paid', 'canceled') NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
CREATE TABLE IF NOT EXISTS tickets (
    id BIGINT PRIMARY KEY,
    code_ticket VARCHAR(255) NOT NULL,
    is_boarding_pass BOOLEAN NOT NULL,
    is_eligible_to_flight BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
-- Indexes
CREATE INDEX code_booking_idx ON bookings (code_booking);
CREATE INDEX departure_time_idx ON flights (departure_time);
CREATE INDEX arrival_time_idx ON flights (arrival_time);
CREATE INDEX city_idx ON airports (city);