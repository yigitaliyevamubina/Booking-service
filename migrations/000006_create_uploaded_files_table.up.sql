CREATE TABLE uploaded_files (
  file_id UUID NOT NULL,
  patient_id UUID REFERENCES patients(id),
  request_id INT NOT NULL,
  file BYTEA NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
