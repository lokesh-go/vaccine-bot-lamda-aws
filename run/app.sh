# !/bin/bash

# run new container
docker run -d --restart=always --env-file $HOME/.env --name vaccine-cowin-telegram-bot-app  vaccine-cowin-telegram-bot-app:latest
