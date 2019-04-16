package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Emilien", "2019-04-16")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v\n", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil\n")
	}
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='GOLANG COURSE', got=%v\n", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Emilien", "2019-04-16")
	if err == nil {
		t.Errorf("Error should be returned on an empty course\n")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklmwxcvbn"
	_, err := New(course, "Emilien", "2019-04-16")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (course=%s", course)
	}
}
