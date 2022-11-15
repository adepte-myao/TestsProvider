package models

type Section struct {
	ID        int32
	Name      string
	CertAreas []CertArea
}

type CertArea struct {
	ID    int32
	Name  string
	Tests []Test
}

type Test struct {
	ID    int32
	Name  string
	Tasks []Task
}

type Task struct {
	ID       int32    `json:"id"`
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Options  []string `json:"options"`
}
