// Package cmd はツールの実際の処理が含まれる
package cmd

import (
	"errors"
	"net/url"
)

// URL を表す型
type URL string

// NewURL は型のコンストラクタ。url.Parserを通すので不正なURLは作れない
func NewURL(s string) (URL, error) {
	_, err := url.Parse(s)
	if err != nil {
		return "", errors.New("url parse error")
	}

	return URL(s), nil
}

// NewURLEmpty は型の空値を返す
func NewURLEmpty() URL {
	return URL("")
}
