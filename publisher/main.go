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

    // ایجاد stream اگر وجود ندارد
    _, err = js.AddStream(&nats.StreamConfig{
        Name:     "EVENTS",
        Subjects: []string{"events.created"},
    })
    if err != nil && err != nats.ErrStreamNameAlreadyInUse {
        log.Fatal(err)
    }

    for {
        msg := []byte(time.Now().Format(time.RFC3339))
        _, err := js.Publish("events.created", msg)
        if err != nil {
            log.Println("publish error:", err)
        } else {
            log.Println("Published:", string(msg))
        }
        time.Sleep(2 * time.Second)
    }
}
