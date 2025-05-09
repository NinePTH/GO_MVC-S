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
 '0876543210', '6543210987654', 'Healthy','None'),
 ( 'P004', 'Michael', 'Brown', 35, '1989-02-18', 'male', 'AB', 
  'michael.brown@example.com', 'yes', '101 Maple St, Capital City', 
  '0654321098', '9876543210123', 'Asthma', 'Smoker'),
( 'P005', 'Emily', 'Davis', 28, '1996-07-05', 'female', 'A', 
  'emily.davis@example.com', 'yes', '202 Birch Ln, Riverside', 
  '0789012345', '1122334455667', 'Allergy', 'None'),
( 'P006', 'William', 'Taylor', 50, '1974-09-30', 'male', 'O', 
  'william.taylor@example.com', 'no', '303 Cedar Dr, Hillside', 
  '0923456781', '7766554433221', 'Heart Disease', 'Drunk'),
( 'P007', 'Sophia', 'Martinez', 40, '1984-03-12', 'female', 'B', 
  'sophia.martinez@example.com', 'yes', '404 Elm St, Lakeside', 
  '0845678910', '3344556677889', 'Obesity', 'Smoker'),
( 'P008', 'James', 'Wilson', 22, '2002-06-25', 'male', 'AB', 
  'james.wilson@example.com', 'no', '505 Cherry Ave, Uptown', 
  '0765432190', '9988776655443', 'Healthy', 'None'),
( 'P009', 'Olivia', 'Anderson', 31, '1993-12-08', 'female', 'O', 
  'olivia.anderson@example.com', 'yes', '606 Willow Rd, Midtown', 
  '0812345678', '5566778899001', 'Hypertension', 'Drunk'),
( 'P010', 'Daniel', 'Thomas', 29, '1995-04-20', 'male', 'B', 
  'daniel.thomas@example.com', 'no', '707 Ash Pl, Downtown', 
  '0743210987', '4433221100998', 'Healthy', 'None');

INSERT INTO Medical_history (patient_id, detail, time, date)
VALUES ('P001', 'Fever and headache and maybe kys', '10:30:00', '2025-04-08'),
('P002', 'Fever and stomachache', '10:30:00', '2025-04-08'),
('P003', 'Mild cough and sore throat', '11:00:00', '2025-04-08'),
('P004', 'Chest pain and shortness of breath', '11:15:00', '2025-04-08'),
('P001', 'Dizziness and fatigue', '11:30:00', '2025-08-08'),
('P006', 'Back pain and muscle cramps', '11:45:00', '2025-04-08');


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

INSERT INTO Employee VALUES
('E001',NULL, 'John', 'Daltin', 'P001', '0812345678', 45000.00, '.wong@example.com', '2022-01-15', NULL, 'yes'),
('E002',NULL, 'Dim', 'Smith', 'P002', '0823456789', 50000.00, 'bob.chan@example.com', '2021-06-10', '2023-08-01', 'no'),
('E003',NULL, 'Jimmy', 'Tompson', 'P003', '0834567890', 52000.00, 'cindy.liu@example.com', '2023-02-20', NULL, 'yes'),
('E004',NULL, 'Brook', 'Sudlor', 'P004', '0845678901', 48000.00, 'david.ng@example.com', '2020-12-01', '2024-03-15', 'no'),
('E005',NULL, 'Nine', 'Ok', 'P011', '0845678908', 48000.00, 'ok@example.com', '2020-12-01', '2024-03-15', 'no'),
('E006', NULL, 'Alice', 'Brown', 'P005', '0856789012', 47000.00, 'alice.brown@example.com', '2021-04-15', NULL, 'yes'),
('E007', NULL, 'Bob', 'Wilson', 'P006', '0867890123', 53000.00, 'bob.wilson@example.com', '2022-05-20', '2024-12-31', 'no'),
('E008', NULL, 'Clara', 'Davis', 'P007', '0878901234', 49000.00, 'clara.davis@example.com', '2023-03-01', NULL, 'yes'),
('E009', NULL, 'David', 'Martinez', 'P008', '0889012345', 46000.00, 'david.martinez@example.com', '2020-09-10', '2023-07-30', 'no'),
('E010', NULL, 'Eva', 'Johnson', 'P009', '0890123456', 51000.00, 'eva.johnson@example.com', '2021-11-05', NULL, 'yes'),
('E011', NULL, 'Frank', 'Taylor', 'P010', '0801234567', 55000.00, 'frank.taylor@example.com', '2022-08-22', NULL, 'yes'),
('E012', NULL, 'Grace', 'Harris', 'P001', '0813456789', 47000.00, 'grace.harris@example.com', '2023-01-10', NULL, 'yes'),
('E013', NULL, 'Henry', 'Lee', 'P002', '0824567890', 52000.00, 'henry.lee@example.com', '2022-03-15', '2024-05-01', 'no'),
('E014', NULL, 'Isabella', 'Clark', 'P003', '0835678901', 53000.00, 'isabella.clark@example.com', '2021-07-20', NULL, 'yes'),
('E015', NULL, 'Jack', 'Robinson', 'P004', '0846789012', 50000.00, 'jack.robinson@example.com', '2020-11-25', '2023-09-30', 'no'),
('E016', NULL, 'Karen', 'Lewis', 'P005', '0857890123', 49000.00, 'karen.lewis@example.com', '2022-02-28', NULL, 'yes'),
('E017', NULL, 'Liam', 'Walker', 'P006', '0868901234', 46000.00, 'liam.walker@example.com', '2021-10-05', '2024-01-20', 'no'),
('E018', NULL, 'Mia', 'Scott', 'P007', '0879012345', 51000.00, 'mia.scott@example.com', '2023-04-18', NULL, 'yes'),
('E019', NULL, 'Noah', 'Green', 'P008', '0880123456', 55000.00, 'noah.green@example.com', '2022-12-12', NULL, 'yes'),
('E020', NULL, 'Olivia', 'Hall', 'P009', '0891234567', 48000.00, 'olivia.hall@example.com', '2021-06-30', '2024-02-28', 'no');

INSERT INTO Patient_Appointment (patient_id, time, date, topic) VALUES
('P001', '11:30:00', '2025-05-15', 'Food restrict'),
('P002', '09:30:00', '2025-05-16', 'Annual check-up'),
('P003', '14:00:00', '2025-05-17', 'Follow-up consultation'),
('P004', '10:00:00', '2025-05-18', 'Cardiac monitoring'),
('P005', '15:30:00', '2025-05-19', 'Dietary consultation'),
('P006', '13:00:00', '2025-05-20', 'Physical therapy session'),
('P007', '08:30:00', '2025-05-21', 'Obesity management'),
('P008', '16:00:00', '2025-05-22', 'General check-up'),
('P009', '11:00:00', '2025-05-23', 'Hypertension follow-up');

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

INSERT INTO Patient_chronic_disease (patient_id, disease_id)
VALUES	('P001', 'I001'),
		('P001', 'I003'),
		('P002', 'I002'),
		('P003', 'I004');

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

INSERT INTO Patient_drug_allergy (patient_id, drug_id)
VALUES	('P001', 'R001'),
		('P002', 'R003'),
		('P003', 'R004');
