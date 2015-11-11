package main

import "github.com/drhayt/contentservice_client"
import "math/rand"

// "github.com/gorilla/rpc/v2/json2"
// "github.com/haisum/rpcexample"

func main() {
	for i := 0; i <= 500; i++ {
		contentservice_client.List(rand.Intn(1000000) + 100000000)
		// utility_client.Ping(0)
	}

}
