package status

import (
	"io/ioutil"
	"log"
	"net/http"
    "encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"wcup/cmd/utils/models"
)


var FlagCountry string


func showStatus(cmd *cobra.Command, args []string) {
	resp, err := http.Get("https://worldcupjson.net/teams")
	if err != nil {
	   log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	sb := string(body)

	var groups models.AllGroups
	json.Unmarshal([]byte(sb), &groups)
	
	
	table := tablewriter.NewWriter(os.Stdout)
	var i int
	i = 1
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
	if FlagCountry == team.Name  || FlagCountry == strings.ToLower(team.Name) {
		table.Append([]string{strconv.Itoa(i) + " " + team.Name,strconv.Itoa(team.GamesPlayed),strconv.Itoa(team.Wins),  strconv.Itoa(team.Draws), strconv.Itoa(team.Losses),  strconv.Itoa(team.GoalsFor), strconv.Itoa(team.GoalsAgainst), strconv.Itoa(team.GoalDifferential), strconv.Itoa(team.GroupPoints)})
		break
	} else if FlagCountry != ""{
		i++
	} else{
		table.Append([]string{strconv.Itoa(i) + " " + team.Name,strconv.Itoa(team.GamesPlayed),strconv.Itoa(team.Wins),  strconv.Itoa(team.Draws), strconv.Itoa(team.Losses),  strconv.Itoa(team.GoalsFor), strconv.Itoa(team.GoalsAgainst), strconv.Itoa(team.GoalDifferential), strconv.Itoa(team.GroupPoints)})
		i++
	}
	
}

// TODO: print message if country not found

}
table.Render()
}


var CmdStatus = &cobra.Command{
	Use:   "status",
	Short: "status",
	Run: showStatus,
}
