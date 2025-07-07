# Cisco IOS Language Server Protocol (LSP) Extension PRD

## Goals and Background Context

### Goals

*   To provide a feature-rich language server for Cisco IOS configurations.
*   To improve the productivity and accuracy of network engineers working with Cisco IOS.
*   To create a tool that supports both Neovim and VSCode for a consistent cross-editor experience.
*   To deliver a powerful and intuitive tool for senior network architects.

### Background Context

Network engineers frequently work with complex Cisco IOS configurations. This project aims to create a Language Server Protocol (LSP) implementation that provides IDE-like features for Cisco IOS configuration files. By offering features like syntax highlighting, autocompletion, and real-time validation, the LSP will reduce errors and speed up the development of network configuration templates. The initial focus is on supporting Jinja2 for templating and providing a core set of features for VSCode and Neovim.

### Change Log

| Date       | Version | Description              | Author |
| :--------- | :------ | :----------------------- | :----- |
| 2025-07-06 | 0.1     | Initial draft of the PRD | Sarah  |

## Requirements

### Functional

*   **FR1:** The LSP shall provide syntax highlighting for Cisco IOS commands.
*   **FR2:** The LSP shall provide syntax highlighting for Jinja2 templates.
*   **FR3:** The LSP shall offer autocompletion for Cisco IOS commands and their parameters.
*   **FR4:** The LSP shall offer autocompletion for Jinja2 syntax.
*   **FR5:** The LSP shall provide real-time diagnostics (linting) for Cisco IOS configurations.
*   **FR6:** The LSP shall provide real-time diagnostics for Jinja2 templates.
*   **FR7:** The LSP shall display hover-over documentation for Cisco IOS commands.
*   **FR8:** The LSP shall provide code formatting for Cisco IOS configuration blocks.
*   **FR9:** The LSP shall provide context-aware completions based on the configuration block.
*   **FR10:** The LSP shall validate configurations against specific IOS/IOS-XE versions.
*   **FR11:** The LSP shall warn users about deprecated Cisco IOS commands.
*   **FR12:** The LSP shall differentiate between different Cisco device types (e.g., routers, switches).
*   **FR13:** The VSCode and Neovim extensions shall provide code snippets for common configurations.
*   **FR14:** The VSCode and Neovim extensions shall feature a tree view to visualize the configuration hierarchy.

### Non Functional

*   **NFR1:** The LSP server should be performant and not introduce noticeable lag during typing.
*   **NFR2:** The extension should be easy to install and configure in both VSCode and Neovim.

## Technical Assumptions

### Repository Structure

*   A single repository (`monorepo`) will be used to house the LSP server and both the VSCode and Neovim extensions.

### Service Architecture

*   The core logic will be a standalone LSP server written in Go.

### Testing requirements

*   Unit tests will be written for all core LSP features.
*   Integration tests will be created to validate the interaction between the LSP and the editor extensions.

## Epics

- Epic 1: Analyze and Integrate with Existing LSP Codebase
- Epic 2: Core LSP Functionality
- Epic 3: Advanced IOS Intelligence & Editor Features

## Epic 1: Analyze and Integrate with Existing LSP Codebase

This epic focuses on understanding the existing LSP codebase, establishing a solid development and testing foundation, and preparing for new feature development.

### Story 1.1: Analyze Existing LSP Implementation

As a developer,
I want to analyze the existing `lsp/` codebase,
so that I can understand its structure, dependencies, and current capabilities to inform our development plan.

#### Acceptance Criteria

- 1: A summary of the existing `lsp/` directory's functionality is created and documented.
- 2: The existing Go dependencies from `go.mod` are reviewed and their purpose is documented.
- 3: A high-level plan for integrating new features with the existing code is drafted.
- 4: The `README.md` is updated with instructions for setting up the development environment for this existing project.

### Story 1.2: Adapt CI/CD Pipeline for LSP

As a developer,
I want to adapt the existing CI/CD pipeline for the LSP server,
so that every change is automatically built and tested.

#### Acceptance Criteria

- 1: The CI/CD pipeline is configured to build the Go LSP server.
- 2: The pipeline is configured to run unit tests on every commit.
- 3: A script is created to build and package the VSCode extension.
- 4: A script is created to package the Neovim extension.
- 5: The process for publishing to the VSCode Marketplace is documented.
- 6: The process for publishing to a Neovim package manager is documented.

### Story 1.3: Configure Testing Framework

As a developer,
I want to configure the testing framework for the LSP server,
so that I can write and run unit tests from the beginning.

#### Acceptance Criteria

- 1: The Go testing framework (`go test`) is configured for the project.
- 2: Initial test files are created for the main packages.
- 3: The CI/CD pipeline is configured to run the tests.

### Story 1.4: Analyze and Document the IOS Command Scraper

As a developer,
I want to analyze the existing `scraper/` application and its data source,
so that I can understand how to maintain and operate it to provide command data to the LSP.

#### Acceptance Criteria

