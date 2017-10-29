package router

import "github.com/pustserg/mbroker/queues_pool"

// Route takes a message and sends it to proper queue from Pool
func Route(msg *Message, pool *queuesPool.Pool) (*queuesPool.Queue) {
	queue := pool.GiveFreeQueue(msg.QueueName)
	queue.InputChan() <- msg.Body
	return &queue
}
