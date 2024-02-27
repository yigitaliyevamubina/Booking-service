--CREATE TYPE consultation_type_enum AS ENUM ('online', 'offline');
CREATE TABLE archive (
  id SERIAL PRIMARY KEY,
  department_id UUID NOT NULL,
  doctor_id UUID NOT NULL,
  patient_id UUID NOT NULL,
  patient_token VARCHAR(10),
  patient_problem TEXT,
  consultation_type consultation_type_enum,
  booked_date DATE,
  booked_time TIME,
  appointment_id INT NOT NULL,
  visits_count INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX appointment_id_idx ON archive(appointment_id);
