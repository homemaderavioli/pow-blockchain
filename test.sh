#!/bin/bash

curl -i http://localhost:9001/public_key

curl -i http://localhost:9001/gossip

curl -i -X POST -H "Content-Type: application/json" \
    -d '{"message": "hello blockchain"}' \
    http://localhost:9001/send_money

curl -i http://localhost:9001/gossip
