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
	"github.com/weimumu/Agenda/agenda/entity/Meeting"
	"github.com/weimumu/Agenda/agenda/entity/User"
)

// delpCmd represents the delp command
var delpCmd = &cobra.Command{
	Use:   "delp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentUser := User.UserState()
		if currentUser == "" {
			AgendaLog.OperateLog("[error]", "delp error => "+"not login")
			return
		}
		if cmd.Flags().NFlag() < 2 {
			fmt.Println("you must provide --title --participator")
			AgendaLog.OperateLog("[error]", "you don't provide all flags --title --participator")
			return
		}
		participator, _ := cmd.Flags().GetStringSlice("participator")
		title, _ := cmd.Flags().GetString("title")
		err := Meeting.DeleteParticipators(currentUser, title, participator)
		if err != nil {
			fmt.Println(err.Error())
			AgendaLog.OperateLog("[error]", "delp error => "+err.Error())
		} else {
			fmt.Println("delete participators successfully")
			AgendaLog.OperateLog("[info]", "delete participators successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(delpCmd)
	delpCmd.Flags().StringP("title", "t", "", "Meeting's title for a specific meeting")
	delpCmd.Flags().StringSliceP("participator", "p", make([]string, 0), "New participator to add")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
