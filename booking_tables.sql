-- Doktorning mavjudligi :
-- Ushbu jadval shifokorlarning mavjudligini ularning guvohnomasi,
-- bo'limi, sanasi va vaqtiga qarab kuzatib boradi. 
-- Bu bemorlarga mavjud bo'shliqlar asosida uchrashuvlarni 
-- rejalashtirish imkonini beradi.
CREATE TABLE doctor_availability (
  id UUID PRIMARY KEY,
  doctor_id UUID NOT NULL,
  department_id UUID NOT NULL,
  availability_date DATE NOT NULL,
  availability_time TIME NOT NULL,
  status BOOLEAN NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


CREATE UNIQUE INDEX unique_availability_idx 
ON doctor_availability(doctor_id, availability_date, availability_time);

CREATE INDEX availability_date_time_idx 
ON doctor_availability(availability_date, availability_time);

-----------------------------------------------------------

-- Bemorlar :
-- Ushbu jadvalda bemorlar to'g'risidagi ma'lumotlar, 
-- jumladan, ularning ismlari, tug'ilgan sanalari, jinsi, 
-- aloqa ma'lumotlari va joylashuvi saqlanadi.
CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');

CREATE TABLE patients (
  id UUID PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  birth_date DATE,
  gender gender_enum,
  city VARCHAR(50),
  country VARCHAR(50),
  phone_number VARCHAR(15) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


CREATE INDEX first_name_idx ON patients(last_name, first_name);


-------------------------------------------------------

-- Belgilangan uchrashuvlar :
-- Ushbu jadval bemorlar, shifokorlar va bo'limlarni ma'lum 
-- sana va vaqtlar bilan bog'lab, bron qilingan uchrashuvlarni boshqaradi. 
-- Shuningdek, u uchrashuv turi, davomiyligi, holati va 
-- tegishli vaqt belgilarini kuzatib boradi.
-- Booked Appointments
CREATE TYPE type_enum AS ENUM ('offline', 'online');
CREATE TYPE status_enum AS ENUM('scheduled', 'completed', 'missed');
CREATE TABLE booked_appointments (
  id UUID PRIMARY KEY,
  department_id UUID NOT NULL,
  doctor_id UUID NOT NULL,
  patient_id UUID,
  appointment_date DATE NOT NULL,
  appointment_time TIME NOT NULL,
  type type_enum NOT NULL,
  duration INTERVAL,
  expires_at DATE,
  token VARCHAR(10) NOT NULL,
  patient_status BOOLEAN NOT NULL,
  status status_enum,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

ALTER TABLE booked_appointments
ADD CONSTRAINT unique_appointment_datetime UNIQUE (appointment_date, appointment_time);

CREATE INDEX patient_id_idx ON booked_appointments(patient_id);
CREATE INDEX patient_status_idx ON booked_appointments(patient_status);

--------------------------------------------------------------

-- Arxiv :
-- Uchrashuv tugallangandan so'ng, 
-- uning tafsilotlari qayd qilish uchun arxivga o'tkaziladi. 
-- Ushbu jadvalda o'tgan uchrashuvlar, jumladan, bemor va shifokor 
-- identifikatorlari, uchrashuv tafsilotlari va 
-- maslahat turi haqidagi ma'lumotlar saqlanadi.
-- Archive Table
CREATE TYPE consultation_type_enum AS ENUM ('online', 'offline');
CREATE TYPE status_archive AS ENUM('completed', 'missed');
CREATE TABLE archive (
  id UUID PRIMARY KEY,
  department_id UUID NOT NULL,
  doctor_id UUID NOT NULL,
  patient_id UUID NOT NULL,
  patient_token VARCHAR(10),
  patient_problem TEXT,
  consultation_type consultation_type_enum,
  booked_date DATE,
  booked_time TIME,
  appointment_id UUID NOT NULL,
  status status_archive,
  visits_count INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX appointment_id_idx ON archive(appointment_id);

-----------------------------------------------------------

-- Yuklangan fayllar :
-- Bemorlar o'zlarining tibbiy holatiga oid fayllarni yuklashlari mumkin. 
-- Ushbu jadval ushbu fayllarni bemor identifikatori, so'rov identifikatori va
-- vaqt belgilari kabi tegishli metama'lumotlar bilan birga saqlaydi.
CREATE TABLE uploaded_files (
  file_id UUID NOT NULL,
  patient_id UUID REFERENCES patients(id),
  request_id UUID NOT NULL,
  file BYTEA NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


-------------------------------------------------------------
-- Bemor to'lovi :
-- Ushbu jadval uchrashuvlar uchun to'lov ma'lumotlarini boshqaradi,
-- to'lovlarni muayyan uchrashuvlarga bog'laydi va to'lov turi, 
-- miqdori va holati kabi ma'lumotlarni saqlaydi.
CREATE TABLE patient_payment (
  id UUID PRIMARY KEY,
  appointment_id UUID REFERENCES booked_appointments(id), 
  patient_id UUID REFERENCES patients(id),  
  type type_enums NOT NULL,
  amount FLOAT NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  paid BOOLEAN NOT NULL DEFAULT FALSE
);


------------------------------------------------------------
-- Doktor eslatmalari :
-- Shifokorlar har bir uchrashuv uchun eslatmalar va retseptlarni yozib olishlari mumkin.
-- Ushbu jadval ushbu eslatmalarni saqlaydi, ularni tegishli uchrashuv va bemor bilan bog'laydi.
-- note_typesiz jadval ichidagi retseptlar va tashxislarni farqlashingiz mumkin doctor_notes. Bu shifokor va
-- bemor o'rtasidagi o'zaro munosabatlar uchun yagona jadvalni saqlab, har xil turdagi 
-- yozuvlarni boshqarishda moslashuvchanlikni ta'minlaydi.
CREATE TYPE note_enum AS ENUM ('prescription', 'diagnosis');
CREATE TABLE doctor_notes (
  id UUID PRIMARY KEY,
  appointment_id UUID REFERENCES booked_appointments(id),
  doctor_id UUID NOT NULL,
  patient_id UUID REFERENCES patients(id),
  note_type note_enum,
  note_text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX appointment_id_idx ON doctor_notes(appointment_id);
CREATE INDEX doctor_patient_id_idx ON doctor_notes(doctor_id, patient_id);


------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION archive_completed_or_missed_appointments()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status = 'completed' OR OLD.status = 'missed' THEN
        INSERT INTO archive (department_id, doctor_id, patient_id, patient_token, patient_problem, consultation_type, booked_date, booked_time, appointment_id, visits_count)
        VALUES (OLD.department_id, OLD.doctor_id, OLD.patient_id, OLD.token, OLD.patient_problem, OLD.type, OLD.appointment_date, OLD.appointment_time, OLD.id, 1);
        
        DELETE FROM booked_appointments WHERE id = OLD.id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_archive_completed_or_missed
AFTER UPDATE ON booked_appointments
FOR EACH ROW
EXECUTE FUNCTION archive_completed_or_missed_appointments();
