doctor_1 d60a149f-d338-4bd6-b4dd-47bb31c7a9d9
doctor_2 13dff975-3e96-40c0-86c3-3337c55057d4
doctor_3 9ed8cfde-036f-43bf-ab9f-028729590d39
doctor_4 711df831-c9d3-4ee5-b26d-a8516d3ce232
doctor_5 e0ffaa83-9044-45c7-8762-1b76b93f1e49
doctor_6 31d2cb03-92af-4e7b-ad0d-a34aa33274bd
doctor_7 2ba2c75a-ebf7-4d11-ac4e-c3e9e9d05ece
doctor_8 28e984ae-b62e-412b-b312-c2fe1910ef85
doctor_9 ce322e85-792b-46b6-8df0-e942c06a0d01
doctor_10 44171e4d-78a5-4c26-8a31-44c982795cdf

-----------------------------------------------

Cardiology       431c9323-7e45-46b9-bfaa-e4644a45c7b2

Dermatology      9eee50c3-1cfd-4843-88a2-96500409f02f

Gastroenterology 721558b6-81e3-4f48-97c9-a207ef6d7da3

Neurology        15c755bc-5fdb-449b-9185-e2a515a14615

Orthopedics      d7091903-2a71-4db4-b049-cfb3bd0c97f1

Pediatrics       0677485e-1f5f-4bf8-ae9d-25b8c819704c

Psychiatry       422d9897-cd78-409e-a571-202aa68e79e9

Radiology        3e2f3a35-8cfb-464d-8746-93c3a75da6b5

Urology          e5d83636-8a26-49ce-9176-e9bc9cba1479

Ophthalmology    428ee5f6-c32d-484e-8ab6-2d22c4535d8d

--------------------------------------------------

INSERT INTO doctor_availability (doctor_id, department_id, availability_date, availability_time, status) VALUES
('d60a149f-d338-4bd6-b4dd-47bb31c7a9d9', '431c9323-7e45-46b9-bfaa-e4644a45c7b2', '2024-02-15', '08:00:00', true),
('13dff975-3e96-40c0-86c3-3337c55057d4', '9eee50c3-1cfd-4843-88a2-96500409f02f', '2024-02-16', '09:00:00', true),
('9ed8cfde-036f-43bf-ab9f-028729590d39', '721558b6-81e3-4f48-97c9-a207ef6d7da3', '2024-02-17', '10:00:00', false),
('711df831-c9d3-4ee5-b26d-a8516d3ce232', '15c755bc-5fdb-449b-9185-e2a515a14615', '2024-02-15', '09:00:00', true),
('e0ffaa83-9044-45c7-8762-1b76b93f1e49', 'd7091903-2a71-4db4-b049-cfb3bd0c97f1', '2024-02-16', '10:00:00', true),
('31d2cb03-92af-4e7b-ad0d-a34aa33274bd', '0677485e-1f5f-4bf8-ae9d-25b8c819704c', '2024-02-17', '11:00:00', true),
('2ba2c75a-ebf7-4d11-ac4e-c3e9e9d05ece', '422d9897-cd78-409e-a571-202aa68e79e9', '2024-02-15', '10:00:00', true),
('28e984ae-b62e-412b-b312-c2fe1910ef85', '3e2f3a35-8cfb-464d-8746-93c3a75da6b5', '2024-02-16', '11:00:00', true),
('ce322e85-792b-46b6-8df0-e942c06a0d01', 'e5d83636-8a26-49ce-9176-e9bc9cba1479', '2024-02-17', '12:00:00', true),
('44171e4d-78a5-4c26-8a31-44c982795cdf', '428ee5f6-c32d-484e-8ab6-2d22c4535d8d', '2024-02-15', '11:00:00', false);

-----------------------------------------------------------------------------------------------------------

