package utils

import (
	"database/sql"
	"fmt"
	"santa-roulette/models"
)

func GetNotOwned(db *sql.DB, idFamille string) []models.Participant {
	query := fmt.Sprintf("call GetNotOwned(%v)", idFamille)
	rows, err := db.Query(query)
	Check(err)

	var participants []models.Participant
	fmt.Println(rows)
	for rows.Next() {
		var id int
		var pseudo string
		var famille int
		err := rows.Scan(&id, &pseudo, &famille)
		Check(err)

		var participant models.Participant
		participant.Id = id
		participant.Pseudo = pseudo
		participant.Famille = famille

		participants = append(participants, participant)
	}

	return participants
}
