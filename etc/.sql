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


CREATE TABLE Department(
department_id VARCHAR(4) PRIMARY KEY,
department_name VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO DEPARTMENT VALUES
('D001', 'Cardiology'),
('D002', 'Dermatology'),
('D003', 'Neurology'),
('D004', 'Pathology'),
('D005', 'Psychiatry'),
('D006', 'Emergency Medicine'),
('D007', 'Gastroenterology'),
('D008', 'Pulmonology'),
('D009', 'Nephrology'),
('D010', 'Endocrinology');

CREATE TABLE Position (
position_id VARCHAR(4) PRIMARY KEY,
department_id VARCHAR(4),
position_name VARCHAR(100) NOT NULL,
FOREIGN KEY (department_id) REFERENCES Department(department_id)
);

INSERT INTO Position Values 
('P001','D001','doctor'),
('P002','D002','Nurse'),
('P003','D001','Nurse'),
('P004','D003','Nurse'),
('P005','D005','Doctor'),
('P006','D007','Nurse'),
('P007','D001','Doctor'),
('P008','D003','Nurse'),
('P009','D009','Doctor'),
('P010','D010','Doctor')

CREATE TYPE status AS ENUM ('yes','no');
CREATE TABLE Employee(
employee_id VARCHAR(4) PRIMARY KEY,
user_id INT,
first_name VARCHAR(100) NOT NULL,
last_name VARCHAR(100) NOT NULL,
position_id VARCHAR(4),
phone_number VARCHAR(15) NOT NULL, 
department_id VARCHAR(4),
salary DECIMAL(10,2) NOT NULL,
email VARCHAR(50) UNIQUE NOT NULL,
hire_date date NOT NULL,
resignation_date date,
work_status status NOT NULL DEFAULT 'yes',
FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL,
FOREIGN KEY (position_id) REFERENCES Position(position_id) ON DELETE SET NULL,
FOREIGN KEY (department_id) REFERENCES Department(department_id) ON DELETE SET NULL,
CHECK (phone_number ~ '^[0-9]+$'),
UNIQUE (first_name, last_name)
);

INSERT INTO Employee VALUES
('E001',1, 'John', 'Daltin', 'P001', '0812345678', 'D001', 45000.00, '.wong@example.com', '2022-01-15', NULL, 'yes'),
('E002',2, 'Dim', 'Smith', 'P002', '0823456789', 'D002', 50000.00, 'bob.chan@example.com', '2021-06-10', '2023-08-01', 'no'),
('E003',3, 'Jimmy', 'Tompson', 'P003', '0834567890', 'D003', 52000.00, 'cindy.liu@example.com', '2023-02-20', NULL, 'yes'),
('E004',4, 'Brook', 'Sudlor', 'P004', '0845678901', 'D004', 48000.00, 'david.ng@example.com', '2020-12-01', '2024-03-15', 'no');
