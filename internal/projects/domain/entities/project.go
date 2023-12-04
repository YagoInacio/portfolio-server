package projects_domain

import (
	"errors"

	technologies_domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID               `bson:"_id" json:"id"`
	Title       string                           `bson:"title" json:"title"`
	Description []string                         `bson:"description" json:"description"`
	Techs       []technologies_domain.Technology `bson:"techs"`
}

func NewProject(
	title string,
	description []string,
	techs []technologies_domain.Technology,
) (*Project, error) {
	project := &Project{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
		Techs:       techs,
	}

	err := project.Validate()
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *Project) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if len(p.Description) == 0 {
		return errors.New("description is required")
	}
	if len(p.Techs) == 0 {
		return errors.New("techs are required")
	}

	return nil
}
