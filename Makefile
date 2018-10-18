include .env

clean:
		@rm -rf dist
		@mkdir -p dist

build: clean
		@for dir in handler/*; do \
			GOOS=linux go build -x -v -o dist/$$dir ./$$dir; \
		done

run:
		sam local start-api

install:
		go get -u ./...

test:
		go test ./... --cover

configure:
		aws s3api create-bucket \
			--bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--create-bucket-configuration LocationConstraint=$(AWS_REGION)

package: build
		@aws cloudformation package \
			--template-file template.yml \
			--s3-bucket $(AWS_BUCKET_NAME) \
			--region $(AWS_REGION) \
			--output-template-file package.yml

deploy:
		@aws cloudformation deploy \
			--template-file package.yml \
			--region $(AWS_REGION) \
			--capabilities CAPABILITY_IAM \
			--stack-name $(AWS_STACK_NAME)

describe:
		@aws cloudformation describe-stacks \
			--region $(AWS_REGION) \
			--stack-name $(AWS_STACK_NAME) \

outputs:
		@make describe | jq -r '.Stacks[0].Outputs'

url:
		@make describe | jq -r ".Stacks[0].Outputs[0].OutputValue" -j
