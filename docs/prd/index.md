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

## Sections

- [Requirements](./requirements.md)
- [Technical Assumptions](./technical-assumptions.md)
- [Epics](./epics.md)
- [Epic 1: Analyze and Integrate with Existing LSP Codebase](./epic-1-analyze-and-integrate-with-existing-lsp-codebase.md)
- [Epic 2: Core LSP Functionality](./epic-2-core-lsp-functionality.md)
- [Epic 3: Advanced IOS Intelligence & Editor Features](./epic-3-advanced-ios-intelligence-and-editor-features.md)
