CREATE TYPE user_office_level AS ENUM ('operador', 'lavador', 'vendedor', 'dono', 'gerente');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(60) NOT NULL,
    office_level user_office_level NOT NULL
);

CREATE INDEX idx_first_name ON users(first_name);
CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_last_name ON users(last_name); -- exclude, if performance becomes an issue

CREATE TABLE user_passwords (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    password_hash VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL
);
-- ---------------------------------------------------

CREATE TABLE parking_lots (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    addr_street VARCHAR(80) NOT NULL,
    addr_number SMALLINT NOT NULL,
    cep VARCHAR(9) NOT NULL,
    owner_id UUID REFERENCES users(id) NOT NULL
);

CREATE INDEX idx_owner_id ON parking_lots(owner_id);
-- ---------------------------------------------------

-- Stores relations of all parking lots to users that are not of type 'dono'
CREATE TABLE parking_lot_employees (
    user_id UUID REFERENCES users(id) NOT NULL UNIQUE,
    parking_lot_id UUID REFERENCES parking_lots(id) NOT NULL

);

CREATE INDEX idx_user_id ON parking_lot_employees(user_id);
CREATE INDEX idx_parkinglot_id ON parking_lot_employees(parking_lot_id);

-- ---------------------------------------------------
CREATE TYPE car_model AS ENUM ('sedã', 'hatch', 'suv', 'cupê', 'conversível', 'picape', 'minivan', 'caminhão');
CREATE TYPE car_color AS ENUM ('branco', 'preto', 'prata', 'azul', 'vermelho', 'verde', 'cinza', 'marrom', 'amarelo');
CREATE TYPE car_brand AS ENUM ('Toyota', 'Honda', 'Ford', 'Chevrolet', 'Nissan', 'Volkswagen', 'BMW', 'Mercedes-Benz', 'Audi', 'Hyundai', 'Kia', 'Fiat', 'Renault', 'Jeep', 'Mitsubishi', 'Subaru', 'Mazda', 'Volvo', 'Jaguar');


CREATE TABLE cars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    model car_model NOT NULL,
    color car_color,
    brand car_brand,
    license_plate VARCHAR(7) NOT NULL UNIQUE
);

CREATE INDEX idx_license_plate ON cars(license_plate);

-- ---------------------------------------------------

CREATE TYPE pay_methods AS ENUM ('cartão de crédito', 'cartão de débito', 'dinheiro em espécie', 'pix');

CREATE TABLE customer_payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_method pay_methods NOT NULL,
    paid_amount INTEGER NOT NULL
);

CREATE TABLE entries_fiscal (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    rps_series SMALLINT NOT NULL,
    rps_number SMALLINT NOT NULL,
    nfe_number SMALLINT NOT NULL,
    cust_payment_id UUID REFERENCES customer_payments(id) NOT NULL UNIQUE
);

CREATE INDEX idx_nfe ON entries_fiscal(nfe_number);

CREATE TYPE op_type AS ENUM ('entry_in', 'entry_out', 'service', 'store');

CREATE TABLE entries_log (
    id BIGSERIAL PRIMARY KEY,
    "current_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    operation_type op_type NOT NULL,
    fiscal_id UUID REFERENCES entries_fiscal(id) UNIQUE, -- optional row, if it's that type of operation
    parking_lot_id UUID REFERENCES parking_lots(id) UNIQUE NOT NULL,
    operator_id UUID REFERENCES users(id) NOT NULL,
    car_license_plate VARCHAR(7) REFERENCES cars(license_plate), -- optional row, if op_type's not store
    cust_payment_id UUID REFERENCES customer_payments(id) UNIQUE,
    time_period INTERVAL,

    CONSTRAINT check_car_license_plate
        CHECK (
            (operation_type IN ('entry_in', 'entry_out', 'service') AND car_license_plate IS NOT NULL)
            OR (operation_type = 'store')
        )
);

CREATE INDEX idx_license_plate_log_non_null ON entries_log(car_license_plate)
    WHERE car_license_plate IS NOT NULL;




