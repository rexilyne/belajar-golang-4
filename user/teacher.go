package user

import "time"

type Teacher struct {
	name      string
	num       int
	timestamp time.Time
}

// constructor method
func NewTeacher(name string, num int) User {
	return &Teacher{
		name: name,
		num:  num,
	}
}

// implementasi interface User
func (t *Teacher) Absent() {}

func (t *Teacher) Attend() {
	t.timestamp = time.Now()
}
