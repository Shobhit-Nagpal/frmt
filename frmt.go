package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/sunshineplan/imgconv"
)

var sameFormats = map[string]string{
    "jpg" : "jpeg",
    "jpeg" : "jpg",
    "png" : "apng",
    "apng" : "png",
    "tiff" : "tif",
    "tif" : "tiff",
    "bmp" : "dib",
    "dib" : "bmp",
}

func printError(err string) {
    fmt.Fprintf(os.Stderr, err)
}

func fileExists(path string) bool {

    info, err := os.Stat(path)
    if err != nil {
        errMsg := fmt.Sprintf("PATH ERROR: %s\n", err)
        printError(errMsg)
        return false
    }

    if info.IsDir() {
        printError("ERROR: Specified path is a directory, not a file\n")
        return false
    }

    fmt.Println(info.Name() + " found!")
    return true

}

func isSameFormat(originalFormat string, desiredFormat string) bool {

    if originalFormat == desiredFormat {
        return true
    }

    if sameFormats[originalFormat] == desiredFormat {
        return true
    }

    return false
}

func main() {
    file := flag.String("file", "", "Original file to convert") 
    desiredFormat := flag.String("format", "jpeg", "Desired format to convert original file to")

    flag.Parse()

    if len(os.Args) < 3 {
        printError("Usage: frmt -file=<FILE_PATH> -format=<DESIRED_FORMAT>\n")
        os.Exit(1)
    }

    *desiredFormat = strings.ToLower(*desiredFormat)

    path, err := filepath.Abs(*file)
    if err != nil {
        errMsg := fmt.Sprintf("PATH ERROR: %s", err)
        printError(errMsg)
        os.Exit(2)
    }

    validFile := fileExists(path)

    if !validFile {
        os.Exit(3)
    }

    sameFormat := isSameFormat(filepath.Ext(path), *desiredFormat)

    if sameFormat {
        printError("FORMAT ERROR: Original file is already present in desired format\n")
        os.Exit(4)
    }

    desiredFile := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + "." + *desiredFormat

    outputFile, err := os.Create(desiredFile) 
    defer outputFile.Close()

    src, err := imgconv.Open(path)
    if err != nil {
        errMsg := fmt.Sprintf("Failed to open file: %s", err)
        printError(errMsg)
        os.Exit(5)
    }

    fileErr := false

    switch *desiredFormat {
    case "png":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.PNG})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "jpg":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.JPEG})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "jpeg":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.JPEG})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "tiff":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.TIFF})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "tif":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.TIFF})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "bmp":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.BMP})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    case "pdf":
        err = imgconv.Write(outputFile, src, &imgconv.FormatOption{Format: imgconv.PDF})
        if err != nil {
            errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
            printError(errMsg)
            os.Exit(6)
        }
        break
    default:
        os.Remove(outputFile.Name())
        errMsg := fmt.Sprintf("%s format is not supported :(\n", *desiredFormat)
        printError(errMsg)
        fileErr = true
        break
    }

    if !fileErr {
        pwd, err := os.Getwd()
        if err != nil {
            errMsg := fmt.Sprintf("PWD ERROR: %s\n", err)
            printError(errMsg)
            os.Exit(7)
        }
        fmt.Printf("File converted! Destination: %s\n", pwd)
    }
}
