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

// querymCmd represents the querym command
var querymCmd = &cobra.Command{
	Use:   "querym",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentUser := User.UserState()
		if currentUser == "" {
			AgendaLog.OperateLog("[error]", "querym error => "+"not login")
			return
		}
		if cmd.Flags().NFlag() < 2 {
			fmt.Println("you must provide --stime --etime")
			AgendaLog.OperateLog("[error]", "querym error => "+"you don't provide all flags --stime --etime")
			return
		}
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		stimedate, err1 := Meeting.StringToDate(stime)
		etimedate, err2 := Meeting.StringToDate(etime)
		if err1 != nil {
			fmt.Println("Time Parse Error! " + err1.Error())
			fmt.Println("you must use --stime like --stime=2016-01-01/21:00")
			AgendaLog.OperateLog("[error]", "querym error => "+err1.Error()+" Usage: you must use --stime like --stime=2016-01-01/21:00")
			return
		}
		if err2 != nil {
			fmt.Println("Time Parse Error! " + err2.Error())
			fmt.Println("you must use --etime like --stime=2016-01-01/21:00")
			AgendaLog.OperateLog("[error]", "querym error => "+err1.Error()+" Usage: you must use --etime like --etime=2016-01-01/21:00")
			return
		}
		meetings := Meeting.QueryMeetingByTime(currentUser, stimedate, etimedate)
		if len(meetings) != 0 {
			AgendaLog.OperateLog("[info]", "querym successfully")
			fmt.Println("The search result can be show as followed:")
			fmt.Println("Initiator	" + "Participator				" + "Title		" + "STime				" + "ETime			")
			for _, value := range meetings {
				fmt.Println(value.Initiator+"       ", value.Participator, "  "+value.Title+"      "+Meeting.DateToString(value.STime)+"          "+Meeting.DateToString(value.ETime))
			}
		} else {
			fmt.Println("No meeting match your search time")
		}
	},
}

func init() {
	RootCmd.AddCommand(querymCmd)
	querymCmd.Flags().StringP("stime", "s", "", "Meeting's start time, Use like as -s = 2006-01-02/15:04")
	querymCmd.Flags().StringP("etime", "e", "", "Meeting's end time, Use like as -e = 2006-01-02/15:04")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
