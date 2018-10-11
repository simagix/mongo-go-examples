#! /bin/bash
# Copyright 2018 Kuei-chun Chen. All rights reserved.

echo ; echo "Spin up mongod"
mongod --version
mkdir -p data/db
rm -rf data/db/*
mongod --port 30097 --dbpath data/db --logpath data/mongod.log --fork --wiredTigerCacheSizeGB .5  --replSet replset
mongo --quiet mongodb://localhost:30097/admin --eval 'rs.initiate()'
sleep 2

mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.drop()'
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.insert({"_id": "30097", "scores": [100]})'

#
# Case 1: prints all oplogs
#   go run argos.go "mongodb://localhost:30097/argos?replicaSet=replset" oplogs
# Case 2: print only updates
#   go run argos.go "mongodb://localhost:30097/argos?replicaSet=replset" oplogs '[{"$match": {"operationType": "update"}}]'
#
export DATABASE_URL="mongodb://localhost:30097/argos?replicaSet=replset"
GOCACHE=off go test ./... -v
sleep 2
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.update({"_id": "30097"}, { "\$push": {"scores": 98}})'
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.drop()'

echo ; echo "Shutdown mongod"
mongo --quiet mongodb://localhost:30097/admin?replicaSet=replset --eval 'db.getSisterDB("admin").adminCommand( { replSetStepDown: 0, secondaryCatchUpPeriodSecs: 0, force: true } )'
mongo --quiet --port 30097 --eval 'db.getSisterDB("admin").shutdownServer()'
rm -rf data/*
