package cmd

import (
	"errors"
	"testing"
)

func TestURLParse(t *testing.T) {
	// 参考: https://future-architect.github.io/articles/20200601/
	tests := []struct {
		name  string
		input URL
		want  URL
	}{
		{
			name:  "normal_test1",
			input: URL("https://rimworld.2game.info/jp_download.php?file_id=1267&id=2194097170"),
			want:  URL("https://img.2game.info/re_archive/l/rimworld/files/up_japanese/2194097170/1267.zip"),
		},
		{
			name:  "normal_test2",
			input: URL("https://rimworld.2game.info/jp_download.php?file_id=1164&id=2354938860"),
			want:  URL("https://img.2game.info/re_archive/l/rimworld/files/up_japanese/2354938860/1164.zip"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := URLParse(tt.input); got.convertedURL != tt.want {
				t.Errorf("URLParse() = %v, want %v", got, tt.want)
			}
		})
	}

	testsFailure := []struct {
		name  string
		input URL
		want  error
	}{
		{
			name:  "failure_test1",
			input: URL("1234"),
			want:  errors.New("id or file_id is None"),
		},
		{
			name:  "failure_test2",
			input: URL("https://rimworld.2game.info/jp_download.php"),
			want:  errors.New("id or file_id is None"),
		},
	}
	for _, tt := range testsFailure {
		t.Run(tt.name, func(t *testing.T) {
			_, err := URLParse(tt.input)

			if err.Error() != tt.want.Error() {
				t.Errorf("URLParse() ErrorMessage = %v, want %v", err, tt.want)
			}
		})
	}
}
