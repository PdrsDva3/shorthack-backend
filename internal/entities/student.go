package entities

type StudentBase struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

type CreateStudent struct {
	StudentBase
	Password string `json:"password"`
}

type StudentLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Student struct {
	StudentBase
	MentorIds []int `json:"mentorids"`
	StudentId int   `json:"id"`
}
