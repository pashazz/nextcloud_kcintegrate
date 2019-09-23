#!/bin/bash
source .env

export NEXTCLOUD_ADMIN_USER
export NEXTCLOUD_ADMIN_PASSWORD

export NEXTCLOUD_LOGIN_NAME=keycloaktest
export NEXTCLOUD_URL=http://localhost:8080
export KEYCLOAK_URL=https://keycloak.isp
export KEYCLOAK_REALM=cloud
export KEYCLOAK_LOGIN_REALM=master
export KEYCLOAK_USER=admin
export KEYCLOAK_PASSWORD=admin
export CLIENT_ID=testClientId0
export CLIENT_NAME=TestClient0
export CLIENT_SECRET=testClientSecret0






if [[ $@ == "docker" ]]
then
   until docker-compose exec app ls /done &> /dev/null;
   do
      echo "waiting for nextcloud installation to end, sleep 1s"
      sleep 1s
   done
   sleep 2s # For installation to end
else
    until curl $NEXTCLOUD_URL &> /dev/null;
    do
        echo "failed to dial $NEXTCLOUD_URL,  sleep 1s"
        sleep 1s
    done
fi

./nextcloud_kcintegrate
