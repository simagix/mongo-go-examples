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

### Case 1: Watch All Changes

```
argos "mongodb://localhost:30097/?replicaSet=replset"
```

### Case 2: Watch Changes From a Database

```
argos "mongodb://localhost:30097/argos?replicaSet=replset"
```

### Case 3: Watch Changes From a Collection

```
argos --collection oplogs "mongodb://localhost:30097/argos?replicaSet=replset"
```

### Case 4: Watch Changes From a Collection With a Pipeline

```
argos --collection --pipeline '[{"$match": {"operationType": "update"}}]' \
  "mongodb://localhost:30097/argos?replicaSet=replset"
```
