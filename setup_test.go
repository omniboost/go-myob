package myob_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/omniboost/go-myob"
	"golang.org/x/oauth2"
)

var client *myob.Client

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	companyFileID := os.Getenv("COMPANY_FILE_ID")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	oauthConfig := myob.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	//t, err := oauthConfig.TokenSource(oauth2.NoContext, token).Token()
	//fmt.Println(t.AccessToken)

	httpClient := oauthConfig.Client(context.Background(), token)

	client = myob.NewClient(httpClient)
	client.SetCompanyFileID(companyFileID)
	client.SetClientID(clientID)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
