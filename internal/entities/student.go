package entities

type StudentBase struct {
	login string `json:"login"`
	name  string `json:"name"`
	level string `json:"level"`
}

type CreateStudent struct {
	StudentBase
	hashed_password string `json:"hashed_password"`
}

type Student struct {
	StudentBase
	mentors   []int `json:"mentors"`
	StudentId int   `json:"id"`
}
