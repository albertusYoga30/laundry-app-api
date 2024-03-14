-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS roles(
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(255) UNIQUE NOT NULL,
    role_desc TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    address VARCHAR(255),
    phone_Number VARCHAR(15),
    role_id INT REFERENCES roles(role_id),
    token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laundry_services(
    service_id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) UNIQUE NOT NULL,
    service_desc TEXT,
    service_price INT,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS durations(
    duration_id SERIAL PRIMARY KEY,
    duration_name VARCHAR(255) UNIQUE NOT NULL,
    duration_days INT,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    quantity INT NOT NULL,
    total_price INT NOT NULL,
    service_id INT REFERENCES laundry_services(service_id),
    duration_id INT REFERENCES durations(duration_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd