package services

import (
	"go_crud/db"
	"go_crud/db/models"
)

func GetPersonService(database db.Database) (*models.PersonList, error) {
	list := &models.PersonList{}
	rows, err := database.Conn.Query("SELECT * FROM persons ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.Id, &person.Name, &person.Email, &person.CreatedAt, &person.Document)
		if err != nil {
			return list, err
		}
		list.PersonList = append(list.PersonList, person)
	}
	return list, nil
}
