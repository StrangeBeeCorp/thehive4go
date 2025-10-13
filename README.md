<div align="center">
  <p>
   <a href="#" target="_blank"><img src="https://strangebee.com/wp-content/uploads/2024/07/Icon4Nav_TheHive.png" alt="TheHive Logo" width="100"></a>
  </p>
	<h1>TheHive4Go</h1>
	<p>
	 <em>Auto-generated Go API client for <a href="https://strangebee.com/thehive/">TheHive 5</a></em>
	</p>
		<!-- Badges are commented out while repository is private; they will work once public -->
		<!--
		<p>
			<a href="https://github.com/StrangeBee/TheHive4Go/releases"><img src="https://img.shields.io/github/v/tag/StrangeBee/TheHive4Go?sort=semver" alt="Latest Tag"></a>
			<a href="https://pkg.go.dev/github.com/StrangeBee/TheHive4Go"><img src="https://pkg.go.dev/badge/github.com/StrangeBee/TheHive4Go.svg" alt="Go Reference"></a>
		</p>
		-->
</div>

---

## Introduction

TheHive4Go is an auto-generated Go client library for interacting with TheHive 5.x API. This project leverages OpenAPI code generation to provide comprehensive, type-safe access to all TheHive API endpoints.

Unlike manually maintained clients, TheHive4Go is automatically generated from TheHive's official OpenAPI specification, ensuring it stays up-to-date with the latest API changes and includes all available endpoints and models.

> **Note**: This is a private repository not affiliated with TheHive Project.

## Features

- **Complete API Coverage**: Auto-generated from TheHive's OpenAPI spec, covering all endpoints
- **Type Safety**: Strongly typed Go structs for all API models and requests
- **Authentication Support**: API key, basic auth, and session-based authentication
- **Organization Context**: Support for multi-organization deployments via `X-Organisation` header
- **Comprehensive Documentation**: Auto-generated documentation for all APIs and models
- **Automated Updates**: Easy regeneration when TheHive API changes

## Installation

**Requirements:** Go 1.18 or later

Add the module to your project:

```bash
go get github.com/StrangeBee/TheHive4Go
```

## Quick Start

### Basic Client Setup

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/StrangeBee/TheHive4Go/thehive"
)

func main() {
    // Create configuration
    config := thehive.NewConfiguration()
    config.Servers = thehive.ServerConfigurations{
        {
            URL:         "http://localhost:9000",
            Description: "Local TheHive instance",
        },
    }
    
    // Create client
    client := thehive.NewAPIClient(config)
    
    // Set up API key authentication
    ctx := context.WithValue(context.Background(), thehive.ContextAccessToken, "your-api-key-here")
    
    // Now you can use the client...
}
```

### Authentication Methods

#### API Key Authentication (Recommended)
```go
ctx := context.WithValue(context.Background(), thehive.ContextAccessToken, "your-api-key")
```

#### Basic Authentication
```go
auth := thehive.BasicAuth{
    UserName: "username",
    Password: "password",
}
ctx := context.WithValue(context.Background(), thehive.ContextBasicAuth, auth)
```

#### Session Authentication
```go
ctx := context.WithValue(context.Background(), thehive.ContextAPIKeys, map[string]thehive.APIKey{
    "Session": {Key: "session-token"},
})
```

### Organization Context

For multi-organization deployments, specify the target organization:

```go
// Add organization header to your requests
ctx = context.WithValue(ctx, thehive.ContextAPIKeys, map[string]thehive.APIKey{
    "Header": {Key: "target-org-name"},
})
```

## Usage Examples

### Querying Alerts

```go
// List all alerts
genericOp := &thehive.InputQueryGenericOperation{
    Name: "listAlert",
}

query := thehive.InputQuery{
    Query: []thehive.InputQueryNamedOperation{
        thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp),
    },
}

result, resp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(query).Execute()
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Found alerts: %+v\n", result)
```

### Filtering and Pagination

```go
// Search for high severity alerts with pagination
genericOp := &thehive.InputQueryGenericOperation{
    Name: "listAlert",
}

filterOp := &thehive.InputQueryFilterOperation{
    Eq: map[string]interface{}{
        "severity": float32(3), // High severity
    },
    Name: "filter",
}

pagingOp := &thehive.InputQueryPagingOperation{
    From: int32(0),
    To:   int32(10),
    Name: "page",
}

query := thehive.InputQuery{
    Query: []thehive.InputQueryNamedOperation{
        thehive.InputQueryGenericOperationAsInputQueryNamedOperation(genericOp),
        thehive.InputQueryFilterOperationAsInputQueryNamedOperation(filterOp),
        thehive.InputQueryPagingOperationAsInputQueryNamedOperation(pagingOp),
    },
}

