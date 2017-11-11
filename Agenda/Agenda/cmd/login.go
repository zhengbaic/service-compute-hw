// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/weimumu/Agenda/agenda/entity/AgendaLog"
	"github.com/weimumu/Agenda/agenda/entity/User"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to enter the system",
	Long:  `login to enter the system. Then you can has more operation on the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() < 2 {
			fmt.Println("you must provide --user --password")
			AgendaLog.OperateLog("[error]", "login error => "+"you don't provide all flags --user --password")
			return
		}
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		err := User.UserLogin(username, password)
		if err != nil {
			fmt.Println(err.Error())
			AgendaLog.OperateLog("[error]", "login error => "+err.Error())
		} else {
			fmt.Println("login successfully")
			AgendaLog.OperateLog("[info]", "login successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("user", "u", "Anonymous", "Use this name to login")
	loginCmd.Flags().StringP("password", "p", "weimumu123", "User this password to login")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
