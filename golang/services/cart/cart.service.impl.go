package services

import (
	"context"
	"errors"

	"abc-pharmacy.com/Upeksha1214/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartServiceImpl struct {
	cartcollection *mongo.Collection
	ctx            context.Context
}

func NewCartService(cartcollection *mongo.Collection, ctx context.Context) CartService {
	return &CartServiceImpl{
		cartcollection: cartcollection,
		ctx:            ctx,
	}
}

// CreateCart implements CartService.
func (c *CartServiceImpl) CreateCart(cart *models.Cart) error {
	_, err := c.cartcollection.InsertOne(c.ctx, cart)
	return err
}

// DeleteCart implements CartService.
func (c *CartServiceImpl) DeleteCart(user_Id*string) error {
	filter := bson.D{primitive.E{Key: "user_Id", Value: user_Id}}
	result, _ := c.cartcollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
// GetAll implements CartService.
func (c *CartServiceImpl) GetAll() ([]*models.Cart, error) {
	var carts []*models.Cart
	cursor, err := c.cartcollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var cart models.Cart
		err := cursor.Decode(&cart)
		if err != nil {
			return nil, err
		}
		carts = append(carts, &cart)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(carts) == 0 {
		return nil, errors.New("documents not found")
	}
	return carts, nil
}


// GetCart implements CartService.
func (c *CartServiceImpl) GetCart(id *string) (*models.Cart, error) {
	var cart *models.Cart
	query := bson.D{bson.E{Key: "user-Id", Value: id}}
	err := c.cartcollection.FindOne(c.ctx, query).Decode(&cart)
	return cart, err
}

// UpdateCart implements CartService.
func (c *CartServiceImpl) UpdateCart(cart *models.Cart) error {
	filter := bson.D{primitive.E{Key: "user-Id", Value: cart.UserId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: cart.Name}, primitive.E{Key: "brand", Value: cart.Brand}, primitive.E{Key: "qty", Value: cart.Qty}}}}
	result, _ := c.cartcollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

