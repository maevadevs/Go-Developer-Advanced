#!/bin/bash

######################################################################################
# SCRIPT:       Allow to quickly generate a new Go project as a Go module            #
# AUTHOR:       github.com/maevadevs                                                 #
# CALL SIGN:    bash init.sh -p <Project_Dir> -o <Out_File> -m <Module_Path_Dir>     #
# CALL EXAMPLE: bash init.sh -p hello_world -o hello -m example.com/proj/hello_world #
######################################################################################

# Get Current Directory
CURRENT_DIR=$PWD;

# Get the project name and module root arguments from the user
while getopts p:o:m: flag
do
    case "${flag}" in
        p) PROJECT_NAME=${OPTARG};;
        o) OUT_FILE=${OPTARG};;
        m) MODULE_ROOT_PATH=${OPTARG};;
    esac
done

# The new project's directory
PROJECT_DIR="$CURRENT_DIR/$PROJECT_NAME"

echo "";
echo "Project Name: $PROJECT_NAME";
echo "Module Root Path: $MODULE_ROOT_PATH";

echo "";
echo "Creating a new Go project: $PROJECT_DIR ..."

# Create directory for the project if it does not exist yet
if test -d $PROJECT_DIR; then
    echo "The project directory already exist. Skipping creating a new directory.";
    echo "";
else
    echo "Creating a new project directory...";
    echo "$PROJECT_DIR";
    mkdir -p "$PROJECT_DIR";
    echo "Done."
    echo "";
fi

# Move into the project directory
cd $PROJECT_DIR;

# Create a new Go module if it does not exist yet
if test -f "$PROJECT_DIR/go.mod"; then
    echo "The go.mod file already exist. Skipping creating a new go module.";
    echo "";
else
    echo "Creating a new Go module...";
    go mod init "$MODULE_ROOT_PATH/$PROJECT_NAME";
    echo "Done."
    echo "";
fi

# Create a readme.md with placeholder
if test -f "$PROJECT_DIR/readme.md"; then
    echo "The readme.md file already exist. Skipping creating a new readme.md.";
    echo "";
else
    echo "Creating a new readme.md file...";
    touch "readme.md";
    # Add default placeholder contents
    echo "\
# $PROJECT_NAME

---
" >> readme.md;

    echo "Done."
    echo "";
fi

# Create a makefile for the build process
if test -f "$PROJECT_DIR/makefile"; then
    echo "The makefile already exist. Skipping creating a new makefile.";
    echo "";
else
    echo "Creating a new makefile...";
    touch "makefile";
    # Add default placeholder contents
    echo "\
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
	go build -o bin/$OUT_FILE src/main.go

# Target
run: build
	# Task: Build module then run
	bin/$OUT_FILE \$(ARGS)

# Target
try: vet
	# Task: Build module, run, then remove built binary
	if test -f bin/$OUT_FILE-Temp; then \\
		rm -f bin/$OUT_FILE-Temp; \\
	fi
	go build -o bin/$OUT_FILE-Temp src/main.go
	bin/$OUT_FILE-Temp \$(ARGS)
	rm -f bin/$OUT_FILE-Temp" >> makefile;

    echo "Done."
    echo "";
fi

# Create a new "src" folder if does not exist yet
if test -d "$PROJECT_DIR/src"; then
    echo "The src directory exist. Skipping creating a new src directory.";
    echo "";
else
    echo "Creating a new src directory...";
    mkdir -p "src";
    echo "Done."
    echo "";
fi

# Create a new "main.go" in "src" if does not exist yet
if test -f "$PROJECT_DIR/src/main.go"; then
    echo "The main.go file exist. Skipping creating a main.go file.";
    echo "";
else
    echo "Creating a new main.go file...";
    touch "src/main.go";
    # Add default placeholder contents
    echo "\
package main

import \"fmt\"

// This is the main entry of the application.
func main() {
    fmt.Println(\"Hello world!\")
}

// AVAILABLE COMMANDS
// ------------------
//  make            Default to \`make try\`
//  make fmt        Format all source files
//  make vet        Verify any possible errors
//  make build      Build module
//  make run        Build module then run
//  make try        Build module, run, then remove built binary

" >> src/main.go;

    echo "Done."
    echo "";
fi


echo "All processes complete. Exiting."
exit
