package services

import (
	"context"
	"errors"

	"abc-pharmacy.com/Upeksha1214/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemServiceImpl struct {
	itemcollection *mongo.Collection
	ctx            context.Context
}

func NewItemService(itemcollection *mongo.Collection, ctx context.Context) ItemService {
	return &ItemServiceImpl{
		itemcollection: itemcollection,
		ctx:            ctx,
	}
}

// CreateItem implements ItemService.
func (i *ItemServiceImpl) CreateItem(item *models.Item) error {
	_, err := i.itemcollection.InsertOne(i.ctx, item)
	return err
}

// DeleteItem implements ItemService.
func (i*ItemServiceImpl) DeleteItem(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := i.itemcollection.DeleteOne(i.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

// GetAll implements ItemService.
func (i *ItemServiceImpl) GetAll() ([]*models.Item, error) {
	var items []*models.Item
	cursor, err := i.itemcollection.Find(i.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(i.ctx) {
		var item models.Item
		err := cursor.Decode(&items)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(i.ctx)

	if len(items) == 0 {
		return nil, errors.New("documents not found")
	}
	return items, nil
}

// GetItem implements ItemService.
func (i*ItemServiceImpl) GetItem(name*string) (*models.Item, error) {
	var item *models.Item
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := i.itemcollection.FindOne(i.ctx, query).Decode(&item)
	return item, err
}

// UpdateItem implements ItemService.
func (*ItemServiceImpl) UpdateItem(*models.Item) error {
	panic("unimplemented")
}


