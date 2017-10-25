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
	queue := Queue{"test queue", make(chan string), make(chan string)}
	pool.AddQueue(queue)
	if len(pool.queues) != 1 {
		t.Error("expect one queue in pool after adding")
	}
}
