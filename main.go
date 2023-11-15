package main

import (
	"fmt"
	pb "protokol_buffer/ProtokolBuffer"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Hello World")

	value := &pb.Users{
		Username:  "John",
		Password:  "123",
		Photo:     "photo",
		CodeUsers: 123,
	}

	data, err := proto.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

}
