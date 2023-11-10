-- Author: Christian Mahardhika
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS passengers (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    date_of_birth TIMESTAMP NOT NULL,
    id_number VARCHAR(255) NOT NULL,
    id_type VARCHAR(255) NOT NULL,
    covid_vaccine_status VARCHAR(255) NOT NULL,
    is_id_verified BOOLEAN NOT NULL,
    case_id INT NOT NULL default 0,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS plane_informations (
    id SERIAL PRIMARY KEY,
    code_plane VARCHAR(255) NOT NULL,
    total_bagage_capacity INT NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    variant VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS plane_information_details (
    id SERIAL PRIMARY KEY,
    class VARCHAR(255) NOT NULL,
    total_seat_capacity INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    plane_id BIGINT NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (plane_id) REFERENCES plane_informations(id)
);
CREATE TABLE IF NOT EXISTS airports (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    province VARCHAR(255) NOT NULL,
    iata VARCHAR(255),
    icao VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS flights (
    id SERIAL PRIMARY KEY,
    code_flight VARCHAR(255) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    departure_airport_name VARCHAR(255) NOT NULL,
    arrival_airport_name VARCHAR(255) NOT NULL,
    departure_airport_id BIGINT NOT NULL,
    arrival_airport_id BIGINT NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    plane_id BIGINT NOT NULL,
    FOREIGN KEY (departure_airport_id) REFERENCES airports(id),
    FOREIGN KEY (arrival_airport_id) REFERENCES airports(id),
    FOREIGN KEY (plane_id) REFERENCES plane_informations(id)
);
CREATE TABLE IF NOT EXISTS flight_prices (
    id SERIAL PRIMARY KEY,
    price FLOAT NOT NULL,
    class VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS flight_reservations (
    id SERIAL PRIMARY KEY,
    class VARCHAR(255) NOT NULL,
    reserved_seat INT NOT NULL,
    total_seat INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    code_booking VARCHAR(255) NOT NULL,
    booking_date TIMESTAMP NOT NULL,
    payment_expired_at TIMESTAMP NOT NULL,
    booking_status VARCHAR(255) NOT NULL,
    case_id INT NOT NULL default 0,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id BIGINT NOT NULL,
    flight_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (flight_id) REFERENCES flights(id)
);
CREATE TABLE IF NOT EXISTS booking_details (
    id SERIAL PRIMARY KEY,
    passenger_id BIGINT NOT NULL,
    seat_number VARCHAR(255) NOT NULL,
    baggage_capacity INT NOT NULL,
    class VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (passenger_id) REFERENCES passengers(id),
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    invoice_number VARCHAR(255) NOT NULL,
    total_payment FLOAT NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    code_ticket VARCHAR(255) NOT NULL,
    is_boarding_pass BOOLEAN NOT NULL,
    is_eligible_to_flight BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    booking_id BIGINT NOT NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
);
CREATE TABLE IF NOT EXISTS workflow_details (
    id SERIAL PRIMARY KEY,
    case_id BIGINT NOT NULL,
    task_name VARCHAR(255) NOT NULL,
    task_id VARCHAR(255) NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    UNIQUE (case_id, task_id)
);
-- Indexes
CREATE INDEX code_booking_idx ON bookings(code_booking);
CREATE INDEX departure_time_idx ON flights(departure_time);
CREATE INDEX arrival_time_idx ON flights(arrival_time);
CREATE INDEX city_idx ON airports(city);