@echo off
set PREDIX_DOMAIN=<predix_domain>
set EC_GATEWAY_NAME=<ecagent_gateway_name>

cf login

cf push
cf enable-diego %EC_GATEWAY_NAME%
cf map-route %EC_GATEWAY_NAME% %PREDIX_DOMAIN% -n %EC_GATEWAY_NAME%
cf start %EC_GATEWAY_NAME%

open https://%EC_GATEWAY_NAME%.%PREDIX_DOMAIN%/health
