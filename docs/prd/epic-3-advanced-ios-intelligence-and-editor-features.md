# Epic 3: Advanced IOS Intelligence & Editor Features

This epic builds on the core LSP by adding advanced intelligence and editor-specific features. This includes context-aware completions, version validation, and a visual tree view of the configuration.

### Story 3.1: Implement Context-Aware Completions

As a network architect,
I want the LSP to provide intelligent completions based on the current configuration context,
so that I only see relevant suggestions.

#### Acceptance Criteria

- 1: Command completions are filtered based on the current configuration block (e.g., `interface`, `router bgp`).
- 2: The LSP suggests appropriate next commands based on the current line.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. The `Completion` function was enhanced with `determineMode` for context awareness, and the `ios.Keyword` struct and scraper were updated to include and populate the `Mode` field.


### Story 3.2: Implement IOS Version Validation

As a network architect,
I want the LSP to validate my configuration against a specific IOS version,
so that I can ensure compatibility with my network devices.

#### Acceptance Criteria

- 1: The user can specify a target IOS/IOS-XE version.
- 2: The LSP flags commands or parameters that are not available in the specified version.
- 3: The LSP warns about the use of deprecated commands in the specified version.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. Version fields were added to `ios.Keyword`, the scraper was updated to populate them, `Diagnose` was modified for version validation, and `lsp/server/server.go` was updated for configuration.


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
