package services

import (
	"context"
	"database/sql"
	"fmt"
	"go-api-tech-challenge/internal/models"
)

type PersonService struct {
	database *sql.DB
}

// NewUserService returns a new UserService struct.
func NewPersonService(db *sql.DB) *PersonService {
	return &PersonService{
		database: db,
	}
}

func (s PersonService) ListPersons(ctx context.Context) ([]models.Person, error) {

	query := `SELECT p.id as person_id, 
	p.first_name, 
	p.last_name,
	p.type,
	p.age,
	Array_AGG(pc.course_id) as course_ids
	FROM person p
	LEFT JOIN person_course pc ON p.id = pc.person_id
	GROUP BY p.id, p.first_name, p.last_name, p.type, p.age
	ORDER BY person_id asc`
	rows, err := s.database.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return []models.Person{}, fmt.Errorf("[in services.ListPersons] failed to get persons: %w", err)
	}
	defer rows.Close()

	var persons []models.Person
	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Type, &person.Age, &person.Courses)
		if err != nil {
			return []models.Person{}, fmt.Errorf("[in services.ListPersons] failed to scan person from row: %w", err)
		}
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return []models.Person{}, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", err)
	}

	return persons, nil

}

func (s PersonService) GetPersonByName(ctx context.Context, name string) (models.Person, error) {
	var person models.Person
	query := `SELECT id as person_id, 
	first_name, 
	last_name,
	type,
	age,
	FROM person p
	WHERE p.last_name = $1`

	err := s.database.QueryRowContext(ctx, query, name).Scan(&person.ID, &person.FirstName,
		&person.LastName, &person.Type, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Person{}, fmt.Errorf("[in services.GetPersonByName] no person found with name: %s", name)
		}
		return models.Person{}, fmt.Errorf("[in services.GetPersonByName] failed to retrieve person: %w", err)
	}
	return person, nil
}
