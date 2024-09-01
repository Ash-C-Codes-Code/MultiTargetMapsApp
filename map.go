// Setup HTTP Connection
// Test using the API Open Source Routing Machine
// Check how many calls can be made and in what timeframe
// Check possibility of adding in an interactive map


package main

import (
	"fmt" 
	// "bytes"
	// "encoding/json"
	// "io"
	"net/http"
	// "errors"
	"os"
	// "time"
)

const serverPort = 3333

func main() {
	requestURL := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/13.388860,52.517037;13.397634,52.529407;13.428555,52.523219?overview=false", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("status code: %d\n", res.StatusCode)
}

