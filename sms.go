package alert

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	sender   string
	username string
	password string
)

// Configure initializes the SMS service credentials
func Configure(apiSender, apiUsername, apiPassword string) {
	sender = apiSender
	username = apiUsername
	password = apiPassword
}

// PhoneFormat validates and formats a phone number to start with '0' and ensures it is not over 10 digits
func PhoneFormat(phone string) (string, error) {
	// Remove whitespace and ensure only digits
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, " ", "")

	// Check if the phone number starts with '0' and has 10 digits
	if len(phone) != 10 || !strings.HasPrefix(phone, "0") {
		return "", fmt.Errorf("invalid phone number: %s, phone must be start with 0 and have 10 digits", phone)
	}
	return phone, nil
}

// sendSMS sends an SMS message to a single phone number using the Arc Innovative API
func sendSMS(phone, message string) error {
	formattedPhone, err := PhoneFormat(phone)
	if err != nil {
		return fmt.Errorf("invalid phone number format: %v", err)
	}

	baseURL := "https://v2.arcinnovative.com/APIConnect.php"
	params := url.Values{}
	params.Add("sender", sender)
	params.Add("username", username)
	params.Add("password", password)
	params.Add("msisdn", formattedPhone)
	params.Add("msg", message)

	// Build the complete URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the HTTP GET request
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send SMS, status code: %d", resp.StatusCode)
	}

	log.Printf("SMS sent successfully to %s", phone)
	return nil
}

// AlertError sends an alert message to a list of phone numbers when a service encounters an error
func AlertError(phoneList, message, service string) {
	alertMessage := fmt.Sprintf("PK Alert for %s: %s", service, message)
	phones := strings.Split(phoneList, ",")

	for _, phone := range phones {
		phone = strings.TrimSpace(phone)
		if err := sendSMS(phone, alertMessage); err != nil {
			log.Printf("Failed to send alert to %s: %v", phone, err)
		}
	}
}
