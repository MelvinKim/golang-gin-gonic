package service

import "github.com/MelvinKim/golang-gin-gonic/entity"

// create an interface, which defines the methods to be implelmented
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// create a struct to implement the interface
type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
