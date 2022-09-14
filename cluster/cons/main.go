package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	if err := setup(); err != nil {
		panic(err)
	}
}

func setup() error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:               []string{"localhost:9091", "localhost:9092", "localhost:9093", "localhost:9094", "localhost:9095"},
		GroupID:               "email-notification",
		Topic:                 "users",
		MaxBytes:              10e6,
		MaxWait:               0,
		HeartbeatInterval:     0,
		CommitInterval:        0,
		WatchPartitionChanges: true,
		RetentionTime:         time.Hour * 168, // 7 days
		StartOffset:           kafka.LastOffset,
	})

	defer r.Close()

	ctx, _ := context.WithCancel(context.Background())
	fmt.Println("starting")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			return err
		}

		fmt.Println("Key", string(m.Key))
		fmt.Println("Value", string(m.Value))
	}
}
