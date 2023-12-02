package experiences_domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Period struct {
	Start string `bson:"start"`
	End   string `bson:"end"`
}

type Experience struct {
	ID       primitive.ObjectID   `bson:"_id"`
	Position string               `bson:"position"`
	Company  string               `bson:"company"`
	Logo     string               `bson:"logo"`
	Period   Period               `bson:"period"`
	Summary  []string             `bson:"summary"`
	Techs    []primitive.ObjectID `bson:"techs"`
}

func NewExperience(
	position string,
	company string,
	logo string,
	start string,
	end string,
	summary []string,
	techs []primitive.ObjectID,
) (*Experience, error) {
	experience := &Experience{
		ID:       primitive.NewObjectID(),
		Position: position,
		Company:  company,
		Logo:     logo,
		Period: Period{
			Start: start,
			End:   end,
		},
		Summary: summary,
		Techs:   techs,
	}

	err := experience.Validate()
	if err != nil {
		return nil, err
	}

	return experience, nil
}

func (e *Experience) Validate() error {
	if e.Position == "" {
		return errors.New("position is required")
	}
	if e.Company == "" {
		return errors.New("company is required")
	}
	if e.Logo == "" {
		return errors.New("logo is required")
	}
	if e.Period.Start == "" {
		return errors.New("period start is required")
	}
	if len(e.Summary) == 0 {
		return errors.New("summary is required")
	}
	if len(e.Techs) == 0 {
		return errors.New("techs are required")
	}

	return nil
}
