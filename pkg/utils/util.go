package utils

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
)

func GenTlsConfig(clientCertPem string, clientKeyPem string, caCertPem string) (tlsConfig *tls.Config) {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(caCertPem)
	if err != nil {
		log.Fatal("read ca-cert.pem error:", err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append ca-cert.pem.")
	}
	clientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair(clientCertPem, clientKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	clientCert = append(clientCert, certs)

	tlsConfig = &tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       clientCert,
		InsecureSkipVerify: true,
	}

	return
}

// 唯一Id
func GetUuid() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}
