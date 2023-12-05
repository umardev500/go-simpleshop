package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Post(url string, data []byte) ([]byte, error) {
	fmt.Println("Processing...")
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	key := os.Getenv("PG_KEY")
	keyEncoded := EncodeBase64(key)
	auth := fmt.Sprintf("Basic %s", keyEncoded)
	// Set headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return nil, err
	}

	return body, nil
}
