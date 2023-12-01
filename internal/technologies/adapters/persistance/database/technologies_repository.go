package repositories

import (
	"context"

	domain "github.com/yagoinacio/portfolio-server/internal/technologies/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TechnologiesRepository struct {
	collection *mongo.Collection
}

func NewTechnologiesRepository(coll *mongo.Collection) *TechnologiesRepository {
	return &TechnologiesRepository{
		collection: coll,
	}
}

func (r *TechnologiesRepository) Save(technology *domain.Technology) error {
	_, err := r.collection.InsertOne(context.TODO(), technology)
	if err != nil {
		return err
	}

	return nil
}

func (r *TechnologiesRepository) FindDisplayable() ([]domain.Technology, error) {
	filter := bson.D{
		primitive.E{Key: "display", Value: true},
	}

	cursor, err := r.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var results []domain.Technology
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *TechnologiesRepository) FindByName(name string) (domain.Technology, error) {
	filter := bson.D{
		primitive.E{Key: "name", Value: name},
	}

	var result domain.Technology
	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TechnologiesRepository) FindByid(id string) (domain.Technology, error) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Technology{}, err
	}
	filter := bson.D{
		primitive.E{Key: "_id", Value: idObject},
	}

	var result domain.Technology
	err = r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TechnologiesRepository) SetDisplay(id string, display bool) error {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": idObject}
	update := bson.M{"$set": bson.M{"display": display}}
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
