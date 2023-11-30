DROP
    DATABASE IF EXISTS santaroulette;
CREATE DATABASE IF NOT EXISTS santaroulette; USE
    santaroulette;
DROP TABLE IF EXISTS
    participant;
CREATE TABLE IF NOT EXISTS participant(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    pseudo VARCHAR(20),
    famille INT,
    photo LONGTEXT,
    CONSTRAINT LONGTEXTparticipant PRIMARY KEY(id)
); DROP TABLE IF EXISTS
    participant_participant;
CREATE TABLE participant_participant(
    id_owner INT NOT NULL UNIQUE,
    id_owned INT NOT NULL UNIQUE,
    CONSTRAINT fk_participant_owner FOREIGN KEY(id_owner) REFERENCES participant(id),
    CONSTRAINT fk_participant_owned FOREIGN KEY(id_owned) REFERENCES participant(id),
    CONSTRAINT pk_participant_participant PRIMARY KEY(id_owner, id_owned)
);