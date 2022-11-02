package services

import (
	"go_crud/db"
	"go_crud/db/models"
)

func CreatePersonService(person *models.Person, database db.Database) error {
	var id int
	var createdAt string
	var name string
	var email string
	var document string
	query := `INSERT INTO persons (name, email, document) VALUES ($1, $2, $3) RETURNING id, created_at, name, email, document`
	err := database.Conn.QueryRow(query, person.Name, person.Email, person.Document).Scan(&id, &createdAt, &name, &email, &document)
	if err != nil {
		return err
	}
	person.Id = id
	person.CreatedAt = createdAt
	person.Name = name
	person.Email = email
	person.Document = document
	return nil
}
