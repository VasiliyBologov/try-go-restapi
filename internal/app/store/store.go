package store

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Store ... DB
type Store struct {
	config          *Config
	logger          *logrus.Logger
	db_context      context.Context
	db_client       *mongo.Client
	user_collection *UserCollection
}

//New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
		logger: logrus.New(),
	}
}

// Open ... connect to DB
func (s *Store) Open() error {

	/*
	   Connect to my mongoDB
	   code by Docs https://www.mongodb.com/languages/golang and https://docs.mongodb.com/drivers/go/
	*/

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.config.DataBaseURL))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	s.db_context = ctx
	s.db_client = client
	s.logger.Info("Successfully connected to DB")

	return nil
}

// Close ...
func (s *Store) Close() {
	defer s.db_client.Disconnect(s.db_context)
}

// User ...
func (s *Store) User() *UserCollection {
	if s.user_collection != nil {
		return s.user_collection
	}

	s.user_collection = &UserCollection{
		store:      s,
		collection: s.db_client.Database("tryGo").Collection("users"),
	}

	return s.user_collection
}
