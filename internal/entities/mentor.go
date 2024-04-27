package entities

type MentorBase struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

type CreateMentor struct {
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
	StudentIDs []int  `json:"student_ids"`
	Level      string `json:"level"`
	Tags       []int  `json:"tags"`
}

type AddNewTag struct {
	MentorID int    `json:"id"`
	Tag      string `json:"new_tag"`
}

type AddTagMt struct {
	MentorID int `json:"id"`
	TagId    int `json:"tag_id"`
}