INSERT INTO patients (id, first_name, last_name, birth_date, gender, city, phone_number) VALUES
('55b237ab-a0ad-4f36-8db0-08d92eea9b9c', 'Gulnoza', 'Ismailova', '1985-09-20', 'female', 'Samarqand', '+998951234567'),
('abcfb629-e89b-477a-a383-203da8a05a22', 'Farrukh', 'Yusupov', '1978-12-10', 'male', 'Jizax', '+998991122334'),
('1de5aec1-37c0-47e8-9a2e-672ec7ab48df', 'Sitora', 'Umarova', '1995-03-25', 'female', 'Andijon', '+998981234567'),
('499c232a-f7f9-4b2f-a09f-75fa5dd98771', 'Akmaljon', 'Ochilov', '1980-07-18', 'male', 'Toshkent', '+998971122334'),
('2276d3c2-8a56-4a79-8037-844700adb7d7', 'Dilfuza', 'Karimova', '1992-11-30', 'female', 'Fargona', '+998961234567'),
('23998682-ec34-4de3-bdac-44ed918882f0', 'Jasurbek', 'Mirzayev', '1983-08-05', 'male', 'Andijon', '+998931112233'),
('356e9eb8-4252-4a32-8b51-66709aea8034', 'Malika', 'Toshmatova', '1987-04-12', 'female', 'Toshkent', '+998901234567'),
('db8f098a-1e63-40b0-b077-5793f40b7494', 'Jahongir', 'Xudoyberganov', '1975-06-28', 'male', 'Toshkent', '+998951122334'),
('0e41e17b-e705-44cd-8706-0f053f2b5f35', 'Nargiza', 'Abdullaeva', '1998-10-08', 'female', 'Toshkent', '+998941234567');


----------------------------------------------------------------------------------------------------------------------
--3 idlar xammasi to'g'ri qo'yilgan
INSERT INTO booked_appointments (department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status, status) VALUES
('431c9323-7e45-46b9-bfaa-e4644a45c7b2', 'd60a149f-d338-4bd6-b4dd-47bb31c7a9d9', '6f6e1ded-734f-453b-bba9-61b0eeb78792', '2024-02-15', '08:00:00', 'offline', '1 hour', NULL, 'ABCDE12345', true, 'scheduled'),
('9eee50c3-1cfd-4843-88a2-96500409f02f', '13dff975-3e96-40c0-86c3-3337c55057d4', '55b237ab-a0ad-4f36-8db0-08d92eea9b9c', '2024-02-16', '11:00:00', 'online', '30 minutes', NULL, 'FGHIJ67890', true, 'scheduled'),
('721558b6-81e3-4f48-97c9-a207ef6d7da3', '9ed8cfde-036f-43bf-ab9f-028729590d39', 'abcfb629-e89b-477a-a383-203da8a05a22', '2024-02-17', '12:00:00', 'offline', '45 minutes', NULL, 'KLMNO54321', false, 'scheduled'),
('15c755bc-5fdb-449b-9185-e2a515a14615', '711df831-c9d3-4ee5-b26d-a8516d3ce232', '1de5aec1-37c0-47e8-9a2e-672ec7ab48df', '2024-02-15', '13:00:00', 'online', '1 hour 15 minutes', NULL, 'PQRST67890', true, 'scheduled'),
('d7091903-2a71-4db4-b049-cfb3bd0c97f1', 'e0ffaa83-9044-45c7-8762-1b76b93f1e49', '499c232a-f7f9-4b2f-a09f-75fa5dd98771', '2024-02-16', '14:00:00', 'offline', '2 hours', NULL, 'UVWXY12345', true, 'scheduled'),
('0677485e-1f5f-4bf8-ae9d-25b8c819704c', '31d2cb03-92af-4e7b-ad0d-a34aa33274bd', '2276d3c2-8a56-4a79-8037-844700adb7d7', '2024-02-17', '15:00:00', 'offline', '1 hour 30 minutes', NULL, 'ZABCD67890', false, 'scheduled'),
('422d9897-cd78-409e-a571-202aa68e79e9', '2ba2c75a-ebf7-4d11-ac4e-c3e9e9d05ece', '23998682-ec34-4de3-bdac-44ed918882f0', '2024-02-15', '16:00:00', 'online', '45 minutes', NULL, 'EFGHI23456', true, 'scheduled'),
('3e2f3a35-8cfb-464d-8746-93c3a75da6b5', '28e984ae-b62e-412b-b312-c2fe1910ef85', '356e9eb8-4252-4a32-8b51-66709aea8034', '2024-02-16', '17:00:00', 'offline', '1 hour', NULL, 'JKLMN98765', true, 'scheduled'),
('e5d83636-8a26-49ce-9176-e9bc9cba1479', 'ce322e85-792b-46b6-8df0-e942c06a0d01', 'db8f098a-1e63-40b0-b077-5793f40b7494', '2024-02-17', '18:00:00', 'online', '2 hours', NULL, 'OPQRS23456', false, 'scheduled'),
('428ee5f6-c32d-484e-8ab6-2d22c4535d8d', '44171e4d-78a5-4c26-8a31-44c982795cdf', '0e41e17b-e705-44cd-8706-0f053f2b5f35', '2024-02-15', '19:00:00', 'offline', '1 hour 15 minutes', NULL, 'TUVWX87654', true, 'scheduled');

