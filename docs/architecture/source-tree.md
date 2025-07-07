# Source Tree Structure

This document outlines the standard source tree structure for the Cisco IOS LSP project.

```plaintext
project-root/
├── .github/                    # GitHub Actions workflows for CI/CD
│   └── workflows/
│       └── go.yml
├── docs/                       # Project documentation
│   ├── prd/                    # Product Requirements Document (sharded)
│   │   ├── epic-1-...
│   │   ├── epic-2-...
│   │   ├── epic-3-...
│   │   ├── epics.md
│   │   ├── index.md
│   │   ├── requirements.md
│   │   └── technical-assumptions.md
│   ├── architecture/           # Architectural documentation
│   │   ├── coding-standards.md
│   │   ├── tech-stack.md
│   │   └── source-tree.md
│   ├── frontend-architecture.md # UI/UX specifications
│   └── PUBLISHING.md           # Extension publishing guide
├── ftplugin/                   # Neovim filetype plugins (for syntax highlighting)
├── lsp/                        # Language Server Protocol implementation (Go module)
│   ├── cmd/                    # Cobra CLI commands
│   ├── ios/                    # Cisco IOS command data and logic
│   │   └── commands.json       # Scraped command data (generated)
│   ├── server/                 # LSP server implementation
│   ├── textdocument/           # LSP text document handlers
│   └── ...                     # Other LSP related files
├── lua/                        # Neovim Lua files (for snippets)
├── scraper/                    # Web scraper for Cisco IOS documentation (Go module)
│   ├── cmd/                    # Cobra CLI commands
│   ├── cache/                  # Scraper cache
│   └── ...                     # Other scraper related files
├── scripts/                    # Utility scripts (packaging, etc.)
├── snippets/                   # VSCode snippets
├── syntaxes/                   # VSCode TextMate grammars
├── .bmad-core/                 # BMad framework core configuration
├── .git/                       # Git repository data
├── .gitignore                  # Git ignore rules
├── go.mod                      # Go module dependencies for root
├── go.sum                      # Go module checksums for root
└── README.md                   # Project README and user guide
```
