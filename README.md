# About
frmt is a CLI format convertor for those who live in the terminal.

It was built for my own convenience and thought it'd be cool to turn this into a tool that anyone can use.
Go was the first choice due to the ease of making CLI tools with the language.

# Usage
To convert your file into another format, execute the following command:

```shell
frmt -file=<FILE_PATH> -format=<DESIRED_FORMAT>
```

- Converted file is created in your present working directory (aka where you executed the command)
- File paths can be relative and/or absolute
- Format flag is case insensitive

# Supported formats
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
