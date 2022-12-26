package controller

import (
	"github.com/MelvinKim/golang-gin-gonic/entity"
	"github.com/MelvinKim/golang-gin-gonic/service"
	"github.com/MelvinKim/golang-gin-gonic/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error //we will use the context to access the values that come with the http request
}

// takes in a service
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	// add a custom validator
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	// extract payload from the struct
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	// validate the data before saving it
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)

	return nil
}
