AWS_ACCOUNT_ID=1234567890
AWS_BUCKET_NAME=your-bucket-name-for-cloudformation-package-data
AWS_STACK_NAME=your-cloudformation-stack-name
AWS_REGION=us-west-1

clean:
		@rm -rf .build
		@mkdir -p .build

build: clean
		@GOOS=linux go build -o .build/main

test:
		go test

configure:
		@aws s3api create-bucket \
			--bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_REGION)

package:
		@aws cloudformation package \
			--template-file template.yml \
			--s3-bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--output-template-file .build/packaged.yml

deploy:
		@aws cloudformation deploy \
			--template-file .build/packaged.yml \
			--region $(AWS_REGION) \
			--capabilities CAPABILITY_IAM \
			--stack-name $(AWS_STACK_NAME)

describe:
		@aws cloudformation describe-stacks \
			--region $(AWS_REGION) \
			--stack-name $(AWS_STACK_NAME) \

outputs:
	 @make describe | jq -r '.Stacks[0].Outputs'