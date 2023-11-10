-- Rollback indexes
DROP INDEX IF EXISTS code_booking_idx;
DROP INDEX IF EXISTS departure_time_idx;
DROP INDEX IF EXISTS arrival_time_idx;
DROP INDEX IF EXISTS city_idx;
-- Drop tables
DROP TABLE IF EXISTS workflow_details;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS booking_details;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS flight_reservations;
DROP TABLE IF EXISTS flight_prices;
DROP TABLE IF EXISTS flights;
DROP TABLE IF EXISTS plane_information_details;
DROP TABLE IF EXISTS plane_informations;
DROP TABLE IF EXISTS airports;
DROP TABLE IF EXISTS passengers;
DROP TABLE IF EXISTS users;