CREATE TABLE Patients (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(20) NOT NULL,
	last_name VARCHAR(20) NOT NULL,
	age INT NOT NULL,
	disease VARCHAR(20) NOT NULL,
	medicine VARCHAR(20) NOT NULL,
	allergies VARCHAR(20) NOT NULL
);

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL
);


INSERT INTO Patients (first_name, last_name, age, disease, medicine, allergies)
VALUES	('John', 'Doe', 45, 'Diabetes', 'Metformin', 'Penicillin'),
	('Alice', 'Smith', 34, 'Hypertension', 'Lisinopril', 'None'),
	('Michael', 'Johnson', 50, 'Asthma', 'Albuterol', 'Peanuts'),
	('Sarah', 'Brown', 28, 'Flu', 'Oseltamivir', 'None'),
	('David', 'Williams', 60, 'Arthritis', 'Ibuprofen', 'Shellfish'),
	('Emma', 'Davis', 25, 'Migraine', 'Sumatriptan', 'None'),
	('James', 'Miller', 40, 'GERD', 'Omeprazole', 'Aspirin'),
	('Olivia', 'Wilson', 55, 'Anemia', 'Ferrous Sulfate', 'None'),
	('Daniel', 'Moore', 33, 'Pneumonia', 'Azithromycin', 'Sulfa drugs'),
	('Sophia', 'Taylor', 48, 'Osteoporosis', 'Alendronate', 'None');