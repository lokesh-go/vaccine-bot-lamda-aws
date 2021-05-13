package dal

import (
	"context"
	"errors"
	"os"
	"time"
	"vaccine-bot-lamda-aws/src/dal/models"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbClient *mongo.Database

func Initialize() (err error) {
	err = connect()
	if err != nil {
		return errors.New("db connection failed" + err.Error())
	}

	// Returns
	return nil
}

func connect() (err error) {
	// Select mongo options
	mongoConnOptions := &options.ClientOptions{
		Hosts: []string{os.Getenv("DBHOST")},
	}

	// Authentication
	mongoConnOptions.Auth = &options.Credential{
		Username:   os.Getenv("DBUSER"),
		Password:   os.Getenv("DBPASS"),
		AuthSource: os.Getenv("DBAUTH"),
	}

	// New client connction
	client, err := mongo.NewClient(mongoConnOptions)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return errors.New("client connection failed ->" + err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.New("client ping failed ->" + err.Error())
	}

	dbClient = client.Database(os.Getenv("DBNAME"))

	// Returns
	return nil
}

func Get(chatID int64) (response *models.Get, err error) {
	// Forms query
	query := map[string]interface{}{
		"_id": chatID,
	}

	// Hits DB
	var res map[string]interface{}
	err = dbClient.Collection(os.Getenv("DBCOLL")).FindOne(nil, query).Decode(&res)
	if err != nil {
		// Handles no record
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
	}

	// Decodes response
	err = mapstructure.Decode(res, &response)
	if err != nil {
		return nil, err
	}

	// Returns
	return response, nil
}

func GetAll() (res *[]models.Get, err error) {
	// Gets all records
	cursor, err := dbClient.Collection(os.Getenv("DBCOLL")).Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}

	// Binds cursor response
	mapRes := []map[string]interface{}{}
	err = cursor.All(nil, &mapRes)
	if err != nil {
		return nil, err
	}
	// Checks records
	if len(mapRes) == 0 {
		return nil, errors.New("NO_DOC_IN_DB")
	}

	// Binds data
	err = mapstructure.Decode(mapRes, &res)
	if err != nil {
		return nil, errors.New("GETALL_BINDS_DATA_ERR")
	}

	// Returns
	return res, nil
}

func Delete(chatID int64) (err error) {
	// Forms query
	query := map[string]interface{}{
		"_id": chatID,
	}

	// Hits DB
	res, err := dbClient.Collection(os.Getenv("DBCOLL")).DeleteOne(nil, query)
	if err != nil {
		return err
	}

	// In case ID not found in db
	if res.DeletedCount == 0 {
		return errors.New("REC_NOT_FOUND")
	}

	// Returns
	return nil
}

func Create(chatID int64, pincode string, name string) (err error) {
	// Forms query
	reqMap := map[string]interface{}{
		"_id":     chatID,
		"pincode": pincode,
		"name":    name,
	}

	// Hits DB
	_, err = dbClient.Collection(os.Getenv("DBCOLL")).InsertOne(nil, &reqMap)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

func BackUp(req map[string]interface{}) (err error) {
	// Hits DB
	_, err = dbClient.Collection("archive").InsertOne(nil, &req)
	if err != nil {
		return err
	}

	// Returns
	return nil
}
