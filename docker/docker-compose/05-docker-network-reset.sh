#!/bin/bash

echo "Restart network"
./network-restart.sh
sleep 3
echo "Restart Docker"
./docker-restart.sh
