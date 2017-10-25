package queuesPool

// Queue struct with name and string channels for input/output
type Queue struct {
	name                        string
	inputChannel, outputChannel chan string
}
