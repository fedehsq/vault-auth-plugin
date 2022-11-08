#!/bin/sh
cd sshwifty/
echo "INSTALLING SSHWIFTY DEPENDENCIES..."
npm install
echo "BUILDING SSHWIFTY APPLICATION..."
npm run build
cd ..
echo "STARTING SSHWIFTY..."
SSHWIFTY_CONFIG=./sshwifty/sshwifty.conf.json ./sshwifty/sshwifty