-- Create `user_role` type if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE user_role AS ENUM ('patient', 'medical_personnel', 'HR');
    END IF;
END $$;

-- Create `blood_group` type if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'blood_group') THEN
        CREATE TYPE blood_group AS ENUM ('A', 'B', 'AB', 'O');
    END IF;
END $$;

-- Create `sex` type if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sex') THEN
        CREATE TYPE sex AS ENUM ('male', 'female');
    END IF;
END $$;

-- Create `status` type if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        CREATE TYPE status AS ENUM ('yes', 'no');
    END IF;
END $$;

-- Create Users table
CREATE TABLE IF NOT EXISTS Users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL
);

-- Create Patient table
CREATE TABLE IF NOT EXISTS Patient (
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

-- Create Medical_history table
CREATE TABLE IF NOT EXISTS Medical_history (
    medical_history_id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    detail TEXT NOT NULL,
    time TIME NOT NULL,
    date date NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE
);

-- Create Department table
CREATE TABLE IF NOT EXISTS Department(
    department_id VARCHAR(4) PRIMARY KEY,
    department_name VARCHAR(100) UNIQUE NOT NULL
);

-- Create Position table
CREATE TABLE IF NOT EXISTS Position (
    position_id VARCHAR(4) PRIMARY KEY,
    department_id VARCHAR(4) NOT NULL,
    position_name VARCHAR(100) NOT NULL,
    FOREIGN KEY (department_id) REFERENCES Department(department_id) ON DELETE CASCADE
);

-- Create Employee table
CREATE TABLE IF NOT EXISTS Employee(
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
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE SET NULL,
    FOREIGN KEY (position_id) REFERENCES Position(position_id) ON DELETE SET NULL,
    CHECK (phone_number ~ '^[0-9]+$'),
    CHECK (resignation_date IS NULL OR resignation_date > hire_date),
    UNIQUE (first_name, last_name)
);

-- Create Patient_Appointment table
CREATE TABLE IF NOT EXISTS Patient_Appointment (
    appointment_id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    time TIME NOT NULL,
    date DATE NOT NULL,
    topic TEXT NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE
);

-- Create Disease table
CREATE TABLE IF NOT EXISTS Disease (
    disease_id VARCHAR(4) PRIMARY KEY, 
    disease_name VARCHAR(100) NOT NULL, 
    UNIQUE(disease_name)
);

-- Create Patient_chronic_disease table
CREATE TABLE IF NOT EXISTS Patient_chronic_disease (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    disease_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (disease_id) REFERENCES Disease(disease_id) ON DELETE CASCADE,
    UNIQUE (patient_id, disease_id)
);

-- Create drug table
CREATE TABLE IF NOT EXISTS drug (
    drug_id VARCHAR(4) PRIMARY KEY, 
    drug_name VARCHAR(100) NOT NULL, 
    UNIQUE(drug_name)
);

-- Create Patient_drug_allergy table
CREATE TABLE IF NOT EXISTS Patient_drug_allergy (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(4) NOT NULL,
    drug_id VARCHAR(4) NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id) ON DELETE CASCADE,
    FOREIGN KEY (drug_id) REFERENCES drug(drug_id) ON DELETE CASCADE,
    UNIQUE (patient_id, drug_id)
);

-- Create indexes
-- For patient search
CREATE INDEX IF NOT EXISTS idx_patient_id ON Patient(patient_id);
CREATE INDEX IF NOT EXISTS idx_patient_first_name ON Patient(first_name);
CREATE INDEX IF NOT EXISTS idx_patient_last_name ON Patient(last_name);

-- For employee search
CREATE INDEX IF NOT EXISTS idx_employee_id ON Employee(employee_id);
CREATE INDEX IF NOT EXISTS idx_employee_first_name ON Employee(first_name);
CREATE INDEX IF NOT EXISTS idx_employee_last_name ON Employee(last_name);

-- For Foreign Keys (use to JOIN tables)
CREATE INDEX IF NOT EXISTS idx_patient_user_id ON Patient(user_id);
CREATE INDEX IF NOT EXISTS idx_medical_history_patient_id ON Medical_history(patient_id);
CREATE INDEX IF NOT EXISTS idx_appointment_patient_id ON Patient_Appointment(patient_id);
CREATE INDEX IF NOT EXISTS idx_chronic_disease_patient_id ON Patient_chronic_disease(patient_id);
CREATE INDEX IF NOT EXISTS idx_drug_allergy_patient_id ON Patient_drug_allergy(patient_id);