// Package disposable provides utilities to detect disposable email addresses.
//
// The package uses an embedded blocklist of known disposable email domains
// and provides efficient O(1) lookup through a hash map.
//
// Example usage:
//
//	import "github.com/bobadilla-tech/is-email-disposable"
//
//	if disposable.IsDisposable("user@tempmail.com") {
//	    fmt.Println("This is a disposable email address")
//	}
//
//	if disposable.IsDisposableDomain("tempmail.com") {
//	    fmt.Println("This domain is in the disposable list")
//	}
package disposable

import (
	_ "embed"
	"strings"
	"sync"
)

//go:embed disposable_email_blocklist.conf
var blocklistData string

var (
	// disposableDomains stores all disposable email domains for O(1) lookup
	disposableDomains map[string]struct{}
	// initOnce ensures the blocklist is loaded only once
	initOnce sync.Once
)

// init loads the embedded blocklist into memory on package initialization
func init() {
	loadBlocklist()
}

// loadBlocklist parses the embedded blocklist and populates the domain map
func loadBlocklist() {
	initOnce.Do(func() {
		disposableDomains = make(map[string]struct{})

		lines := strings.Split(blocklistData, "\n")
		for _, line := range lines {
			domain := strings.TrimSpace(line)
			// Skip empty lines
			if domain == "" {
				continue
			}
			// Store domain in lowercase for case-insensitive matching
			disposableDomains[strings.ToLower(domain)] = struct{}{}
		}
	})
}

// IsDisposable checks if the given email address uses a disposable domain.
// The email parameter should be a complete email address (e.g., "user@example.com").
// Returns true if the domain part of the email is in the disposable list.
//
// The check is case-insensitive. For example:
//
//	IsDisposable("user@TempMail.com") == IsDisposable("user@tempmail.com")
func IsDisposable(email string) bool {
	if email == "" {
		return false
	}

	// Extract domain from email
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := strings.TrimSpace(parts[1])
	return IsDisposableDomain(domain)
}

// IsDisposableDomain checks if the given domain is in the disposable list.
// The domain parameter should be just the domain part (e.g., "example.com").
// Returns true if the domain is in the disposable list.
//
// The check is case-insensitive. For example:
//
//	IsDisposableDomain("TempMail.com") == IsDisposableDomain("tempmail.com")
func IsDisposableDomain(domain string) bool {
	if domain == "" {
		return false
	}

	domain = strings.ToLower(strings.TrimSpace(domain))
	_, exists := disposableDomains[domain]
	return exists
}

// GetAllDomains returns a slice containing all disposable email domains.
// The returned slice is a copy, so modifications won't affect the internal list.
//
// Note: This can return a large list (potentially thousands of domains).
// Consider using IsDisposable() or IsDisposableDomain() for most use cases.
func GetAllDomains() []string {
	domains := make([]string, 0, len(disposableDomains))
	for domain := range disposableDomains {
		domains = append(domains, domain)
	}
	return domains
}

// Count returns the total number of disposable email domains in the blocklist.
func Count() int {
	return len(disposableDomains)
}
