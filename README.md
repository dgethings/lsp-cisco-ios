# Language Server for Cisco IOS Configuration

Language Server for Cisco IOS is an [LSP](https://microsoft.github.io/language-server-protocol/) for Cisco IOS configuration. It is to be used with editors like VSCode or NeoVIM that support the Langurage Server Protocol defined by Microsoft.

This app is at Alpha stage meaning it has very limited functionality and requires developer expertise to get working. As this app matures these issues will be addressed making much more user friendly.

The target audience, eventually, is Network Engineers tasked with creating and maintaining network configurations.

## Getting Started (for Developers)

This project contains two main Go modules: the Language Server itself (`lsp/`) and a web scraper (`scraper/`) that provides the command data.

### Prerequisites

- Go 1.22.1 or later

### Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/dgethings/lsp-cisco-ios.git
    cd lsp-cisco-ios
    ```

2.  **Install LSP dependencies:**
    ```sh
    cd lsp
    go mod download
    ```

3.  **Install Scraper dependencies:**
    ```sh
    cd ../scraper
    go mod download
    ```

### Building

To build the LSP executable:

```sh
cd lsp
go build
```

This will create an `lsp` executable in the `lsp/` directory.

## User Guide

This Language Server Protocol (LSP) extension provides IDE-like features for Cisco IOS configuration files, enhancing productivity and accuracy for network engineers.

### Features

-   **Syntax Highlighting:** Provides distinct color highlighting for Cisco IOS commands and Jinja2 templating syntax.
-   **Autocompletion:** Offers intelligent suggestions for Cisco IOS commands and their parameters, filtered by context and device type.
-   **Real-time Diagnostics (Linting):** Flags unknown or invalid Cisco IOS commands as you type.
-   **Hover Documentation:** Displays detailed documentation for Cisco IOS commands when you hover over them.
-   **Code Formatting:** Provides basic formatting for Cisco IOS configuration blocks.
-   **Configuration Tree View:** (Editor-specific) Visualizes the configuration hierarchy for easier navigation.

### Configuration

You can configure the LSP extension through your editor's settings. The following settings are available under the `ios-lsp` scope:

-   `ios-lsp.targetIOSVersion`: Specifies the target Cisco IOS/IOS-XE version for validation (e.g., `15.2`, `16.9`). Commands not available in the specified version will be flagged.
-   `ios-lsp.targetDeviceType`: Specifies the target device type for command suggestions (e.g., `router`, `switch`). This helps filter completions to only show relevant commands for your device.

**Example VSCode `settings.json`:**

```json
{
    "ios-lsp.targetIOSVersion": "15.4",
    "ios-lsp.targetDeviceType": "router"
}
```

### Code Snippets

Common Cisco IOS configurations are available as code snippets to speed up your workflow. You can trigger these snippets by typing their prefix and pressing `Tab` (or your editor's snippet trigger key).

**Available Snippets:**

-   `interface`: Inserts a basic interface configuration block.
-   `router bgp`: Inserts a basic BGP router configuration block.
-   `ip route`: Inserts a static route configuration.

### Examples

**Autocompletion in action:**

Type `int` and press `Tab` to get suggestions for `interface`.

```cisco-ios
interface GigabitEthernet0/1
 ip address 192.168.1.1 255.255.255.0
 no shutdown
```

**Hovering over a command:**

Hover your mouse over `ip address` to see its documentation.

**Diagnostics for invalid commands:**

Typing an unknown command like `invalid-command` will show an error.

```cisco-ios
invalid-command
! ^^^^^^^^^^^^^^ Unknown or invalid Cisco IOS command
```

**Using a snippet:**

Type `interface` and press `Tab` to insert the interface configuration snippet.

```cisco-ios
interface GigabitEthernet0/1
 description Link to Core
 ip address 10.0.0.1 255.0.0.0
 no shutdown
```
