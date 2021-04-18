/*
Copyright © 2021 ADITYA DAS

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/olekukonko/tablewriter"
)


func tablemaker(coins [][]string) {
	tabler := tablewriter.NewWriter(os.Stdout)
	tabler.SetHeader([]string{"Coin-name" })
	tabler.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiRedColor})
	for _,v := range coins {
		tabler.Append(v)
	}
	
	tabler.Render()
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the supported coins by this CLI",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		coins := [][]string {
			[]string{"dogeinr"},
			[]string{"ethinr"},
			[]string{"wrxinr"},
			[]string{"usdtinr"},
			[]string{"bttinr"},
			[]string{"xrpinr"},
			[]string{"btcinr"},
	

		
		}
		fmt.Println(Green+"List of all supported coins"+Reset)
		tablemaker(coins)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)


}
