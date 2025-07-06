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

- Epic1 Core LSP Functionality & Setup: Establish the foundational features of the language server, including syntax highlighting, autocompletion, and basic linting.
- Epic2 Advanced IOS Intelligence & Editor Features: Implement advanced, context-aware features and editor-specific integrations like the configuration tree view.

## Epic 1 Core LSP Functionality & Setup

This epic focuses on establishing the core functionality of the language server. It includes setting up the project, implementing fundamental LSP features like syntax highlighting, autocompletion, and diagnostics for both Cisco IOS and Jinja2.

### Story 1.1 Setup Project and Basic LSP Server

As a developer,
I want to set up the initial project structure and a basic LSP server,
so that we have a foundation for implementing features.

#### Acceptance Criteria

- 1: A Go project for the LSP server is created.
- 2: A basic LSP server is implemented that can communicate with a client.
- 3: The repository is initialized with a `.gitignore` file and a `README.md`.

### Story 1.2 Implement Syntax Highlighting

As a network architect,
I want to see syntax highlighting for Cisco IOS and Jinja2,
so that I can easily read and understand the configuration files.

#### Acceptance Criteria

- 1: Cisco IOS commands are highlighted with distinct colors.
- 2: Jinja2 templating syntax is highlighted correctly within the configuration files.

### Story 1.3 Implement Autocompletion

As a network architect,
I want the editor to suggest completions for IOS commands and Jinja2 syntax,
so that I can write configurations faster and with fewer errors.

#### Acceptance Criteria

- 1: The LSP suggests completions for Cisco IOS commands as I type.
- 2: The LSP suggests completions for Jinja2 syntax.

### Story 1.4 Implement Diagnostics (Linting)

As a network architect,
I want the editor to show me errors in my configuration in real-time,
so that I can fix them before deployment.

#### Acceptance Criteria

- 1: The LSP flags unknown or invalid Cisco IOS commands.
- 2: The LSP flags syntax errors in Jinja2 templates.

### Story 1.5 Implement Hover Documentation

As a network architect,
I want to see documentation for Cisco IOS commands when I hover over them,
so that I can quickly understand their purpose and usage.

#### Acceptance Criteria

- 1: Hovering over a Cisco IOS command displays a popup with its documentation.

### Story 1.6 Implement Code Formatting

As a network architect,
I want to be able to automatically format my Cisco IOS configuration,
so that it is clean and consistently styled.

#### Acceptance Criteria

- 1: A command is available to format the current Cisco IOS configuration file.
- 2: Formatting is applied consistently according to a predefined style.

## Epic 2 Advanced IOS Intelligence & Editor Features

This epic builds on the core LSP by adding advanced intelligence and editor-specific features. This includes context-aware completions, version validation, and a visual tree view of the configuration.

### Story 2.1 Implement Context-Aware Completions

As a network architect,
I want the LSP to provide intelligent completions based on the current configuration context,
so that I only see relevant suggestions.

#### Acceptance Criteria

- 1: Command completions are filtered based on the current configuration block (e.g., `interface`, `router bgp`).
- 2: The LSP suggests appropriate next commands based on the current line.

### Story 2.2 Implement IOS Version Validation

As a network architect,
I want the LSP to validate my configuration against a specific IOS version,
so that I can ensure compatibility with my network devices.

#### Acceptance Criteria

- 1: The user can specify a target IOS/IOS-XE version.
- 2: The LSP flags commands or parameters that are not available in the specified version.
- 3: The LSP warns about the use of deprecated commands in the specified version.

### Story 2.3 Implement Device Type Awareness

As a network architect,
I want the LSP to be aware of the device type I am configuring,
so that I only see commands that are valid for that platform.

#### Acceptance Criteria

- 1: The user can specify a target device type (e.g., router, switch).
- 2: The LSP filters command suggestions based on the selected device type.

### Story 2.4 Implement Code Snippets

As a network architect,
I want to have access to code snippets for common configurations,
so that I can quickly insert boilerplate code.

#### Acceptance Criteria

- 1: A collection of snippets is available for common Cisco IOS configurations.
- 2: Snippets can be inserted into the editor via a command or by typing a prefix.

### Story 2.5 Implement Configuration Tree View

As a network architect,
I want to see a tree view of my configuration,
so that I can easily navigate and understand its structure.

#### Acceptance Criteria

- 1: A tree view is available in the editor's side panel.
- 2: The tree view accurately represents the hierarchy of the Cisco IOS configuration.
- 3: Clicking on a node in the tree view navigates to the corresponding line in the editor.

### Story 2.6 Implement Deployment Workflow

As a developer,
I want to have a defined process for building and publishing the extensions,
so that we can reliably release new versions to users.

#### Acceptance Criteria

- 1: A script is created to build and package the VSCode extension.
- 2: A script is created to package the Neovim extension.
- 3: The process for publishing to the VSCode Marketplace is documented.
- 4: The process for publishing to a Neovim package manager is documented.

### Story 2.7 Create User Documentation

As a network architect,
I want clear documentation on how to use the extension,
so that I can understand its features and configure it correctly.

#### Acceptance Criteria

- 1: The project's main `README.md` is updated with comprehensive user documentation.
- 2: The documentation covers all features, including syntax highlighting, autocompletion, and validation.
- 3: The documentation explains how to configure the extension, including setting the IOS version and device type.
- 4: The documentation includes clear examples of how to use the extension.
