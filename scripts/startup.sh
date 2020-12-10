#!/bin/bash


set -e


cp -f conf/config-local.toml   conf/config.toml

if [ -z $NODE_ENV ]; then
   NODE_ENV="local"
fi

echo "NODE_ENV = $NODE_ENV"
echo "NEW_RELIC_APP_NAME = $NEW_RELIC_APP_NAME"

if [ $NODE_ENV = "testing" ];then
   cp -f conf/config-testing.toml   conf/config.toml
fi

if [ $NEW_RELIC_APP_NAME = "lmstfy.automizelyapi.org_couriers-scheduler-production" ];then
    cp -f conf/config-couriers-scheduler-production.toml   conf/config.toml
fi

exec ./lmstfy-server -c  conf/config.toml
