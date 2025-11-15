package main

import (
    "log"
    "time"

    "github.com/nats-io/nats.go"
)

func connectNATS(url string) *nats.Conn {
    var nc *nats.Conn
    var err error

    for i := 0; i < 10; i++ {
        nc, err = nats.Connect(url)
        if err == nil {
            log.Println("Connected to NATS")
            return nc
        }
        log.Println("Waiting for NATS to be ready...", err)
        time.Sleep(2 * time.Second)
    }

    log.Fatal("Failed to connect to NATS after 10 attempts:", err)
    return nil
}

func main() {
    nc := connectNATS("nats://nats:4222")
    defer nc.Close()

    js, err := nc.JetStream()
    if err != nil {
        log.Fatal(err)
    }

    // PullSubscribe با durable درست
    sub, err := js.PullSubscribe(
        "events.created",
        "worker-consumer", // durable name
    )
    if err != nil {
        log.Fatal(err)
    }

    for {
        msgs, err := sub.Fetch(10, nats.MaxWait(5*time.Second))
        if err != nil && err != nats.ErrTimeout {
            log.Println("fetch error:", err)
            continue
        }

        for _, m := range msgs {
            log.Println("Received:", string(m.Data))
            m.Ack()
        }
    }
}
