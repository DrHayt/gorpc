package main

import (
	"bytes"

	"github.com/drhayt/bomstripper"
	"github.com/gorilla/rpc/v2/json2"

	// "github.com/gorilla/rpc/v2/json2"
	// "github.com/haisum/rpcexample"

	"log"
	"net/http"
)

func main() {
	var i int
	count := 3
	for i = 1; i < count; i++ {
		// test_multiply(i, i+2)
		// test_ping(i)
		test_prod_ping(i)
	}
}

func test_ping(a int) {
	// Where are we going to send this puppy?
	url := "http://localhost:1234/rpc"

	//  Procedure Arguments Struct
	type Args struct {
		A int
	}

	//  Procedure Arguments Struct
	type Result string

	//  Procedure Arguments Filled out
	args := &Args{
		A: a,
	}

	// Encode the message
	message, err := json2.EncodeClientRequest("Utility.Ping", args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	log.Printf("Sending Request: %s\n", message)
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	// Make a client
	client := new(http.Client)

	// Do the post
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result Result
	err = json2.DecodeClientResponse(resp.Body, &result)
	log.Printf("Received result: %s\n", resp.Body)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("Remote side asked to sleep %d, slept %d\n", args.A, result)

}

func test_multiply(a int, b int) {
	// Where are we going to send this puppy?
	url := "http://localhost:1234/rpc"

	//  Procedure Arguments Struct
	type Args struct {
		A, B int
	}

	//  Procedure Arguments Struct
	type Result int

	//  Procedure Arguments Filled out
	args := &Args{
		A: a,
		B: b,
	}

	// Encode the message
	message, err := json2.EncodeClientRequest("Arith.Multiply", args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	// Make a client
	client := new(http.Client)

	// Do the post
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result Result
	err = json2.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("%d*%d=%d\n", args.A, args.B, result)

}
func test_prod_ping(a int) {
	// Where are we going to send this puppy?
	url := "http://servicebus/Execute.svc/Execute"

	//  Procedure Arguments Struct
	type Args struct {
		Delaymilliseconds int
	}

	//  Procedure Arguments Struct
	type Result struct {
		PingResponse string
	}

	//  Procedure Arguments Filled out
	args := &Args{
		Delaymilliseconds: a,
	}

	// Encode the message
	message, err := json2.EncodeClientRequest("Utility.Ping", args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("Sending Request: %s\n", message)
	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	// Make a client
	client := new(http.Client)

	// Do the post
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	bomstripper := &bomstripper.BomStrip{R: resp.Body}

	var result Result
	// err = json2.DecodeClientResponse(resp.Body, &result)
	err = json2.DecodeClientResponse(bomstripper, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("Result was %s\n", result)

}
