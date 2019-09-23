package utils

import (
	"errors"
	"fmt"
	"github.com/pashazz/gocloak"
)

type Keycloak struct {
	Client gocloak.GoCloak
	Token string
	Realm string
}


func CreateNextcloudClient (kc *Keycloak, nextcloudUrl string, clientId string, clientSecret string,  clientName string, loginName string) error {
	err := kc.Client.CreateClient(kc.Token, kc.Realm, gocloak.Client{
		AdminURL:                           nextcloudUrl,
		AuthorizationServicesEnabled:       false,
		BearerOnly:                         false,
		ClientAuthenticatorType:            "client-secret",
		ClientID:                           clientId,
		ConsentRequired:                    false,
		DirectAccessGrantsEnabled:          true,
		Enabled:                            true,
		FrontChannelLogout:                 false,
		FullScopeAllowed:                   true,
		ImplicitFlowEnabled:                false,
		Name:                               clientName,
		NodeReRegistrationTimeout:          -1,
		Origin:                             "",
		Protocol:                           "openid-connect",
		ProtocolMappers:                    nil,
		PublicClient:                       false,
		RedirectURIs:                       []string{fmt.Sprintf("%v/apps/sociallogin/custom_oidc/%v", nextcloudUrl, loginName)},
		RootURL:                            nextcloudUrl,
		Secret:                             clientSecret,
		ServiceAccountsEnabled:             false,
		StandardFlowEnabled:                true,
		SurrogateAuthRequired:              false,
		WebOrigins:                         []string{nextcloudUrl},
	})
	return err
}

func ConnectToKeycloak (keycloakUrl string, user string, password string, realm string, loginRealm string )  (*Keycloak, error) {
	client := gocloak.NewClient(keycloakUrl)

	token, err := client.LoginAdmin(user, password, loginRealm)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to keycloak: %v", err))
	}
	return &Keycloak{
		Client: client,
		Token:  token.AccessToken,
		Realm:  realm,
	}, nil
	
}