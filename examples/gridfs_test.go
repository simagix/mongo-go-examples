// Copyright 2019 Kuei-chun Chen. All rights reserved.

package examples

import (
	"bufio"
	"bytes"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func TestGridFS(t *testing.T) {
	var err error
	var client *mongo.Client
	var bucket *gridfs.Bucket
	var ustream *gridfs.UploadStream

	str := "This is a test file"
	if client, err = getMongoClient(); err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	if bucket, err = gridfs.NewBucket(client.Database(dbName), options.GridFSBucket().SetName("myFiles")); err != nil {
		t.Fatal(err)
	}

	opts := options.GridFSUpload()
	opts.SetMetadata(bsonx.Doc{{Key: "content-type", Value: bsonx.String("application/json")}})
	if ustream, err = bucket.OpenUploadStream("test.txt", opts); err != nil {
		t.Fatal(err)
	}

	if _, err = ustream.Write([]byte(str)); err != nil {
		t.Fatal(err)
	}

	fileID := ustream.FileID
	ustream.Close()
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	if _, err = bucket.DownloadToStream(fileID, w); err != nil {
		t.Fatal(err, ustream.FileID)
	}

	if b.String() != str {
		t.Fatal("expected", str, "but got", b.String())
	}
}
