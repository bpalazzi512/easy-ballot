package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Example API calls for the positions endpoint
func main() {
	baseURL := "http://localhost:8080"

	// Example 1: Create a position
	fmt.Println("=== Creating a Position ===")
	createPositionData := map[string]string{
		"organization_id": "org123",
		"title":           "Software Engineer",
	}
	createPositionJSON, _ := json.Marshal(createPositionData)

	resp, err := http.Post(baseURL+"/positions", "application/json", bytes.NewBuffer(createPositionJSON))
	if err != nil {
		fmt.Printf("Error creating position: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Create Position Response: %s\n\n", string(body))

	// Example 2: Get all positions
	fmt.Println("=== Getting All Positions ===")
	resp, err = http.Get(baseURL + "/positions")
	if err != nil {
		fmt.Printf("Error getting positions: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Positions Response: %s\n\n", string(body))

	// Example 3: Get positions by organization
	fmt.Println("=== Getting Positions by Organization ===")
	resp, err = http.Get(baseURL + "/organizations/org123/positions")
	if err != nil {
		fmt.Printf("Error getting positions by organization: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Positions by Organization Response: %s\n\n", string(body))

	// Example 4: Update a position (assuming position ID is "pos123")
	fmt.Println("=== Updating a Position ===")
	updatePositionData := map[string]string{
		"organization_id": "org123",
		"title":           "Senior Software Engineer",
	}
	updatePositionJSON, _ := json.Marshal(updatePositionData)

	req, _ := http.NewRequest("PUT", baseURL+"/positions/pos123", bytes.NewBuffer(updatePositionJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error updating position: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Update Position Response: %s\n\n", string(body))

	// Example 5: Delete a position
	fmt.Println("=== Deleting a Position ===")
	req, _ = http.NewRequest("DELETE", baseURL+"/positions/pos123", nil)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error deleting position: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Delete Position Response: %s\n", string(body))
}
