-- Add the SUBSIDIARY_ID column to the employees table
ALTER TABLE employees
ADD COLUMN subsidiary_id INTEGER NOT NULL DEFAULT 1;

-- Add a new primary key constraint on EMPLOYEE_ID and SUBSIDIARY_ID
ALTER TABLE employees
ADD CONSTRAINT employees_pk PRIMARY KEY (employee_id, subsidiary_id);
