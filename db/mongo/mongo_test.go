package mongo_

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDB         = "test"
	testCollection = "collection"
)

var (
	mongoAddress *string
)

func init() {
	mongoAddress = flag.String("mongo", "localhost:27017", "mongodb server address")
}

func TestMongo(t *testing.T) {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", *mongoAddress)))
	require.Nil(t, err)
	session, err := client.StartSession()
	require.Nil(t, err)

	// 插入数据
	fmt.Println("insert")
	collection := session.Client().Database(testDB).Collection(testCollection)
	ret, err := collection.InsertOne(context.Background(), bson.D{{"name", "value"}, {"smiecj", "123"}})
	require.Nil(t, err)
	fmt.Println(ret.InsertedID)

	// 查询数据: FindOne
	findRet := collection.FindOne(context.Background(), bson.D{{"name", "value"}})
	require.Nil(t, findRet.Err())
	raw, err := findRet.DecodeBytes()
	require.Nil(t, err)
	fmt.Println("FindOne ret: " + raw.String())
}
