# NOTE: Make sure all indentations use tabs

# DEFAULT_GOAL specifies the default target
# This is run when no target is provided during the call
.DEFAULT_GOAL := try

# Target definitions
# .PHONY helps avoid possible name-collisions with other directory or file names on the computer
.PHONY: fmt vet build run try

# Target
fmt:
	# Task: Format all source files
	cd src
	go fmt ./...
	cd ..

# Target
vet: fmt
	# Task: Verify any possible errors
	cd src
	go vet ./...
	cd ..

# Target
build: vet
	# Task: Build module
	go build -o bin/Types src/main.go

# Target
run: build
	# Task: Build module then run
	bin/Types

# Target
try: vet
	# Task: Build module, run, then remove built binary
	if test -f bin/Types-Temp; then \
		rm -f bin/Types-Temp; \
	fi
	go build -o bin/Types-Temp src/main.go
	bin/Types-Temp
	rm -f bin/Types-Temp
