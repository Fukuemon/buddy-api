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


INSERT INTO service_codes (id, code, service_time_range_start, service_time_range_end, created_at, updated_at) VALUES
('01JBVE7Z0H0E0M6BX3FV1DK69A', '訪看I２', 20, 29, NOW(), NOW()),
('01JBVE7Z0J2V4RK19SA57E6Z05', '予防看I２', 20, 29, NOW(), NOW()),
('01JBVE7Z0JQ9VG0E9SBSAQPYK3', '予訪看I２', 20, 29, NOW(), NOW()),
('01JBVE7Z0JBNQTF4ZYM229SRJ2', '予防訪看I２', 20, 29, NOW(), NOW()),
('01JBVE7Z0JRTMVMMNPKTEHSYY7', '訪看I３', 30, 59, NOW(), NOW()),
('01JBVE7Z0JFW5ZBQ89FZ80JE3D', '予防看I３', 30, 59, NOW(), NOW()),
('01JBVE7Z0JFW8NYWCPTSBGTREV', '予訪看I３', 30, 59, NOW(), NOW()),
('01JBVE7Z0JGC8VKTAS8MEV7C42', '予防訪看I３', 30, 59, NOW(), NOW()),
('01JBVE7Z0JVGPNAZBD6QEES9XP', '訪看I４', 60, 89, NOW(), NOW()),
('01JBVE7Z0JW6KRZ5VEQDN9MPSN', '予防看I４', 60, 89, NOW(), NOW()),
('01JBVE7Z0JMM2H0VVKB07295KR', '予訪看I４', 60, 89, NOW(), NOW()),
('01JBVE7Z0JC1R1NCVXW5QD50VY', '予防訪看I４', 60, 89, NOW(), NOW()),
('01JBVE7Z0J7173WPCNTXQVKA2C', '訪看I５', 21, 40, NOW(), NOW()),
('01JBVE7Z0JH4VE7NW3R73PAH0M', '予防看I５', 21, 40, NOW(), NOW()),
('01JBVE7Z0JB1EA1AFH3DREZZAR', '予訪看I５', 21, 40, NOW(), NOW()),
('01JBVE7Z0J656V3S0KKQXGWRBM', '予防訪看I５', 21, 40, NOW(), NOW()),
('01JBVE7Z0JTE372AGDJ1629P30', '訪看I５・２超', 41, 60, NOW(), NOW()),
('01JBVE7Z0KTMH989KS7SWDTS4N', '訪看I５２超', 41, 60, NOW(), NOW()),
('01JBVE7Z0KFT13Y8M1VTX37E4H', '予防看I５・２超', 41, 60, NOW(), NOW()),
('01JBVE7Z0K7TMZ0NHBHDTVKAEA', '予訪看I５・２超', 41, 60, NOW(), NOW()),
('01JBVE7Z0KB3NPEWT5M09JBDHT', '予防訪看I５２超', 41, 60, NOW(), NOW()),
('01JBVE7Z0K4KJ7G9C4Y8AWSVB7', '基本療養費', 30, 90, NOW(), NOW()),
('01JBVE7Z0K9VT9ZYJRE6YFARFX', '医', 30, 90, NOW(), NOW()),
('01JBVE7Z0KCS5C4ATYW3WPDHYN', '難病等複数回訪問加算(２回)', 30, 90, NOW(), NOW());
