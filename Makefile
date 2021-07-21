.PHONY: build

build:
	sam build

deploy:
	sam deploy

plan:
	sam deploy --no-execute-changeset
