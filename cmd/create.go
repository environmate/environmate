// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

  "environmate/libs/envutils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Creating env (%v)...", env))
    envutils.CreateEnv(env, key)
	},
}

func init() {
	createCmd.Flags().StringVarP(&env, "env", "e", "", "The name of the environment that will be created")
	createCmd.Flags().StringVarP(&key, "key", "k", "", "The key used to decrypt and encrypt the environments")
	createCmd.MarkFlagRequired("env")
	createCmd.MarkFlagRequired("key")
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
