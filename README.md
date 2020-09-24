<div align="center">
	<h1>Meeseeks.</h1>
    <a href="https://github.com/melbahja/goph">
    </a>
    <h4 align="center">
	   Helper to run repeating multiple cli commands    
	</h4>
</div>

<p align="center">
    <a href="#installation">Overview</a> |
    <a href="#features">Features</a> |
    <a href="#usage">Usage</a> |
    <a href="#license">License</a>
</p>


## Overview
Meeseeks is a tool to run cli commands that you find yourself running all the time. Meeseeks helps out by providing a configurable file where such commands can be provided and meeseeks runs the commands for you displaying the output and status of running commands.

## Installation

You can download the binary from the release page or If you have a working Go installation, you can also build from source
```bash
go build main.go
```
then you can run the binary using 
```bash 
./main
```

## Features

- Easy to use.
- Configurable using **JSON** 
- Cross Platform
- Supports concurrent and synchronous running of commands.

## Usage

modify config.json with required commands, run meeseeks, sit back and watch meeseeks play ☺️
```json
{
  "commands": {
    "concurrent": true,  
    "list": [
      {
        "directory": "C:\\Users\\username\\work\\myproject",
        "concurrent": true,
        "list": [ { "cmd": "npm run build -- --spa"}, {"cmd":"go build main.go"} ]
      },
      {
        "directory": "C:\\Users\\username\\project\\prjectFolder",
        "concurrent": false,
         "list": [{ "cmd": "npm run build"}, {"cmd":"7z a .nuxt.zip .nuxt"}]
      }
    ]
  }
}

```
This is a work in progress, more features coming.

## License

Meeseeks is provided under the [MIT License](https://github.com/Popoola-Opeyemi/meeseeks/master/LICENSE).