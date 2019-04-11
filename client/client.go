package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/beevik/etree"
)

// Default Path to send REST requests.
const DefaultPath = "/nuova"

const LoginPayLoad = `
<aaaLogin
	inName="%s"
	inPassword="%s" />
`

const RefreshPayLoad = `
<aaaRefresh 
    inName="%s" 
    inPassword="%s" 
	inCookie="%s"/>
`

type Client struct {
	BaseURL    *url.URL
	UserName   string
	Password   string
	insecure   bool
	httpClient *http.Client
	AuthToken  *Auth
	proxyUrl   string
}

type Option func(*Client)

// singleton implementation of a client
var clientImpl *Client

func Insecure(insecure bool) Option {
	return func(client *Client) {
		client.insecure = insecure
	}
}
func ProxyUrl(pUrl string) Option {
	return func(client *Client) {
		client.proxyUrl = pUrl
	}
}
func Password(password string) Option {
	return func(client *Client) {
		client.Password = password
	}
}
func initClient(clientUrl, username string, options ...Option) *Client {
	var transport *http.Transport
	bUrl, err := url.Parse(clientUrl)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{
		BaseURL:    bUrl,
		UserName:   username,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	if client.insecure {
		transport = client.useInsecureHTTPClient()
	}
	if client.proxyUrl != "" {
		transport = client.configProxy(transport)
	}
	client.httpClient = &http.Client{
		Transport: transport,
	}
	return client
}

func (c *Client) configProxy(transport *http.Transport) *http.Transport {
	pUrl, err := url.Parse(c.proxyUrl)
	if err != nil {
		log.Fatal(err)
	}
	transport.Proxy = http.ProxyURL(pUrl)
	return transport
}

func (c *Client) useInsecureHTTPClient() *http.Transport {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       true,
			MinVersion:               tls.VersionTLS11,
			MaxVersion:               tls.VersionTLS11,
		},
	}

	return transport

}

func GetClient(clientUrl, username string, options ...Option) *Client {
	if clientImpl == nil {
		clientImpl = initClient(clientUrl, username, options...)
	}
	return clientImpl
}

func (c *Client) MakeRestRequest(method string, body *etree.Document, authenticated bool) (*http.Request, error) {

	url, err := url.Parse(DefaultPath)
	if err != nil {
		return nil, err
	}
	if authenticated {

		body, err = c.InjectAuthenticationHeader(body)
		if err != nil {
			return nil, err
		}
	}
	fURL := c.BaseURL.ResolveReference(url)
	bodyBytes, err := body.WriteToBytes()
	req, err := http.NewRequest(method, fURL.String(), bytes.NewBuffer(bodyBytes))

	if err != nil {
		return nil, err
	}

	return req, nil

}
func (c *Client) Do(req *http.Request) (*etree.Document, *http.Response, error) {

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("\nHTTP Request: %s %s \n", req.Method, req.URL.String())
	fmt.Printf("\nHTTP Response: %d %s \n", resp.StatusCode, resp.Status)

	doc := etree.NewDocument()
	_, err = doc.ReadFrom(resp.Body)

	respStr, _ := doc.WriteToString()
	fmt.Println("\n\n\n ****** %s \n\n\n", respStr)
	if err != nil {
		fmt.Println("Error occurred.")
		return nil, resp, err
	}

	return doc, resp, err
}

func StrtoInt(s string, base int, bitSize int) (int64, error) {
	return strconv.ParseInt(s, base, bitSize)
}
