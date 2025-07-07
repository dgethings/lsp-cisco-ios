#!/bin/bash

# This is a placeholder script for packaging the Neovim extension.
# In a real-world scenario, this would involve creating a zip file or a LuaRock.

echo "Packaging Neovim extension..."
mkdir -p dist
zip -r dist/cisco-ios-lsp-nvim.zip lsp/
echo "Neovim extension packaged successfully."
