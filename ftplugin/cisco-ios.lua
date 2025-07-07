-- Neovim syntax file for Cisco IOS configurations

vim.cmd [[syntax clear]]

-- Comments
vim.cmd [[syntax match ciscoIosComment "!.*$"]]

-- Jinja templating
vim.cmd [[syntax region jinjaBlock start="{\%" end="\%}" keepend contains=@jinja]]
vim.cmd [[syntax region jinjaVar start="{{" end="}}" keepend contains=@jinja]]

-- Keywords
vim.cmd [[syntax keyword ciscoIosKeyword interface router ip access-list banner enable configure line crypto snmp-server ntp]]
vim.cmd [[syntax keyword ciscoIosFlowKeyword no exit end]]
vim.cmd [[syntax keyword ciscoIosFunction permit deny log shutdown negotiation duplex speed]]

-- Strings
vim.cmd [[syntax region ciscoIosString start='"' end='"']]

-- Highlight links
vim.cmd [[hi def link ciscoIosComment Comment]]
vim.cmd [[hi def link jinjaBlock Delimiter]]
vim.cmd [[hi def link jinjaVar Delimiter]]
vim.cmd [[hi def link ciscoIosKeyword Keyword]]
vim.cmd [[hi def link ciscoIosFlowKeyword Statement]]
vim.cmd [[hi def link ciscoIosFunction Function]]
vim.cmd [[hi def link ciscoIosString String]]

-- Include Jinja syntax highlighting
vim.cmd [[runtime! syntax/jinja.vim]]
