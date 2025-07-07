# Epic 2: Core LSP Functionality

This epic focuses on implementing the core LSP features, building upon the existing codebase.

### Story 2.1: Implement Syntax Highlighting

As a network architect,
I want to see syntax highlighting for Cisco IOS and Jinja2,
so that I can easily read and understand the configuration files.

#### Acceptance Criteria

- 1: Cisco IOS commands are highlighted with distinct colors.
- 2: Jinja2 templating syntax is highlighted correctly within the configuration files.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. Syntax highlighting files for VSCode (`syntaxes/cisco-ios.tmLanguage.json`) and Neovim (`ftplugin/cisco-ios.lua`) have been created.


### Story 2.2: Implement Autocompletion

As a network architect,
I want the editor to suggest completions for IOS commands and Jinja2 syntax,
so that I can write configurations faster and with fewer errors.

#### Acceptance Criteria

- 1: The LSP suggests completions for Cisco IOS commands as I type.
- 2: The LSP suggests completions for Jinja2 syntax.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. The data loading mechanism was refactored to use JSON, the scraper was updated, and the `Completion` function was enhanced for context-aware suggestions.


### Story 2.3: Implement Diagnostics (Linting)

As a network architect,
I want the editor to show me errors in my configuration in real-time,
so that I can fix them before deployment.

#### Acceptance Criteria

- 1: The LSP flags unknown or invalid Cisco IOS commands.
- 2: The LSP flags syntax errors in Jinja2 templates.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. The `Diagnose` function was implemented, integrated into `DidChange`, and the LSP server was updated to advertise diagnostics capability.


### Story 2.4: Implement Hover Documentation

As a network architect,
I want to see documentation for Cisco IOS commands when I hover over them,
so that I can quickly understand their purpose and usage.

#### Acceptance Criteria

- 1: Hovering over a Cisco IOS command displays a popup with its documentation.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. The `hover.go` file was refactored to improve the accuracy of hover documentation.


### Story 2.5: Implement Code Formatting

As a network architect,
I want to be able to automatically format my Cisco IOS configuration,
so that it is clean and consistently styled.

#### Acceptance Criteria

- 1: A command is available to format the current Cisco IOS configuration file.
- 2: Formatting is applied consistently according to a predefined style.

**Status:** ✅ Completed
**Completion Notes:** This story was completed during the PRD refinement phase. The `formatting.go` file was created with a basic `Format` function, and integrated into the LSP server to handle `TextDocumentFormatting` requests.

