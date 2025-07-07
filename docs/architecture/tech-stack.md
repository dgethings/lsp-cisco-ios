# Technology Stack

This document defines the definitive technology stack for the Cisco IOS LSP project.

## Technology Stack Table

| Category           | Technology         | Version     | Purpose                      | Rationale                                      |
| :----------------- | :----------------- | :---------- | :--------------------------- | :--------------------------------------------- |
| **Language**       | Go                 | 1.22.1      | Primary development language | Performance, concurrency, cross-platform       |
| **LSP Framework**  | tliron/glsp        | 0.2.2       | Language Server Protocol     | Robust, well-maintained LSP library for Go     |
| **Web Scraper**    | gocolly/colly      | 1.2.0       | Web scraping                 | Efficient and flexible web scraping library    |
| **CLI Framework**  | spf13/cobra        | 1.8.0       | Command-line interface       | Standard for Go CLI applications               |
| **Testing**        | stretchr/testify   | 1.3.0       | Testing utilities            | Assertions and mocking for Go tests            |
| **Logging**        | tliron/commonlog   | 0.2.17      | Logging                      | Flexible logging for Go applications           |
| **CI/CD**          | GitHub Actions     | N/A         | Continuous Integration/Deployment | Integrated with GitHub, widely used            |
| **Editor Ext.**    | VSCode, Neovim     | N/A         | User Interface               | Target editors for network engineers           |
