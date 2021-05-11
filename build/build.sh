# !/bin/bash

# stop and remove old running container
docker rm -f vaccine-cowin-telegram-bot-app
# remove old build image
docker rmi -f vaccine-cowin-telegram-bot-app:latest
# build new app image
docker build --tag "vaccine-cowin-telegram-bot-app:latest" --force-rm=true --no-cache=true --file ../docker/Dockerfile ../