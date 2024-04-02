vim.lsp.start({
	name = "IOS LSP",
	cmd = { "ios-lsp" },
	root_dir = vim.fn.getcwd(),
})
