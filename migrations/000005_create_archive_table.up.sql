CREATE TYPE status_archive AS ENUM('completed', 'missed');
CREATE TABLE archive (
  id UUID PRIMARY KEY,
  department_id UUID NOT NULL,
  doctor_id UUID NOT NULL,
  patient_id UUID NOT NULL,
  patient_token VARCHAR(10),
  consultation_type type_enum,
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
