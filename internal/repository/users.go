package repository

import (
	"mazehKav/internal/entity"

	"gorm.io/gorm"
)

type usersRepository struct {
	CommonBehaviourRepository[entity.User]
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return &usersRepository{
		CommonBehaviourRepository: NewCommonBehaviour[entity.User](db),
	}
}
