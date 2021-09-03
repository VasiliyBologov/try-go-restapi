package store

import (
	"github.com/VasiliyBologov/try-go-restapi/internal/app/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	store      *Store
	collection *mongo.Collection
}

func (c *UserCollection) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (c *UserCollection) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
