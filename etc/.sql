CREATE TYPE user_role AS ENUM ('patient', 'medical_personnel', 'HR');
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL
);

CREATE TYPE blood_group AS ENUM ('A', 'B', 'AB', 'O');
CREATE TYPE sex AS ENUM ('male', 'female');
CREATE TABLE Patient (
    patient_id VARCHAR(4) PRIMARY KEY,
    user_id INT UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    age SMALLINT NOT NULL,
    date_of_birth DATE NOT NULL,
    gender sex NOT NULL,
    blood_type blood_group NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    health_insurance BOOLEAN NOT NULL,
    address TEXT NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    id_card_number VARCHAR(13) NOT NULL,
    ongoing_treatment VARCHAR(50) NOT NULL,
     FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL,
    CHECK (phone_number ~ '^[0-9]+$'),
	UNIQUE (first_name, last_name)
);

CREATE TABLE Medical_history (
	medical_history_id SERIAL PRIMARY KEY,
	patient_id VARCHAR(4) NOT NULL,
	detail TEXT NOT NULL,
	time TIME NOT NULL,
	date date NOT NULL,
	FOREIGN KEY (patient_id) REFERENCES Patient(patient_id)
);

INSERT INTO Users (username, password, role)
VALUES
('john_doe', 'hashed_password_1', 'patient'),
('dr_smith', 'hashed_password_2', 'medical_personnel'),
('hr_admin', 'hashed_password_3', 'HR');

INSERT INTO Patient (
    patient_id, first_name, last_name, age, date_of_birth, gender,
    blood_type, email, health_insurance, address, phone_number,
    id_card_number, ongoing_treatment
)
VALUES
( 'P001', 'John', 'Doe', 30, '1994-05-15', 'male', 'A', 
 'john.doe@example.com', TRUE, '123 Main St, Cityville', 
 '0123456789', '1234567890123', 'Hypertension'),
( 'P002', 'Jane', 'Smith', 45, '1979-11-22', 'female', 'B',
 'jane.smith@example.com', FALSE, '456 Oak Ave, Townsville', 
 '0987654321', '3210987654321', 'Diabetes'),
( 'P003', 'Mary', 'Johnson', 25, '1999-08-10', 'female', 'O',
 'mary.johnson@example.com', TRUE, '789 Pine Rd, Villagetown', 
 '0876543210', '6543210987654', 'Healthy');

 INSERT INTO Medical_history (patient_id, detail, time, date)
VALUES ('P001', 'Fever and headache and maybe kys', '10:30:00', '2025-04-08');