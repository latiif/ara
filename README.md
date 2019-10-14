# Ara

A command line tool that correctly displays Arabic text in terminals. ([screencap](https://i.imgur.com/d8OGWux.gifv))

### Installation 

Check [releases](https://github.com/llusx/ara/releases) and download the binary directly.

Get the binary and add it to your `$PATH`
```bash
wget https://github.com/llusx/ara/releases/download/v0.1-alpha/ara && chmod +x ara
```
### Key Features
- Improved Non-Arabic and punctuation interpolation with Arabic text.
- RTL line wrap and RTL line padding. (Depending on terminal width)
- RTL rendering in terminals. (Right alignment)
- Glyphes and ligatures for لآ,لا,لأ,لإ.

### Font support

Unfortunately, not all monospace fonts support Arabic text fully. 
Here is a list of known fonts with good support for Arabic.
* DejaVu Sans Mono
* Kawkab Mono
* Liberation Mono
* Monospace
* Nimbus Mono L

### Usage
When displaying Arabic text in terminals, it is shown incorrectly, see image.
![alt text](http://i.imgur.com/ygbSxHq.png "Original behaviour")

With *Ara*, you can pipe in Arabic text and it will display correctly.

![alt text](http://i.imgur.com/0mVt1km.png "Piped with Ara")

In addition, the output is right-aligned with smart line breaks accomodating to terminal width.

![alt text](http://i.imgur.com/VGd16kM.png "Line-wrap with right alignment") 


### Building from source
* Clone the repo.
* Run `make`.
* Run the produced binary `ara`.


### Disclaimer
This tool is based on [goarabic](https://github.com/01walid/goarabic) under MIT license, with additional functionalities.