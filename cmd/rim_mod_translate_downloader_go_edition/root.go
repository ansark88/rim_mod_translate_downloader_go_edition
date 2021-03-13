// Package cmd はツールの実際の処理が含まれる。root.goはその起点であり
// CLI作成ライブラリのcobraの処理が中心
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rim_mod_translate_downloader_go_edition",
	Short: "rim_mod_translate_downloader",
	Long:  "Download the Japanese translation patch for the Rimworld Mod to the appropriate filepath",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			url := args[0]

			// ダウンロード処理
			userpath := NewUserPath()
			downloader := NewDownloader(userpath, url)
			_, err := downloader.download()
			if err != nil {
				cmd.Println("Download Error!!!", err)
			}

		} else {
			cmd.Println("No input url!") // stderrに出る
		}
	},
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
