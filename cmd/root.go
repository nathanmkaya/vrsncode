/*
Copyright © 2020 Nathan Mkaya <nathanmkaya@gmail.com>

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
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var packageName string
var keyFile string

var rootCmd = &cobra.Command{
	Use:   "vrsncode",
	Short: "Android Version Code manipulation tool",
	Long: `Vrsncode is a tool to fetch and update Android version code with respect to
the latest version code for an app deployed to play store`,
}

const (
	defaultConfigName = ".vrsncode"

	configFlag          = "config"
	configFlagShortHand = "c"
	configFlagUsageDesc = "config file (default is .vrsncode)"

	keyFlag          = "key"
	keyFlagShortHand = "k"
	keyFlagUsageDesc = "key file (e.g service account json file)"

	packageFlag          = "package"
	packageFlagShortHand = "p"
	packageFlagUsageDesc = "package name"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, configFlag, configFlagShortHand, "", configFlagUsageDesc)
	rootCmd.PersistentFlags().StringVarP(&keyFile, keyFlag, keyFlagShortHand, "", keyFlagUsageDesc)
	err := viper.BindPFlag(keyFlag, rootCmd.PersistentFlags().Lookup(keyFlag))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVarP(&packageName, packageFlag, packageFlagShortHand, "", packageFlagUsageDesc)
	err = viper.BindPFlag(packageFlag, rootCmd.PersistentFlags().Lookup(packageFlag))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(cwd)
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
