package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 20
var MaxLenName = 30

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name, MaxLenName)
	if err != nil {
		return nil, err
	}
	d := date

	cert := &Cert{
		Course:          c,
		Name:            n,
		LabelTitle:      fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion: "Certification of Completion",
		LabelPresented:  "This Certification is Presented To",
		LabelDate:       fmt.Sprintf("Date: %v", d),
	}
	return cert, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, "course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string. got='%s', len=%d", c, len(c))
	} else if len(c) >= maxLen {
		return c, fmt.Errorf("Invalid string. got='%s', len=%d", c, len(c))
	}
	return c, nil
}

func validateName(name string, maxLen int) (string, error) {
	n, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}