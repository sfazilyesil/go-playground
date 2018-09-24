package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sfazilyesil/go-playground/encoding-formats/protocol-buffers"
	"github.com/sfazilyesil/go-playground/encoding-formats/protocol-buffers/sm"
	"io/ioutil"
)

func main() {

	addressBook()
	simpleMsg()
}

func addressBook() {
	fmt.Println("-------------------")
	book := &pb.AddressBook{People: []*pb.Person{
		{
			Name: "ali",
		},
	}}

	out, err := proto.Marshal(book)
	if err != nil {
		fmt.Println("Failed to encode address book:", err)
	}

	fmt.Printf("Decimal: %d\n", out)
	fmt.Printf("Hex: %x\n", out)
	fmt.Printf("Binary: %b\n", out)
	fmt.Printf("# of bytes: %d\n", len(out))

	if err := ioutil.WriteFile("/tmp/proto_out", out, 0644); err != nil {
		fmt.Println("Failed to write address book:", err)
	}
}

func simpleMsg() {
	fmt.Println("---------simple msg---------")

	var a int32 = 150
	simpleMsg := &sm.Message{A: &a}

	out, err := proto.Marshal(simpleMsg)
	if err != nil {
		fmt.Println("Failed to encode simple msg:", err)
	}

	fmt.Printf("Decimal: %d\n", out)
	fmt.Printf("Hex: %x\n", out)
	fmt.Printf("Binary: %b\n", out)
	fmt.Printf("# of bytes: %d\n", len(out))

	fmt.Println("----------------------------")
}
