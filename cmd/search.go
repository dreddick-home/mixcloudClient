/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"runtime"

	c "github.com/dreddick-home/mixcloudclient/client"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search mixcloud for a term",
	Long:  `Search mixcloud using the term specified`,
	Run: func(cmd *cobra.Command, args []string) {
		term, _ := cmd.Flags().GetString("term")
		max, _ := cmd.Flags().GetInt32("max")
		workers, _ := cmd.Flags().GetInt32("workers")
		excludes, _ := cmd.Flags().GetStringSlice("excludes")

		includes, _ := cmd.Flags().GetStringSlice("includes")
		log.Printf(`Search for term '%s' using %d workers`, term, workers)
		search := c.NewClient(term, createFilter(excludes, includes))

		search.SearchAsync(max, workers)
		search.PrintResults()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	cores := int32(runtime.NumCPU())

	searchCmd.Flags().StringP("term", "t", "", "Search Term")
	searchCmd.MarkFlagRequired("term")
	searchCmd.Flags().Int32P("max", "m", 20, "Max results (in multiples of 100). Default 20 (2000).")
	searchCmd.Flags().Int32P("workers", "w", cores, "The max number of concurrent workers. Defaults to number of cores of system.")
	searchCmd.Flags().StringSliceP("excludes", "e", []string{}, "Must exclude term, multiple items accepted.")
	searchCmd.Flags().StringSliceP("includes", "i", []string{}, "Must include term, multiple items accepted.")

}

func createFilter(excludes []string, includes []string) *c.Filter {
	if len(excludes) == 0 && len(includes) == 0 {
		return nil
	}
	return c.NewFilter(excludes, includes)
}
