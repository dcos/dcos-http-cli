package pluginutil

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/dcos/dcos-cli/pkg/log"
	"github.com/sirupsen/logrus"
)

// HTTPClient returns an HTTP client from a plugin runtime.
func HTTPClient(baseURL string, opts ...httpclient.Option) *httpclient.Client {
	if baseURL == "" {
		baseURL, _ = os.LookupEnv("DCOS_URL")
	}
	var baseOpts []httpclient.Option

	if acsToken, _ := os.LookupEnv("DCOS_ACS_TOKEN"); acsToken != "" {
		baseOpts = append(baseOpts, httpclient.ACSToken(acsToken))
	}

	if verbosity, _ := os.LookupEnv("DCOS_VERBOSITY"); verbosity != "" {
		baseOpts = append(baseOpts, httpclient.Logger(Logger()))
	}

	tlsInsecure, _ := os.LookupEnv("DCOS_TLS_INSECURE")
	if tlsInsecure == "1" {
		baseOpts = append(baseOpts, httpclient.TLS(&tls.Config{InsecureSkipVerify: true}))
	} else {
		tlsCAPath, _ := os.LookupEnv("DCOS_TLS_CA_PATH")
		if tlsCAPath != "" {
			rootCAsPEM, err := ioutil.ReadFile(tlsCAPath)
			if err == nil {
				certPool := x509.NewCertPool()
				if certPool.AppendCertsFromPEM(rootCAsPEM) {
					baseOpts = append(baseOpts, httpclient.TLS(&tls.Config{RootCAs: certPool}))
				}
			}
		}
	}
	return httpclient.New(strings.TrimRight(baseURL, "/"), append(baseOpts, opts...)...)
}

// Logger returns a logger for a given plugin runtime.
func Logger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &log.Formatter{},
		Hooks:     make(logrus.LevelHooks),
	}
	verbosity, _ := os.LookupEnv("DCOS_VERBOSITY")
	if verbosity == "1" {
		logger.SetLevel(logrus.InfoLevel)
	} else if verbosity == "2" {
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}
