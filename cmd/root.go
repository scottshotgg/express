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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/scottshotgg/express/lex"
	"github.com/scottshotgg/express/parse"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("json-indent", "j", "\t", "Whether or not to output C++ transpilation")

	// TODO: make a debug logger for every level, or just make our own logger that checks the level
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "Whether or not to output logs during compilation")
	RootCmd.PersistentFlags().BoolP("emit-all", "a", false, "Whether or not to output C++ transpilation")
	RootCmd.PersistentFlags().BoolP("emit-lex", "l", false, "Whether or not to output C++ transpilation")
	RootCmd.PersistentFlags().BoolP("emit-syn", "s", false, "Whether or not to output C++ transpilation")
	RootCmd.PersistentFlags().BoolP("emit-sem", "x", false, "Whether or not to output C++ transpilation")
	RootCmd.PersistentFlags().BoolP("emit-cpp", "c", false, "Whether or not to output C++ transpilation")

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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("ERROR: You must provide an input program")
			return
		}

		//fmt.Println("args", args)

		// This is where I would need an env variable
		var err error
		parse.LibBase, err = filepath.Abs("lib/")
		if err != nil {
			os.Exit(9)
		}

		jsonIndent := viper.GetString("json-indent")

		// Replace the \t and \n string literals with their non-escaped counterparts
		jsonIndent = strings.Replace(jsonIndent, `\n`, "\n", -1)
		jsonIndent = strings.Replace(jsonIndent, `\t`, "\t", -1)

		// TODO: need to check it for all the available characters
		filenameArg := args[len(args)-1]
		// filenameFull, err := filepath.Abs()
		stat, err := os.Stat(filenameArg)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if stat.IsDir() {
			fmt.Println("Directories compilation is not currently supported.")
			os.Exit(0)
		}
		filename := stat.Name()
		filenameSplit := strings.Split(filename, ".")
		// TODO: not sure why I cut out having en-extensioned files
		if len(filenameSplit) < 1 {
			os.Exit(9)
		}
		filenameNoExt := filenameSplit[0]

		input, err := ioutil.ReadFile(filenameArg)
		if err != nil {
			fmt.Printf("ERROR: Cannot read input program: %s\n", args[len(args)-1])
			os.Exit(9)
		}

		lexTokens, err := lex.New(string(input)).Lex()
		if err != nil {
			fmt.Println("error lexing", err)
			os.Exit(9)
		}

		if viper.GetBool("emit-lex") || viper.GetBool("emit-all") {
			lexTokensJSON, err := json.MarshalIndent(lexTokens, "", jsonIndent)
			if err != nil {
				// TODO:
				return
			}
			err = writeTokensJSONToFile(lexTokensJSON, filenameNoExt+".lex.json")
			if err != nil {
				// TODO:
				return
			}
		}

		// p := parse.New(lexTokens)
		// tokens, err := p.Parse()
		// if err != nil {
		// 	fmt.Println("error in syntactic parsing", err)
		// 	os.Exit(9)
		// }
		// fmt.Println("tokens", tokens)
		// // PrintTokens(tokens, jsonIndent)
		// // fmt.Println("\n\n")

		// // p = parse.New(syntacticTokens)

		syntacticTokens, err := parse.New(lexTokens).Syntactic()
		if err != nil {
			fmt.Println("error in syntactic parsing", err)
			os.Exit(9)
		}

		if viper.GetBool("emit-syn") || viper.GetBool("emit-all") {
			syntacticTokensJSON, err := json.MarshalIndent(syntacticTokens, "", jsonIndent)
			if err != nil {
				// TODO:
				return
			}
			err = writeTokensJSONToFile(syntacticTokensJSON, filenameNoExt+".syn.json")
			if err != nil {
				// TODO:
				return
			}
		}

		pNew := parse.New(syntacticTokens)
		semanticTokens, err := pNew.Semantic()
		if err != nil {
			fmt.Println("error in semantic parsing", err)
			os.Exit(9)
		}

		if viper.GetBool("emit-sem") || viper.GetBool("emit-all") {
			semanticTokensJSON, err := json.MarshalIndent(semanticTokens, "", jsonIndent)
			if err != nil {
				// TODO:
				return
			}
			err = writeTokensJSONToFile(semanticTokensJSON, filenameNoExt+".sem.json")
			if err != nil {
				// TODO:
				return
			}
		}

		cpp, err := pNew.Transpile(semanticTokens)
		if err != nil {
			os.Exit(9)
		}

		cppFilename := filename + ".cpp"
		if !viper.GetBool("emit-cpp") && !viper.GetBool("emit-all") {
			tempDir := os.TempDir()
			cppFilename = tempDir + cppFilename
		}

		f, err := os.Create(cppFilename)
		if err != nil {
			fmt.Println("got an err creating file", err)
			return
		}
		n, err := f.WriteString(cpp)
		if err != nil {
			// TODO:
			return
		}
		if n != len(cpp) {
			// TODO:
		}
		err = f.Close()
		if err != nil {
			// TODO:
		}

		// FIXME: write this to a temp dir/file using Go and then move it if we need it
		output, err := exec.Command("clang++", "-std=gnu++2a", cppFilename, "-o", filenameNoExt+".exe").CombinedOutput()
		if err != nil {
			// TODO:
			fmt.Println("err compile", err, string(output))
			os.Exit(9)
		}

		if viper.GetBool("emit-cpp") || viper.GetBool("emit-all") {
			output, err = exec.Command("clang-format", "-i", cppFilename).CombinedOutput()
			if err != nil {
				// TODO:
				fmt.Println("err compile", err, string(output))
				os.Exit(9)
				return
			}
		} else {
			output, err = exec.Command("rm", cppFilename).CombinedOutput()
			if err != nil {
				// TODO:
				fmt.Println("err delete", err, string(output))
				os.Exit(9)
			}
		}
	},
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
