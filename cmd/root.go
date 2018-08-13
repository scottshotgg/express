// Copyright © 2018 Scott Gaydos, scgaydos@gmail.com
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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("json-indent", "j", "\t", "output lex tokens in json format")

	// TODO: make a debug logger for every level, or just make our own logger that checks the level
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "output log messges during operation")
	RootCmd.PersistentFlags().BoolP("emit-all", "a", false, "output tokens from all stages and final transpiled C++ with binary")
	RootCmd.PersistentFlags().BoolP("emit-lex", "l", false, "output tokens from lex stage in json format")
	RootCmd.PersistentFlags().BoolP("emit-syn", "y", false, "output tokens from syntactic stage in json format")
	RootCmd.PersistentFlags().BoolP("emit-sem", "s", false, "output tokens from semantic stage in json format")
	RootCmd.PersistentFlags().BoolP("emit-cpp", "c", false, "output transpiled C++ program")

	_ = viper.BindPFlag("json-indent", RootCmd.PersistentFlags().Lookup("json-indent"))
	_ = viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag("emit-all", RootCmd.PersistentFlags().Lookup("emit-all"))
	_ = viper.BindPFlag("emit-lex", RootCmd.PersistentFlags().Lookup("emit-lex"))
	_ = viper.BindPFlag("emit-syn", RootCmd.PersistentFlags().Lookup("emit-syn"))
	_ = viper.BindPFlag("emit-sem", RootCmd.PersistentFlags().Lookup("emit-sem"))
	_ = viper.BindPFlag("emit-cpp", RootCmd.PersistentFlags().Lookup("emit-cpp"))
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

		// Search config in home directory with name ".express" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".express")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "express",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

func writeTokensJSONToFile(tokensJSON []byte, pathOfFile string) error {
	lexFile, err := os.Create(pathOfFile)
	if err != nil {
		// TODO:
		return err
	}

	n, err := lexFile.Write(tokensJSON)
	if err != nil {
		// TODO:
		return err
	}
	if n != len(tokensJSON) {
		// TODO:
		// need to rewrite the lexFile
		return errors.New("Not all bytes were written")
	}

	err = lexFile.Close()
	if err != nil {
		// TODO:
		return err
	}

	return nil
}
