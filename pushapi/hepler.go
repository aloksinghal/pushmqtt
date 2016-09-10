package pushapi

import (
    "fmt"
)

func PublishMessageonMQTT(requestBody map[string]interface{}) string {
	fmt.Println(requestBody)
	message := requestBody["message"]
	topic := requestBody["topic"].(string)
	_ = topic
	_ = message

	client.Publish(topic, 0, false, message)
	return "done"

	}