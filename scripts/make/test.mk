.PHONY: test 

# Run all tests
test: ## Run all tests
	@echo "Running all tests..."
	@TEST_USE_MOCKS=true go test $(TEST_FLAGS) ./...
