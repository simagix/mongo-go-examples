// Copyright 2020 Kuei-chun Chen. All rights reserved.

package examples

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
 * copied from keyhole to remove dependency
 */

// Feeder seeds feeder
type Feeder struct {
	collection string
	database   string
	isDrop     bool
	total      int
}

// Model - robot model
type Model struct {
	ID          string `json:"_id" bson:"_id"`
	Name        string
	Description string
	Year        int
}

// Task - robot task
type Task struct {
	For         string `json:"for" bson:"for"`
	MinutesUsed int    `json:"minutesUsed" bson:"minutesUsed"`
}

// Robot -
type Robot struct {
	ID         string  `json:"_id" bson:"_id"`
	ModelID    string  `json:"modelId,omitempty" bson:"modelId,omitempty"`
	Notes      string  `json:"notes" bson:"notes"`
	BatteryPct float32 `json:"batteryPct,omitempty" bson:"batteryPct,omitempty"`
	Tasks      []Task  `json:"tasks" bson:"tasks"`
}

// NewFeeder establish seeding parameters
func NewFeeder() *Feeder {
	return &Feeder{isDrop: false, total: 1000}
}

// SetCollection set collection
func (f *Feeder) SetCollection(collection string) {
	f.collection = collection
}

// SetDatabase set database
func (f *Feeder) SetDatabase(database string) {
	f.database = database
}

// SetIsDrop set isDrop
func (f *Feeder) SetIsDrop(isDrop bool) {
	f.isDrop = isDrop
}

// SetTotal set total
func (f *Feeder) SetTotal(total int) {
	f.total = total
}

// SeedFavorites seeds demo data of collection favorites
func (f *Feeder) SeedFavorites(client *mongo.Client) error {
	var err error
	var ctx = context.Background()
	c := client.Database(f.database).Collection("lookups")
	favoritesCollection := client.Database(f.database).Collection("favorites")
	if f.isDrop {
		if err = c.Drop(ctx); err != nil {
			return err
		}
		if err = favoritesCollection.Drop(ctx); err != nil {
			return err
		}
	}

	for i := 0; i < 10; i++ {
		c.InsertOne(ctx, bson.M{"_id": i + 1000, "type": "sports", "name": Favorites.Sports[i]})
		c.InsertOne(ctx, bson.M{"_id": i + 1100, "type": "book", "name": Favorites.Books[i]})
		c.InsertOne(ctx, bson.M{"_id": i + 1200, "type": "movie", "name": Favorites.Movies[i]})
		c.InsertOne(ctx, bson.M{"_id": i + 1300, "type": "city", "name": Favorites.Cities[i]})
		c.InsertOne(ctx, bson.M{"_id": i + 1400, "type": "music", "name": Favorites.Music[i]})
	}
	f.seedCollection(favoritesCollection, 2)
	return err
}

// SeedCars seeds cars collection
func (f *Feeder) SeedCars(client *mongo.Client) error {
	var err error
	var ctx = context.Background()
	carsCollection := client.Database(f.database).Collection("cars")
	dealersCollection := client.Database(f.database).Collection("dealers")
	employeesCollection := client.Database(f.database).Collection("employees")
	if f.isDrop {
		carsCollection.Drop(ctx)
		dealersCollection.Drop(ctx)
		employeesCollection.Drop(ctx)
	}

	// Upsert examples
	for i := 0; i < len(dealers); i++ {
		dealerID := fmt.Sprintf("DEALER-%d", 1+i)
		opts := options.Update()
		opts.SetUpsert(true)
		if _, err := dealersCollection.UpdateOne(ctx, bson.M{"_id": dealerID}, bson.M{"$set": bson.M{"name": dealers[i]}}, opts); err != nil {
			log.Fatal(err)
		}
	}

	var emp bson.M
	opts := options.Replace()
	opts.SetUpsert(true)
	var empID = int(1001)
	emp = getEmployee(empID, 0)
	empID++
	employeesCollection.ReplaceOne(ctx, bson.M{"_id": emp["_id"]}, emp, opts)
	for i := 0; i < 2; i++ {
		emp = getEmployee(empID, 1001)
		parent := empID
		employeesCollection.ReplaceOne(ctx, bson.M{"_id": emp["_id"]}, emp, opts)
		empID++
		for j := 0; j < 3; j++ {
			emp = getEmployee(empID, parent)
			pID := empID
			employeesCollection.ReplaceOne(ctx, bson.M{"_id": emp["_id"]}, emp, opts)
			empID++
			for k := 0; k < 5; k++ {
				emp = getEmployee(empID, pID)
				employeesCollection.ReplaceOne(ctx, bson.M{"_id": emp["_id"]}, emp, opts)
				empID++
			}
		}
	}

	// create index example
	indexView := carsCollection.Indexes()
	indexView.CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{Key: "color", Value: 1}}})
	indexView.CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{Key: "color", Value: 1}, {Key: "brand", Value: 1}}})
	indexView.CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{Key: "filters.k", Value: 1}, {Key: "filters.v", Value: 1}}})
	if _, err = dealersCollection.CountDocuments(ctx, bson.M{}); err != nil {
		return err
	}
	f.seedCollection(carsCollection, 1)
	fopts := options.Find()
	filter := bson.D{{Key: "color", Value: "Red"}}
	fopts.SetSort(bson.D{{Key: "brand", Value: -1}})
	fopts.SetProjection(bson.D{{Key: "_id", Value: 0}, {Key: "color", Value: 1}, {Key: "brand", Value: 11}})
	carsCollection.Find(ctx, filter, fopts)
	// fmt.Printf("Seeded cars: %d, dealers: %d\n", carsCount, dealersCount)
	return err
}

