package main

import (
	"brdbth/graph"
	"brdbth/graph/generated"
	"context"
	"fmt"
	"log"
	"time"

	"net/http"
	"os"

	"brdbth/auth"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	tw_co_ke := os.Getenv("TWITTER_CONSUMER_KEY")
	tw_co_se := os.Getenv("TWITTER_CONSUMER_SECRET")
	tw_ac_to := os.Getenv("TWITTER_ACCESS_TOKEN")
	tw_ac_se := os.Getenv("TWITTER_ACCESS_SECRET")

	creds := auth.TwitterCredentials{AccessToken: tw_ac_to, AccessTokenSecret: tw_ac_se, ConsumerKey: tw_co_ke, ConsumerSecret: tw_co_se}

	fmt.Printf("%+v\n", creds)

	client, err := auth.GetClient(&creds)

	if err != nil {
		log.Println("Error getting twitter client")
		log.Println(err)
	}

	fmt.Printf("%+v\n", client)

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	_, e := GetMongoConnection()

	if e != nil {
		log.Fatal(e)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func GetMongoConnection() (*mongo.Client, error) {

	password := os.Getenv("DB_PASS")
	// connect to the database
	uri := fmt.Sprintf("mongodb+srv://admin:%s@baldwin-dev.fbv4t.mongodb.net/baldwin-dev?retryWrites=true&w=majority", password)
	log.Println(uri)
	clientOptions := options.Client().
		ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
