package sender

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Payload represents the structure of the data expected to be sent as a webhook
type Payload struct {
	Event   string
	Date    string
	Id      string
	Payment string
}

// SendWebhook sends a JSON POST request to the specified URL and updates the event status in the database
func SendWebhook(data interface{}, url string, webhookId string) error {
	// Marshal the data into JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Retrieve database connection information from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Prepare the webhook request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the webhook request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// Determine the status based on the response code
	status := "failed"
	if resp.StatusCode == http.StatusOK {
		status = "delivered"
	}

	// Update the status in the database
	if err := updateStatus(webhookId, status, dbUser, dbPassword, dbHost, dbPort, dbName); err != nil {
		return err
	}

	return nil
}

// updateStatus connects to the database and updates the status of the event with the given webhookId
func updateStatus(webhookId, status, dbUser, dbPassword, dbHost, dbPort, dbName string) error {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		}
	}(db)

	// Update the event status
	_, err = db.Exec("UPDATE events SET status=$1 WHERE id=$2", status, webhookId)
	if err != nil {
		return err
	}

	return nil
}
