package repository

import (
	"database/sql"
	"fmt"
	"people-api/internal/domain"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(person *domain.Person) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO people (name, surname, patronymic, age, gender, nationality)
         VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality,
	).Scan(&id)
	return id, err
}

func (r *PersonRepository) List(name, surname, nationality string, page, limit int) ([]domain.Person, error) {
	var people []domain.Person
	query := `SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE 1=1`
	args := []interface{}{}
	i := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", i)
		args = append(args, "%"+name+"%")
		i++
	}
	if surname != "" {
		query += fmt.Sprintf(" AND surname ILIKE $%d", i)
		args = append(args, "%"+surname+"%")
		i++
	}
	if nationality != "" {
		query += fmt.Sprintf(" AND nationality = $%d", i)
		args = append(args, nationality)
		i++
	}

	offset := (page - 1) * limit
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality); err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func (r *PersonRepository) GetByID(id string) (*domain.Person, error) {
	var p domain.Person
	err := r.db.QueryRow(`SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PersonRepository) Update(id string, person *domain.Person) error {
	_, err := r.db.Exec(
		`UPDATE people SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6, updated_at=now() WHERE id=$7`,
		person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, id,
	)
	return err
}

func (r *PersonRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM people WHERE id=$1`, id)
	return err
}
