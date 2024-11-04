CREATE TABLE IF NOT EXISTS facilities (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS departments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS teams (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS policies (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS positions (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS position_policies (
    position_id VARCHAR(255),
    policy_id VARCHAR(255),
    PRIMARY KEY (position_id, policy_id),
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (policy_id) REFERENCES policies(id)
);

CREATE TABLE IF NOT EXISTS addresses (
    id VARCHAR(255) PRIMARY KEY,
    zip_code VARCHAR(255),
    prefecture VARCHAR(255),
    city VARCHAR(255),
    address_line1 VARCHAR(255),
    address_line2 VARCHAR(255),
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS areas (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    facility_id VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);


CREATE TABLE IF NOT EXISTS area_addresses (
    address_id VARCHAR(255),
    area_id VARCHAR(255),
    PRIMARY KEY (address_id, area_id),
    FOREIGN KEY (address_id) REFERENCES addresses(id),
    FOREIGN KEY (area_id) REFERENCES areas(id)
);


CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255),
    phone_number VARCHAR(255),
    position_id VARCHAR(255) NOT NULL,
    department_id VARCHAR(255) NOT NULL,
    team_id VARCHAR(255) NOT NULL,
    area_id VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (area_id) REFERENCES areas(id),
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS user_policies (
    user_id VARCHAR(255),
    policy_id VARCHAR(255),
    PRIMARY KEY (user_id, policy_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (policy_id) REFERENCES policies(id)
);

CREATE TABLE IF NOT EXISTS service_codes (
    id VARCHAR(255) PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    service_time_range_start int,
    service_time_range_end int,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    preferred_time VARCHAR(255) NOT NULL,
    preferred_gender VARCHAR(255) NOT NULL,
    service_code_id VARCHAR(255) NOT NULL,
    address_id VARCHAR(255) NOT NULL,
    area_id VARCHAR(255) NOT NULL,
    assigned_staff_id VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (service_code_id) REFERENCES service_codes(id),
    FOREIGN KEY (address_id) REFERENCES addresses(id),
    FOREIGN KEY (area_id) REFERENCES areas(id),
    FOREIGN KEY (assigned_staff_id) REFERENCES users(id),
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS routes (
    id VARCHAR(255) PRIMARY KEY,
    travel_time int,
    address_id VARCHAR(255) NOT NULL,
    destination_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (address_id) REFERENCES addresses(id),
    FOREIGN KEY (destination_id) REFERENCES addresses(id)
);

CREATE TABLE IF NOT EXISTS visit_infos (
    id VARCHAR(255) PRIMARY KEY,
    patient_id VARCHAR(255) NOT NULL,
    assigned_staff_id VARCHAR(255) NOT NULL,
    companion_id VARCHAR(255),
    route_id VARCHAR(255) NOT NULL,
    service_code_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    FOREIGN KEY (assigned_staff_id) REFERENCES users(id),
    FOREIGN KEY (companion_id) REFERENCES users(id),
    FOREIGN KEY (route_id) REFERENCES routes(id),
    FOREIGN KEY (service_code_id) REFERENCES service_codes(id)
);


-- Insert data
INSERT INTO facilities (id, name, created_at, updated_at)
VALUES ('01J6SMYDSKKKNJCR2Y3242T7YX', 'テスト訪問看護ステーション', NOW(), NOW());

-- departments（看護・リハビリ）
INSERT INTO departments (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ07BD8MRGBPTHA489', '看護', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

INSERT INTO departments (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ0JP7H3TVVP0G02TM', 'リハビリ', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

-- teams（A, B, C）
INSERT INTO teams (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQZXH934C0KSJBGTM4', 'A', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

INSERT INTO teams (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ3HA481017KTWE7B9', 'B', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

INSERT INTO teams (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQQF95SJ96ZGC2RBX3', 'C', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

-- positions（manager, reader, member）
INSERT INTO positions (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ0ZKYEHBRADF1Q8B4', 'manager', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

INSERT INTO positions (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ2JCBA6SAFTQ5MC87', 'reader', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

INSERT INTO positions (id, name, facility_id, created_at, updated_at)
VALUES ('01J71685CQ5GXXEE7EN2ECVQR6', 'member', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

-- policies（すべてのカレンダーの閲覧）
INSERT INTO policies (id, name, created_at, updated_at)
VALUES ('01J76AN7GHXYRSAKW8014CHRKE', 'すべてのカレンダーの閲覧', NOW(), NOW()),
       ('01JAYEF1X00F61QQZT51NMK4QH', 'すべてのカレンダーの編集', NOW(), NOW());

-- position_policies
INSERT INTO position_policies (position_id, policy_id)
VALUES ('01J71685CQ0ZKYEHBRADF1Q8B4', '01J76AN7GHXYRSAKW8014CHRKE'),
       ('01J71685CQ0ZKYEHBRADF1Q8B4', '01JAYEF1X00F61QQZT51NMK4QH'),
       ('01J71685CQ2JCBA6SAFTQ5MC87', '01J76AN7GHXYRSAKW8014CHRKE');

-- address
INSERT INTO addresses (id, zip_code, prefecture, city, address_line1, address_line2, latitude, longitude, created_at, updated_at)
VALUES ('01JBBCQ314307YD69N89XT9WBN', '000-0000', '東京都', '千代田区', '千代田1-1-1', '', 35.682839, 139.759455, NOW(), NOW()),
        ('01JBBCPQ5SKAMZNDE9PY940ZZN', '000-0000', '東京都', '千代田区', '千代田1-1-2', '', 35.682838, 139.759454, NOW(), NOW()),
        ('01JBBCPX6RJPK18T4B7CAJG6F3', '000-0000', '東京都', '千代田区', '千代田1-1-3', '', 35.682837, 139.759453, NOW(), NOW());

-- areas
INSERT INTO areas (id, name, facility_id, created_at, updated_at)
VALUES ('01JBBCQ867SR9Y001CR5HD6DNJ', 'A', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW()),
       ('01JBBCQP71J81VFQB2NRX77VDJ', 'B', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW()),
       ('01JBBCQP71KE6PETRQQZG2658N', 'C', '01J6SMYDSKKKNJCR2Y3242T7YX', NOW(), NOW());

-- area_address
INSERT INTO area_addresses (address_id, area_id)
VALUES ('01JBBCQ314307YD69N89XT9WBN', '01JBBCQ867SR9Y001CR5HD6DNJ'),
       ('01JBBCPQ5SKAMZNDE9PY940ZZN', '01JBBCQ867SR9Y001CR5HD6DNJ'),
       ('01JBBCPX6RJPK18T4B7CAJG6F3', '01JBBCQP71J81VFQB2NRX77VDJ');