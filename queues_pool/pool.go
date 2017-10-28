package queuesPool

// Pool slice with queues
type Pool struct {
	queues []Queue
}

func (pool *Pool) addQueue(q Queue) {
	pool.queues = append(pool.queues, q)
}

// Return free queue by name, or creates new queue with given name
func (pool *Pool) GiveFreeQueue(name string) Queue {
	var queue Queue
	for i:=0; i < len(pool.queues); i++ {
		if pool.queues[i].name == name {
			queue = pool.queues[i]
		}
	}
	if queue.name == "" {
		queue = Queue{
			name: name,
			inputChannel: make(chan []byte),
			outputChannel: make(chan []byte),
		}
		pool.addQueue(queue)
	}
	return queue
}
