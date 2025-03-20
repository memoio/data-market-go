#!/bin/bash

ps -ef | grep data-market | grep -v 'color' | awk '{print $2}' | xargs kill -9

nohup ./data-market daemon run -e=$1 > log 2>&1 &