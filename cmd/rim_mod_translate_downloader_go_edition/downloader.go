// Package cmd はツールの実際の処理が含まれる
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Downloader のDL処理に使う情報
type Downloader struct {
	userPath *UserPath
	url      URL
	destPath FilePath
}

// NewDownloader はDownloader構造体のコンストラクタ
func NewDownloader(userPath *UserPath, url URL) *Downloader {
	return &Downloader{userPath, url, NewFilePathEmpty()}
}

// 参考: https://golangcode.com/download-a-file-from-a-url/
func (d Downloader) fetch(url URL) error {
	// Get the data
	resp, err := http.Get(string(url))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	err = d.copy(resp.Body)

	return err
}

func (d Downloader) copy(reader io.ReadCloser) error {
	// Create the file
	out, err := os.Create(string(d.destPath))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	return err
}

func (d Downloader) download() (FilePath, error) {
	parseURL := d.url

	u, err := URLParse(parseURL)
	if err != nil {
		return NewFilePathEmpty(), err
	}

	fetchURL := u.convertedURL
	fmt.Println("fetchURL:", fetchURL)

	d.destPath, err = d.userPath.workshopDir.join(u.id, "download.zip")
	if err != nil {
		return NewFilePathEmpty(), err
	}

	err = d.fetch(fetchURL)
	if err != nil {
		return NewFilePathEmpty(), err
	}

	return d.destPath, nil
}
