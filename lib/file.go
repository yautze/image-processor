package lib

import (
	"os"
	"path/filepath"
)

// FileInfo -
type FileInfo struct {
	// Path -
	Path string

	// Name -
	Name string
}

// FileProducer -
func FileProducer(filePath string, out chan *FileInfo) {
	filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		out <- &FileInfo{
			Path: path,
			Name: f.Name(),
		}

		return nil
	})

	close(out)
}

// GetInputImgPaths - 取得全部輸入圖片路徑
func GetInputImgPaths(filePath string) []*FileInfo {
	imagePaths := make([]*FileInfo, 0)

	filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		imagePaths = append(imagePaths, &FileInfo{
			Path: path,
			Name: f.Name(),
		})
		return nil
	})

	return imagePaths
}
