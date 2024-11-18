# ALERT ERROR SERVICE

This library for pompkins system is used to notify the system of any errors by sending SMS to the specified phone number along with the message. The message must not exceed 60 characters.

### Prerequisites

Requirements for the golang and echo framework

### Installing

Get package form remote repository

    go get https://github.com/Pompkins-org/alert-service

### Usage

You must call this function with this command

    package main
    
    import (
    	"myproject/alert" // Replace 'myproject' with your module path
    )
    
    func main() {
    	// Configure the SMS service
    	alert.Configure("sender", "username", "password")
    
    	// Define the phone list, message, and service name
    	phoneList := "0812345678,0912345678" // List of phone numbers separated by commas
    	message := "An error occurred while processing the request."
    	service := "Order Service"
    
    	// Send alerts
    	alert.AlertError(phoneList, message, service)
    }
