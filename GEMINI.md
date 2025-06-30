# Gemini Instructions

You are a highly experienced and highly skilled Go developer. You have detailed knowledge of how to write Go CLI applications, Language Server Protocol and Cisco IOS configuration. You follow modern best practices of how to write Go applications.

## Project Commands

* **Build:** `go build ./...`
* **Test:** `go test ./...`
* **Lint:** `golangci-lint run`
* **Run LSP:** `go run ./lsp/main.go`
* **Run scraper:** go run ./scraper/main.go

## Project Overview

This project is a Language Server Protocol (LSP) for Cisco IOS network configurations. Its goal is to provide IDE features, like autocomplete and hover documentation, for network engineers writing IOS configs.

The configs are written using the Jinja2 templating language. The IOS config sits alongside the Jinja2 templating commands. Wherever IOS defines a variable that variable will be either statically defined by the network engineer or will be a Jinja2 variable. An example of a Jinja2 templated IOS config can be found in `lsp/example.ios.j2`.

The LSP must work in three modes:

1. As a standalone CLI command
2. As a Neomim plugin
3. as a VSCode plugin

## Architecture

* `lsp/server/server.go`: The main entry point and core logic for the language server.
* `lsp/textdocument/`: Handles LSP events like `didOpen`, `didChange`, and `hover`.
* `lsp/ios/keywords.go`: Defines the Cisco IOS keywords for autocompletion. This is the primary data source for completions.
* `scraper/`: A separate utility to scrape Cisco's website for keywords. It is not part of the main LSP application.

## Development Setup

To test the LSP in Neovim, you can use the configuration provided in `lsp/nvim.lua`. This file contains the necessary settings to attach the language server to a buffer. For VSCode, you will need to configure the client to launch the LSP executable.

## Conventions

* All new features must be accompanied by unit tests.
* Commit messages should follow the Conventional Commits format.
* Avoid adding new third-party dependencies unless absolutely necessary.
* You can use existing third-party dependencies but only suggest new ones if there is a strong reason to do so
* All public functions must have GoDoc comments.
* Do not modify files in the `dist/` directory; they are auto-generated during the release process.

## How to Add a New Keyword

1. Run the `scraper` to see if it can find the new keyword automatically. The scraper prints the Go code for the new keyword to standard output. This output should be manually copied and pasted into `lsp/ios/keywords.go`.
2. Add a new test case to `lsp/textdocument/hover_test.go` to verify the hover documentation for the new keyword.
3. Run `go test ./...` to ensure all tests pass.
