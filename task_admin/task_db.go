package task_admin

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

var collection *mongo.Collection
var ctx = context.TODO()

type Task struct {
	TaskID      int    `bson:"taskid"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
}

func DbCreateTask(title string, desc string) *Task {
	return &Task{
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(desc),
		Status:      "Pending",
	}
}

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("task_manager").Collection("tasks")
}

func getNextTaskID() int {
	options := options.FindOne().SetSort(bson.M{"taskid": -1})
	result := collection.FindOne(ctx, bson.M{}, options)
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

func DbAddTask(task *Task) error {
	taskID := getNextTaskID()
	task.TaskID = taskID

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func DbDeleteTask(taskID int) error {
	filter := bson.M{"taskid": taskID}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func DbUpdateTask(taskID int, task Task) error {

	filter := bson.M{"taskid": taskID}
	update := bson.M{"$set": task}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DbGetTask(taskID int) (Task, error) {
	filter := bson.M{"taskid": taskID}
	var task Task

	err := collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func DbGetTasks() ([]Task, error) {
	var tasks []Task

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