result, resp, err := client.QueryAndExportAPI.QueryAPI(ctx).InputQuery(query).Execute()
```

### Working with Cases

```go
// Create a new case
inputCase := &thehive.InputCreateCase{
    Title:       "Security Incident",
    Description: "Detected suspicious activity",
    Severity:    thehive.PtrInt32(2),
    Tlp:         thehive.PtrInt32(2),
}

createdCase, resp, err := client.CaseAPI.CreateCase(ctx).InputCreateCase(*inputCase).Execute()
if err != nil {
    fmt.Printf("Error creating case: %v\n", err)
    return
}

fmt.Printf("Created case: %+v\n", createdCase)
```

## Available APIs

The generated client includes comprehensive support for all TheHive APIs:

- **AlertAPI**: Alert management and operations
- **CaseAPI**: Case lifecycle management  
- **ObservableAPI**: Observable handling and analysis
- **TaskAPI**: Task management within cases
- **CommentAPI**: Comments on alerts and cases
- **UserAPI**: User management and authentication
- **OrganisationAPI**: Multi-organization support
- **DashboardAPI**: Dashboard and reporting
- **QueryAndExportAPI**: Advanced querying and data export
- **AdminAPI**: System administration
- And many more...

For complete API documentation, see the [generated docs](thehive/docs/) directory.

## Code Generation

This client is automatically generated from TheHive's OpenAPI specification using a fully containerized build process with Docker.

### Prerequisites

- Docker (for code generation)
- Go 1.18+ (for using the generated client)

### Generation Process

The build process is completely containerized and consists of:

1. **Download**: Fetches the latest OpenAPI spec from TheHive's documentation
2. **Preprocessing**: Applies necessary fixes for Go code generation compatibility  
3. **Generation**: Uses OpenAPI Generator in Docker to create Go client code
4. **Postprocessing**: Applies Go-specific optimizations and fixes

### Regenerating the Client

To update the client with the latest API changes:

```bash
make generate
```

This single command will:
- Build the Docker image with all required tools
- Download the latest OpenAPI specification
- Apply preprocessing fixes for Go compatibility
- Generate fresh client code using OpenAPI Generator
- Apply postprocessing optimizations for Go best practices

All generation happens inside Docker containers, ensuring consistent results across different development environments.

### Build Scripts

The generation process uses these containerized scripts:

- `scripts/download_openapi.sh`: Downloads the OpenAPI spec
- `scripts/preprocess_openapi.sh`: Fixes OpenAPI spec issues for Go generation
- `scripts/generate_client.sh`: Generates the Go client using OpenAPI Generator
- `scripts/postprocess_client.sh`: Applies Go-specific fixes and optimizations

## Development

### Testing

Run the test suite (containerized):

```bash
make test
```

### Code Quality

All quality checks run in Docker containers for consistency:

```bash
make fmt           # Format code using Go in Docker
make security      # Run security scans in Docker
make vulncheck     # Check for vulnerabilities using Go in Docker
make vetlint       # Run linter in Docker
make check-generate # Verify generated code is up-to-date
```

### Dependencies

Update dependencies:

```bash
make updatedep     # Update Go dependencies in Docker
```

### Available Make Targets

Run `make help` to see all available targets with descriptions.

## Project Structure

```
├── README.md                 # This file
├── Makefile                  # Build automation
├── go.mod                    # Go module definition
├── .github/
│   └── workflows/
│       ├── ci.yml           # Continuous integration
│       └── release.yml      # Release automation
├── scripts/                 # Code generation scripts
│   ├── download_openapi.sh
│   ├── preprocess_openapi.sh
│   ├── generate_client.sh
│   └── postprocess_client.sh
├── thehive/                 # Generated client code
│   ├── docs/               # API documentation
│   ├── *.go               # Generated Go client files
│   └── model_*.go         # Generated model definitions
└── tmp/                    # Temporary files and examples
    ├── examples.go         # Usage examples
    └── *.yaml             # OpenAPI specifications
```

## Versioning Policy

This project follows Semantic Versioning:

- **v0.x.y**: API is not yet stable; breaking changes may appear in minor releases
- **v1.0.0+**: Backwards compatibility preserved; breaking changes require major version bump

## Contributing

Contributions are welcome! Please:

1. Open an issue to discuss proposed changes
2. Create a feature branch with a descriptive name
3. Implement changes with appropriate tests
4. Ensure all quality checks pass (`make security test`)
5. Submit a pull request

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## Acknowledgments

- **TheHive Project**: For the excellent security incident response platform
- **StrangeBee**: For TheHive development and maintenance
- **OpenAPI Generator**: For the code generation framework

---

<div align="center">
  <p>Generated with ❤️ for the security community</p>
</div>