----------------------------------------------------------------------------------------------------------------------

INSERT INTO archive (department_id, doctor_id, patient_id, patient_token, patient_problem, consultation_type, booked_date, booked_time, appointment_id, visits_count) VALUES
('1a6a48a4-f3f7-48cd-8e70-97b8da27a424', 'd6e8a92f-ae4e-4f5d-b15e-52131e3d9f72', '8f3a6cd5-33c4-4da2-9aef-bbf8762c9fd1', 'ABCDE12345', 'Fever and headache', 'offline', '2024-02-15', '10:00:00', 1, 1),
('6b12c1a0-17f3-4e17-af2b-9b0b8fd3d4de', 'a6eeb1e3-543a-4b02-8b2c-0e7c454b24e1', 'e43820b4-47e6-42a3-8948-b64897bbfa6b', 'FGHIJ67890', 'Stomachache', 'online', '2024-02-16', '11:00:00', 2, 1),
('df4f6512-6b7f-4d4a-92d0-5b24b2cfaf41', 'b51c3324-7d13-4d3b-aeb7-ff76c775358a', '02a90e18-54d5-4eef-9414-c785eef92a3d', 'KLMNO54321', 'Back pain', 'offline', '2024-02-17', '12:00:00', 3, 1),
('30c805d0-bbd8-4cd4-83c7-bf63ac54246d', 'c6c5b5d2-ec62-4a0d-99ef-1a3635f9a70f', 'f01996b1-1ec2-46fd-b0c0-d43741b88609', 'PQRST67890', 'Allergy', 'online', '2024-02-15', '13:00:00', 4, 1),
('7a2e0017-70cf-49fd-ba1f-79d9f56f90c5', 'd36b4908-889b-4f1f-b525-5418da0984e5', 'fcd67142-76db-44e3-9145-64e9ccf5b623', 'UVWXY12345', 'Fever', 'offline', '2024-02-16', '14:00:00', 5, 1),
('1a6a48a4-f3f7-48cd-8e70-97b8da27a424', 'd6e8a92f-ae4e-4f5d-b15e-52131e3d9f72', '531b3d62-1d8f-4a38-b18b-b9203d0ab330', 'ZABCD67890', 'Cough', 'offline', '2024-02-17', '15:00:00', 6, 1),
('6b12c1a0-17f3-4e17-af2b-9b0b8fd3d4de', 'a6eeb1e3-543a-4b02-8b2c-0e7c454b24e1', 'c645b3d4-cc9d-4e52-862d-9eb3b0c09b60', 'EFGHI23456', 'Headache', 'online', '2024-02-15', '16:00:00', 7, 1),
('df4f6512-6b7f-4d4a-92d0-5b24b2cfaf41', 'b51c3324-7d13-4d3b-aeb7-ff76c775358a', 'd0b8e541-2496-485d-ae90-8b251b4ad49c', 'JKLMN98765', 'Joint pain', 'offline', '2024-02-16', '17:00:00', 8, 1),
('30c805d0-bbd8-4cd4-83c7-bf63ac54246d', 'c6c5b5d2-ec62-4a0d-99ef-1a3635f9a70f', '71d0b29f-b425-4a49-92b6-4c228574f0c4', 'OPQRS23456', 'Fever and flu', 'online', '2024-02-17', '18:00:00', 9, 1),
('7a2e0017-70cf-49fd-ba1f-79d9f56f90c5', 'd36b4908-889b-4f1f-b525-5418da0984e5', 'ca1d6215-e7ff-4688-95e7-af0530ce13af', 'TUVWX87654', 'High blood pressure', 'offline', '2024-02-15', '19:00:00', 10, 1);


----------------------------------------------------------------------------------------------------------------------

