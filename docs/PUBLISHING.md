# Publishing the Cisco IOS LSP Extensions

This document outlines the process for publishing the VSCode and Neovim extensions.

## VSCode Extension

### Prerequisites

- [Node.js](https://nodejs.org/) and `npm` installed.
- `vsce` (the Visual Studio Code Extension Manager) installed globally: `npm install -g vsce`
- A Personal Access Token (PAT) for the Visual Studio Marketplace.

### Packaging

1.  Run the `package-vscode.sh` script:
    ```sh
    ./scripts/package-vscode.sh
    ```
2.  This will create a `.vsix` file in the `dist/` directory.

### Publishing

1.  Log in to the Visual Studio Marketplace:
    ```sh
    vsce login <publisher-name>
    ```
2.  Publish the extension:
    ```sh
    vsce publish -p <token>
    ```

## Neovim Extension

### Packaging

1.  Run the `package-nvim.sh` script:
    ```sh
    ./scripts/package-nvim.sh
    ```
2.  This will create a `cisco-ios-lsp-nvim.zip` file in the `dist/` directory.

### Publishing

Publishing a Neovim extension typically involves creating a Git tag and letting users install it via a plugin manager like `packer` or `lazy.nvim`. Alternatively, it could be published to a package manager like [LuaRocks](https://luarocks.org/).

**Example (Git Tag):**

```sh
git tag v0.1.0
git push origin v0.1.0
```

Users can then install it by referencing the tag in their Neovim configuration.
