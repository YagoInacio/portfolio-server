package repositories

import (
	"context"
	"fmt"

	domain "github.com/yagoinacio/portfolio-server/internal/experiences/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type databaseExperience struct {
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
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "technologies"},
			{Key: "localField", Value: "techs"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "techs"},
		}},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage})
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
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "technologies"},
			{Key: "localField", Value: "techs"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "techs"},
		}},
	}
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{{Key: "_id", Value: idObject}}},
	}

	var result domain.Experience
	cursor, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage, matchStage})
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
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

func (r *ExperiencesRepository) convertExperience(xp *domain.Experience, techs []primitive.ObjectID) *databaseExperience {
	return &databaseExperience{
		ID:       xp.ID,
		Position: xp.Position,
		Company:  xp.Company,
		Logo:     xp.Logo,
		Period:   xp.Period,
		Summary:  xp.Summary,
		Techs:    techs,
	}
}
