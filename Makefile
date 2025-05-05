# Usage:
#   make run DIR=01-loop-reference-pitfall
#   make build DIR=01-loop-reference-pitfall
#   make clean

GO=go

run:
	@echo "👉 Running $(DIR)/main.go..."
	cd $(DIR) && $(GO) run main.go

build:
	@echo "🔧 Building binary from $(DIR)/main.go..."
	cd $(DIR) && $(GO) build -o ../bin/$(notdir $(DIR))

clean:
	@echo "🧹 Cleaning up built binaries..."
	rm -rf bin/
