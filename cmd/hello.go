/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name2 string
	name3 string
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Helloと出力します",
	Long:  `このコマンドはHelloと出力するシンプルなコマンドです。`,
	Run: func(cmd *cobra.Command, args []string) {
		if name1, err := cmd.PersistentFlags().GetString("name1"); err == nil {
			fmt.Println("Hello, ", name1)
		}

		fmt.Printf("Hello, %s!\n", name2)
		fmt.Printf("Hello, %s!\n", name3)
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// String(フラグ名, デフォルト値, 説明文)
	helloCmd.Flags().String("name1", "Cobra1", "フラグを設定しました")

	// StringVar(変数ポインタ, フラグ名, デフォルト値, 説明文)
	helloCmd.Flags().StringVar(&name2, "name2", "Cobra2", "挨拶の対象となる名前")
	// 関数名の最後が"P"だと短縮フラグを指定できる
	// StringVarP(変数ポインタ, フラグ名, 短縮フラグ名, デフォルト値, 説明文)
	helloCmd.Flags().StringVarP(&name3, "name3", "n", "Cobra3", "挨拶の対象となる名前")
}
