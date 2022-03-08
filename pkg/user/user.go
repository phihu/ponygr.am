package user

import (
	"fmt"
	"time"

	"golang.org/x/text/language"

	"github.com/google/uuid"
)

type (
	User struct {
		ID      uuid.UUID
		Handle  string
		Email   string
		Created time.Time

		Status status

		Setting Setting
	}

	Setting struct {
		Lang language.Tag
	}

	status string
)

const (
	StatusInvalid status = "invalid"
	StatusNew     status = "new"
	StatusTest    status = "test"
	StatusActive  status = "active"
)

func (s status) String() string {
	return string(s)
}

func (s *status) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*s = StatusFromString(v)
	case []byte:
		*s = StatusFromString(string(v))

	default:
		return fmt.Errorf("unexpected type %T for status", src)
	}
	return nil
}

func StatusFromString(str string) status {
	switch str {
	case StatusNew.String():
		return StatusNew

	case StatusTest.String():
		return StatusTest

	case StatusActive.String():
		return StatusActive

	default:
		return StatusInvalid
	}
}
