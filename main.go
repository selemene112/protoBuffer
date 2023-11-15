package main

import (
	"fmt"
	"log"
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

	dataencode := &pb.Users{}

	if err = proto.Unmarshal(data, dataencode); err != nil {
		log.Fatal("Unmarshal error ", err)
	}
	// code := proto.Unmarshal(data, dataencode)

	fmt.Println(dataencode)

}
