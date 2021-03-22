package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ZipArchiver のDL処理に使う情報
type ZipArchiver struct {
	zipFile     FilePath
	destDirPath FilePath
}

// NewZipArchiver はDownloader構造体のコンストラクタ
func NewZipArchiver(filePath FilePath, destDirPath FilePath) *ZipArchiver {
	return &ZipArchiver{filePath, destDirPath}
}

// 参考: https://golangcode.com/unzip-files-in-go/
func (z ZipArchiver) extract() error {
	var destDirPath string = z.destDirPath.String()

	r, err := zip.OpenReader(z.zipFile.String())
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(destDirPath, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(destDirPath)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// deferだとforループが終わるまでcloseされないので1回ごとに閉じてる
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}

		fmt.Println("Unzip success: ", fpath)
	}

	return nil
}
