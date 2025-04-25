.PHONY: run 

# Run the application
run: ## Run the application locally
	@echo "Running application..."
	@go run $(MAIN_PATH)
