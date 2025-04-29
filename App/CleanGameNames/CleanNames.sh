#!/bin/sh
cd $(dirname "$0")
export HOME=/mnt/SDCARD
sleep 1
#export LD_LIBRARY_PATH=$(dirname "$0")/lib:$LD_LIBRARY_PATH
./clean_names > "${PWD}"/log.txt
