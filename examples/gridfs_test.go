// Copyright 2019 Kuei-chun Chen. All rights reserved.

package examples

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/gridfs"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

func TestGridFS(t *testing.T) {
	var err error
	var client *mongo.Client
	var bucket *gridfs.Bucket
	var ustream *gridfs.UploadStream

	str := "This is a test file"
	client = getMongoClient()

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
