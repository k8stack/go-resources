package main

import (
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go-protobuf/src/user"
	"io/ioutil"
	"log"
)

func doUser() *userpb.User {
	u := userpb.User{
		Id:       121,
		Username: "zillani",
		Email:    "shaikzillani",
		Address: []*userpb.Address{
			&userpb.Address{
				AddressLane_1: "a-36 begumpet",
				AddressLane_2: "near airport",
				City:          "hyderabad",
				State:         "telangana",
			},
		},
		EyeColor: userpb.EyeColor_BROWN_EYE_COLOR,
		Password: "root",
	}
	return &u
}
func protoToJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("unable to marshall", err)
		return ""
	}
	return out
}

func jsonToProto(in string, pb proto.Message) proto.Message {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("unable to marshall!")
	}
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatal("unable to write the file!", err)
		return err
	}
	fmt.Println("Successfully written to a file!")
	return nil
}

func readFile(filename string, pb proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Unable to read the file!", err)
		return err
	}
	if err := proto.Unmarshal(data, pb); err != nil {
		log.Fatalln("Something went wrong, unable to read file", err)
		return err
	}
	return nil
}
func main() {
	u := doUser()
	fmt.Println(u)
	if err := writeToFile("user.data", u); err != nil {
		log.Fatalln("Failed to write to file!", err)
	}
	u2 := &userpb.User{}
	if err := readFile("user.data", u2); err != nil {
		log.Fatalln("Failed", err)
	}
	fmt.Println("###########  READING THE FILE TO OUT  #########")
	fmt.Println(u2)
	fmt.Println("###########  PROTO TO JSON  #########")
	result := protoToJson(u)
	fmt.Println(result)
	u3 := &userpb.User{}
	fmt.Println("###########  JSON TO PROTO #########")

	jsonToProto(result, u3)
	fmt.Println(u3)

}
