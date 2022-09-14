package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	if err := setup(); err != nil {
		panic(err)
	}
}

func setup() error {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9091", "localhost:9092", "localhost:9093", "localhost:9094", "localhost:9095"),
		Topic:                  "users",
		BatchSize:              0,
		BatchBytes:             0,
		BatchTimeout:           0,
		ReadTimeout:            0,
		WriteTimeout:           0,
		RequiredAcks:           1,
		AllowAutoTopicCreation: true,
	}

	defer w.Close()

	for i := 40; i < 80; i++ {
		msg := fmt.Sprintf("check%d@check.com", i)
		msgs := []kafka.Message{
			{
				Value: []byte(msg),
			},
		}

		err := w.WriteMessages(context.Background(), msgs...)
		if err != nil {
			return err
		}
	}

	return nil
}
