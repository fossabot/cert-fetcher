package cmd

import (
	"github.com/bakito/cert-fetcher/cert/jks"
	"github.com/spf13/cobra"
)

var (
	jksPassword string
	jksSource   string
)

// jksCmd represents the jks command
var jksCmd = &cobra.Command{
	Version: version,
	Use:     "jks",
	Short:   "store the certificates into an java keystore",
	Long:    "store the certificates into an java keystore",
	RunE: func(cmd *cobra.Command, args []string) error {
		return jks.Export(targetURL, certIndexes, jksSource, jksPassword, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(jksCmd)
	jksCmd.PersistentFlags().StringVarP(&jksPassword, "password", "p", "changeit", "the password to be used for the java keystore")
	jksCmd.PersistentFlags().StringVarP(&jksSource, "source", "s", "", "the source keystore to add the certs to")

}
