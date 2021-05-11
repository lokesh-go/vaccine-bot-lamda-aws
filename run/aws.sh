# !/bin/bash

# run new container
docker run -d --restart=always --env-file /home/ubuntu/tbot.env --name vaccine-cowin-telegram-bot-app  vaccine-cowin-telegram-bot-app:latest
