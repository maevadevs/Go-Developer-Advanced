# NOTE: Make sure all indentations use tabs

# DEFAULT_GOAL specify the default target
# This is run when no target is specified
.DEFAULT_GOAL := build

# Target definitions
# .PHONY helps avoid possible name-collisions with other directory or file names on the computer
.PHONY: fmt vet build

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
	# Task: Build in debug mode in Linux
	go build -o bin/Composite src/main.go

# Target
build-win: vet
	# Task: Build in debug mode in Windows
	go build -o bin\Composite.exe src\main.go

# Target
try: vet
	# Task: Build in debug mode, run, then remove built binary
	if test -f bin/Composite-Temp; then \
		rm -f bin/Composite-Temp; \
	fi
	go build -o bin/Composite-Temp src/main.go
	bin/Composite-Temp
	rm -f bin/Composite-Temp

# Target
try-win: vet
	# Task: Build in debug mode, run, then remove built binary in Windows
	if test -f bin\Composite-Temp; then \
		rm -f bin\Composite-Temp; \
	fi
	go build -o bin\Composite-Temp src\main.go
	bin\Composite-Temp
	rm -f bin\Composite-Temp

# Target
run:
	# Task: Run the application in Linux
	go run src/main.go

# Target
run-win:
	# Task: Run the application in Windows
	go run src\main.go
