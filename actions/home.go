package actions

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}

func CreateHealthData() error {
	fmt.Println("Test")
	database := GetDatabase()
	collection := database.Collection(Name)
	dummyHealth := HealthModel{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Test Title",
		Description: "Test Description",
	}
	_, err := collection.InsertOne(context.TODO(), dummyHealth)
	if err != nil {
		return err
	}
	return nil
}

func CreateHealthHandler(c buffalo.Context) error {
	err := CreateHealthData()
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(map[string]string{"message": "Error occured while creating records!"}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Created records successfully!"}))
}
