package product

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var RepoErrreponoql = errors.New("Unable to handle Repo Request")

type reponoql struct {
	db     *mongo.Client
	logger log.Logger
}

func NewReponoSql(db *mongo.Client, logger log.Logger) Repository {
	return &reponoql{
		db:     db,
		logger: log.With(logger, "repo", "nosql"),
	}
}

func (repo *reponoql) CreateProduct(ctx context.Context, product Product) error {
	if product.Name == "" || product.Sku == "" || product.Description == "" || product.Price == 0 {
		return RepoErrreponoql
	}
	_, err := repo.db.Database("product").Collection("products").InsertOne(ctx, bson.D{
		{"id", product.ID},
		{"sku", product.Sku},
		{"name", product.Name},
		{"description", product.Description},
		{"price", product.Price},
	})
	if err != nil {
		return err
	}
	return nil
}
func (repo *reponoql) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	cursor, err := repo.db.Database("product").Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, RepoErrSql
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, RepoErrSql
	}
	return products, nil
}
