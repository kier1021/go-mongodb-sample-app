package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/kier1021/go-mongodb-sample-app/api/apierrors"
	"github.com/kier1021/go-mongodb-sample-app/api/models"
	"github.com/kier1021/go-mongodb-sample-app/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	mongoDB        *db.MongoDB
	mongoClient    *mongo.Client
	todoCollection *mongo.Collection
}

func NewTodoRepository(mongoDB *db.MongoDB) *TodoRepository {

	client := mongoDB.GetClient()
	todoCollection := client.Database(os.Getenv("DB_MONGODB_NAME")).Collection("todos")

	return &TodoRepository{
		mongoDB:        mongoDB,
		mongoClient:    client,
		todoCollection: todoCollection,
	}
}

func (repo *TodoRepository) GetTodos() (todos []models.Todo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := repo.todoCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var todo models.Todo
		err := cur.Decode(&todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (repo *TodoRepository) GetTodoByID(todoID string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	res := repo.todoCollection.FindOne(ctx, bson.M{"_id": idPrimitive})

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, res.Err()
	}

	res.Decode(&todo)
	return &todo, nil
}

func (repo *TodoRepository) AddTodo(todo models.Todo) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var data map[string]interface{}
	b, _ := json.Marshal(todo)
	json.Unmarshal(b, &data)

	data["_id"] = primitive.NewObjectID()

	res, err := repo.todoCollection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *TodoRepository) DeleteTodoByID(todoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return err
	}

	res, err := repo.todoCollection.DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return apierrors.NO_ENTITY_DELETED_ERROR
	}

	return nil
}

func (repo *TodoRepository) UpdateTodoByID(todoID string, data map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idPrimitive, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return err
	}

	res, err := repo.todoCollection.UpdateOne(
		ctx, bson.M{"_id": idPrimitive},
		bson.D{
			{Key: "$set", Value: data},
		},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return apierrors.NO_ENTITY_UPDATED_ERROR
	}

	return nil
}
