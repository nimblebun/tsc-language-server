# Client Guides

These guides will show you how you can take advantage of the TSC Language Server
in your favorite editor/IDE.

Most of these setups will require you to download a [release][release-url].

- [Visual Studio Code](#visual-studio-code)
- [Sublime Text](#sublime-text)
- [IntelliJ IDEA](#intellij-idea)
- [neovim](#neovim)

## Visual Studio Code

You can download the [Cave Story TSC extension][vscode-extension] for Visual
Studio Code. You don't have to have the language server installed - the
extension will the installation and the updates of the language server for you.

## Sublime Text

### Language Client

1. Install Package Control (<kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd>, then
type `Install Package Control` and press <kbd>Enter</kbd>).

2. Install the [Sublime LSP][sublime-lsp] package.

3. Add the following code to your User settings (`LSP.sublime-settings`) -
`Preferences -> Package Settings -> LSP -> Settings`:

```json
{
	"clients": {
		"tsc": {
      "command": ["/path/to/your/tsc-ls", "start"],
      "enabled": true,
      "languageId": "tsc",
      "scopes": ["source.tsc"],
      "syntaxes": ["Packages/User/tsc.sublime-syntax"]
    }
  }
}
```

### Syntax Highlighting

Create a new syntax file (`Tools -> Developer -> New Syntax`) with the following
contents:

```yml
%YAML 1.2
---
file_extensions:
  - tsc
scope: source.tsc
contexts:
  main:
    - match: '#.+\n'
      scope: constant.numeric.tsc

    - match: '<([A-Z0-9+-]{3})'
      scope: keyword.control.tsc

    - match: '(?<=(<([A-Z0-9+-]{3}))|[^#0-9])([0-9V])([0-9]){3}'
      scope: string.quoted.single.tsc

    - match: '\/\/.+\n'
      scope: comment.line.tsc
```

## IntelliJ IDEA

### Language Client

- Install the [LSP Support][intellij-lsp] plugin.
- `File -> Settings -> Language & Frameworks -> Language Server Protocol`, then
select `Server Definitions`
  - Mode: `Executable`
  - Extension: `tsc`
  - Path: (insert path to tsc-ls)
  - Args: `start`
- Click on `Apply`

### Syntax Highlighter

Currently, there is no IDEA extension that can provide syntax highlighting for
the TSC language. In the meantime, you can try creating a tmbundle using the
[syntaxes from the Visual Studio Code extension][vsc-syntaxes].

## neovim

We recommend using [coc.nvim][coc-nvim]. Once installed, you can add the
following to your `coc-setting.json` (`:CocConfig` in nvim):

```json
{
  "languageserver": {
    "tsc": {
      "command": "/path/to/tsc-ls",
      "args": [ "start" ],
      "filetypes": [ "tsc" ]
    }
  }
}
```

[release-url]: https://github.com/nimblebun/tsc-language-server/releases/latest
[vscode-extension]: https://marketplace.visualstudio.com/items?itemName=jozsefsallai.vscode-tsc
[sublime-lsp]: https://github.com/sublimelsp/LSP#installation
[intellij-lsp]: https://plugins.jetbrains.com/plugin/10209-lsp-support
[vsc-syntaxes]: https://github.com/jozsefsallai/vscode-tsc/blob/master/syntaxes/tsc.tmLanguage.json
[coc-nvim]: https://github.com/neoclide/coc.nvim
