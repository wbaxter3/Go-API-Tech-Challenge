package services

import (
	"context"
	"database/sql"
	"fmt"
	"go-api-tech-challenge/internal/models"
)

type CourseService struct {
	database *sql.DB
}

func NewCourseService(db *sql.DB) *CourseService {
	return &CourseService{
		database: db,
	}
}

func (s CourseService) ListCourses(ctx context.Context) ([]models.Course, error) {
	rows, err := s.database.QueryContext(
		ctx,
		`SELECT * FROM "courses"`,
	)

	if err != nil {
		return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to get courses: %w", err)
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		if err = rows.Scan(&course.ID, &course.Name); err != nil {
			return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to scan course from row: %w", err)
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", err)
	}
	return courses, nil

}
