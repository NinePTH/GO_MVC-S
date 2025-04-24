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
	FOREIGN KEY (patient_id) REFERENCES Patient(patient_id)
);

CREATE TABLE Department(
department_id VARCHAR(4) PRIMARY KEY,
department_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Position (
position_id VARCHAR(4) PRIMARY KEY,
department_id VARCHAR(4),
position_name VARCHAR(100) NOT NULL,
FOREIGN KEY (department_id) REFERENCES Department(department_id)
);

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
    hire_date DATE NOT NULL,
    resignation_date DATE,
    work_status status NOT NULL DEFAULT 'yes',
    
    -- Constraints
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL,
    FOREIGN KEY (position_id) REFERENCES Position(position_id) ON DELETE SET NULL,
    FOREIGN KEY (department_id) REFERENCES Department(department_id) ON DELETE SET NULL,
    
    CHECK (phone_number ~ '^[0-9]+$'),
    CHECK (
        resignation_date IS NULL OR resignation_date > hire_date
    ),
    
    UNIQUE (first_name, last_name)
);

CREATE TABLE Patient_Appointment (
    appointment_id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    time TIME NOT NULL,
    date DATE NOT NULL,
    topic TEXT NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id)
);

CREATE TABLE Disease (
    disease_id VARCHAR(4) PRIMARY KEY, 
    disease_name VARCHAR(100) NOT NULL, 
    UNIQUE(disease_name) 
);

CREATE TABLE Patient_chronic_disease (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    disease_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
    FOREIGN KEY (disease_id) REFERENCES Disease(disease_id),
    UNIQUE (patient_id, disease_id)
);

CREATE TABLE drug (
    drug_id VARCHAR(4) PRIMARY KEY, 
    drug_name VARCHAR(100) NOT NULL, 
    UNIQUE(drug_name) 
);

CREATE TABLE Patient_drug_allergy (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    drug_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
    FOREIGN KEY (drug_id) REFERENCES drug(drug_id),
    UNIQUE (patient_id, drug_id)
);
