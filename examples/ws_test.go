// Copyright 2018 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestConnectSimple(t *testing.T) {
	uri := "mongodb://user:password@localhost/?replicaSet=replset&authSource=admin"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	count, _ := client.Database("test").Collection("tutorial").CountDocuments(context.TODO(), bson.D{{}})
	t.Log("count:", count)
}

func TestConnectBuild(t *testing.T) {
	opts := options.Client()
	opts.SetHosts([]string{"localhost:27017"})
	auth := options.Credential{Username: "user", Password: "password", AuthSource: "admin"}
	opts.SetAuth(auth)
	opts.SetReplicaSet("replset")
	opts.SetMaxPoolSize(5)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		t.Fatal(err)
	}
	count, _ := client.Database("test").Collection("tutorial").CountDocuments(context.TODO(), bson.D{{}})
	t.Log("count:", count)
}

func TestConnectTLS(t *testing.T) {
	uri := "mongodb://user:password@localhost/?replicaSet=replset&authSource=admin"
	opts := options.Client().ApplyURI(uri)
	caBytes, _ := ioutil.ReadFile("/etc/ssl/certs/ca.pem")
	clientBytes, _ := ioutil.ReadFile("/etc/ssl/certs/client.pem")
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(caBytes); !ok {
		t.Fatal(errors.New("failed to parse root certificate"))
	}
	certs, err := tls.X509KeyPair(clientBytes, clientBytes)
	if err != nil {
		t.Fatal(err)
	}
	cfg := &tls.Config{RootCAs: roots, Certificates: []tls.Certificate{certs}}
	opts.SetTLSConfig(cfg)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		t.Fatal(err)
	}
	count, _ := client.Database("test").Collection("tutorial").CountDocuments(context.TODO(), bson.D{{}})
	t.Log("count:", count)
}