- 1: The functionality of the `scraper/` application is documented, including its inputs (the Cisco website) and outputs (`keywords.go`).
- 2: The process for running the scraper to refresh the command database is documented in the `README.md`.
- 3: The risks associated with depending on the structure of the Cisco documentation website are identified and documented.
- 4: The `keywords.tmpl` file is reviewed to understand how the Go code is generated.

### Story 1.5: Define UI/UX Specifications for Editor Extensions

As a user,
I want a clear and consistent user interface for the editor extensions,
so that I can easily understand and interact with the LSP's features.

#### Acceptance Criteria

- 1: A `frontend-architecture.md` document is created.
- 2: Wireframes or mockups for the "tree view" are created and added to the new document.
- 3: The interaction model for all major UI features (hover, diagnostics, tree view) is defined.
- 4: The user workflow for configuring the extension (e.g., setting IOS version) is documented.
- 5: A plan for a frontend testing strategy is included.

## Epic 2: Core LSP Functionality

This epic focuses on implementing the core LSP features, building upon the existing codebase.

### Story 2.1: Implement Syntax Highlighting

As a network architect,
I want to see syntax highlighting for Cisco IOS and Jinja2,
so that I can easily read and understand the configuration files.

#### Acceptance Criteria

- 1: Cisco IOS commands are highlighted with distinct colors.
- 2: Jinja2 templating syntax is highlighted correctly within the configuration files.

### Story 2.2: Implement Autocompletion

As a network architect,
I want the editor to suggest completions for IOS commands and Jinja2 syntax,
so that I can write configurations faster and with fewer errors.

#### Acceptance Criteria

- 1: The LSP suggests completions for Cisco IOS commands as I type.
- 2: The LSP suggests completions for Jinja2 syntax.

### Story 2.3: Implement Diagnostics (Linting)

As a network architect,
I want the editor to show me errors in my configuration in real-time,
so that I can fix them before deployment.

#### Acceptance Criteria

- 1: The LSP flags unknown or invalid Cisco IOS commands.
- 2: The LSP flags syntax errors in Jinja2 templates.

### Story 2.4: Implement Hover Documentation

As a network architect,
I want to see documentation for Cisco IOS commands when I hover over them,
so that I can quickly understand their purpose and usage.

#### Acceptance Criteria

- 1: Hovering over a Cisco IOS command displays a popup with its documentation.

### Story 2.5: Implement Code Formatting

As a network architect,
I want to be able to automatically format my Cisco IOS configuration,
so that it is clean and consistently styled.

#### Acceptance Criteria

- 1: A command is available to format the current Cisco IOS configuration file.
- 2: Formatting is applied consistently according to a predefined style.

## Epic 3: Advanced IOS Intelligence & Editor Features

This epic builds on the core LSP by adding advanced intelligence and editor-specific features. This includes context-aware completions, version validation, and a visual tree view of the configuration.

### Story 3.1: Implement Context-Aware Completions

As a network architect,
I want the LSP to provide intelligent completions based on the current configuration context,
so that I only see relevant suggestions.

#### Acceptance Criteria

- 1: Command completions are filtered based on the current configuration block (e.g., `interface`, `router bgp`).
- 2: The LSP suggests appropriate next commands based on the current line.

### Story 3.2: Implement IOS Version Validation

As a network architect,
I want the LSP to validate my configuration against a specific IOS version,
so that I can ensure compatibility with my network devices.

#### Acceptance Criteria

- 1: The user can specify a target IOS/IOS-XE version.
- 2: The LSP flags commands or parameters that are not available in the specified version.
- 3: The LSP warns about the use of deprecated commands in the specified version.

### Story 3.3: Implement Device Type Awareness

As a network architect,
I want the LSP to be aware of the device type I am configuring,
so that I only see commands that are valid for that platform.

#### Acceptance Criteria

- 1: The user can specify a target device type (e.g., router, switch).
- 2: The LSP filters command suggestions based on the selected device type.

### Story 3.4: Implement Code Snippets

As a network architect,
I want to have access to code snippets for common configurations,
so that I can quickly insert boilerplate code.

#### Acceptance Criteria

- 1: A collection of snippets is available for common Cisco IOS configurations.
- 2: Snippets can be inserted into the editor via a command or by typing a prefix.

### Story 3.5: Implement Configuration Tree View

As a network architect,
I want to see a tree view of my configuration,
so that I can easily navigate and understand its structure.

#### Acceptance Criteria

- 1: A tree view is available in the editor's side panel.
- 2: The tree view accurately represents the hierarchy of the Cisco IOS configuration.
- 3: Clicking on a node in the tree view navigates to the corresponding line in the editor.

### Story 3.6: Create User Documentation

As a network architect,
I want clear documentation on how to use the extension,
so that I can understand its features and configure it correctly.

#### Acceptance Criteria

- 1: The project's main `README.md` is updated with comprehensive user documentation.
- 2: The documentation covers all features, including syntax highlighting, autocompletion, and validation.
- 3: The documentation explains how to configure the extension, including setting the IOS version and device type.
- 4: The documentation includes clear examples of how to use the extension.
