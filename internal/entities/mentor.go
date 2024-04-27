package entities

type MentorBase struct {
	name  string `json:"name"`
	login string `json:"login"`
}

type MentorCreate struct {
	MentorBase
	hashed_password string `json:"hashed_password"`
}

type Mentor struct {
	MentorBase
	MentorID  int      `json:"id"`
	StudentID []int    `json:"student_id"`
	level     string   `json:"level"`
	knowledge []string `json:"knowledge"`
}
