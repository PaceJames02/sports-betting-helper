-- +migrate Up
CREATE TABLE bookmakers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(50),
    enabled BOOLEAN
);

INSERT INTO bookmakers (id, name, enabled)
VALUES ('betonlineag', 'BetOnline.ag', false),
       ('betmgm', 'BetMGM', true),
       ('betrivers', 'BetRivers', true),
       ('betus', 'BetUS', false),
       ('bovada', 'Bovada', false),
       ('williamhill_us', 'Caesars (William Hill)', true),
       ('draftkings', 'DraftKings', true),
       ('fanatics', 'Fanatics', true),
       ('fanduel', 'FanDuel', true),
       ('mybookieag', 'MyBookie.ag', false),
       ('ballybet', 'Bally Bet', false),
       ('betanysports', 'BetAnySports', false),
       ('betparx', 'betPARX', false),
       ('espnbet', 'theScore Bet', true),
       ('fliff', 'Fliff', false),
       ('hardrockbet', 'Hard Rock Bet', false),
       ('rebet', 'Rebet', false)
ON CONFLICT (id) DO NOTHING;

-- +migrate Down
DROP TABLE bookmakers;