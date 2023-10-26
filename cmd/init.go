/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		argss := cmd.Args
		fmt.Println(argss)

		if _, err := os.Stat(Config.InitFolder); !os.IsNotExist(err) {
			fmt.Printf("%v folder already exist \n", Config.InitFolder)
			return
		}

		if err := createInitAndObjectFolders(); err != nil {
			panic(err)
		}

		if err := createRefFile(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createInitAndObjectFolders() error {
	path := fmt.Sprintf("%v/objects", Config.InitFolder)
	return os.MkdirAll(path, os.ModePerm)
}

func createRefFile() error {
	path := fmt.Sprintf("./%v/ref.txt", Config.InitFolder)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("0")
	if err != nil {
		return err
	}

	return nil
}
