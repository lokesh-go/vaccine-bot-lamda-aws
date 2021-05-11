# !/bin/bash

# run new build container
docker run -d --restart=always --env-file /home/pi/tbot.env --name vaccine-cowin-telegram-bot-app  vaccine-cowin-telegram-bot-app:latest
