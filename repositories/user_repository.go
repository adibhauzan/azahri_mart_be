package repositories

import "github.com/adibhauzan/azahri_mart_be/model/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
}