package queuesPool

import "testing"

func TestInitialize(t *testing.T) {
	pool := Pool{}

	if len(pool.queues) != 0 {
		t.Error("expect empty queues after init pool")
	}
}

func TestAdd(t *testing.T) {
	pool := Pool{}
	queue := Queue{
		name: "test queue",
		inputChannel: make(chan []byte),
		outputChannel: make(chan []byte),
	}
	pool.addQueue(queue)

	if len(pool.queues) != 1 {
		t.Error("expect one queue in pool after adding")
	}
}

// free queue from empty pool
func TestGiveFreeQueueWithEmptyPool(t *testing.T) {
	pool := Pool{}
	queue := pool.GiveFreeQueue("test queue name")

	if queue.name == "" {
		t.Error("It does not return queue")
	}
	if len(pool.queues) != 1 {
		t.Error("It does not add queue to pool")
	}
	if queue.name != "test queue name" {
		t.Error("It creates queue with wrong name")
	}
}

// free queue if pool has queue with given name
func TestGiveFreeQueueWithNotEmptyPool(t *testing.T) {
	pool := Pool{}
	queue := Queue{
		name: "test queue",
		inputChannel: make(chan []byte),
		outputChannel: make(chan []byte),
	}
	pool.addQueue(queue)
	elseQueue := pool.GiveFreeQueue("test queue")

	if len(pool.queues) != 1 {
		t.Error("It created new queue")
	}

	if elseQueue != queue {
		t.Error("It changes old queue")
	}
}

// free queue if pool is not empty but has not queue with given name
func TestPoolGiveFreeQueueWithOtherNameQueue(t *testing.T) {
	pool := Pool{}
	queue := Queue{
		name: "test queue",
		inputChannel: make(chan []byte),
		outputChannel: make(chan []byte),
	}
	pool.addQueue(queue)
	elseQueue := pool.GiveFreeQueue("new name queue")

	if len(pool.queues) != 2 {
		t.Error("It does not create new queue")
	}

	if elseQueue.name != "new name queue" {
		t.Error("it assing wrong name for queue")
	}

	if elseQueue == queue {
		t.Error("It creates a copy of old queue")
	}
}
