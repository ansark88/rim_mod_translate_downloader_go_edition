package cmd

// Archiver は圧縮ファイルを扱う構造体のインターフェイス(圧縮形式ごとに用意)
type Archiver interface {
	extract() error
}
