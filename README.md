# Is Email Disposable

A Go package to detect disposable email addresses using an embedded, automatically-updated blocklist.

[![Go Reference](https://pkg.go.dev/badge/github.com/bobadilla-tech/is-email-disposable.svg)](https://pkg.go.dev/github.com/bobadilla-tech/is-email-disposable)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobadilla-tech/is-email-disposable)](https://goreportcard.com/report/github.com/bobadilla-tech/is-email-disposable)
[![License](https://img.shields.io/github/license/bobadilla-tech/is-email-disposable)](LICENSE)

## Features

- 🚀 **Zero dependencies** - Pure Go implementation
- 📦 **Embedded blocklist** - No external files needed at runtime
- ⚡ **Fast lookups** - O(1) hash map-based detection
- 🔄 **Auto-updates** - Monthly automated updates via GitHub Actions
- 🧪 **Well tested** - Comprehensive test coverage
- 💻 **Simple API** - Easy to integrate

## Installation

```bash
go get github.com/bobadilla-tech/is-email-disposable
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/bobadilla-tech/is-email-disposable"
)

func main() {
    // Check if an email is disposable
    if disposable.IsDisposable("user@tempmail.com") {
        fmt.Println("This is a disposable email address")
    }

    // Check if a domain is disposable
    if disposable.IsDisposableDomain("10minutemail.com") {
        fmt.Println("This domain is disposable")
    }

    // Get the total count of disposable domains
    count := disposable.Count()
    fmt.Printf("Tracking %d disposable email domains\n", count)
}
```

### API Reference

#### `IsDisposable(email string) bool`

Checks if the given email address uses a disposable domain.

```go
disposable.IsDisposable("user@tempmail.com") // returns true
disposable.IsDisposable("user@gmail.com")    // returns false
```

The check is case-insensitive:
```go
disposable.IsDisposable("USER@TEMPMAIL.COM") // returns true
```

#### `IsDisposableDomain(domain string) bool`

Checks if the given domain is in the disposable list.

```go
disposable.IsDisposableDomain("tempmail.com") // returns true
disposable.IsDisposableDomain("gmail.com")    // returns false
```

#### `GetAllDomains() []string`

Returns a slice containing all disposable email domains.

```go
domains := disposable.GetAllDomains()
fmt.Printf("Total disposable domains: %d\n", len(domains))
```

Note: This returns a large list (thousands of domains). Use `IsDisposable()` or `IsDisposableDomain()` for most use cases.

#### `Count() int`

Returns the total number of disposable email domains in the blocklist.

```go
count := disposable.Count()
fmt.Printf("Blocking %d disposable domains\n", count)
```

## How It Works

1. **Embedded Data**: The blocklist is embedded directly into the compiled binary using Go's `embed` directive
2. **Efficient Lookup**: Domains are stored in a hash map for O(1) lookup performance
3. **Auto-Updates**: A GitHub Action runs monthly to fetch the latest blocklist and create a PR
4. **Source**: The blocklist comes from [disposable-email-domains](https://github.com/disposable-email-domains/disposable-email-domains)

## Testing

Run the tests:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench=.
```

Example benchmark results:
```
BenchmarkIsDisposable-8              	10000000	       120 ns/op
BenchmarkIsDisposableDomain-8        	20000000	        85 ns/op
BenchmarkIsDisposableNotFound-8      	10000000	       115 ns/op
```

## Automated Updates

This package includes a GitHub Action that:
- Runs on the 1st of every month
- Downloads the latest blocklist from the upstream repository
- Runs tests to ensure compatibility
- Creates a pull request with the changes
- Can also be triggered manually via workflow dispatch

See [.github/workflows/update-blocklist.yml](.github/workflows/update-blocklist.yml) for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

- Blocklist maintained by [disposable-email-domains](https://github.com/disposable-email-domains/disposable-email-domains)
- Inspired by the need for simple, efficient disposable email detection in Go

## Related Projects

- [disposable-email-domains](https://github.com/disposable-email-domains/disposable-email-domains) - The upstream blocklist source
