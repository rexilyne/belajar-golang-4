package user

type Student struct {
	no           int
	name         string
	absentReason string
}

// constructor method
func NewStudent(no int, name string) User {
	return &Student{
		no:   no,
		name: name,
	}
}

// implementasi interface User
func (s *Student) Absent() {}

func (s *Student) Attend() {
	s.absentReason = "bolos bos"
}
