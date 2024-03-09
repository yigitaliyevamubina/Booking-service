--CREATE TYPE type_enums AS ENUM ('online', 'offline');
CREATE TABLE patient_payment (
  id UUID PRIMARY KEY,
  appointment_id UUID REFERENCES booked_appointments(id),
  type type_enums NOT NULL,
  amount FLOAT NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
