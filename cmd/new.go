/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/aifaniyi/bootstrap/pkg/generator"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new golang web application",
	Long: `Create a new golang application from the specified app.spec.json file.
	
For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		project, _ := cmd.Flags().GetString("project")
		output, _ := cmd.Flags().GetString("output")

		if input == "" {
			fmt.Printf("\ninput spec file must be provided\n\n")
			return
		}

		if project == "" {
			fmt.Printf("\nproject name must be provided\n\n")
			return
		}

		newProject(input, project, output)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().StringP("input", "i", "", "input filename e.g app.spec.json")
	newCmd.Flags().StringP("project", "p", "", "project name")
	newCmd.Flags().StringP("output", "o", "", "output directory name. Defaults to current directory")
}

func newProject(input, project, output string) {
	err := generator.GenerateGolang(input, output, project)
	if err != nil {
		fmt.Println(err)
	}
}
