package mock

import (
	"time"

	"lincoln.boris/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID: 1,
	Name: "Alice",
	Email: "alice@example.com",
	Created: time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
