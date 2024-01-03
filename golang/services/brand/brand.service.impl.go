package services

import (
	"context"
	"errors"

	"abc-pharmacy.com/Upeksha1214/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BrandServiceImpl struct {
	brandcollection *mongo.Collection
	ctx             context.Context
}

func NewBrandService(brandcollection *mongo.Collection, ctx context.Context) BrandService {
	return &BrandServiceImpl{
		brandcollection: brandcollection,
		ctx:             ctx,
	}
}

// CreateBrand implements BrandService.
func (b *BrandServiceImpl) CreateBrand(brand *models.Brand) error {
	_, err := b.brandcollection.InsertOne(b.ctx, brand)
	return err
}

// DeleteBrand implements BrandService.
func (b *BrandServiceImpl) DeleteBrand(brand *string) error {
	filter := bson.D{primitive.E{Key: "brand", Value: brand}}
	result, _ := b.brandcollection.DeleteOne(b.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}


// GetAll implements BrandService.
func (b *BrandServiceImpl) GetAll() ([]*models.Brand, error) {
	var brands []*models.Brand
	cursor, err := b.brandcollection.Find(b.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(b.ctx) {
		var brand models.Brand
		err := cursor.Decode(&brand)
		if err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(b.ctx)

	if len(brands) == 0 {
		return nil, errors.New("documents not found")
	}
	return brands, nil
}

// GetBrand implements BrandService.
func (b *BrandServiceImpl) GetBrand(brandName *string) (*models.Brand, error) {
	var brand *models.Brand
	query := bson.D{bson.E{Key: "brand-name", Value: brandName}}
	err := b.brandcollection.FindOne(b.ctx, query).Decode(&brand)
	return brand, err
}

// UpdateBrand implements BrandService.
func (b *BrandServiceImpl) UpdateBrand(brand *models.Brand) error {
	filter := bson.D{primitive.E{Key: "brand-name", Value: brand.BrandName}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "brand-name", Value: brand.BrandName}, primitive.E{Key: "category", Value: brand.Category}}}}
	result, _ := b.brandcollection.UpdateOne(b.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}


