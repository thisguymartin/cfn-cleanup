# cfn-cleanup-go

A lightweight CLI tool to automatically clean up AWS CloudFormation stacks matching a specified prefix. Perfect for cleaning up development or testing stacks to avoid unnecessary AWS costs.

## Features

- 🚀 Single binary deployment
- 🔍 Prefix-based stack matching
- 📝 Comprehensive logging
- ⏱️ Waits for stack deletion completion
- 🔄 Handles pagination for large numbers of stacks
- 🛡️ Safe error handling
- 🌎 Region configurable

## Installation

### Using Go

```bash
go install github.com/yourusername/cfn-cleanup-go@latest
```

### From Binary Releases

Download the latest binary from the [releases page](https://github.com/yourusername/cfn-cleanup-go/releases).

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/cfn-cleanup-go.git

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
cfn-cleanup -prefix="your-stack-prefix" -region="us-west-2"
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
3. IAM role when running on AWS services (EC2, ECS, etc.)

Required IAM permissions:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudformation:ListStacks",
                "cloudformation:DeleteStack",
                "cloudformation:DescribeStacks"
            ],
            "Resource": "*"
        }
    ]
}
```

## Logging

The tool logs all operations to:
- Standard output
- A local file (`cf_cleanup.log`)


### Prerequisites

- Go 1.19 or higher
- AWS account and credentials
- Make (optional, for using Makefile)

### Running Tests

```bash
go test ./...
```
