package queuesPool

// Queue struct with name and string channels for input/output
type Queue struct {
	name                        string
	inputChannel, outputChannel chan []byte
}

func (q *Queue) InputChan() chan []byte {
	return q.inputChannel
}

func (q *Queue) OutputChan() chan []byte {
	return q.outputChannel
}
