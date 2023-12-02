package repositories

import (
	"context"

	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseExperience struct {
	ID       primitive.ObjectID   `bson:"_id"`
	Position string               `bson:"position"`
	Company  string               `bson:"company"`
	Logo     string               `bson:"logo"`
	Period   domain.Period        `bson:"period"`
	Summary  []string             `bson:"summary"`
	Techs    []primitive.ObjectID `bson:"techs"`
}

type ExperiencesRepository struct {
	collection *mongo.Collection
}

func NewExperiencesRepository(coll *mongo.Collection) *ExperiencesRepository {
	return &ExperiencesRepository{
		collection: coll,
	}
}

func (r *ExperiencesRepository) Save(experience *domain.Experience) error {
	var techs []primitive.ObjectID

	for _, tech := range experience.Techs {
		techs = append(techs, tech.ID)
	}

	experienceToInsert := r.convertExperience(experience, techs)

	_, err := r.collection.InsertOne(context.TODO(), experienceToInsert)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) ListAll() ([]domain.Experience, error) {
	pipeline := []bson.M{
		{
			"$unwind": "$techs",
		},
		{
			"$lookup": bson.M{
				"from":         "technologies",
				"localField":   "techs",
				"foreignField": "_id",
				"as":           "technology",
			},
		},
		{
			"$unwind": "$technology",
		},
	}
	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}

	var results []domain.Experience
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *ExperiencesRepository) FindById(id string) (domain.Experience, error) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Experience{}, err
	}
	filter := bson.D{
		primitive.E{Key: "_id", Value: idObject},
	}

	var result domain.Experience
	err = r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *ExperiencesRepository) Update(experience *domain.Experience) error {
	var techs []primitive.ObjectID

	for _, tech := range experience.Techs {
		techs = append(techs, tech.ID)
	}

	experienceToUpdate := r.convertExperience(experience, techs)

	filter := bson.M{"_id": experience.ID}
	update := bson.M{"$set": experienceToUpdate}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) convertExperience(xp *domain.Experience, techs []primitive.ObjectID) *DatabaseExperience {
	return &DatabaseExperience{
		ID:       xp.ID,
		Position: xp.Position,
		Company:  xp.Company,
		Logo:     xp.Logo,
		Period:   xp.Period,
		Summary:  xp.Summary,
		Techs:    techs,
	}
}
