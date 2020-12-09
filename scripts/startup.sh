#!/bin/sh


set -e


cp -f conf/config-local.toml   conf/config.toml

if [ $NODE_ENV = "testing" ];then
   cp -f conf/config-testing.toml   conf/config.toml
fi

if [ $NODE_ENV = "production" ];then
    cp -f conf/config-production.toml   conf/config.toml
fi

exec ./lmstfy-server -c  conf/config.toml
