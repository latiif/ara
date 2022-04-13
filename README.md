# Ara
![Validate](https://github.com/latiif/ara/workflows/Validate/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/latiif/ara)](https://goreportcard.com/report/github.com/latiif/ara)
![GitHub last commit](https://img.shields.io/github/last-commit/latiif/ara)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/latiif/ara)

A command line tool that correctly displays Arabic text in terminals.

### Installation

Check [releases](https://github.com/latiif/ara/releases) and download the binary directly.

Get the binary and add it to your `$PATH`
```bash
wget https://github.com/latiif/ara/releases/download/0.7/ara && chmod +x ara
```
### Key Features
- 🆕 Undotting of Arabic letters (Rasm).
- 🆕 Revamped command line interface.
- 🆕 Add option to use files as input and/or output in the tool.
- Improved Non-Arabic and punctuation interpolation with Arabic text.
- RTL line wrap and RTL line padding. (Depending on terminal width).
- RTL rendering in terminals. (Right alignment).
- Glyphes and ligatures for لآ,لا,لأ,لإ.
- Full support (including RTL) for Tashkeels (One tashkeel per letter).

### Usage
When displaying Arabic text in terminals, it is shown incorrectly.

![alt text](https://i.ibb.co/wcYTjty/original.png "Original behaviour")

With *Ara*, you can send in Arabic text and it will display correctly. In addition, the output can be right-aligned with smart line breaks accomodating to terminal width.

![alt text](https://i.ibb.co/JrdRbNK/wara.png "sent with Ara")

You can use *Ara* to undot letters according to Rasm rules.

![alt text](https://i.ibb.co/Rv9ky8J/undot.png "undot")

### Synopsis
```
Correctly displays Arabic text in terminals
Usage:
  ara [flags]
  ara [command]

Available Commands:
  help        Help about any command
  undot       Removes the dots of Arabic letters
  version     Print the version number of ara

Flags:
  -a, --adjust-right    Adjust output text to be rtl (useful when in shell, less so if you want to pipe into a file)
  -h, --help            help for ara
  -i, --input string    Apply command on the contents of the input file.
  -o, --output string   Write output to specified file.

Use "ara [command] --help" for more information about a command.

```

### Font support
Unfortunately, not all monospace fonts support Arabic text fully.
Here is a list of known fonts with good support for Arabic.
* DejaVu Sans Mono
* Kawkab Mono
* Liberation Mono
* Monospace
* Nimbus Mono L

### Building from source
* Clone the repo.
* Run `make`.
* Run the produced binary `ara`.


### Disclaimer
This tool is based on [goarabic](https://github.com/01walid/goarabic) under MIT license, with additional functionalities.
