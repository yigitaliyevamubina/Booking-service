CREATE TYPE type_enum AS ENUM ('offline', 'online');

CREATE TABLE patient_payment (
  id UUID PRIMARY KEY,
  appointment_id UUID REFERENCES booked_appointments(id), 
  patient_id UUID REFERENCES patients(id),  
  type type_enum NOT NULL,
  amount FLOAT NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  paid BOOLEAN NOT NULL DEFAULT FALSE
);
