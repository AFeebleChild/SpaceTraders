package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/afeeblechild/SpaceTraders/lib"
	"github.com/spf13/cobra"
)

var (
	CallSign string
	Faction  string
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agent called")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		CallSign = strings.ToUpper(CallSign)
	},
}

// var newAgentCmd = &cobra.Command{
// 	Use:   "new",
// 	Short: "Create a new agent",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if CallSign == "" {
// 			fmt.Println("CallSign is required")
// 			return
// 		}
// 		if Faction == "" {
// 			fmt.Println("Faction is required")
// 			return
// 		}
// 		newAgentData := lib.NewAgent(CallSign, Faction)
// 		agent := newAgentData.Data.Agent
// 		agent.Token = newAgentData.Data.Token
// 		err := agent.Save()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Agent created\n\n")
// 		out, err := json.MarshalIndent(newAgentData, "", "  ")
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(out))
// 	},
// }

type (
	AgentResp struct {
		Agent lib.Agent `json:"data"`
	}
)

var viewAgentCmd = &cobra.Command{
	Use:   "view",
	Short: "View an agent",
	Run: func(cmd *cobra.Command, args []string) {
		if CallSign == "" {
			fmt.Println("CallSign is required")
			return
		}
		client, agent := lib.NewClientFromCallsign(CallSign)

		body, err := lib.HandleResp(client.GetMyAgent(context.TODO()))
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

		a := AgentResp{}
		err = json.Unmarshal(body, &a)
		if err != nil {
			panic(err)
		}
		a.Agent.Token = agent.Token
		a.Agent.Save()

		// agent.Save()
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// agentCmd.AddCommand(newAgentCmd)
	agentCmd.AddCommand(viewAgentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	agentCmd.PersistentFlags().StringVarP(&CallSign, "callsign", "c", "", "Callsign of the agent")
	agentCmd.PersistentFlags().StringVarP(&Faction, "faction", "f", "", "Faction of the agent")
}
