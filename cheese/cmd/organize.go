// Copyright © 2018 Alex Haslehurst <alex.haslehurst@gmail.com>
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
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

// organizeCmd represents the organize command
var organizeCmd = &cobra.Command{
	Use:   "organize [path to photo library]",
	Short: "Organize your photo library",
	Long: `Sort photos into folders, remove duplicates, generate a summary`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheeseArgs = args

		service, err := InitializeOrganize()
		if err != nil {
			log.Fatal(err)
		}

		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}


