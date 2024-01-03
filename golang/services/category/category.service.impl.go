package services

import (
	"context"
	"errors"

	"abc-pharmacy.com/Upeksha1214/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryServiceImpl struct {
	categorycollection *mongo.Collection
	ctx                context.Context
}

func NewCategoryService(categorycollection *mongo.Collection, ctx context.Context) CategoryService {
	return &CategoryServiceImpl{
		categorycollection: categorycollection,
		ctx:                ctx,
	}
}

// CreateCategory implements CategoryService.
func (ca *CategoryServiceImpl) CreateCategory(category *models.Category) error {
	_, err := ca.categorycollection.InsertOne(ca.ctx, category)
	return err
}

// DeleteCategory implements CategoryService.
func (ca *CategoryServiceImpl) DeleteCategory(category *string) error {
	filter := bson.D{primitive.E{Key: "category", Value: category}}
	result, _ := ca.categorycollection.DeleteOne(ca.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

// GetAll implements CategoryService.
func (ca *CategoryServiceImpl) GetAll() ([]*models.Category, error) {
	var categorys []*models.Category
	cursor, err := ca.categorycollection.Find(ca.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ca.ctx) {
		var category models.Category
		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}
		categorys = append(categorys, & category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ca.ctx)

	if len(categorys) == 0 {
		return nil, errors.New("documents not found")
	}
	return categorys, nil
}

// GetCategory implements CategoryService.
func (ca *CategoryServiceImpl) GetCategory(categoryN *string) (*models.Category, error) {
	var category *models.Category
	query := bson.D{bson.E{Key: "category", Value: categoryN}}
	err := ca.categorycollection.FindOne(ca.ctx, query).Decode(&category)
	return category, err
}

// UpdateCategory implements CategoryService.
func (ca *CategoryServiceImpl) UpdateCategory(category *models.Category) error {
	filter := bson.D{primitive.E{Key: "category", Value: category.Category}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "category", Value: category.Category}, primitive.E{Key: "type", Value: category.Type}}}}
	result, _ := ca.categorycollection.UpdateOne(ca.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
