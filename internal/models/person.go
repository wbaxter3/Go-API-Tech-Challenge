package models

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Age       int    `json:"age"`
	//Courses   pq.Int64Array `json:"courses"`
	Courses []int `json:"courses"`
}
