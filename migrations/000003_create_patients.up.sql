--CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');

CREATE TABLE patients (
  id UUID PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  birth_date DATE,
  gender gender_enum,
  city VARCHAR(50),
  country VARCHAR(50),
  phone_number VARCHAR(15) NOT NULL
);

CREATE INDEX first_name_idx ON patients(last_name, first_name);