INSERT INTO uploaded_files (file_id, patient_id, request_id, file) VALUES
('9e6fe0ff-d83e-4a46-9987-0e15bb4560a0', '8f3a6cd5-33c4-4da2-9aef-bbf8762c9fd1', 1, 'SomeBinaryData1'),
('3bb4c065-38fd-45c7-b8ee-af130e2a3b23', 'e43820b4-47e6-42a3-8948-b64897bbfa6b', 2, 'SomeBinaryData2'),
('4c739f86-0e36-45e4-8fa0-89c2143f9d2f', '02a90e18-54d5-4eef-9414-c785eef92a3d', 3, 'SomeBinaryData3'),
('cc75f5d1-1f70-4057-8f7d-7d9d3c8e3459', 'f01996b1-1ec2-46fd-b0c0-d43741b88609', 4, 'SomeBinaryData4'),
('1e8efda6-86a4-420f-9469-6f62d4b6a7e2', 'fcd67142-76db-44e3-9145-64e9ccf5b623', 5, 'SomeBinaryData5'),
('a0955c35-4b4d-4d7f-9e46-90a1d0a55d3f', '531b3d62-1d8f-4a38-b18b-b9203d0ab330', 6, 'SomeBinaryData6'),
('2cf8a8da-4646-4385-8146-fb997b6ec13c', 'c645b3d4-cc9d-4e52-862d-9eb3b0c09b60', 7, 'SomeBinaryData7'),
('b35b10c0-4453-4ab7-a5c8-511ae640f38b', 'd0b8e541-2496-485d-ae90-8b251b4ad49c', 8, 'SomeBinaryData8'),
('3f7a9d4d-7469-4823-bf7c-996dbf3fe9af', '71d0b29f-b425-4a49-92b6-4c228574f0c4', 9, 'SomeBinaryData9'),
('46af26f5-7f5a-4046-a88b-fc7dbf52e4b0', 'ca1d6215-e7ff-4688-95e7-af0530ce13af', 10, 'SomeBinaryData10');

----------------------------------------------------------------------------------------------------------------------

INSERT INTO patient_payment (appointment_id, type, amount, status) VALUES
(1, 'online', 50.00, 'paid'),
(2, 'offline', 75.00, 'pending'),
(3, 'online', 60.00, 'paid'),
(4, 'offline', 100.00, 'pending'),
(5, 'online', 80.00, 'paid'),
(6, 'offline', 70.00, 'pending'),
(7, 'online', 65.00, 'paid'),
(8, 'offline', 90.00, 'pending'),
(9, 'online', 120.00, 'paid'),
(10, 'offline', 85.00, 'pending');

----------------------------------------------------------------------------------------------------------------------

INSERT INTO doctor_notes (appointment_id, doctor_id, patient_id, note_type, note_text) VALUES
(1, 'd6e8a92f-ae4e-4f5d-b15e-52131e3d9f72', '8f3a6cd5-33c4-4da2-9aef-bbf8762c9fd1', 'prescription', 'Take medicine X twice a day for 3 days.'),
(2, 'a6eeb1e3-543a-4b02-8b2c-0e7c454b24e1', 'e43820b4-47e6-42a3-8948-b64897bbfa6b', 'diagnosis', 'Diagnosed with gastritis.'),
(3, 'b51c3324-7d13-4d3b-aeb7-ff76c775358a', '02a90e18-54d5-4eef-9414-c785eef92a3d', 'prescription', 'Apply hot compresses to the affected area.'),
(4, 'c6c5b5d2-ec62-4a0d-99ef-1a3635f9a70f', 'f01996b1-1ec2-46fd-b0c0-d43741b88609', 'diagnosis', 'Diagnosed with allergic rhinitis.'),
(5, 'd36b4908-889b-4f1f-b525-5418da0984e5', 'fcd67142-76db-44e3-9145-64e9ccf5b623', 'prescription', 'Take medicine Y once a day for 5 days.'),
(6, 'd6e8a92f-ae4e-4f5d-b15e-52131e3d9f72', '531b3d62-1d8f-4a38-b18b-b9203d0ab330', 'diagnosis', 'Diagnosed with acute bronchitis.'),
(7, 'a6eeb1e3-543a-4b02-8b2c-0e7c454b24e1', 'c645b3d4-cc9d-4e52-862d-9eb3b0c09b60', 'prescription', 'Rest and drink plenty of fluids.'),
(8, 'b51c3324-7d13-4d3b-aeb7-ff76c775358a', 'd0b8e541-2496-485d-ae90-8b251b4ad49c', 'diagnosis', 'Diagnosed with osteoarthritis.'),
(9, 'c6c5b5d2-ec62-4a0d-99ef-1a3635f9a70f', '71d0b29f-b425-4a49-92b6-4c228574f0c4', 'prescription', 'Physical therapy sessions recommended.'),
(10, 'd36b4908-889b-4f1f-b525-5418da0984e5', 'ca1d6215-e7ff-4688-95e7-af0530ce13af', 'diagnosis', 'Diagnosed with hypertension.');

----------------------------------------------------------------------------------------------------------------------