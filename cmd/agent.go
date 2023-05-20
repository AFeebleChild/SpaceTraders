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
		if CallSign == "" {
			fmt.Println("CallSign is required")
			return
		}
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

var getAgentCmd = &cobra.Command{
	Use:   "get-agent",
	Short: "Get an agent",
	Run: func(cmd *cobra.Command, args []string) {
		agent, err := lib.GetAgent(CallSign)
		if err != nil {
			panic(err)
		}

		a, _ := json.Marshal(agent)
		lib.JsonPrettyPrint(a)

		agent.Save()
	},
}

var getContractsCmd = &cobra.Command{
	Use:   "get-contracts",
	Short: "Get all contracts",
	Run: func(cmd *cobra.Command, args []string) {
		contracts, err := lib.GetContracts(CallSign)
		if err != nil {
			panic(err)
		}

		c, _ := json.Marshal(contracts)
		lib.JsonPrettyPrint(c)
	},
}

var getLocationCmd = &cobra.Command{
	Use:   "get-location",
	Short: "Get location",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := lib.NewClientFromCallsign(CallSign)
		if err != nil {
			panic(err)
		}

		agent, err := lib.LoadAgent(CallSign)
		if err != nil {
			panic(err)
		}
		split := strings.Split(agent.Headquarters, "-")

		system := split[0] + "-" + split[1]
		waypoint := agent.Headquarters

		body, err := lib.HandleResp(client.GetWaypoint(context.TODO(), system, waypoint))
		if err != nil {
			panic(err)
		}
		lib.JsonPrettyPrint(body)
	},
}

var getWaypointsCmd = &cobra.Command{
	Use:   "get-waypoints",
	Short: "Get all system waypoints",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := lib.NewClientFromCallsign(CallSign)
		if err != nil {
			panic(err)
		}

		agent, err := lib.LoadAgent(CallSign)
		if err != nil {
			panic(err)
		}
		split := strings.Split(agent.Headquarters, "-")
		system := split[0] + "-" + split[1]

		params := &lib.GetSystemWaypointsParams{}

		body, err := lib.HandleResp(client.GetSystemWaypoints(context.TODO(), system, params))
		if err != nil {
			panic(err)
		}
		lib.JsonPrettyPrint(body)
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	// agentCmd.AddCommand(newAgentCmd)
	agentCmd.AddCommand(getAgentCmd)
	agentCmd.AddCommand(getContractsCmd)
	agentCmd.AddCommand(getLocationCmd)
	agentCmd.AddCommand(getWaypointsCmd)

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
