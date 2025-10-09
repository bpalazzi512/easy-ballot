package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Example API calls for the nominations endpoint
func main() {
	baseURL := "http://localhost:8080"

	// Example 1: Create a nomination
	fmt.Println("=== Creating a Nomination ===")
	createNominationData := map[string]string{
		"position_id":  "pos123",
		"nominee_id":   "user456",
		"nominator_id": "user789",
	}
	createNominationJSON, _ := json.Marshal(createNominationData)

	resp, err := http.Post(baseURL+"/nominations", "application/json", bytes.NewBuffer(createNominationJSON))
	if err != nil {
		fmt.Printf("Error creating nomination: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Create Nomination Response: %s\n\n", string(body))

	// Example 2: Get all nominations
	fmt.Println("=== Getting All Nominations ===")
	resp, err = http.Get(baseURL + "/nominations")
	if err != nil {
		fmt.Printf("Error getting nominations: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Nominations Response: %s\n\n", string(body))

	// Example 3: Get nominations by position
	fmt.Println("=== Getting Nominations by Position ===")
	resp, err = http.Get(baseURL + "/positions/pos123/nominations")
	if err != nil {
		fmt.Printf("Error getting nominations by position: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Nominations by Position Response: %s\n\n", string(body))

	// Example 4: Get nominations by nominee
	fmt.Println("=== Getting Nominations by Nominee ===")
	resp, err = http.Get(baseURL + "/users/user456/nominations")
	if err != nil {
		fmt.Printf("Error getting nominations by nominee: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Nominations by Nominee Response: %s\n\n", string(body))

	// Example 5: Get nominations by nominator
	fmt.Println("=== Getting Nominations by Nominator ===")
	resp, err = http.Get(baseURL + "/users/user789/nominations-made")
	if err != nil {
		fmt.Printf("Error getting nominations by nominator: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Get Nominations by Nominator Response: %s\n\n", string(body))

	// Example 6: Update nomination status (assuming nomination ID is "nom123")
	fmt.Println("=== Updating Nomination Status ===")
	updateStatusData := map[string]string{
		"status": "accepted",
	}
	updateStatusJSON, _ := json.Marshal(updateStatusData)

	req, _ := http.NewRequest("PATCH", baseURL+"/nominations/nom123/status", bytes.NewBuffer(updateStatusJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error updating nomination status: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Update Nomination Status Response: %s\n\n", string(body))

	// Example 7: Update a nomination (assuming nomination ID is "nom123")
	fmt.Println("=== Updating a Nomination ===")
	updateNominationData := map[string]interface{}{
		"position_id":  "pos123",
		"nominee_id":   "user456",
		"nominator_id": "user789",
		"status":       "declined",
	}
	updateNominationJSON, _ := json.Marshal(updateNominationData)

	req, _ = http.NewRequest("PUT", baseURL+"/nominations/nom123", bytes.NewBuffer(updateNominationJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error updating nomination: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Update Nomination Response: %s\n\n", string(body))

	// Example 8: Delete a nomination
	fmt.Println("=== Deleting a Nomination ===")
	req, _ = http.NewRequest("DELETE", baseURL+"/nominations/nom123", nil)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error deleting nomination: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Delete Nomination Response: %s\n", string(body))
}
