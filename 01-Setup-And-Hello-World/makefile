# NOTE: Make sure all indentations use actual tabs

# DEFAULT_GOAL specifies the default target
# This is run when no target is provided during the call
.DEFAULT_GOAL := try

# Target definitions
# .PHONY helps avoid possible name-collisions with other directory or file names on the computer
.PHONY: fmt vet build bld build-release bld-rls run run-release run-rls try-build-run tbr try-run try

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
	go build -o bin/debug/Hello-World src/main.go

# Target: Alias to build
bld: vet
	# Task: Build module
	go build -o bin/debug/Hello-World src/main.go

# Target
build-release: vet
	# Task: Build module for release, strip symbols
	go build -ldflags="-s -w" -o bin/release/Hello-World src/main.go

# Target: ALias to build-release
bld-rls: vet
	# Task: Build module for release, strip symbols
	go build -ldflags="-s -w" -o bin/release/Hello-World src/main.go

# Target
run: build
	# Task: Build module then run
	bin/debug/Hello-World

# Target
run-release: build-release
	# Task: Build module for release then run
	bin/release/Hello-World

# Target: ALias to run-release
run-rls: build-release
	# Task: Build module for release then run
	bin/release/Hello-World

# Target
try-build-run: vet
	# Task: Build module, run, then remove built binary
	if test -f bin/debug/Hello-World-Temp; then \
		rm -f bin/debug/Hello-World-Temp; \
	fi
	go build -o bin/debug/Hello-World-Temp src/main.go
	bin/debug/Hello-World-Temp
	rm -f bin/debug/Hello-World-Temp

# Target: Alias to try-build-run
tbr: vet
	# Task: Build module, run, then remove built binary
	if test -f bin/debug/Hello-World-Temp; then \
		rm -f bin/debug/Hello-World-Temp; \
	fi
	go build -o bin/debug/Hello-World-Temp src/main.go
	bin/debug/Hello-World-Temp
	rm -f bin/debug/Hello-World-Temp

# Target
try-run: vet
	# Task: Test-Run the module without building anything
	go run src/main.go

# Target: Alias to try-run
try: vet
	# Task: Test-Run the module without building anything
	go run src/main.go
