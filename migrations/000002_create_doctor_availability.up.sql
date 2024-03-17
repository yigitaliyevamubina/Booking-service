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
