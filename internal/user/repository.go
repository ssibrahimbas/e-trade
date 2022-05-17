package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ssibrahimbas.com/e-trade/databases"
	"ssibrahimbas.com/e-trade/services"
)

type Repository struct {
	c      *mongo.Collection
	ctx    context.Context
	cipher *services.Cipher
}

func NewRepository(m *databases.MongoDB, cipher *services.Cipher) *Repository {
	c := m.GetCollection("users")
	return &Repository{
		c:      c,
		ctx:    context.TODO(),
		cipher: cipher,
	}
}

func (r *Repository) Create(u *CreateUserDto, ip string) (interface{}, error) {
	pw, err := r.cipher.Encrypt(u.Password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Id:          "",
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		UserName:    u.UserName,
		Email:       u.Email,
		Password:    pw,
		Active:      true,
		IpAddress:   ip,
		DateOfBirth: u.DateOfBirth,
		Phone:       u.Phone,
	}
	id, err := r.c.InsertOne(r.ctx, user)
	return id.InsertedID, err
}

func (r *Repository) GetByEmail(e string) (User, error) {
	result := r.c.FindOne(r.ctx, bson.M{"email": e})
	user := User{}
	err := result.Decode(&user)
	return user, err
}

func (r *Repository) GetByUserName(u string) (User, error) {
	result := r.c.FindOne(r.ctx, bson.M{"userName": u})
	user := User{}
	err := result.Decode(&user)
	return user, err
}

func (r *Repository) GetById(id primitive.ObjectID) (User, error) {
	result := r.c.FindOne(r.ctx, bson.M{"_id": id})
	user := User{}
	err := result.Decode(&user)
	return user, err
}

func (r *Repository) CheckPassword(u *User, p string) (bool, error) {
	return r.cipher.Compare(p, []byte(u.Password))
}
