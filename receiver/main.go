package main

import (
	"log"
	"net/http"
	"tolling/common"
)

const (
	kafkaTopic = "test-topic"
	kafkaAddrs = "127.0.0.1:29092"
)

func main() {
	logger := common.NewCustomLogger()
	l := logger.New()

	receiver, err := NewDataReceiver(logger)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", receiver.handler)

	if err := http.ListenAndServe(":30000", nil); err != nil {
		l.Errorf("failed to start receiver server: \n%v", err)
	} else {
		l.Infof("receiver server started on address %s", kafkaAddrs)
	}
}
