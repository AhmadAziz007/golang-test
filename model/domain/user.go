package domain

import "time"

type User struct {
	Id         int
	RoleId     int
	Name       string
	Email      string
	Password   string
	LastAccess time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
