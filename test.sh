#! /bin/bash
# Copyright 2018 Kuei-chun Chen. All rights reserved.

echo ; echo "Spin up mongod"
mongod --version
mkdir -p data/db
rm -rf data/db/*
mongod --port 30097 --dbpath data/db --logpath data/mongod.log --fork --wiredTigerCacheSizeGB .5  --replSet replset
mongo --quiet mongodb://localhost:30097/admin --eval 'rs.initiate()'
sleep 2

# Case 1: prints all oplogs
# Case 2: print only updates
#   '[{"$match": {"operationType": "update"}}]'
export DATABASE_URL="mongodb://localhost:30097/argos?replicaSet=replset"
GOCACHE=off go test ./... -v

echo ; echo "Shutdown mongod"
mongo --quiet --port 30097 --eval 'db.getSisterDB("admin").shutdownServer()' > /dev/null 2>&1
rm -rf data/*
