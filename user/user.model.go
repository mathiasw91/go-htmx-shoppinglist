package user

import (
	"log"
	"shoppinglist/database"
)

type user struct {
	id      int
	mail    string
	created string
}

func getIdByMail(mail string) (int, error) {
	row := database.DB.QueryRow("SELECT id FROM users WHERE mail = ?", mail)
	if err := row.Err(); err != nil {
		log.Print(err.Error())
		return 0, err
	}
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	return id, nil
}
