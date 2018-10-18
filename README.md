# Lambda Go Example

## Dependecies

(tip: if you append to the path `...` when you `go get` this package you will already have all dependencies)

```bash
$ > go get -u ./...
```

## Build

```bash
# Build binary
$ > make build

# Test Go Code
$ > make test
```

## Deploy

### Create .env

```bash
AWS_ACCOUNT_ID=1234567890
AWS_BUCKET_NAME=your-bucket-name-for-cloudformation-package-data
AWS_STACK_NAME=your-cloudformation-stack-name
AWS_REGION=us-west-1
```

### Install AWS CLI

```bash
$ > brew install awscli
```

### Command

```bash
# Create S3 Bucket
$ > make configure

# Upload data to S3 Bucket
$ > make package

# Deploy CloudFormation Stack
$ > make deploy
```

## Usage

```bash
$ > make outputs

[
  {
    "OutputKey": "URL",
    "OutputValue": "https://random-id.execute-api.us-west-1.amazonaws.com/Prod",
    "Description": "URL for HTTPS Endpoint"
  }
]

$ > curl https://random-id.execute-api.us-west-1.amazonaws.com/Stage/people

{"data":[{"id":"d1","name":"Anton","age":31},{"id":"c2","name":"Frank","age":28},{"id":"b1","name":"Horst","age":42}]}

$ > curl https://random-id.execute-api.us-west-1.amazonaws.com/Stage/person/b1

{"data":{"id":"b1","name":"Horst","age":42}}
```
