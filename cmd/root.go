// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/autom8ter/api"
	"github.com/autom8ter/backend"
	"github.com/autom8ter/backend/cache"
	"github.com/autom8ter/backend/config"
	"github.com/autom8ter/backend/contact"
	"github.com/autom8ter/backend/payment"
	"github.com/stripe/stripe-go"

	"github.com/autom8ter/backend/user"
	"github.com/autom8ter/backend/utility"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var credspath string
var port int
var debug bool
var email string
var name string

func init() {
	api.Util.DotEnv()
	rootCmd.Flags().IntVarP(&port, "port", "p", 3000, "port to serve on")
	rootCmd.Flags().StringVarP(&credspath, "creds", "c", "credentials.json", "path to gcp service account credentials (JSON)")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debugging mode for development")
	rootCmd.Flags().DurationVarP(&cache.SYNC_FREQUENCY, "sync", "s", 1*time.Minute, "time to wait inbetween cache sync")
	rootCmd.Flags().StringVarP(&email, "email", "e", os.Getenv("SENDGRID_EMAIL"), "sendgrid email for admin->user emails")
	rootCmd.Flags().StringVarP(&name, "name", "n", "Admin", "name to user in admin emails")

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "The Autom8ter gprc API Backend",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg = config.FromEnv(credspath)
		if err := cfg.Validate(); err != nil {
			api.Util.Entry().Fatalln("Set Env: SENDGRID_KEY, TWILIO_ACCOUNT, TWILIO_KEY, AUTH0_DOMAIN, AUTH0_CLIENT_SECRET, AUTH0_CLIENT_ID, STRIPE_KEY", err.Error())
		}
		stripe.Key = cfg.StripeKey
		cache.Init()
		b := backend.NewBackend(
			utility.NewUtility(cfg).PluginFunc,
			contact.NewConatact(cfg).PluginFunc,
			user.NewUser().PluginFunc,
			payment.NewSubscriber().PluginFunc,
		)
		err := b.Serve(fmt.Sprintf(":%v", port), debug)
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
