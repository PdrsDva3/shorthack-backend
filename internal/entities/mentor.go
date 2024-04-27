package entities

type MentorBase struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

type MentorCreate struct {
	MentorBase
	Password string `json:"hashed_password"`
}

type MentorLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Mentor struct {
	MentorBase
	MentorID   int    `json:"id"`
	StudentIDs []int  `json:"studentids"`
	Level      string `json:"level"`
	Tags       []int  `json:"tags"`
}
