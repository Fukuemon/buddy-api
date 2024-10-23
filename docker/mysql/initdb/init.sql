CREATE TABLE IF NOT EXISTS facilities (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS departments (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS teams (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS policies (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE IF NOT EXISTS positions (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    facility_id VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS position_policies (
    position_id VARCHAR(255),
    policy_id VARCHAR(255),
    PRIMARY KEY (position_id, policy_id),
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (policy_id) REFERENCES policies(id)
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255) UNIQUE,
    position_id VARCHAR(255),
    department_id VARCHAR(255),
    team_id VARCHAR(255),
    facility_id VARCHAR(255),
    area_id VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (department_id) REFERENCES departments(id),
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TABLE IF NOT EXISTS user_policies (
    user_id VARCHAR(255),
    policy_id VARCHAR(255),
    PRIMARY KEY (user_id, policy_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (policy_id) REFERENCES policies(id)
);

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
VALUES ('01J76AN7GHXYRSAKW8014CHRKE', 'すべてのカレンダーの閲覧', NOW(), NOW());

-- position_policies
INSERT INTO position_policies (position_id, policy_id)
VALUES ('01J71685CQ0ZKYEHBRADF1Q8B4', '01J76AN7GHXYRSAKW8014CHRKE');
