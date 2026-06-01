package main

import (
	"fmt"

	disposable "github.com/bobadilla-tech/is-email-disposable"
)

func main() {
	fmt.Println("=== Disposable Email Checker Demo ===")

	// Test various emails
	emails := []string{
		"user@gmail.com",
		"test@tempmail.com",
		"admin@0-mail.com",
		"contact@yahoo.com",
		"fake@10minutemail.com",
		"real@mycompany.com",
	}

	fmt.Println("Checking emails:")
	
	for _, email := range emails {
		isDisp := disposable.IsDisposable(email)
		status := "✅ legitimate"
	
		if isDisp {
			status = "❌ disposable"
		}
	
		fmt.Printf("  %s → %s\n", email, status)
	}

	// Check just domains
	fmt.Println("\nChecking domains:")
	
	domains := []string{
		"gmail.com",
		"tempmail.com",
		"guerrillamail.com",
		"microsoft.com",
	}

	for _, domain := range domains {
		isDisp := disposable.IsDisposableDomain(domain)
	
		status := "✅ legitimate"
	
		if isDisp {
			status = "❌ disposable"
		}
	
		fmt.Printf("  %s → %s\n", domain, status)
	}

	// Show statistics
	fmt.Printf("\n📊 Statistics:\n")
	fmt.Printf("  Total disposable domains tracked: %d\n", disposable.Count())

	// Case sensitivity demo
	fmt.Println("\n🔤 Case sensitivity test:")
	testCases := []string{
		"USER@TEMPMAIL.COM",
		"User@TempMail.Com",
		"user@tempmail.com",
	}
	
	for _, email := range testCases {
		fmt.Printf("  %s → disposable: %v\n", email, disposable.IsDisposable(email))
	}
}
