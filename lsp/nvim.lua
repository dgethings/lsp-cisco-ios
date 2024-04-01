vim.lsp.start({
	name = "lsp-cisco-ios",
	cmd = { "./bin/lsp-cisco-ios" },
	root_dir = vim.fn.getcwd(),
})
