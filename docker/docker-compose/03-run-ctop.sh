#!/bin/bash

# https://github.com/bcicen/ctop

if sudo which ctop; then
  echo "ctop found, can proceed"
else
  echo "Unable to find ctop, please install it."
  echo "Check for details:"
  echo "https://github.com/bcicen/ctop"
  echo ""
  echo "Install script is available as well as:"
  echo "./install-ctop.sh"
  exit 1
fi

sudo cp .ctop /root/ctop
sudo ctop
