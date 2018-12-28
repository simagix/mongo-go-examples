# Argos

A MongoDB Change Streams implementation using [mongodb-go-driver](https://github.com/mongodb/mongo-go-driver).

## Demo

Set up a replica set

```
mkdir -p data/db
rm -rf data/db/*
mongod --port 30097 --dbpath data/db --logpath data/mongod.log --fork --wiredTigerCacheSizeGB .5  --replSet replset
mongo --quiet mongodb://localhost:30097/admin --eval 'rs.initiate()'
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.insert({"_id": "30097", "scores": [100]})'
```

### Case 1: prints all oplogs

```
argos "mongodb://localhost:30097/argos?replicaSet=replset" oplogs
```

### Case 2: print only updates

```
argos "mongodb://localhost:30097/argos?replicaSet=replset" oplogs '[{"$match": {"operationType": "update"}}]'
```

#### Generate oplogs

```
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.insert({"_id": "90210", "scores": [85]})'
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.update({"_id": "90210"}, { "\$push": {"scores": 98}})'
mongo --quiet mongodb://localhost:30097/argos?replicaSet=replset --eval 'db.oplogs.remove({"_id": "90210"})'
```
