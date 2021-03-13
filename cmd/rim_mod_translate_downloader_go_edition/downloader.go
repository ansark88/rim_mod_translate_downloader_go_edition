// Package cmd はツールの実際の処理が含まれる
package cmd

import (
	"fmt"
)

// Downloader のDL処理に使う情報
type Downloader struct {
	userPath *UserPath
	url      string
	destPath FilePath
}

// NewDownloader はDownloader構造体のコンストラクタ
func NewDownloader(userPath *UserPath, url string) *Downloader {
	return &Downloader{userPath, url, FilePath{}}
}

func (d Downloader) fetch(url string, id string) error {
	return nil
}

func (d Downloader) download() (FilePath, error) {
	parseURL := d.url

	u, err := URLParse(parseURL)
	if err != nil {
		return FilePath{}, err
	}

	fetchURL := u.convertedURL
	id := u.id
	fmt.Println("fetchURL:", fetchURL)

	err = d.fetch(fetchURL, id)

	if err != nil {
		return FilePath{}, err
	}

	return d.destPath, nil
}
