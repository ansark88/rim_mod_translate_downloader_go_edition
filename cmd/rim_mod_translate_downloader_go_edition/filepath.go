// Package cmd はツールの実際の処理が含まれる
package cmd

// FilePath はファイルパスを表す型
type FilePath string

// NewFilePath は型のコンストラクタ
func NewFilePath(s string) (FilePath, error) {
	return FilePath(s), nil
}

// NewFilePathEmpty は型の空値を返す
func NewFilePathEmpty() FilePath {
	return FilePath("")
}
