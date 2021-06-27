// Package cmd はツールの実際の処理が含まれる。root.goはその起点であり
// CLI作成ライブラリのcobraの処理が中心
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rim_mod_translate_downloader_go_edition",
	Short: "rim_mod_translate_downloader",
	Long:  "Download the Japanese translation patch for the Rimworld Mod to the appropriate filepath",

	Run: func(cmd *cobra.Command, args []string) {
		lenarg := len(args)

		// TODO: 複数URL対応をしたい
		if lenarg > 0 {
			for i := 0; i < lenarg; i++ {
				wg.Add(1)
				go mainDownload(args[i], cmd)
			}

			wg.Wait()
			cmd.Print("Done !!!!!!!")
		} else {
			cmd.PrintErrln("No input URL!!!") // stderrに出る
		}
	},
}

func mainDownload(newurl string, cmd *cobra.Command) {
	url, err := NewURL(newurl)
	defer wg.Done()

	if err != nil {
		cmd.PrintErrln("Invalid URL!!!", err)
	}

	// ダウンロード処理
	userpath := NewUserPath()
	downloader := NewDownloader(userpath, url)
	err = downloader.download()
	if err != nil {
		cmd.PrintErrln("Download Error!!!", err)
	}

	// 解凍処理
	destDir, err := NewFilePath(filepath.Dir(downloader.destPath.String()))
	if err != nil {
		cmd.PrintErrln("Dest Directory Error!!!", err)
	}

	var archiver Archiver = NewZipArchiver(downloader.destPath, destDir)

	if err := archiver.extract(); err != nil {
		cmd.PrintErrln("Decompress Error!!!", err)
	} else {
		fmt.Println("Complete!!!")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rim_mod_translate_downloader_go_edition.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
