-- +migrate Up
CREATE TABLE sports (
    id VARCHAR(50) PRIMARY KEY,
    "group" VARCHAR(50) NOT NULL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    active BOOLEAN DEFAULT true,
    has_outrights BOOLEAN DEFAULT false
);
CREATE TABLE events (
    id VARCHAR(32) PRIMARY KEY ,
    sport VARCHAR(50) NOT NULL,
    commence_time TIMESTAMP NOT NULL,
    home_team varchar(50) NOT NULL,
    away_team varchar(50) NOT NULL,
    FOREIGN KEY (sport) REFERENCES sports(id)

);
CREATE TABLE bookmakers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(50),
    enabled BOOLEAN
);
CREATE TABLE markets (
    id VARCHAR(50) PRIMARY KEY
);

CREATE TABLE lines (
    event VARCHAR(32),
    book VARCHAR(50),
    market VARCHAR(50),
    outcome VARCHAR(50),
    price DECIMAL(10, 2),
    FOREIGN KEY (event) REFERENCES events(id),
    FOREIGN KEY (book) REFERENCES bookmakers(id),
    FOREIGN KEY (market) REFERENCES markets(id),
    PRIMARY KEY (event, book, market, outcome)
);
-- +migrate Down
DROP TABLE sports;
DROP TABLE events;
DROP TABLE bookmakers;