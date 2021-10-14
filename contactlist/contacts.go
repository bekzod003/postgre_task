package contactlist

import (
	"database/sql"
	"errors"
)

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type ContactList struct {
	Contacts *sql.DB
}

func (cl *ContactList) Create(c Contact) error {
	_, err := cl.Contacts.Exec(`insert into contact_list (first_name, last_name, phone, email)
	values ($1, $2, $3, $4);`, c.FirstName, c.LastName, c.Phone, c.Email)

	return err
}

func (cl *ContactList) Update(c Contact) error {
	_, err := cl.Contacts.Exec(`update contact_list set first_name = $1, last_name = $2, phone = $3, email = $4
	where id = $5;`, c.FirstName, c.LastName, c.Phone, c.Email, c.Id)

	return err
}

func (cl *ContactList) Get(id int) (Contact, error) {
	row := cl.Contacts.QueryRow(`select * from contact_list where id = $1;`, id)

	var answer, empty Contact

	row.Scan(&answer.Id, &answer.FirstName, &answer.LastName, &answer.Phone, &answer.Email)

	if answer == empty {
		return empty, errors.New("contact with such id does not exist")
	}

	return answer, nil
}

func (cl *ContactList) GetAll() ([]Contact, error) {
	rows, err := cl.Contacts.Query(`select * from contact_list;`)
	if err != nil {
		return make([]Contact, 0), err
	}
	var (
		tempContact Contact
		answer      []Contact
	)

	for rows.Next() {
		rows.Scan(&tempContact.Id, &tempContact.FirstName, &tempContact.LastName, &tempContact.Phone, &tempContact.Email)
		answer = append(answer, tempContact)
	}

	return answer, nil
}

func (cl *ContactList) Delete(id int) error {
	_, err := cl.Contacts.Exec(`delete from contact_list where id = $1;`, id)
	return err
}
