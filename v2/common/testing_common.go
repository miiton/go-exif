package exifcommon

import (
	"os"
	"path"

	"encoding/binary"
	"io/ioutil"

	"github.com/dsoprea/go-logging"
)

var (
	assetsPath        = ""
	testImageFilepath = ""

	testExifData = make([]byte, 0)

	// EncodeDefaultByteOrder is the default byte-order for encoding operations.
	EncodeDefaultByteOrder = binary.BigEndian

	// Default byte order for tests.
	TestDefaultByteOrder = binary.BigEndian
)

func GetModuleRootPath() string {
	moduleRootPath := os.Getenv("EXIF_MODULE_ROOT_PATH")
	if moduleRootPath != "" {
		return moduleRootPath
	}

	currentWd, err := os.Getwd()
	log.PanicIf(err)

	currentPath := currentWd
	visited := make([]string, 0)

	for {
		tryStampFilepath := path.Join(currentPath, ".MODULE_ROOT")

		_, err := os.Stat(tryStampFilepath)
		if err != nil && os.IsNotExist(err) != true {
			log.Panic(err)
		} else if err == nil {
			break
		}

		visited = append(visited, tryStampFilepath)

		currentPath = path.Dir(currentPath)
		if currentPath == "/" {
			log.Panicf("could not find module-root: %v", visited)
		}
	}

	return currentPath
}

func init() {
	moduleRootPath := GetModuleRootPath()
	assetsPath = path.Join(moduleRootPath, "assets")

	testImageFilepath = path.Join(assetsPath, "NDM_8901.jpg")

	// Load test EXIF data.

	filepath := path.Join(assetsPath, "NDM_8901.jpg.exif")

	var err error
	testExifData, err = ioutil.ReadFile(filepath)
	log.PanicIf(err)
}
