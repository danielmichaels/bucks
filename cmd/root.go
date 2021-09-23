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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danielmichaels/bucks/lib/currencies"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bucks",
	Short: "A simple currency converter made with Go",
	Long: `Bucks is a dead simple currency converter for the commandline

Bucks requires the use of third-party currency APIs and must be set with at least
one API key from the vendors listed.

Supported currency APIs:

- https://www.currencyconverterapi.com/

A valid API key must be provided to use this. type 'bucks config --help' for how
to get started.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("at least one argument must be supplied")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bucks starting")
		amount, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatalln("first argument must be a valid int or float")
		}

		// validate the currency is valid
		from := strings.ToUpper(args[1])
		err = checkValidCurrency(from)
		if err != nil {
			log.Fatalln(err)
		}

		// validate the currency is valid
		to := strings.ToUpper(args[2])
		err = checkValidCurrency(to)
		if err != nil {
			log.Fatalln(err)
		}

		key := checkAPIKeySet()
		if err != nil {
			log.Fatalln(err)
		}

		respBytes := freeCurrencyRequest(amount, from, to, key)

		formatter(to, from, respBytes.Total, amount)
	},
}

// formatter prints to screen the results
func formatter(to, from string, total, amount float64) {
	c := currencies.FreeCurrencyConverterList
	curTo := currencies.CurrencyList{}
	curFrom := currencies.CurrencyList{}
	for _, v := range c {
		if v.Id == to {
			curTo = v
		}
		if v.Id == from {
			curFrom = v
		}
	}
	output := fmt.Sprintf(`
Total: %s%.2f %s

Conversion: %.2f %s to %s
`, curTo.CurrencySymbol, total, curTo.CurrencyName, amount, curFrom.Id, curTo.Id)
	fmt.Println(output)
}

// checkValidCurrency validates that the entered currency is valid
func checkValidCurrency(from string) error {
	configCurrencyAPI := currencies.FreeCurrencyConverterList
	for _, v := range configCurrencyAPI {
		if v.Id == from {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("currency %q is not valid for this API", from))
}

// checkAPIKeySet validates that an API key has been set
func checkAPIKeySet() string {
	key := viper.GetString("providers.currencyconverterapi.key")
	if key == "" {
		log.Fatalln("no api key found")
	}
	return key

	// todo check for a default flag

}

func freeCurrencyRequest(amount float64, from string, to string, key string) currencies.Currency {
	currency := currencies.Currency{}

	// use a map because the key is dynamically created based on user supplied
	// currency pair.
	var resp map[string]json.RawMessage

	apiKey := key

	url := currencies.FreeCurrencyURL{
		Url: fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s_%s&compact=ultra&apiKey=%s", from, to, apiKey),
	}
	request, err := http.NewRequest(
		http.MethodGet,
		url.Url,
		nil,
	)
	if err != nil {
		log.Fatalln(err, "failed to retrieve data from API")
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "'Bucks' currency conversion CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalln(err, "could not establish connection with API server.")
	}

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		log.Fatalln(err, "could not read data received from API")
	}
	// Free Currency API returns a single dynamic key made up of the two currencies
	// being queried. To query the map for those keys, make a dynamic lookup based
	// on the 'from' and 'to' values.
	jsonKey := fmt.Sprintf("%s_%s", from, to)

	apiAmount, err := strconv.ParseFloat(string(resp[jsonKey]), 64)
	if err != nil {
		log.Fatalln(err, "failed to convert returned currency into human-readable string")
	}

	currency.Total = Converter(apiAmount, amount)
	return currency
}

func Converter(value, amount float64) float64 {
	total := value * amount
	return total
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bucks.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bucks" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".bucks")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "No config file found. See 'bucks --help' for more information about how to set one")
		os.Exit(1)
	}
}
