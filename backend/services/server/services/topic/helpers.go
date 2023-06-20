package topicservices

import (
	"fmt"
	"server/entities"
	"server/services"

	"github.com/go-playground/validator"
)

func validateTopic(topic entities.Topic) (error) {
	if topic.Name == OTHERS_TOPIC_NAME || topic.ID == OTHERS_TOPIC_ID {
		return fmt.Errorf("can not change 'Others' topics")
	}
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