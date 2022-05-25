#!/bin/bash
dependency=$(npm list -g | grep nodemon)

RunNodemon() {
   nodemon --exec go run ./src/main.go --signal SIGTERM
}

if [ "$dependency" != "" ]; then
  echo "nodemon already installed"
  RunNodemon
else
  echo "nodemon not installed, installing..."
  npm i -g nodemon
  RunNodemon
fi
