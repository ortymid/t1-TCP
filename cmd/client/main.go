package main

import (
	"fmt"
	"os"

	"github.com/ortymid/t1-tcp/market"
	"github.com/ortymid/t1-tcp/mmtp"
)

func main() {

	fmt.Println("Enter a number representing a message type:")
	// fmt.Println("1 - Request Product")
	// fmt.Println("2 - Send Product")
	fmt.Println("3 - Request Product List")
	// fmt.Println("4 - Send Product List")
	fmt.Println("5 - Add Product")

	fmt.Print("Message type: ")
	typ := readInt()

	msg := &mmtp.Message{Type: mmtp.MessageType(typ)}

	switch msg.Type {
	case mmtp.MessageProductListRequest:
		msg.Payload = nil
	case mmtp.MessageProductAdd:
		fmt.Print("\nEnter product name: ")
		name := readString()
		fmt.Print("Enter product price: ")
		price := readInt()
		msg.Payload = &market.Product{Name: name, Price: price}
	default:
		fmt.Println("Unknown message type. Exiting.")
		os.Exit(0)
	}

	c, err := mmtp.Dial(":8080")
	if err != nil {
		panic(err)
	}
	c.SendMessage(msg)

	res, err := c.ReceiveMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nServer response:")
	fmt.Println(res)
}

func readInt() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return i
}

func readString() string {
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return s
}
