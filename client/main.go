package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tommy-sho/memory-store-sample/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("34.85.103.2:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := proto.NewSampleServiceClient(conn)

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

L:
	for s.Scan() {
		var key, value string
		n := s.Text()
		switch n {
		case "set":
			fmt.Print("key: > ")
			if s.Scan() {
				key = s.Text()
			}

			fmt.Print("value: > ")
			if s.Scan() {
				value = s.Text()
			}

			_, err := client.Register(ctx, &proto.RegisterRequest{Key: key, Value: value})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("set key: %v, value: %v \n", key, value)
			fmt.Print("> ")

		case "get":
			if s.Scan() {
				key = s.Text()
			}
			r, err := client.Get(ctx, &proto.GetRequest{Key: key})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("key: %v, value: %v \n", r.Key, r.Value)
			fmt.Print("> ")
		case "ex":
			break L
		default:
		}
	}

}
