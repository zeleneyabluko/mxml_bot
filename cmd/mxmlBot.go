/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

var (
	appVersion = "Version"
)

// mxmlBotCmd represents the mxmlBot command
var mxmlBotCmd = &cobra.Command{
	Use:        "mxmlBot",
	Aliases:    []string{"start"},
	SuggestFor: []string{},
	Short:      "A brief description of your command",
	GroupID:    "",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example:   "",
	ValidArgs: []string{},
	//ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// },
	//Args: func(cmd *cobra.Command, args []string) error {
	//},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            map[string]string{},
	Version:                "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	//PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
	//},
	//PreRun: func(cmd *cobra.Command, args []string) {
	//},
	//PreRunE: func(cmd *cobra.Command, args []string) error {
	//},
	Run: func(cmd *cobra.Command, args []string) {
		//var appVersion string
		fmt.Println("mxmlBot " + appVersion + " started!")

		mxmlBot, err := telebot.NewBot(telebot.Settings{URL: "", Token: TeleToken, Poller: &telebot.LongPoller{Timeout: 10 * time.Second}})
		if err != nil {
			log.Fatalf("Please check the TELE_TOKEN!", err)
			return
		}
		mxmlBot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			return err
		})
		mxmlBot.Start()
	},
	//RunE: func(cmd *cobra.Command, args []string) error {
	//},
	//PostRun: func(cmd *cobra.Command, args []string) {
	//},
	//PostRunE: func(cmd *cobra.Command, args []string) error {
	//},
	//PersistentPostRun: func(cmd *cobra.Command, args []string) {
	//},
	//PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
	//},
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}

func init() {
	rootCmd.AddCommand(mxmlBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mxmlBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mxmlBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
