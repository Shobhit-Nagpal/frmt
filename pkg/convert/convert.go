package convert

import (
	"fmt"
	"github.com/Shobhit-Nagpal/frmt/pkg/utils"
	"github.com/sunshineplan/imgconv"
	"os"
)

func ConvertMedia(path string, destFormat string, destFile *os.File) {

	src, err := imgconv.Open(path)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to open file: %s", err)
		utils.PrintError(errMsg)
		os.Exit(5)
	}

	fileErr := false

	switch destFormat {
	case "png":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.PNG})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "jpg":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.JPEG})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "jpeg":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.JPEG})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "tiff":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.TIFF})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "tif":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.TIFF})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "bmp":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.BMP})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	case "pdf":
		err = imgconv.Write(destFile, src, &imgconv.FormatOption{Format: imgconv.PDF})
		if err != nil {
			errMsg := fmt.Sprintf("Failed to write file: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(6)
		}
		break
	default:
		os.Remove(destFile.Name())
		errMsg := fmt.Sprintf("%s format is not supported :(\n", destFormat)
		utils.PrintError(errMsg)
		fileErr = true
		break
	}

	if !fileErr {
		pwd, err := os.Getwd()
		if err != nil {
			errMsg := fmt.Sprintf("PWD ERROR: %s\n", err)
			utils.PrintError(errMsg)
			os.Exit(7)
		}
		fmt.Printf("File converted! Destination: %s\n", pwd)
	}
}
