# !/bin/bash

# code push to aws server
scp -i techy-lokesh.pem -r /Users/lokeshchandra/Desktop/go-projects/src/vaccine-bot-lamda-aws ubuntu@13.234.202.219:"/home/ubuntu/telegram_bot"