package models

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Document  string `json:"document"`
}

type PersonList struct {
	PersonList []Person `json:"persons"`
}
