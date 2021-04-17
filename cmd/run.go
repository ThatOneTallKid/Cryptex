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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"net/url"
	"time"
	"github.com/olekukonko/tablewriter"
	
	"github.com/spf13/cobra"
)



type Coin struct {
	At     int64  `json:"at"`    
	Ticker Ticker `json:"ticker"`
}

type Ticker struct {
	Buy  string `json:"buy"` 
	Sell string `json:"sell"`
	Low  string `json:"low"` 
	High string `json:"high"`
	Last string `json:"last"`
	Vol  string `json:"vol"` 
}


func runner() {
	var str = `

	░█████╗░██████╗░██╗░░░██╗██████╗░████████╗███████╗██╗░░██╗
	██╔══██╗██╔══██╗╚██╗░██╔╝██╔══██╗╚══██╔══╝██╔════╝╚██╗██╔╝
	██║░░╚═╝██████╔╝░╚████╔╝░██████╔╝░░░██║░░░█████╗░░░╚███╔╝░
	██║░░██╗██╔══██╗░░╚██╔╝░░██╔═══╝░░░░██║░░░██╔══╝░░░██╔██╗░
	╚█████╔╝██║░░██║░░░██║░░░██║░░░░░░░░██║░░░███████╗██╔╝╚██╗
	░╚════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░░░░░░░╚═╝░░░╚══════╝╚═╝░░╚═╝v1.0.0	



Copyright 2021, Aditya Das
	`

	fmt.Println(Yellow+str+Reset)
}


func parser(coin string) {
	runner()
	coinname  := url.QueryEscape(coin)
	url := fmt.Sprintf("https://api.wazirx.com/api/v2/tickers/%s", coinname);
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var record Coin
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	data := [][]string {
		[]string{record.Ticker.Buy,record.Ticker.Sell,record.Ticker.Low,record.Ticker.High,record.Ticker.Last,record.Ticker.Vol },
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Buy", "Sell", "Low","High","Last","Vol" })
	for _, v := range data {
		table.Append(v)
	}
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgHiRedColor})
	fmt.Println(Green+"[*] ",coinname, " ( Current price:",record.Ticker.Buy,")"+Reset )
	table.Render()
	dt := time.Now()
	fmt.Println(Blue+"Fetched at : ", dt.Format("01-02-2006 15:04:05")+Reset)
}


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [arg]",
	Short: "runs the cli",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		var mine string = args[0]
		//checker implementation
		sup := []string {"dogeinr","ethinr", "wrxinr","usdtinr","bttinr","xrpinr","btcinr"}
		var flag bool = stringInSlice( mine,sup)
		
		if flag == true {
			parser(mine)
		} else {
			fmt.Println(Red+"Unsupported Coin : use 'cryptex list' to list all supported coins"+Reset)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
