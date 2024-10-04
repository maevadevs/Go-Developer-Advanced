# NOTE: Make sure all indentations use tabs

# DEFAULT_GOAL specifies the default target
# This is run when no target is provided during the call
.DEFAULT_GOAL := try

# Target definitions
# .PHONY helps avoid possible name-collisions with other directory or file names on the computer
.PHONY: fmt vet build build-release run run-release try

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
	go build -o bin/debug/Pointers src/main.go

# Target
build-release: vet
	# Task: Build module for release, strip symbols
	go build -ldflags="-s -w" -o bin/release/Pointers src/main.go

# Target
run: build
	# Task: Build module then run
	bin/debug/Pointers

# Target
run-release: build-release
	# Task: Build module for release then run
	bin/release/Pointers

# Target
try: vet
	# Task: Build module, run, then remove built binary
	if test -f bin/debug/Pointers-Temp; then \
		rm -f bin/debug/Pointers-Temp; \
	fi
	go build -o bin/debug/Pointers-Temp src/main.go
	bin/debug/Pointers-Temp
	rm -f bin/debug/Pointers-Temp
