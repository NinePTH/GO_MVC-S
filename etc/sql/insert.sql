INSERT INTO Users (username, password, role)
VALUES
('john_doe', 'hashed_password_1', 'patient'),
('dr_smith', 'hashed_password_2', 'medical_personnel'),
('hr_admin', 'hashed_password_3', 'HR');

INSERT INTO Patient (
    patient_id, first_name, last_name, age, date_of_birth, gender,
    blood_type, email, health_insurance, address, phone_number,
    id_card_number, ongoing_treatment,unhealthy_habits
)
VALUES
( 'P001', 'John', 'Doe', 30, '1994-05-15', 'male', 'A', 
 'john.doe@example.com', TRUE, '123 Main St, Cityville', 
 '0123456789', '1234567890123', 'Hypertension','Drunk'),
( 'P002', 'Jane', 'Smith', 45, '1979-11-22', 'female', 'B',
 'jane.smith@example.com', FALSE, '456 Oak Ave, Townsville', 
 '0987654321', '3210987654321', 'Diabetes','None'),
( 'P003', 'Mary', 'Johnson', 25, '1999-08-10', 'female', 'O',
 'mary.johnson@example.com', TRUE, '789 Pine Rd, Villagetown', 
 '0876543210', '6543210987654', 'Healthy','None');

 INSERT INTO Medical_history (patient_id, detail, time, date)
VALUES ('P001', 'Fever and headache and maybe kys', '10:30:00', '2025-04-08');

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
('P010','D010','Doctor');

INSERT INTO Employee VALUES
('E001',NULL, 'John', 'Daltin', 'P001', '0812345678', 'D001', 45000.00, '.wong@example.com', '2022-01-15', NULL, 'yes'),
('E002',NULL, 'Dim', 'Smith', 'P002', '0823456789', 'D002', 50000.00, 'bob.chan@example.com', '2021-06-10', '2023-08-01', 'no'),
('E003',NULL, 'Jimmy', 'Tompson', 'P003', '0834567890', 'D003', 52000.00, 'cindy.liu@example.com', '2023-02-20', NULL, 'yes'),
('E004',NULL, 'Brook', 'Sudlor', 'P004', '0845678901', 'D004', 48000.00, 'david.ng@example.com', '2020-12-01', '2024-03-15', 'no');

INSERT INTO Patient_Appointment (appointment_id, patient_id, time, date, topic) VALUES
(1, 'P001', '11:30:00', '2025-05-15', 'Food restrict'),
(2, 'P002', '09:30:00', '2025-05-16', 'Annual check-up'),
(3, 'P003', '14:00:00', '2025-05-17', 'Follow-up consultation');

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

INSERT INTO Patient_chronic_disease VALUES
(1, 'P001', 'I001'),
(2, 'P001', 'I003'),
(3, 'P002', 'I002'),
(4, 'P003', 'I004');

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

INSERT INTO Patient_drug_allergy VALUES
(1, 'P001', 'R001'),
(2, 'P002', 'R003'),
(3, 'P003', 'R004');
