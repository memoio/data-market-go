#!/bin/bash
ps -ef | grep data-market | grep -v 'color' | awk '{print $2}' | xargs kill -9
