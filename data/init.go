package data

import (
	"os"
	"path/filepath"
)

//need later
// func CreateFileWithDirs(path string) (*os.File, error) {
// 	// Create parent directories if they don't exist
// 	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
// 		return nil, err
// 	}

// 	// Create or open the file
// 	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
// }

func GetDefaultDataPath() string {
	xdgData := os.Getenv("XDG_DATA_HOME")
	if xdgData == "" {
		home := os.Getenv("HOME")
		xdgData = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(xdgData, "beaver-task", "data.db")
}
