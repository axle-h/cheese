package store

import (
	"github.com/axle-h/cheese/store/models"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type PhotoRepository struct {
	context MongoContext
}

func NewPhotoRepository(context MongoContext) PhotoRepository {
	return PhotoRepository{context}
}

func (r PhotoRepository) getCollection() *mongo.Collection {
	return r.context.GetCollection("photos")
}

func (r PhotoRepository) GetAll(photos *[]models.Photo) error {
	collection := r.getCollection()

	cursor, err := collection.Find(r.context.timeout(), nil)
	if err != nil {
		return err
	}

	defer cursor.Close(r.context.timeout())

	for cursor.Next(r.context.timeout()) {
		var entity = models.Photo{}
		err := cursor.Decode(&entity)
		if err != nil {
			return err
		}

		*photos = append(*photos, entity)
	}

	return nil
}

func (r PhotoRepository) Add(photo models.Photo) (string, error) {
	collection := r.getCollection()
	result, err := collection.InsertOne(r.context.timeout(), photo)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).String(), nil
}

func (r PhotoRepository) AddMany(photos []models.Photo) error {
	collection := r.getCollection()

	entities := make([]interface{}, len(photos))
	for i, v := range photos {
		entities[i] = v
	}

	_, err := collection.InsertMany(r.context.timeout(), entities)
	if err != nil {
		return err
	}
	return nil
}