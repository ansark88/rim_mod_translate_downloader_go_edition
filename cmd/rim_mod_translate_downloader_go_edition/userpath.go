// Package cmd はツールの実際の処理が含まれる
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// UserPath にはRimworldのワークショップフォルダのパスが格納される
type UserPath struct {
	workshopDir FilePath
}

// NewUserPath はUserPath構造体のコンストラクタ
func NewUserPath() *UserPath {
	workshopDir := decideWorkshopDir()
	return &UserPath{workshopDir}
}

func decideWorkshopDir() FilePath {
	var path string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	switch runtime.GOOS {
	case "windows":
		path = `C:\Program Files (x86)\Steam\steamapps\workshop\content\294100`
	case "darwin":
		path = filepath.Join(homeDir, "Library/Application Support/Steam/steamapps/workshop/content/294100")
	default:
		path = "/hoge" // 未対応
	}

	fmt.Println("workshopDir:", path)

	return FilePath(path)
}
