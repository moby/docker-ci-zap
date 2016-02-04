// +build windows
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microsoft/hcsshim"
)

func folderexists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	var folder string
	flag.StringVar(&folder, "folder", "", "Folder to zap.")
	flag.Parse()
	if folder == "" {
		fmt.Println("Error: folder must be supplied")
		return
	}
	if folderexists(folder) {
		location, foldername := filepath.Split(folder)
		info := hcsshim.DriverInfo{
			HomeDir: location,
			Flavour: 0,
		}
		err := hcsshim.DestroyLayer(info, foldername)
		fmt.Println(err)
	} else {
		fmt.Println("Folder does not exist")
	}
}
