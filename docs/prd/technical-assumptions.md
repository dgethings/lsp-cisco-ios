# Technical Assumptions

### Repository Structure

*   A single repository (`monorepo`) will be used to house the LSP server and both the VSCode and Neovim extensions.

### Service Architecture

*   The core logic will be a standalone LSP server written in Go.

### Testing requirements

*   Unit tests will be written for all core LSP features.
*   Integration tests will be created to validate the interaction between the LSP and the editor extensions.
