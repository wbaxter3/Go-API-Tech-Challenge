package models

import "github.com/lib/pq"

type Person struct {
	ID        int           `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Type      string        `json:"type"`
	Age       string        `json:"age"`
	Courses   pq.Int64Array `json:"courses"`
}
