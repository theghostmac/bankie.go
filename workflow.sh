#!/bin/bash

# Set the GitHub repository name and owner
REPO_OWNER="theghostmac"
REPO_NAME="bankie.go"

# Set the workflow file path and content
WORKFLOW_FILE=".github/workflows/build.yml"
WORKFLOW_CONTENT="
name: Build

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v2

      - name: Set up Go environment
        uses: actions/setup-go@v2
        with:
          go-version: 1.16

      - name: Build
        run: go build ./cmd/app

      - name: Run tests
        run: go test -v ./...

      - name: Create coverage report
        run: go test -coverprofile=coverage.out ./...

      - name: Upload coverage report
        uses: codecov/codecov-action@v2
        with:
          file: coverage.out
          flags: unittests
"

# Create the workflow file
echo "$WORKFLOW_CONTENT" > "$WORKFLOW_FILE"

say "GitHub workflow created successfully!"
