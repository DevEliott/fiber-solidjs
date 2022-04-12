#! /bin/bash

cd client && yarn build && cd ../server && go run main.go