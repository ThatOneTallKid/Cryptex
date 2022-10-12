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
	
	"fmt"
	"os"
	"github.com/spf13/cobra"

)

var coin string

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"







var rootCmd = &cobra.Command{
	Use:   "cryptex",
	Short: ``,
	Long: `

	░█████╗░██████╗░██╗░░░██╗██████╗░████████╗███████╗██╗░░██╗
	██╔══██╗██╔══██╗╚██╗░██╔╝██╔══██╗╚══██╔══╝██╔════╝╚██╗██╔╝
	██║░░╚═╝██████╔╝░╚████╔╝░██████╔╝░░░██║░░░█████╗░░░╚███╔╝░
	██║░░██╗██╔══██╗░░╚██╔╝░░██╔═══╝░░░░██║░░░██╔══╝░░░██╔██╗░
	╚█████╔╝██║░░██║░░░██║░░░██║░░░░░░░░██║░░░███████╗██╔╝╚██╗
	░╚════╝░╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░░░░░░░╚═╝░░░╚══════╝╚═╝░░╚═╝v1.0.0	
	

A command line Crypto-currency ticker made using golang and Wazirx Api to help keep you updated.


COPYRIGHT 2021, Aditya Das
	`,

	// Run: func(cmd *cobra.Command, args []string) { },
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }
}

func init() {
	//cobra.OnInitialize(initConfig)


	//rootCmd.Flags().StringVar(&coin, "run", "Unsupported Coin", "runs the CLI",)

	
}


