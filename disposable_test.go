package disposable

import (
	"testing"
)

func TestIsDisposable(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "disposable email - 0-mail.com",
			email:    "test@0-mail.com",
			expected: true,
		},
		{
			name:     "disposable email - 10mail.org",
			email:    "user@10mail.org",
			expected: true,
		},
		{
			name:     "disposable email - uppercase",
			email:    "TEST@0-MAIL.COM",
			expected: true,
		},
		{
			name:     "disposable email - mixed case",
			email:    "Test@10Mail.Org",
			expected: true,
		},
		{
			name:     "legitimate email - gmail.com",
			email:    "user@gmail.com",
			expected: false,
		},
		{
			name:     "legitimate email - yahoo.com",
			email:    "test@yahoo.com",
			expected: false,
		},
		{
			name:     "legitimate email - company domain",
			email:    "employee@mycompany.com",
			expected: false,
		},
		{
			name:     "empty email",
			email:    "",
			expected: false,
		},
		{
			name:     "invalid email - no @",
			email:    "notanemail",
			expected: false,
		},
		{
			name:     "invalid email - multiple @",
			email:    "test@@example.com",
			expected: false,
		},
		{
			name:     "email with spaces",
			email:    "user@0-mail.com ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDisposable(tt.email)
			if result != tt.expected {
				t.Errorf("IsDisposable(%q) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsDisposableDomain(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{
			name:     "disposable domain - 0-mail.com",
			domain:   "0-mail.com",
			expected: true,
		},
		{
			name:     "disposable domain - uppercase",
			domain:   "0-MAIL.COM",
			expected: true,
		},
		{
			name:     "disposable domain - mixed case",
			domain:   "0-Mail.Com",
			expected: true,
		},
		{
			name:     "legitimate domain - gmail.com",
			domain:   "gmail.com",
			expected: false,
		},
		{
			name:     "legitimate domain - yahoo.com",
			domain:   "yahoo.com",
			expected: false,
		},
		{
			name:     "empty domain",
			domain:   "",
			expected: false,
		},
		{
			name:     "domain with spaces",
			domain:   " 0-mail.com ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsDisposableDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsDisposableDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestGetAllDomains(t *testing.T) {
	domains := GetAllDomains()

	if len(domains) == 0 {
		t.Error("GetAllDomains() returned empty slice, expected non-empty")
	}

	// Check that the returned slice contains some known disposable domains
	knownDomains := []string{"0-mail.com", "10mail.org", "027168.com"}
	found := make(map[string]bool)

	for _, domain := range domains {
		for _, known := range knownDomains {
			if domain == known {
				found[known] = true
			}
		}
	}

	for _, known := range knownDomains {
		if !found[known] {
			t.Errorf("GetAllDomains() missing expected domain %q", known)
		}
	}
}

func TestCount(t *testing.T) {
	count := Count()

	if count == 0 {
		t.Error("Count() returned 0, expected positive number")
	}

	// The count should match the number of domains returned by GetAllDomains
	domains := GetAllDomains()
	if count != len(domains) {
		t.Errorf("Count() = %d, but GetAllDomains() returned %d domains", count, len(domains))
	}
}

func TestBlocklistLoaded(t *testing.T) {
	// Verify that the blocklist was loaded correctly
	if len(disposableDomains) == 0 {
		t.Error("disposableDomains map is empty, blocklist was not loaded")
	}

	// Test that we can find domains that should be in the first 30 lines
	// based on the file format shown earlier
	firstLineDomains := []string{
		"0-mail.com",
		"027168.com",
		"062e.com",
		"0815.ru",
		"10mail.org",
	}

	for _, domain := range firstLineDomains {
		if !IsDisposableDomain(domain) {
			t.Errorf("Expected domain %q to be in disposable list", domain)
		}
	}
}

// Benchmark tests
func BenchmarkIsDisposable(b *testing.B) {
	email := "test@0-mail.com"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsDisposable(email)
	}
}

func BenchmarkIsDisposableDomain(b *testing.B) {
	domain := "0-mail.com"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsDisposableDomain(domain)
	}
}

func BenchmarkIsDisposableNotFound(b *testing.B) {
	email := "test@gmail.com"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsDisposable(email)
	}
}
