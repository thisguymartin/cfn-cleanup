# cfn-cleanup-go

A lightweight CLI tool to automatically clean up AWS CloudFormation stacks matching a specified prefix. Perfect for cleaning up development or testing stacks to avoid unnecessary AWS costs.

<!-- ## Features -->

## Features

- List CloudFormation stacks filtered by prefix
- Batch delete stacks matching a prefix
- Configurable AWS region support
- Interactive deletion confirmation
- Detailed logging of operations
- Timeout handling for long-running deletions

## Installation

### Using Go

```bash
go install github.com/thisguymartin/cfn-cleanup-go@latest
```

### From Binary Releases

Download the latest binary from the [releases page](https://github.com/yourusername/cfn-cleanup-go/releases).

### From Source

```bash
# Clone the repository
git clone https://github.com/thisguymartin/cfn-cleanup-go.git

# Change into directory
cd cfn-cleanup-go

# Build
go build -o cfn-cleanup

# Optional: Move to a directory in your PATH
sudo mv cfn-cleanup /usr/local/bin/
```

## Usage

```bash
# Basic usage with default prefix
cfn-cleanup

# Specify a custom prefix
cfn-cleanup -prefix="your-stack-prefix"

# Specify AWS region
cfn-cleanup -prefix="your-stack-prefix" -region="us-west-1"

# Delete Cloudformations based on region and prefix
cfn-cleanup -prefix="your-stack-prefix" -region="us-west-1" -delete=true

```

### Command Line Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-prefix` | Stack name prefix to match | "thisguymartin-pit" |
| `-region` | AWS region | AWS_REGION env variable |

## AWS Credentials

The tool uses the standard AWS SDK credential chain. You can provide credentials in several ways:

1. Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
2. AWS credentials file (`~/.aws/credentials`)


<!-- ## Logging

The tool logs all operations to:
- Standard output
- A local file (`cf_cleanup.log`) -->


