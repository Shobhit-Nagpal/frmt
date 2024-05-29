package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/Shobhit-Nagpal/frmt/pkg/convert"
	"github.com/Shobhit-Nagpal/frmt/pkg/utils"
)

func main() {
    path := flag.String("p", "", "Path of original file to convert") 
    desiredFormat := flag.String("f", "jpeg", "Desired format to convert original file to")

    flag.Parse()

    if len(os.Args) < 3 {
        utils.PrintError("Usage: frmt -p <FILE_PATH> -f <DESIRED_FORMAT>\n")
        os.Exit(1)
    }

    *desiredFormat = strings.ToLower(*desiredFormat)

    absPath, err := utils.ConvertToAbsolutePath(*path)
    if err != nil {
      os.Exit(2)
    }

    validFile := utils.FileExists(absPath)

    if !validFile {
        os.Exit(3)
    }

    sameFormat := utils.IsSameFormat(filepath.Ext(absPath), *desiredFormat)

    if sameFormat {
        utils.PrintError("FORMAT ERROR: Original file is already present in desired format\n")
        os.Exit(4)
    }

    desiredFile := strings.TrimSuffix(filepath.Base(absPath), filepath.Ext(absPath)) + "." + *desiredFormat

    outputFile, err := os.Create(desiredFile) 
    defer outputFile.Close()

    convert.ConvertMedia(absPath, *desiredFormat, outputFile)
}
