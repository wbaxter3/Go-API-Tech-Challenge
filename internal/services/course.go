package services

import (
	"context"
	"fmt"
	"go-api-tech-challenge/internal/models"

	"gorm.io/gorm"
)

type CourseService struct {
	database *gorm.DB
}

func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{
		database: db,
	}
}

func (s CourseService) ListCourses(ctx context.Context) ([]models.Course, error) {
	//rows, err := s.database.QueryContext(
	//ctx,
	//`SELECT * FROM course`,
	//)

	//if err != nil {
	//return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to get courses: %w", err)
	//}
	//defer rows.Close()

	//var courses []models.Course
	//for rows.Next() {
	//var course models.Course
	//if err = rows.Scan(&course.ID, &course.Name); err != nil {
	//return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to scan course from row: %w", err)
	//}
	//courses = append(courses, course)
	//}

	//if err = rows.Err(); err != nil {
	//return []models.Course{}, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", err)
	//}
	var courses []models.Course

	result := s.database.WithContext(ctx).Find(&courses)
	if result.Error != nil {
		return nil, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", result.Error)
	}

	return courses, nil

}

func (s CourseService) GetCourseByID(ctx context.Context, id int) (models.Course, error) {
	var course models.Course
	result := s.database.WithContext(ctx).First(&course, id)

	if result.Error != nil {
		return models.Course{}, fmt.Errorf("[in services.ListCourses] failed to scan courses: %w", result.Error)
	}

	return course, nil
}
