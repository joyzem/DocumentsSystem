CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    account VARCHAR(20) NOT NULL,
    bank_name TEXT NOT NULL,
    bank_identity_number VARCHAR(9) NOT NULL
);

CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    account_id INT NOT NULL REFERENCES accounts(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    chief TEXT NOT NULL,
    financial_chief TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS employees (
	id SERIAL PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	middle_name TEXT,
	post TEXT NOT NULL,
	passport_series VARCHAR(4) NOT NULL,
	passport_number VARCHAR(6) NOT NULL,
	passport_issued_by TEXT NOT NULL,
	passport_date_of_issue DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS units (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL
);

INSERT INTO units (
    id, name
) VALUES (
    0, 'Не указано'
)
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    unit_id INT NOT NULL DEFAULT 0 REFERENCES units(id) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS proxies (
    id SERIAL NOT NULL PRIMARY KEY,
    organization_id INT NOT NULL REFERENCES organizations(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    customer_id INT NOT NULL REFERENCES customers(id) ON DELETE RESTRICT ON UPDATE CASCADE,
	employee_id INT NOT NULL REFERENCES employees(id) ON DELETE RESTRICT ON UPDATE CASCADE,
	date_of_issue DATE NOT NULL,
	is_valid_until DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS proxy_bodies	(
    id SERIAL NOT NULL PRIMARY KEY,
	product_id INT NOT NULL REFERENCES products(id) ON DELETE RESTRICT ON UPDATE CASCADE,
	proxy_id INT NOT NULL REFERENCES proxies(id) ON DELETE CASCADE ON UPDATE CASCADE,
	product_amount INT NOT NULL
);