package curl

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/lucas-clemente/quic-go/h2quic"
	log "github.com/sirupsen/logrus"
)

func getRoundTripper(cfg Config) (roundTripper *h2quic.RoundTripper, err error) {
	roundTripper = &h2quic.RoundTripper{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if strings.TrimSpace(cfg.CACert) == "" {
		log.Debug("No CA file provided: Skipping TLS Verification")
		return
	}

	fileBytes, err := ioutil.ReadFile(cfg.CACert)
	if err != nil {
		log.Debug("Failed reading CA File: ", err)
		return
	}

	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		log.Debug("Failed getting system cert pool: ", err)
		return
	}

	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if ok := rootCAs.AppendCertsFromPEM(fileBytes); !ok {
		return nil, errors.New("Failed appending CA file cert to cert pool")
	}

	roundTripper.TLSClientConfig = &tls.Config{RootCAs: rootCAs}
	return
}

// Curl is the entrypoint when executing curl against QUIC server
func Curl(cfg Config) {
	roundTripper, err := getRoundTripper(cfg)
	if err != nil {
		log.Error("Failed getting roundTripper: ", err)
		return
	}
	defer roundTripper.Close()

	client := &http.Client{
		Transport: roundTripper,
	}

	log.Debug("CURL: ", cfg.URL)
	res, err := client.Get(cfg.URL)
	if err != nil {
		log.Error("failed getting url: ", err)
		return
	}

	io.Copy(os.Stdout, res.Body)
}
