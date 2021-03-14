// Package cmd はツールの実際の処理が含まれる
package cmd

import "strings"

// FilePath はファイルパスを表す型
type FilePath string

// NewFilePath は型のコンストラクタ
func NewFilePath(s string) (FilePath, error) {
	// Todo: パスとして妥当かチェックしたい
	return FilePath(s), nil
}

// NewFilePathEmpty は型の空値を返す
func NewFilePathEmpty() FilePath {
	return FilePath("")
}

func (f FilePath) join(s ...string) (FilePath, error) {
	slicePath := append([]string{string(f)}, s...)

	joinPath := strings.Join(slicePath, "")

	finalPath, err := NewFilePath(joinPath)
	if err != nil {
		return "", err
	}

	return finalPath, nil
}
