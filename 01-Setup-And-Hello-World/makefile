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
	# Task: Build module in Linux
	go build -o bin/Hello-World src/main.go

# Target
build-win: vet
	# Task: Build module in Windows
	go build -o bin\Hello-World.exe src\main.go

# Target
run: build
	# Task: Build module then run in Linux
	bin/Hello-World

# Target
run-win: build-win
	# Task: Build module then run in Windows
	bin\Hello-World.exe

# Target
try: vet
	# Task: Build module, run, then remove built binary in Linux
	if test -f bin/Hello-World-Temp; then \
		rm -f bin/Hello-World-Temp; \
	fi
	go build -o bin/Hello-World-Temp src/main.go
	bin/Hello-World-Temp
	rm -f bin/Hello-World-Temp

# Target
try-win: vet
	# Task: Build module, run, then remove built binary in Windows
	if test -f bin\Hello-World-Temp; then \
		rm -f bin\Hello-World-Temp; \
	fi
	go build -o bin\Hello-World-Temp src\main.go
	bin\Hello-World-Temp
	rm -f bin\Hello-World-Temp
