package databases

import "ssibrahimbas.com/e-trade/config"

type Database struct {
	MongoDB *MongoDB
}

func New(c *config.Config) *Database {
	m := NewMongoDB(c.MongoDB.Uri, c.MongoDB.DbName)
	return &Database{
		MongoDB: m,
	}
}
