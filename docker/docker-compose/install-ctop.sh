#!/bin/bash

# Installation instructions are from here:
# https://github.com/bcicen/ctop

echo "deb http://packages.azlux.fr/debian/ buster main" | sudo tee /etc/apt/sources.list.d/azlux.list
wget -qO - https://azlux.fr/repo.gpg.key | sudo apt-key add -
sudo apt update
sudo apt install -y docker-ctop

# Install binary
# sudo wget https://github.com/bcicen/ctop/releases/download/0.7.6/ctop-0.7.6-linux-amd64 -O /usr/local/bin/ctop
# sudo chmod +x /usr/local/bin/ctop
