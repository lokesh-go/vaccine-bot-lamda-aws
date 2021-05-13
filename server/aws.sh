# !/bin/bash

# code push to aws server
scp -i techy-lokesh.pem -r /Users/lokeshchandra/Desktop/go-projects/src/vaccine-bot-lamda-aws ubuntu@192.168.0.100:"/home/ubuntu/telegram_bot"