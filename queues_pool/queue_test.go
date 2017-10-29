package queuesPool

import "testing"

func TestInputChan(t *testing.T) {
	inputChannel := make(chan []byte)
	queue := Queue{
		name: "test queue",
		inputChannel: inputChannel,
		outputChannel: make(chan []byte),
	}

	if queue.InputChan() != inputChannel {
		t.Error("it returns wrong channel")
	}
}

func TestOutputChan(t *testing.T) {
	outputChannel := make(chan []byte)
	queue := Queue{
		name: "test queue",
		inputChannel: make(chan []byte),
		outputChannel: outputChannel,
	}

	if queue.OutputChan() != outputChannel {
		t.Error("it returns wrong channel")
	}
}
