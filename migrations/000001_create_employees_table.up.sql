CREATE TABLE IF NOT EXISTS employees (
  employee_id SERIAL PRIMARY KEY,
  first_name VARCHAR(1000) NOT NULL,
  last_name VARCHAR(1000) NOT NULL,
  date_of_birth DATE NOT NULL,
  phone_number VARCHAR(1000) NOT NULL
);

