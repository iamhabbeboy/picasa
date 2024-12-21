# Makefile
WORKER=scheduler/runner.go
WORKER_BIN=picasa_scheduler
WORKER_PATH=scheduler/picasa_scheduler

build-worker:
	@echo "Building worker..."
	go build -o $(WORKER_PATH) $(WORKER)

release-app:
	@echo "Building desktop app for release..."
	wails build

# Build both the worker and the desktop app
build: build-worker #build-app

# Run both the worker and the desktop app
run:
	@echo "Running desktop app with background worker..."
	# Run worker in the background
	./$(WORKER_PATH) & \
	# Run desktop app
	wails dev

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(WORKER_PATH)

lint:
	@echo "Linting..."
	golangci-lint run
