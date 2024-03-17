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
