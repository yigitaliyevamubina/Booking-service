--CREATE TYPE note_enum AS ENUM ('prescription', 'diagnosis');

CREATE TABLE doctor_notes (
  id UUID PRIMARY KEY,
  appointment_id INT REFERENCES booked_appointments(id),
  doctor_id UUID NOT NULL,
  patient_id UUID REFERENCES patients(id),
  note_type note_enum, -- Added a note_type column
  note_text TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX appointment_id_idx ON doctor_notes(appointment_id);
CREATE INDEX doctor_patient_id_idx ON doctor_notes(doctor_id, patient_id);
