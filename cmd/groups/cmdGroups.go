package groups

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

	"wcup/cmd/utils/models"
	"wcup/cmd/utils/constants"
)


var FlagGroup string

var CmdGroups = &cobra.Command{
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

		var groups models.AllGroups
		json.Unmarshal([]byte(sb), &groups)


	for _, group := range groups.Groups {

		if FlagGroup != "" && ( FlagGroup != group.Letter  && FlagGroup != strings.ToLower(group.Letter) ){
			// check if the group is not in the available groups
			found := false
			for _, g := range constants.AvailableGroups {
				if g == FlagGroup {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("WARNING!\nGroup " + FlagGroup + " is not a valid group.")
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


