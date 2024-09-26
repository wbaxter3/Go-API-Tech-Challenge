package services

import (
	"context"
	"database/sql"
	"fmt"
	"go-api-tech-challenge/internal/models"

	"github.com/lib/pq"
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
		var dbCourseIDs pq.Int64Array // Temp var to scan in postgres array into []int
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Type, &person.Age, &dbCourseIDs)
		if err != nil {
			return []models.Person{}, fmt.Errorf("[in services.ListPersons] failed to scan person from row: %w", err)
		}
		// Convert pg array into []int
		intCourseIDs := make([]int, len(dbCourseIDs))
		for i, id := range dbCourseIDs {
			intCourseIDs[i] = int(id) // Convert int64 to int
		}
		person.Courses = intCourseIDs
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return []models.Person{}, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", err)
	}

	return persons, nil

}

func (s PersonService) GetPersonByName(ctx context.Context, name string) (models.Person, error) {
	var person models.Person
	var dbCourseIDs pq.Int64Array
	query := `SELECT p.id as person_id, 
	p.first_name, 
	p.last_name,
	p.type,
	p.age,
	Array_AGG(pc.course_id) as course_ids
	FROM person p
	LEFT JOIN person_course pc ON p.id = pc.person_id
	WHERE LOWER(p.last_name) = LOWER($1)
	GROUP BY p.id, p.first_name, p.last_name, p.type, p.age`

	err := s.database.QueryRowContext(ctx, query, name).Scan(&person.ID, &person.FirstName,
		&person.LastName, &person.Type, &person.Age, &dbCourseIDs)
	intCourseIDs := make([]int, len(dbCourseIDs))
	for i, id := range dbCourseIDs {
		intCourseIDs[i] = int(id) // Convert int64 to int
	}
	person.Courses = intCourseIDs
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Person{}, fmt.Errorf("[in services.GetPersonByName] no person found with name: %s", name)
		}
		return models.Person{}, fmt.Errorf("[in services.GetPersonByName] failed to retrieve person: %w", err)
	}
	return person, nil
}

func (s PersonService) UpdatePerson(ctx context.Context, lastName string, updatedPerson models.Person) (models.Person, error) {
	var person models.Person
	var updatedCourseIDs []int

	tx, err := s.database.Begin()
	if err != nil {
		return models.Person{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Update the person
	query := `
	UPDATE person
	SET first_name = $1,
	last_name = $2,
	type = $3,
	age = $4
	WHERE LOWER(last_name) = LOWER($5)
	RETURNING id, first_name, last_name, type, age;
	`

	// Execute the update query
	err = s.database.QueryRow(query, updatedPerson.FirstName, updatedPerson.LastName,
		updatedPerson.Type, updatedPerson.Age, lastName).Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.Type,
		&person.Age,
	)

	if err != nil {
		tx.Rollback() // Rollback transaction on failure
		return models.Person{}, fmt.Errorf("failed to update person: %w", err)
	}

	if len(updatedPerson.Courses) > 0 {

		// Delete current course IDs
		deleteCoursesQuery := `DELETE FROM person_course WHERE person_id = $1`
		_, err = tx.Exec(deleteCoursesQuery, person.ID)
		if err != nil {
			tx.Rollback()
			return models.Person{}, fmt.Errorf("failed to delete existing courses: %w", err)
		}
		// Add new course IDs
		insertCoursesQuery := `INSERT INTO person_course (person_id, course_id) VALUES ($1, $2)`
		for _, courseID := range updatedPerson.Courses {
			_, err := tx.Exec(insertCoursesQuery, person.ID, courseID)
			if err != nil {
				tx.Rollback()
				return models.Person{}, fmt.Errorf("failed to insert new courses: %w", err)
			}
		}

	}

	// Get course IDs from person_course
	selectCoursesQuery := `SELECT course_id FROM person_course WHERE person_id = $1`
	rows, err := tx.Query(selectCoursesQuery, person.ID)
	if err != nil {
		tx.Rollback()
		return models.Person{}, fmt.Errorf("failed to retrieve updated courses: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var courseID int
		if err := rows.Scan(&courseID); err != nil {
			tx.Rollback()
			return models.Person{}, fmt.Errorf("failed to scan course ID: %w", err)
		}
		updatedCourseIDs = append(updatedCourseIDs, courseID)
	}
	// Attach updated courses to the person struct
	person.Courses = updatedCourseIDs
	err = tx.Commit()
	if err != nil {
		return models.Person{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return person, nil

}
