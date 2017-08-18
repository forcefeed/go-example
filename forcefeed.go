package main

import (
	"fmt"
	"github.com/r3labs/sse"
)

func main() {
	client := sse.NewClient("https://forcefeed.ir/sse")
	client.Subscribe("post", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
	})
}
