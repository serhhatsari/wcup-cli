package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "encoding/json"
	"os"
	"strconv"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AllGroups struct {
	Groups []struct {
		Letter string `json:"letter"`
		Teams  []struct {
			Country          string `json:"country"`
			Name             string `json:"name"`
			GroupLetter      string `json:"group_letter"`
			GroupPoints      int    `json:"group_points"`
			Wins             int    `json:"wins"`
			Draws            int    `json:"draws"`
			Losses           int    `json:"losses"`
			GamesPlayed      int    `json:"games_played"`
			GoalsFor         int    `json:"goals_for"`
			GoalsAgainst     int    `json:"goals_against"`
			GoalDifferential int    `json:"goal_differential"`
		} `json:"teams"`
	} `json:"groups"`
}

var flagVerbose bool
var flagGroup string
var availableGroups = []string{"A", "B", "C", "D", "E", "F", "G", "H", "a", "b", "c", "d", "e","f","g","h"}

var cmdRoot = &cobra.Command{
	Use:   "wcup",
	Short: "wcup CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

var cmdHelp = &cobra.Command{
	Use:   "--help",
	Short: "-h",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Help")
	},
}

var cmdGroups = &cobra.Command{
	Use:   "groups",
	Short: "groups",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://worldcupjson.net/teams")
		if err != nil {
		   log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		   log.Fatalln(err)
		}
		sb := string(body)

		var groups AllGroups
		json.Unmarshal([]byte(sb), &groups)


	for _, group := range groups.Groups {

		if flagGroup != "" && ( flagGroup != group.Letter  && flagGroup != strings.ToLower(group.Letter) ){
			// check if the group is not in the available groups
			found := false
			for _, g := range availableGroups {
				if g == flagGroup {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("WARNING!\nGroup " + flagGroup + " is not a valid group.")
				os.Exit(1)
			}

			continue
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "PTS"})
		
		// colorize table
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},

		)

		// colorize columns
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
	)

		// sort group by points
		sort.Slice(group.Teams, func(i, j int) bool {
			return group.Teams[i].GroupPoints > group.Teams[j].GroupPoints
		})

		for i, team := range group.Teams {
			if team.Name == "Uruguay" {
				group.Teams[i], group.Teams[i+1] = group.Teams[i+1], group.Teams[i]
				break
			}
		}

		for i, team := range group.Teams {
			if team.Name == "Mexico" {
				group.Teams[i], group.Teams[i+1] = group.Teams[i+1], group.Teams[i]
				break
			}
		}


		var i int
		i = 0
		for _, team := range group.Teams {
			i++
			table.Append([]string{strconv.Itoa(i) + " " + team.Name, strconv.Itoa(team.GroupPoints)})
		}

		fmt.Println(color.HiMagentaString("\nGroup " + group.Letter))

		table.Render()
	}
	},
}

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "status",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://worldcupjson.net/teams")
		if err != nil {
		   log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		   log.Fatalln(err)
		}
		sb := string(body)

		var groups AllGroups
		json.Unmarshal([]byte(sb), &groups)

		table := tablewriter.NewWriter(os.Stdout)
		var i int
		i = 0
	for _, group := range groups.Groups {

		
		table.SetHeader([]string{"Name", "G", "W",   "D", "L", "GF", "GA", "A","Group PTS"})
		
		// colorize table
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		)

		// colorize columns
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor},
		)


		for _, team := range group.Teams {
			i++
			table.Append([]string{strconv.Itoa(i) + " " + team.Name,strconv.Itoa(team.GamesPlayed),strconv.Itoa(team.Wins),  strconv.Itoa(team.Draws), strconv.Itoa(team.Losses),  strconv.Itoa(team.GoalsFor), strconv.Itoa(team.GoalsAgainst), strconv.Itoa(team.GoalDifferential), strconv.Itoa(team.GroupPoints)})
		}

	}
	table.Render()
	},
}

func Execute() error {
	cmdRoot.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output")

	cmdRoot.MarkPersistentFlagRequired("port")
	//if err := cmdRoot.Execute(); err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
		
	

	cmdGroups.PersistentFlags().StringVarP(&flagGroup, "group", "g", "", "group name")


	cmdRoot.AddCommand(cmdHelp)
	cmdRoot.AddCommand(cmdStatus)
	cmdRoot.AddCommand(cmdGroups)

	return cmdRoot.Execute()

}

func init() {
	// Flags
	//cmdRoot.PersistentFlags().Int16P("port", "p", 8080, "listening port")

	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("wcup")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	// Set flags
	//viper.BindPFlag("port", cmdRoot.PersistentFlags().Lookup("port"))

	// Set undefined variables
	viper.SetDefault("version", "0.1.0")

	// If config file exists
	//if err := viper.ReadInConfig(); err != nil {
	//	fmt.Println("fatal error config file: default \n", err)
	//	os.Exit(1)
	//}

}
