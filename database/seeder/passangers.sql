INSERT INTO passengers (full_name, gender, date_of_birth, id_number, id_type, covid_vaccine_status, is_id_verified, case_id, created_at, updated_at, deleted_at)
VALUES 
('John Doe', 'Male', '1990-01-01', '1234567890123456', 'National ID', 'Vaccine I', true, 0, NOW(), NOW(), NULL),
('Jane Doe', 'Female', '1995-05-05', '2345678901234567', 'Passport', 'Vaccine II', true, 0, NOW(), NOW(), NULL),
('Bob Smith', 'Male', '1985-12-25', '3456789012345678', 'National ID', 'Not Vaccinated', false, 0, NOW(), NOW(), NULL),
('Alice Johnson', 'Female', '2000-06-15', '4567890123456789', 'Driver License', 'Vaccine III', true, 0, NOW(), NOW(), NULL),
('Tom Wilson', 'Male', '1978-09-30', '5678901234567890', 'National ID', 'Not Vaccinated', false, 0, NOW(), NOW(), NULL);
