package product

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ssibrahimbas.com/e-trade/databases"
)

type Repository struct {
	c   *mongo.Collection
	ctx context.Context
}

func NewRepository(m *databases.MongoDB) *Repository {
	c := m.GetCollection("products")
	return &Repository{
		c:   c,
		ctx: context.TODO(),
	}
}

func (r *Repository) Create(p *CreateProductDto) (interface{}, error) {
	product := &Product{
		Id:         "",
		Name:       p.Name,
		Price:      p.Price,
		StockCount: p.StockCount,
		Date:       time.Now(),
	}
	id, err := r.c.InsertOne(r.ctx, product)
	return id.InsertedID, err
}

func (r *Repository) GetAll() ([]Product, error) {
	result, err := r.c.Find(r.ctx, bson.D{{}})
	products := make([]Product, 0)
	if err := result.All(r.ctx, &products); err != nil {
		return nil, err
	}
	return products, err
}

func (r *Repository) GetDetail(p *GetDetailProductDto) (Product, error) {
	result := r.c.FindOne(r.ctx, bson.M{"_id": p.Id})
	product := Product{}
	err := result.Decode(&product)
	return product, err
}
