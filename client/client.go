package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	chat "github.com/jedwards1230/go-web/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// Parse cli arguments
	ip := flag.String("ip", "127.0.0.1", "IP for the client to connect to")
	port := flag.Int("port", 8090, "Port for the client to connect to")
	flag.Parse()

	// Connect to server
	var conn *grpc.ClientConn
	fmt.Println(fmt.Sprintf("Connecting to %v:%d", *ip, *port))
	fmt.Println("Sending request...")
	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", *ip, *port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect")
	}
	defer conn.Close()

	// Get client message
	fmt.Print("Message: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	// Format message to send to server
	msg := chat.Request{
		Message: input,
	}

	// Send message to server
	c := chat.NewChatServiceClient(conn)
	response, err := c.Hello(context.Background(), &msg)
	if err != nil {
		log.Fatal("Error when making request")
	}
	fmt.Println("Response from server:", response.Message)
}
