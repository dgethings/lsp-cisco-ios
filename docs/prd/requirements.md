# Requirements

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
