package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
	"log"
)

type NextcloudOIDCEntry struct {
	Name string `json:"name"`
	Title string `json:"title"`
	AuthorizeUrl string `json:"authorizeUrl"`
	TokenUrl string `json:"tokenUrl"`
	UserInfoUrl string `json:"userInfoUrl"`
	LogoutUrl string `json:"logoutUrl"`
	ClientID string `json:"clientId"`
	ClientSecret  string `json:"clientSecret"`
	Scope string `json:"scope"`
	GroupsClaim string `json:"groupsClaim"`
	style string `json:"style"`
	defaultGroup string `json:"defaultGroup"`
}

func ConnectToNextcloud (nextcloudURL string, user string, password string) (*gonextcloud.Client, error) {
	c, err := gonextcloud.NewClient(nextcloudURL)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to nextcloud %v: %v", nextcloudURL, err))
	}
	if err := c.Login(user, password); err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to log in nextcloud %v: %v", nextcloudURL, err))
	}
	return c, nil
}

func ConfigureSocialLogin (client *gonextcloud.Client, loginName string, loginTitle string, keycloakURL string, realm string, clientId string, clientSecret string) error {
	const APP_NAME = "sociallogin"
	const KEY_NAME = "custom_oidc_providers"
	keys, err := client.AppsConfig().Keys(APP_NAME)
	if err != nil || len(keys) == 0 {
		return errors.New(fmt.Sprintf("Can't find application '%v'", APP_NAME))
	}
	log.Print(keys)
	data, err := client.AppsConfig().Details(APP_NAME)
	if err != nil {
		return errors.New(fmt.Sprintf("Cant find setting %v in %v", KEY_NAME, APP_NAME))
	}
	value := data[KEY_NAME]
	log.Print(value)
	//Unmarshal a JSON string value
	var providers []*NextcloudOIDCEntry
	err = json.Unmarshal([]byte(value), &providers)
	if err != nil {
		return errors.New("Unable to read oidc_providers setting")
	}
	oidcPath := fmt.Sprintf("auth/realms/%v/protocol/openid-connect", realm)
	newProvider := &NextcloudOIDCEntry{
		Name:         loginName,
		Title:        loginTitle,
		AuthorizeUrl: fmt.Sprintf("%v/%v/auth", keycloakURL, oidcPath),
		TokenUrl:     fmt.Sprintf("%v/%v/token", keycloakURL, oidcPath),
		UserInfoUrl:  fmt.Sprintf("%v/%v/userinfo", keycloakURL, oidcPath),
		LogoutUrl:    fmt.Sprintf("%v/%v/logout", keycloakURL, oidcPath),
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scope:        "openid",
		GroupsClaim:  "",
		style:        "",
		defaultGroup: "",
	}
	providers = append(providers, newProvider)
	newData, err :=  json.Marshal(providers)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to marshal JSON OIDC provider data: %v", err))
	}
	err = client.AppsConfig().SetValue(APP_NAME, KEY_NAME, string(newData))
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to set setting %v/%v: %v", APP_NAME, KEY_NAME, err))
	}
	return nil
}