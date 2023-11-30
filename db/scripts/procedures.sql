DELIMITER
    //
DROP
PROCEDURE IF EXISTS GetOwning //
CREATE PROCEDURE GetOwning(id INT)
BEGIN
    SELECT
        p.id,
        p.pseudo,
        p.famille,
        p.photo
    FROM
        participant p
    INNER JOIN participant_participant pp ON
        p.id = pp.id_owned
    WHERE
        pp.id_owner = id ;
END //
DROP
PROCEDURE IF EXISTS GetOwnedBy //
CREATE PROCEDURE GetOwnedBy(id INT)
BEGIN
    SELECT
        p.id,
        p.pseudo
    FROM
        participant p
    INNER JOIN participant_participant pp ON
        p.id = pp.id_owner
    WHERE
        pp.id_owned = id ;
END //
DROP
PROCEDURE IF EXISTS GetNotOwned //
CREATE PROCEDURE GetNotOwned(idFamille INT)
BEGIN
    SELECT
        id,
        pseudo,
        famille
    FROM
        participant
    WHERE
        id NOT IN(
        SELECT
            id_owned
        FROM
            participant_participant
    ) AND famille != idFamille ;
END //
DROP
PROCEDURE IF EXISTS SetOwnership //
CREATE PROCEDURE SetOwnership(ownerId INT, ownedId INT)
BEGIN
    INSERT INTO participant_participant(id_owner, id_owned)
VALUES(ownerId, ownedId) ;
END //
DROP
PROCEDURE IF EXISTS GetUserByPseudo //
CREATE PROCEDURE GetUserByPseudo(pseudoP VARCHAR(20))
BEGIN
    SELECT
        p.id,
        p.pseudo,
        p.famille,
        p.photo
    FROM
        participant p
    WHERE
        LOWER(p.pseudo) = LOWER(pseudoP) ;
END //
DROP
PROCEDURE IF EXISTS GetImages //
CREATE PROCEDURE GetImages()
BEGIN
    SELECT
        photo
    FROM
        participant ;
END //
DELIMITER
    ;