package cmd

import (
	"github.com/rms1000watt/hello-world-quic/curl"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var curlCmd = &cobra.Command{
	Use:     "curl",
	Short:   "curl a QUIC server",
	Long:    `curl a QUIC server`,
	Example: `./hello-world-quic curl https://localhost:7100/hello-world`,
	Run:     curlFunc,
}

var curlCfg curl.Config

func init() {
	rootCmd.AddCommand(curlCmd)

	curlCmd.Flags().StringVar(&curlCfg.CACert, "cacert", "", "CA crt to trust")

	setFlagsFromEnv(curlCmd)
}

func curlFunc(cmd *cobra.Command, args []string) {
	configureLogging()

	if len(args) == 0 {
		log.Error("No URL provided")
		return
	}

	curlCfg.URL = args[0]
	curl.Curl(curlCfg)
}
