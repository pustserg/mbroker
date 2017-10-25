package queuesPool

// Pool slice with queues
type Pool struct {
	queues []Queue
}

// AddQueue adds queue to pool
func (pool *Pool) AddQueue(q Queue) {
	pool.queues = append(pool.queues, q)
}
