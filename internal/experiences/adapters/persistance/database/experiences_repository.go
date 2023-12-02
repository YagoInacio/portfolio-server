package repositories

import (
	"context"

	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperiencesRepository struct {
	collection *mongo.Collection
}

func NewExperiencesRepository(coll *mongo.Collection) *ExperiencesRepository {
	return &ExperiencesRepository{
		collection: coll,
	}
}

func (r *ExperiencesRepository) Save(experience *domain.Experience) error {
	_, err := r.collection.InsertOne(context.TODO(), experience)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExperiencesRepository) ListAll() ([]domain.Experience, error) {
	filter := bson.D{}
	cursor, err := r.collection.Find(context.TODO(), filter)
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
	filter := bson.M{"_id": experience.ID}
	update := bson.M{"$set": experience}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
