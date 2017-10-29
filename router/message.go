package router

type Message struct {
	Body 	  []byte
	QueueName string
}

func NewMessage(name string, body []byte) *Message {
	message := Message{QueueName: name, Body: body}
	return &message
}
