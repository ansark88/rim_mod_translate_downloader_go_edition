// Package cmd はツールの実際の処理が含まれる
package cmd

import (
	"errors"
	"fmt"
	"net/url"
)

// ParseValue にはURLのクエリ部分をパースした結果が格納される
type ParseValue struct {
	convertedURL URL
	id           string
}

// URLParse は標準ライブラリurl.Parseのラッパーとなっている。URLのクエリを元にダウンロード用のURLとIDを返す
func URLParse(input URL) (ParseValue, error) {
	u, err := url.Parse(string(input))
	if err != nil {
		return ParseValue{}, errors.New("url parse error")
	}

	q := u.Query()
	id := q.Get("id")
	fileID := q.Get("file_id")

	if len(id) == 0 || len(fileID) == 0 {
		return ParseValue{}, errors.New("id or file_id is None")
	}

	convertedURL, nil := convert(id, fileID)
	if err != nil {
		return ParseValue{}, errors.New("url convert error")
	}

	return ParseValue{convertedURL, id}, nil
}

// リンクのURLに対する実際のzipファイルのURLは以下のようになるので、変換する必要がある
// https://img.2game.info/re_archive/l/rimworld/files/up_japanese/2205980094/935.zip
func convert(id string, fileID string) (URL, error) {
	u := fmt.Sprintf("https://img.2game.info/re_archive/l/rimworld/files/up_japanese/%s/%s.zip", id, fileID)

	newURL, err := NewURL(u)
	if err != nil {
		return "", err
	}

	return newURL, nil
}
