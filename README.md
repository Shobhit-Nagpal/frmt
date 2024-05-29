# About
frmt is a CLI format convertor for those who live in the terminal.

It was built for my own convenience and thought it'd be cool to turn this into a tool that anyone can use.
Go was the first choice due to the ease of making CLI tools with the language.

Pronounced as format.

![Frmt working ss](https://shobhit-nagpal.github.io/portfolio/static/media/frmt.4a5dd7b2c1cff4087995.png)

# Installation

## 1. Install with Go
```shell
go install github.com/Shobhit-Nagpal/frmt@latest
```

## 2. Install from Homebrew
### Tap the repo
``` shell
brew tap Shobhit-Nagpal/frmt
```
### Install frmt
```shell
brew install frmt
```
- To check the directory where frmt is installed:
```shell
which frmt
```
## 3. Install manually
### Clone repo
```shell
git clone git@github.com:Shobhit-Nagpal/frmt.git
```

### Move binary to /bin directory
```shell
sudo mv frmt /bin/
```

- Once the binary file has been moved to /bin/ directory, you can start using frmt as you'd like :)

# Uninstallation
## 1. With Homebrew
```shell
brew uninstall frmt
```
## 2. Manually
To remove the binary from your bin directory:
```shell
which frmt
```

```shell
sudo rm $(which frmt) 
```

# Usage
To convert your file into another format, execute the following command:

```shell
frmt -p <FILE_PATH> -f <DESIRED_FORMAT>
```

- Converted file is created in your present working directory (aka where you executed the command)
- File paths can be relative and/or absolute
- Format flag is case insensitive

- For information on flags, run
``` shell
frmt -h
```

# Supported output formats
- PNG
- JPG / JPEG
- BMP
- PDF
- TIF / TIFF

# Working
The way frmt works is pretty simple.

- First, it takes your file path and desired format as arguments/flags
- It converts the file path to an absolute path and then checks if the file exists
- If the file exists, it checks if the file is already present in the desired format
- If the file is not in its desired format, it creates a new file which is the converted format of the original file
- Error handling is done at each stage

# To add:
- Give pdf as input and convert to desired format
