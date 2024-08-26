// Setup HTTP Connection
// Test using the API Open Source Routing Machine
// Check how many calls can be made and in what timeframe
// Check possibility of adding in an interactive map


package main

import {
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"errors"
	"os"
	"time"
}

const serverPort = 3333

func main() {
	requestURL := fmt.Sprintf("", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("status code: %d\n", res.StatusCode)
}

