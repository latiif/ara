# Ara

A command line tool that correctly displays Arabic text in terminals.
### Requirements
* The Go programming language
* `dep` Dependency manager

### Install
* Clone the repo.
* Run `make`.
* Run the produced binary `ara`.

### Key Features
- Improved Non-Arabic and punctuation interpolation with Arabic text.
- RTL line wrap and RTL line padding. (Depending on terminal width)
- RTL rendering in terminals. (Right alignment)
- Glyphes and ligatures for لآ,لا,لأ,لإ.

### Usage
When displaying Arabic text in terminals, it is shown incorrectly, see image.
![alt text](http://i.imgur.com/ygbSxHq.png "Original behaviour")

With *Ara*, you can pipe in Arabic text and it will display correctly.

![alt text](http://i.imgur.com/0mVt1km.png "Piped with Ara")

In addition, the output is right-aligned with smart line breaks accomodating to terminal width.

![alt text](http://i.imgur.com/VGd16kM.png "Line-wrap with right alignment") 

### Disclaimer
This tool is based on [goarabic](https://github.com/01walid/goarabic) under MIT license, with additional functionalities.