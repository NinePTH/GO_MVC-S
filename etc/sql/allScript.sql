CREATE TYPE user_role AS ENUM ('patient', 'medical_personnel', 'HR');
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL
);

CREATE TYPE blood_group AS ENUM ('A', 'B', 'AB', 'O');
CREATE TYPE sex AS ENUM ('male', 'female');
CREATE TYPE status AS ENUM ('yes','no');
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
    health_insurance status NOT NULL,
    address TEXT NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    id_card_number VARCHAR(13) UNIQUE NOT NULL,
    ongoing_treatment VARCHAR(50) NOT NULL,
    unhealthy_habits VARCHAR(50) NOT NULL,
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
	FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE
);

INSERT INTO Users (username, password, role)
VALUES
('john_doe', 'hashed_password_1', 'patient'),
('dr_smith', 'hashed_password_2', 'medical_personnel'),
('hr_admin', 'hashed_password_3', 'HR');

INSERT INTO Patient (
    patient_id, first_name, last_name, age, date_of_birth, gender,
    blood_type, email, health_insurance, address, phone_number,
    id_card_number, ongoing_treatment, unhealthy_habits
)
VALUES
( 'P001', 'John', 'Doe', 30, '1994-05-15', 'male', 'A', 
 'john.doe@example.com', 'yes', '123 Main St, Cityville', 
 '0123456789', '1234567890123', 'Hypertension','Drunk'),
( 'P002', 'Jane', 'Smith', 45, '1979-11-22', 'female', 'B',
 'jane.smith@example.com', 'yes', '456 Oak Ave, Townsville', 
 '0987654321', '3210987654321', 'Diabetes','Drunk'),
( 'P003', 'Mary', 'Johnson', 25, '1999-08-10', 'female', 'O',
 'mary.johnson@example.com', 'no', '789 Pine Rd, Villagetown', 
 '0876543210', '6543210987654', 'Healthy','None');

 INSERT INTO Medical_history (patient_id, detail, time, date)
VALUES ('P001', 'Fever and headache and maybe kys', '10:30:00', '2025-04-08');
VALUES ('P002', 'Fever and stomachache', '10:30:00', '2025-04-08');


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
('D010', 'Endocrinology'),
('D011', 'Human Resource');

CREATE TABLE Position (
position_id VARCHAR(4) PRIMARY KEY,
department_id VARCHAR(4),
position_name VARCHAR(100) NOT NULL,
FOREIGN KEY (department_id) REFERENCES Department(department_id)
);

INSERT INTO Position Values 
('P001','D001','Doctor'),
('P002','D002','Nurse'),
('P003','D001','Nurse'),
('P004','D003','Nurse'),
('P005','D005','Doctor'),
('P006','D007','Nurse'),
('P007','D001','Doctor'),
('P008','D003','Nurse'),
('P009','D009','Doctor'),
('P010','D010','Doctor'),
('P011', 'D011', 'HR');



CREATE TABLE Employee(
    employee_id VARCHAR(4) PRIMARY KEY,
    user_id INT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    position_id VARCHAR(4),
    phone_number VARCHAR(15) UNIQUE NOT NULL, 
    salary DECIMAL(10,2) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    hire_date DATE NOT NULL,
    resignation_date DATE,
    work_status status NOT NULL DEFAULT 'yes',
    
    -- Constraints
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL,
    FOREIGN KEY (position_id) REFERENCES Position(position_id) ON DELETE SET NULL,
    
    CHECK (phone_number ~ '^[0-9]+$'),
    CHECK (
        resignation_date IS NULL OR resignation_date > hire_date
    ),
    
    UNIQUE (first_name, last_name)
);

INSERT INTO Employee VALUES
('E001',NULL, 'John', 'Daltin', 'P001', '0812345678', 45000.00, '.wong@example.com', '2022-01-15', NULL, 'yes'),
('E002',NULL, 'Dim', 'Smith', 'P002', '0823456789', 50000.00, 'bob.chan@example.com', '2021-06-10', '2023-08-01', 'no'),
('E003',NULL, 'Jimmy', 'Tompson', 'P003', '0834567890', 52000.00, 'cindy.liu@example.com', '2023-02-20', NULL, 'yes'),
('E004',NULL, 'Brook', 'Sudlor', 'P004', '0845678901', 48000.00, 'david.ng@example.com', '2020-12-01', '2024-03-15', 'no'),
('E005',NULL, 'Nine', 'Ok', 'P011', '0845678901', 48000.00, 'ok@example.com', '2020-12-01', '2024-03-15', 'no');;


CREATE TABLE Patient_Appointment (
    appointment_id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    time TIME NOT NULL,
    date DATE NOT NULL,
    topic TEXT NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE SET NULL
    --  Check date ให้มากกว่าวันที่ปัจจุบัน
);

INSERT INTO Patient_Appointment (appointment_id, patient_id, time, date, topic) VALUES
(1, 'P001', '11:30:00', '2025-05-15', 'Food restrict'),
(2, 'P002', '09:30:00', '2025-05-16', 'Annual check-up'),
(3, 'P003', '14:00:00', '2025-05-17', 'Follow-up consultation');

CREATE TABLE Disease (
    disease_id VARCHAR(4) PRIMARY KEY, 
    disease_name VARCHAR(100) NOT NULL, 
    UNIQUE(disease_name) 
);

INSERT INTO Disease VALUES
('I001', 'streptococcus pneumoniae'),
('I002', 'tuberculosis'),
('I003', 'hepatitis B'),
('I004', 'malaria'),
('I005', 'dengue fever'),
('I006', 'measles'),
('I007', 'influenza'),
('I008', 'cholera'),
('I009', 'typhoid fever'),
('I010', 'rabies'),
('I011', 'meningitis');


CREATE TABLE Patient_chronic_disease (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    disease_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (disease_id) REFERENCES Disease(disease_id) ON DELETE CASCADE,
    UNIQUE (patient_id, disease_id)
);

INSERT INTO Patient_chronic_disease (patient_id, disease_id)
VALUES	('P001', 'I001'),
		('P001', 'I003'),
		('P002', 'I002'),
		('P003', 'I004');


CREATE TABLE drug (
    drug_id VARCHAR(4) PRIMARY KEY, 
    drug_name VARCHAR(100) NOT NULL, 
    UNIQUE(drug_name) 
);

INSERT INTO drug VALUES
('R001', 'anti bacteria'),
('R002', 'paracetamol'),
('R003', 'amoxicillin'),
('R004', 'ibuprofen'),
('R005', 'azithromycin'),
('R006', 'ciprofloxacin'),
('R007', 'metformin'),
('R008', 'omeprazole'),
('R009', 'atorvastatin'),
('R010', 'insulin'),
('R011', 'lisinopril');

CREATE TABLE Patient_drug_allergy (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    drug_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (drug_id) REFERENCES drug(drug_id) ON DELETE CASCADE,
    UNIQUE (patient_id, drug_id)
);

INSERT INTO Patient_drug_allergy (patient_id, drug_id)
VALUES	('P001', 'R001'),
		('P002', 'R003'),
		('P003', 'R004');

