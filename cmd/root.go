/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tayfun8/scanwp/dirb"
	"github.com/tayfun8/scanwp/scraper"
	"github.com/tayfun8/scanwp/utils"
	"github.com/tayfun8/scanwp/wpscanapi"
)

var (
	target   string
	wordlist string
	scanwp   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scanwp",
	Short: "Wordpress Vuln Scan Tool",
	Long: `

    _____________  _  __          
   / __/ ___/ _ | / |/ /    _____ 
  _\ \/ /__/ __ |/    / |/|/ / _ \
 /___/\___/_/ |_/_/|_/|__,__/ .__/
			    /_/			  
  
Developed By Tayfun Ürkut
Github: github.com/tayfun8				
`,
	Example: "./scanwp -t http://test.com/ -w /path/to/wordlist -s",
	Run: func(cmd *cobra.Command, args []string) {
		target := utils.UrlFormat(target)
		scraper.Scrape(target)
		dirb.Dirb(wordlist, target)
		if scanwp == true {
			wpscanapi.Scan()
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&scanwp, "scanwp", "s", true, "Scanning parameter")
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "Target Scan")
	rootCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "wordlist/test.txt", "Wordlist")
	rootCmd.MarkFlagRequired("target")
}
