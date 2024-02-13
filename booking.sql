CREATE DATABASE IF NOT EXISTS bookingdb;

USE DATABASE bookingdb; /*  \c bookingdb  */


/*calculate doctors' time*/
CREATE TABLE doctor_availability (
  id SERIAL PRIMARY KEY,
  doctor_id UUID,
  department_id UUID,
  doctor_date DATE,
  doctor_time TIME,
  status BOOLEAN
);

CREATE UNIQUE INDEX unique_day_time_idx ON doctor_availability(doctor_id, doctor_date, doctor_time); --unique doctor_id, doctor_day, doctor_time
CREATE INDEX doctor_id_idx ON doctor_availability(doctor_date, doctor_time);

/*patiens*/
CREATE TABLE patients (
  id UUID PRIMARY KEY,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  birth_date DATE,
  gender ENUM('male', 'female'),
  city VARCHAR(50),
  country VARCHAR(50),
  patient_problem TEXT,
  phone_number VARCHAR(15)
);

CREATE INDEX first_name_idx ON patients(last_name, first_name);

/*patients' booked appointment*/
CREATE TABLE booked_appointments (
  id SERIAL PRIMARY KEY,
  department_id UUID,
  doctor_id UUID,
  patient_id UUID,
  appointment_date DATE,
  appointment_time TIME,
  type ENUM('offline', 'online'),
  duration INTERVAL,
  expires_at TIMESTAMP,
  token VARCHAR(10),
  patient_status BOOLEAN,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


CREATE INDEX patient_id_idx ON booked_appointments(patient_id);
CREATE INDEX patient_status_idx ON booked_appointments(patient_status);


/*stored in the archive after appointment ended*/
CREATE TABLE archive (
  id SERIAL PRIMARY KEY,
  department_id UUID,
  doctor_id UUID,
  patient_id UUID REFERENCES patients(id),
  patient_token VARCHAR(10),
  patient_problem TEXT,
  consultation_type ENUM('online', 'offline'),
  booked_date DATE,
  booked_time TIME,
  appointment_id INT REFERENCES booked_appointments(id),
  visits_count INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


/*any files patient uploaded*/
CREATE TABLE uploaded_files (
  file_id UUID,
  patient_id UUID REFERENCES patients(id),
  request_id INT,
  file BYTEA,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

/*payment for patient's booked appointments*/
CREATE TABLE patient_payment (
  id SERIAL PRIMARY KEY,
  appointment_id INT REFERENCES booked_appointments(id),
  type ENUM('offline', 'online'),
  amount FLOAT,
  status VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

/*doctors' presscriptions*/
CREATE TABLE doctor_notes (
  id SERIAL PRIMARY KEY,
  appointment_id INT REFERENCES booked_appointments(id),
  doctor_id UUID,
  patient_id UUID REFERENCES patients(id),
  prescription TEXT,
  datetime TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX appointment_id_idx ON doctor_notes(appointment_id);
CREATE INDEX doctor_patient_id_idx ON doctor_notes(doctor_id, patient_id);