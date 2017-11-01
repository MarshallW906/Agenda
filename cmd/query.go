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

	"github.com/FideoJ/Agenda/service"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query [-uUsername] [-pPassword]",
	Short: "Query a User by Username and Password.",
	Long: `Query a User by Username and Password.

If the username exists and password is correct,
Agenda query will print that user.
specify -p without -u will be ignored.
If neither -u nor -p is specified, query will simply print all the registered users' username.`,

	Run: func(cmd *cobra.Command, args []string) {
		// log implemented in service/user.go
		fmt.Println("query called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		// log implemented in service/user.go
		fmt.Printf("query Users by Username:[%+v] and Password:[%+v]\n", username, password)
		service.Query(username, password)
	},
}

func init() {
	queryCmd.Flags().StringP("username", "u", "", "the target username")
	queryCmd.Flags().StringP("password", "p", "", "the related password")

	RootCmd.AddCommand(queryCmd)
}