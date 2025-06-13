-- 中古マンション
DROP TABLE IF EXISTS used_condos CASCADE;

CREATE TABLE used_condos (
    used_condo_id SERIAL PRIMARY KEY,
    url VARCHAR(2048) UNIQUE,
    used_condo_name VARCHAR(100),
    price_text VARCHAR(100),
    price_num INTEGER,
    layout VARCHAR(100),
    total_units_text VARCHAR(100),
    total_units_num INTEGER,
    occupied_area_text VARCHAR(100),
    occupied_area_num NUMERIC,
    other_area_text VARCHAR(100),
    other_area_num NUMERIC,
    floor VARCHAR(100),
    built_at DATE,
    address_id INTEGER,
    address VARCHAR(100),
    sales_schedule TEXT,
    event_info TEXT,
    most_price_range VARCHAR(100),
    maintenance_fee TEXT,
    repair_reserve_fund TEXT,
    intial_repair_reserve_fund TEXT,
    additional_costs TEXT,
    direction VARCHAR(100),
    energy_efficiency TEXT,
    insulation_performance TEXT,
    estimated_utility_cost TEXT,
    reform TEXT,
    structure VARCHAR(100),
    site_area VARCHAR(100),
    land_right TEXT,
    zoning TEXT,
    parking TEXT,
    contractor TEXT,
    listed_flag BOOLEAN,
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100)
);

-- 住所
DROP TABLE IF EXISTS address CASCADE;

CREATE TABLE address (
    address_id serial PRIMARY KEY,
    postal_code CHAR(7),
    prefecture VARCHAR(50),
    city VARCHAR(100),
    town VARCHAR(100),
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100)
);

-- 路線
DROP TABLE IF EXISTS train_line CASCADE;

CREATE TABLE train_line (
    train_line_id CHAR(5) PRIMARY KEY,
    train_line_name VARCHAR(100),
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100)
);

-- 駅
DROP TABLE IF EXISTS station CASCADE;

CREATE TABLE station (
    station_id CHAR(7),
    train_line_id CHAR(5),
    station_name VARCHAR(100),
    postal_code CHAR(7),
    address_id INTEGER,
    address VARCHAR(100),
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100),
    FOREIGN KEY (train_line_id) REFERENCES train_line (train_line_id),
    FOREIGN KEY (address_id) REFERENCES address (address_id)
);

-- 中古マンション-駅
DROP TABLE IF EXISTS used_condos_station CASCADE;

CREATE TABLE used_condos_stations (
    id serial PRIMARY KEY,
    used_condo_id INTEGER,
    station_id CHAR(7),
    walking_minutes INTEGER,
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100),
    FOREIGN KEY (used_condo_id) REFERENCES used_condos (used_condo_id),
    FOREIGN KEY (station_id) REFERENCES station (station_id)
);

-- 中古マンション-バス停
DROP TABLE IF EXISTS used_condos_bus_stop CASCADE;

CREATE TABLE used_condos_bus_stop (
    id serial PRIMARY KEY,
    used_condo_id INTEGER,
    bus_stop_name VARCHAR(100),
    station_name VARCHAR(100),
    train_line_name VARCHAR(100),
    bus_minutes INTEGER,
    walking_minutes INTEGER,
    delete_flag BOOLEAN,
    register_date_time TIMESTAMP,
    update_date_time TIMESTAMP,
    register_function VARCHAR(100),
    update_function VARCHAR(100),
    process_id VARCHAR(100),
    FOREIGN KEY (used_condo_id) REFERENCES used_condos (used_condo_id)
);