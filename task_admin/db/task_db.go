package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Task struct {
	TaskID      int    `bson:"taskid"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
}

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}

	collection = client.Database("task_manager").Collection("tasks")
}

func getNextTaskID() int {
	options := options.FindOne().SetSort(bson.M{"taskid": -1})
	result := collection.FindOne(context.TODO(), bson.M{}, options)
	task := Task{}
	err := result.Decode(&task)
	if err == mongo.ErrNoDocuments {
		return 1
	} else if err != nil {
		fmt.Println("Failed to retrieve the maximum task ID:", err)
		panic(err)
	}

	return task.TaskID + 1
}

func AddTask(task *Task) error {
	taskID := getNextTaskID()
	task.TaskID = taskID

	_, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskID int) error {
	filter := bson.M{"taskid": taskID}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(taskID int, task Task) error {
	filter := bson.M{"taskid": taskID}
	update := bson.M{"$set": task}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GetTask(taskID int) (Task, error) {
	filter := bson.M{"taskid": taskID}
	var task Task

	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func GetTasks() ([]Task, error) {
	var tasks []Task

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
