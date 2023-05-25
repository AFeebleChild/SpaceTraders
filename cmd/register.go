package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/afeeblechild/SpaceTraders/lib"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new agent",
	Run: func(cmd *cobra.Command, args []string) {
		if CallSign == "" {
			fmt.Println("CallSign is required")
			return
		}
		if Faction == "" {
			fmt.Println("Faction is required")
			return
		}
		var err error
		Client, err = lib.NewClientBase()
		if err != nil {
			panic(err)
		}
		newAgentData, err := lib.NewAgent(Client, CallSign, Faction)
		if err != nil {
			panic(err)
		}
		err = newAgentData.Agent.Save()
		if err != nil {
			panic(err)
		}
		token := &lib.Token{
			Token:  newAgentData.Token,
			Symbol: newAgentData.Agent.Symbol,
		}
		err = token.Save()
		if err != nil {
			panic(err)
		}
		fmt.Println("Agent created\n\n")
		out, err := json.MarshalIndent(newAgentData, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.PersistentFlags().StringVarP(&CallSign, "callsign", "c", "", "Callsign of the agent")
	registerCmd.PersistentFlags().StringVarP(&Faction, "faction", "f", "", "Faction of the agent")
}
