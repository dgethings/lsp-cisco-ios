vim.filetype.add({
	pattern = { ["ios.*j2"] = "IOS-JINJA" },
})
local client = vim.lsp.start_client({
	cmd = { "ios-lsp" },
	name = "iosLSP",
	root_dir = vim.fn.getcwd(),
})

if not client then
	vim.notify("Failed to start test client")
end
vim.notify("Starting IOS LSP " .. client)

vim.api.nvim_create_autocmd("FileType", {
	pattern = "IOS-JINJA",
	callback = function()
		vim.lsp.buf_attach_client(0, client)
		local suc = vim.lsp.buf_is_attached(0, client)
		vim.notify_once("LSP attached to buf is " .. tostring(suc))
	end,
})
