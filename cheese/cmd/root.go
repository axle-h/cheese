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
	"errors"
	"fmt"
	"github.com/axle-h/cheese/config"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var cfgFile string
var verbose bool
var cheeseArgs []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cheese",
	Short: "Organize and sync photos",
	Long: `Cheese is a simple CLI for managing a library of photos and keeping it in sync with cloud providers.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)

	// Persistent, global flags.
	persistentFlags := rootCmd.PersistentFlags()
	persistentFlags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cheese.yaml)")
	persistentFlags.BoolVarP(&verbose,"verbose", "v", false, "increases verbosity")
}

func initLogging() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cheese" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cheese")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func NewCheeseConfig() (config.CheeseConfig, error) {
	path, err := getPath(cheeseArgs[0])
	return config.CheeseConfig{Path: path}, err
}

func getPath(rawPath string) (string, error) {
	path, err := filepath.Abs(rawPath)

	if err != nil {
		log.Errorf("Cannot find path: %s", rawPath)
		return "", err
	}

	fi, err := os.Stat(rawPath)
	if err != nil {
		log.Errorf("Cannot read path: %s", path)
		return "", err
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		log.Debugf("Using path: %s", path)
		return path, nil

	default:
		log.Errorf("Not a directory: %s", path)
		return path, errors.New("must specify a directory")
	}
}

