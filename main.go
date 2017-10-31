package main

import (
	"fmt"
	"github.com/pustserg/mbroker/router"
	"github.com/pustserg/mbroker/queues_pool"
	"github.com/pustserg/mbroker/mbroker_log"
)

func main() {
	logFile, err := mbroker_log.InitLog()
	if err != nil { panic("cannot open logfile") }
	defer logFile.Close()

	pool := queuesPool.NewPool()
	messages := []*router.Message{
		router.NewMessage("odd queue",  []byte("first message")),
		router.NewMessage("even queue", []byte("second message")),
		router.NewMessage("odd queue",  []byte("third message")),
		router.NewMessage("even queue", []byte("fourth message")),
		router.NewMessage("odd queue",  []byte("fifth message")),
		router.NewMessage("even queue", []byte("sixth message")),
		router.NewMessage("odd queue",  []byte("seventh message")),
		router.NewMessage("even queue", []byte("eighth message")),
		router.NewMessage("odd queue",  []byte("ninth message")),
		router.NewMessage("even queue", []byte("tenth message")),
	}

	firstQ := pool.GiveFreeQueue("odd queue")
	secondQ := pool.GiveFreeQueue("even queue")

	go receiver(firstQ.OutputChan(), "odd queue")
	go receiver(secondQ.OutputChan(), "even queue")

	for i := 0; i < len(messages); i++ {
		go router.Route(messages[i], pool)
	}

	var input string
	fmt.Scanln(&input)
}

func receiver(channel chan []byte, name string) {
	for {
		msg := <- channel
		fmt.Println(string(msg), "from", name)
	}
}
