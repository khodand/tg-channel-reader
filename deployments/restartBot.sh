#!/bin/bash

cd $(dirname "$0")

while getopts "b" option; do
   case $option in
      b) # tbot
         killall tbot
         echo "tg_bot killed"
         cp bot.log "logs/bot_$(date +%F).log"
         ./yt-dlp -U
         nohup ./tbot 1>bot.log 2>&1 &
         echo "tg_bot restarted"
         ;;
   esac
done

echo "script finished"
