package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var sameFormats = map[string]string{
	"jpg":  "jpeg",
	"jpeg": "jpg",
	"png":  "apng",
	"apng": "png",
	"tiff": "tif",
	"tif":  "tiff",
	"bmp":  "dib",
	"dib":  "bmp",
}

func PrintError(err string) {
	fmt.Fprintf(os.Stderr, err)
}

func FileExists(path string) bool {

	info, err := os.Stat(path)
	if err != nil {
		errMsg := fmt.Sprintf("PATH ERROR: %s\n", err)
		PrintError(errMsg)
		return false
	}

	if info.IsDir() {
		PrintError("ERROR: Specified path is a directory, not a file\n")
		return false
	}

	fmt.Println(info.Name() + " found!")
	return true

}

func IsSameFormat(originalFormat string, desiredFormat string) bool {

	if originalFormat == desiredFormat {
		return true
	}

	if sameFormats[originalFormat] == desiredFormat {
		return true
	}

	return false
}

func ConvertToAbsolutePath(relPath string) (string, error) {
    absPath, err := filepath.Abs(relPath)
    if err != nil {
        errMsg := fmt.Sprintf("PATH ERROR: %s", err)
        PrintError(errMsg)
        return "", err
    }

    return absPath, nil
}