var dealers = []string{"Atlanta Auto", "Buckhead Auto", "Johns Creek Auto"}
var brands = []string{"Acura", "Alfa Romeo", "Audi", "Bentley", "BMW", "Buick", "Cadillac", "Chevrolet", "Chrysler", "Dodge",
	"Fiat", "Ford", "GMC", "Genesis", "Honda", "Hyundai", "Infiniti", "Jaguar", "Jeep", "Kia",
	"Land Rover", "Lexus", "Lincoln", "Maserati", "Mazda", "Mercedes-Benz", "Nissan", "Porsche", "Toyota", "Volkswagen"}
var styles = []string{"Sedan", "Coupe", "Convertible", "Minivan", "SUV", "Truck"}
var colors = []string{"Beige", "Black", "Blue", "Brown", "Gold",
	"Gray", "Green", "Orange", "Pink", "Purple",
	"Red", "Silver", "White", "Yellow"}

func getVehicle() bson.M {
	curYear := time.Now().Year()
	delta := rand.Intn(8)
	year := curYear - delta
	used := true
	if delta == 0 {
		used = false
	}
	brand := brands[rand.Intn(len(styles))]
	color := colors[rand.Intn(len(colors))]
	style := styles[rand.Intn(len(styles))]

	return bson.M{
		"dealer": fmt.Sprintf("DEALER-%d", 1+rand.Intn(len(dealers))),
		"brand":  brand,
		"color":  color,
		"style":  style,
		"year":   year,
		"used":   used,
		"filters": []bson.M{
			bson.M{"k": "brand", "v": brand},
			bson.M{"k": "color", "v": color},
			bson.M{"k": "style", "v": style},
			bson.M{"k": "year", "v": year},
			bson.M{"k": "used", "v": used}},
	}
}

func (f *Feeder) seedCollection(c *mongo.Collection, fnum int) int {
	var ctx = context.Background()
	var bsize = 100
	var remaining = f.total

	for threadNum := 0; threadNum < f.total; threadNum += bsize {
		num := bsize
		if remaining < bsize {
			num = remaining
		}
		remaining -= num
		var contentArray []interface{}
		for n := 0; n < num; n++ {
			if fnum == 1 {
				contentArray = append(contentArray, getVehicle())
			} else if fnum == 2 {
				contentArray = append(contentArray, getDemoDoc())
			}
		}
		opts := options.InsertMany()
		opts.SetOrdered(false) // ignore duplication errors
		c.InsertMany(ctx, contentArray, opts)
	}
	cnt, _ := c.CountDocuments(ctx, bson.M{})
	return int(cnt)
}

func getEmployee(id int, supervisor int) bson.M {
	dealerID := "DEALER-1"
	email := getEmailAddress()
	s := strings.Split(strings.Split(email, "@")[0], ".")
	doc := bson.M{"_id": int32(id), "dealer": dealerID, "email": email, "name": s[0] + " " + s[2]}
	if supervisor != 0 {
		doc["manager"] = int32(supervisor)
	}
	return doc
}

var domains = []string{"gmail.com", "me.com", "yahoo.com", "outlook.com", "google.com",
	"simagix.com", "aol.com", "mongodb.com", "example.com", "cisco.com",
	"microsoft.com", "facebook.com", "apple.com", "amazon.com", "oracle.com"}
var fnames = []string{"Andrew", "Ava", "Becky", "Brian", "Cindy",
	"Connie", "David", "Dawn", "Elizabeth", "Emma",
	"Felix", "Frank", "George", "Grace", "Hector",
	"Henry", "Ian", "Isabella", "Jennifer", "John",
	"Kate", "Kenneth", "Linda", "Logan", "Mary",
	"Michael", "Nancy", "Noah", "Olivia", "Otis",
	"Patricia", "Peter", "Quentin", "Quinn", "Richard",
	"Robert", "Samuel", "Sophia", "Todd", "Tom",
	"Ulysses", "Umar", "Vincent", "Victoria", "Wesley",
	"Willaim", "Xavier", "Xena", "Yosef", "Yuri", "Zach", "Zoey",
}
var lnames = []string{"Smith", "Johnson", "Williams", "Brown", "Jones",
	"Miller", "Davis", "Garcia", "Rodriguez", "Chen",
	"Adams", "Arthur", "Bush", "Carter", "Clinton",
	"Eisenhower", "Ford", "Grant", "Harrison", "Hoover"}
