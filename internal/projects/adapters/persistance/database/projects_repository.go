package repositories

import (
	"context"

	domain "github.com/yagoinacio/portfolio-server/internal/projects/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type databaseProject struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Title       string               `bson:"title"`
	Description []string             `bson:"description"`
	Techs       []primitive.ObjectID `bson:"techs"`
	Display     bool                 `bson:"display"`
}

type ProjectsRepository struct {
	collection *mongo.Collection
}

func NewProjectsRepository(coll *mongo.Collection) *ProjectsRepository {
	return &ProjectsRepository{
		collection: coll,
	}
}

func (r *ProjectsRepository) Save(project *domain.Project) error {
	var techs []primitive.ObjectID

	for _, tech := range project.Techs {
		techs = append(techs, tech.ID)
	}

	projectToInsert := r.convertProject(project, techs)

	_, err := r.collection.InsertOne(context.TODO(), projectToInsert)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectsRepository) ListDisplayable() ([]domain.Project, error) {
	lookupStage := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "technologies"},
			{Key: "localField", Value: "techs"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "techs"},
		}},
	}
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{{Key: "display", Value: true}}},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage, matchStage})
	if err != nil {
		return nil, err
	}

	var results []domain.Project
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (r *ProjectsRepository) FindById(id string) (domain.Project, error) {
	idObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Project{}, err
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

	var result domain.Project
	cursor, err := r.collection.Aggregate(context.TODO(), mongo.Pipeline{lookupStage, matchStage})
	if err != nil {
		return result, err
	}

	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func (r *ProjectsRepository) Update(project *domain.Project) error {
	var techs []primitive.ObjectID

	for _, tech := range project.Techs {
		techs = append(techs, tech.ID)
	}

	projectToUpdate := r.convertProject(project, techs)

	filter := bson.M{"_id": project.ID}
	update := bson.M{"$set": projectToUpdate}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectsRepository) convertProject(project *domain.Project, techs []primitive.ObjectID) *databaseProject {
	return &databaseProject{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		Display:     project.Display,
		Techs:       techs,
	}
}
