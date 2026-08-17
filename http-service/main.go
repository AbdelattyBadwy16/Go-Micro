package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"test"

	pb "github.com/abdelatty/grpc-service/proto"

	"google.golang.org/grpc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer conn.Close()
	// Hello

	client := pb.NewMessageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SendMessage(ctx, &pb.MessageRequest{
		Text: "Hello from HTTP 🚀",
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(resp.Reply))
}

func main() {
	http.HandleFunc("/send", handler)
	log.Println("HTTP service running on :8080")
	http.ListenAndServe(":8080", nil)
}
