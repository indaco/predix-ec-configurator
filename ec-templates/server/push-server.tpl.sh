#!/bin/bash

PREDIX_DOMAIN=<predix_domain>
EC_SERVER_NAME=<ecagent_server_name>

cf login

cf push
cf enable-diego $EC_SERVER_NAME
cf map-route $EC_SERVER_NAME $PREDIX_DOMAIN -n $EC_SERVER_NAME
cf start $EC_SERVER_NAME

open https://$EC_SERVER_NAME.$PREDIX_DOMAIN/health
