package models

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Age       string `json:"age"`
	Courses   []int  `json:"courses"`
}
