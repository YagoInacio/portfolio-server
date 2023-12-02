package technologies_domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Technology struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Src     string             `bson:"src" json:"src"`
	Display bool               `bson:"display" json:"display"`
}

func NewTechnology(name string, src string, display bool) (*Technology, error) {
	tech := &Technology{
		ID:      primitive.NewObjectID(),
		Name:    name,
		Src:     src,
		Display: display,
	}

	err := tech.Validate()
	if err != nil {
		return nil, err
	}

	return tech, nil
}

func (t *Technology) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}
	if t.Src == "" {
		return errors.New("src is required")
	}

	return nil
}
