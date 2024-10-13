package test

import (
	"fmt"
	"project1/app/models"
	"testing"

	"github.com/bxcodec/faker/v4"
)

func TestMain(t *testing.T) {
	// assert := assert.New(t)

	a := models.User{
		Email:    faker.Email(),
		Password: faker.Word(),
	}
	fmt.Println(a)
}
