package repositories

import (
	"context"

	"escort-book-delete-customers/config"
	"escort-book-delete-customers/db"
	"escort-book-delete-customers/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var customerRemovalCollection = config.InitializeMongo().Collections.SchedulerCustomerRemoval

type CustomerRemovalRepository struct {
	Data *db.MongoClient
}

func (r CustomerRemovalRepository) Get(ctx context.Context, filter interface{}) (models.CustomerRemoval, error) {
	collection := r.Data.SchedulerDB.Collection(customerRemovalCollection)
	singleResult := collection.FindOne(ctx, filter)

	var removal models.CustomerRemoval

	if err := singleResult.Decode(&removal); err != nil {
		return removal, err
	}

	return removal, nil
}

func (r CustomerRemovalRepository) GetMany(ctx context.Context, filter bson.M, offset, limit int64) ([]models.CustomerRemoval, error) {
	collection := r.Data.SchedulerDB.Collection(customerRemovalCollection)

	skip := offset * limit
	findOptions := options.FindOptions{Limit: &limit, Skip: &skip}
	cursor, err := collection.Find(ctx, filter, &findOptions)

	var removals []models.CustomerRemoval

	if err != nil {
		return removals, err
	}

	if err := cursor.All(ctx, &removals); err != nil {
		return removals, err
	}

	return removals, nil
}

func (r CustomerRemovalRepository) Create(ctx context.Context, removal models.CustomerRemoval) (models.CustomerRemoval, error) {
	collection := r.Data.SchedulerDB.Collection(customerRemovalCollection)

	if _, err := collection.InsertOne(ctx, removal); err != nil {
		return removal, err
	}

	return removal, nil
}

func (r CustomerRemovalRepository) Update(ctx context.Context, filter interface{}, document interface{}) (models.CustomerRemoval, error) {
	collection := r.Data.SchedulerDB.Collection(customerRemovalCollection)
	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	updateResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": document}, &updateOptions)
	removal := models.CustomerRemoval{}

	if err := updateResult.Decode(&removal); err != nil {
		return removal, err
	}

	return removal, nil
}

func (r CustomerRemovalRepository) Delete(ctx context.Context, filter interface{}) error {
	collection := r.Data.SchedulerDB.Collection(customerRemovalCollection)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
