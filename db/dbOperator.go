package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Task struct {
	TaskID      int    `bson:"taskid"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
}

type TaskBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTask(title string, desc string) *Task {
	return &Task{
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(desc),
		Status:      "Pending",
	}
}

func nextTaskID() int {
	setSort := options.FindOne().SetSort(bson.M{"taskid": -1})
	result := Collection.FindOne(Ctx, bson.M{}, setSort)
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
	taskID := nextTaskID()
	task.TaskID = taskID

	_, err := Collection.InsertOne(Ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func DbDeleteTask(taskID int) error {
	filter := bson.M{"taskid": taskID}

	_, err := Collection.DeleteOne(Ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func DbUpdateTask(taskID int, task Task) error {

	filter := bson.M{"taskid": taskID}
	update := bson.M{"$set": task}

	_, err := Collection.UpdateOne(Ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DbGetTask(taskID int) (Task, error) {
	filter := bson.M{"taskid": taskID}
	var task Task

	err := Collection.FindOne(Ctx, filter).Decode(&task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func DbGetTasks() ([]Task, error) {
	var tasks []Task

	cursor, err := Collection.Find(Ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, Ctx)

	for cursor.Next(Ctx) {
		var task Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
