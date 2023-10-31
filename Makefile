.PHONY: build clean development and deploy

run-tools:
	@echo "Running tools..."
	docker compose -f infrastructure-devops/docker-compose.yml up -d

clean-tools:
	@echo "Cleaning tools..."
	docker compose -f infrastructure-devops/docker-compose.yml down --rmi all

lint:
	@echo "Running lint..."
	golangci-lint run ./internal/...

test:
	@echo "Running tests..."
	go test ./internal/... -v -cover

scan:
	@echo "Running scann..."
	gosec ./internal/...

