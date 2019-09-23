package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/Nerzal/gocloak/v3"
	gocloak2 "github.com/pashazz/gocloak"
	"github.com/pashazz/nextcloud_kcintegrate/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	//Disable security checks
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Get Nextcloud URL
	sNextcloudUrl := os.Getenv("NEXTCLOUD_URL")
	if err := utils.CheckUrl(sNextcloudUrl, "NEXTCLOUD_URL"); err != nil {
		log.Fatal(err)
	}
	// Get Keycloak URL
	sKeycloakUrl := os.Getenv("KEYCLOAK_URL")
	if err := utils.CheckUrl(sKeycloakUrl, "KEYCLOAK_URL"); err != nil {
		log.Fatal(err)
	}

	sUser := utils.GetenvNonEmpty("KEYCLOAK_USER")
	sPassword := utils.GetenvNonEmpty("KEYCLOAK_PASSWORD")
	sRealm := utils.GetenvNonEmpty("KEYCLOAK_REALM")
	sLoginRealm := utils.GetenvNonEmpty("KEYCLOAK_LOGIN_REALM")

	sClientId := utils.GetenvNonEmpty("CLIENT_ID")
	sClientName := utils.GetenvNonEmpty("CLIENT_NAME")
	sClientSecret := utils.GetenvNonEmpty("CLIENT_SECRET")



	sNextcloudUser := utils.GetenvNonEmpty("NEXTCLOUD_USER")
	sNextcloudPassword := utils.GetenvNonEmpty("NEXTCLOUD_PASSWORD")
	sLoginName := utils.GetenvNonEmpty("NEXTCLOUD_LOGIN_NAME")

	// Connect to Keycloak
	kc, err := utils.ConnectToKeycloak(sKeycloakUrl, sUser, sPassword, sRealm, sLoginRealm)
	if err != nil {
		log.Fatalf("Can't connect to keycloak: %v", err)
	}
	err = utils.CreateNextcloudClient(kc, sNextcloudUrl, sClientId, sClientSecret, sClientName, sLoginName)
	if err != nil {
		log.Fatalf("Can't create client on keycloak: %v", err)
	}
	clients, err := kc.Client.GetClients(kc.Token, kc.Realm, gocloak2.GetClientsParams(gocloak.GetClientsParams{
		ClientID: sClientId}))
	if err != nil {
		log.Fatal(err)
	}
	for i, client := range clients {
		log.Print(i)
		marshalled, _ := json.MarshalIndent(client, "", "\t")
	log.Println(string(marshalled))
	//	utils.PrintClient(client)

	}
	// Change Nextcloud settings
	nc, err := utils.ConnectToNextcloud(sNextcloudUrl,sNextcloudUser, sNextcloudPassword, )
	if err != nil {
		log.Fatal(err)
	}
	err = utils.ConfigureSocialLogin(nc, sLoginName, sLoginName, sKeycloakUrl, sRealm, sClientId, sClientSecret)
	if err != nil {
		log.Fatal(err)
	}

}

