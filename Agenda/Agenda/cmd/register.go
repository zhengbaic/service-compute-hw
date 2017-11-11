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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "regist for a account",
	Long:  `regist for a account`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() < 4 {
			fmt.Println("you must provide --user --password --email --mobilephone")
			AgendaLog.OperateLog("[error]", "register error => "+"you don't provide all flags --user --password --email --mobilephone")
			return
		}
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("mobilephone")
		body := User.User{username, password, phone, email}
		err := User.UserRegitser(body)
		if err != nil {
			fmt.Println(err.Error())
			AgendaLog.OperateLog("[error]", "register error => "+err.Error())
		} else {
			fmt.Println("register successfully")
			AgendaLog.OperateLog("[info]", "register successfully")
		}
	},
	Args: cobra.ExactArgs(0),
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "Use this name to register a account")
	registerCmd.Flags().StringP("password", "p", "weimumu123", "User this password to login later")
	registerCmd.Flags().StringP("email", "e", "934657014@qq.com", "User's emial")
	registerCmd.Flags().StringP("mobilephone", "m", "15692434434", "User's mobilephone")
}
