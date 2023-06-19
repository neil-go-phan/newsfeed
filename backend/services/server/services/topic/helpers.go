package topicservices

import (
	"server/entities"
	"server/services"

	"github.com/go-playground/validator"
)

func validateTopic(topic entities.Topic) (error) {
	validate := validator.New()
	err := validate.Struct(topic)
	if err != nil {
		return err
	}

	return nil
}

func castEntityTopicToTopicResponse(entityTopic entities.Topic) (services.TopicResponse) {
	return services.TopicResponse{
		Name: entityTopic.Name,
		ID: entityTopic.ID,
		CategoryID: entityTopic.CategoryID,
	}
}