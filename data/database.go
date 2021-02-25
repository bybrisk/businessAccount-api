package data

import (
	"fmt"
	"github.com/bybrisk/structs"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/sirupsen/logrus"
	"database/sql"
	_ "github.com/lib/pq"
)

var resultID string

const (
	host     = "arjuna.db.elephantsql.com"
	port     = 5432
	user     = "adypejae"
	password = "yq8yuFQvGnynPwzjZMbFdfQyV2sa5HFm"
	dbname   = "adypejae"
  )

//Database Funcs
func createBusinessAccount (account *BusinessAccountRequest) string {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, account)
	if insertErr != nil {
		log.Error("Create Business Account ERROR:")
		log.Error(insertErr)
	} else {
		fmt.Println("createBusinessAccount() API result:", result)

		newID := result.InsertedID
		fmt.Println("createBusinessAccount() newID:", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}

func getProfileConfig (account *BusinessAccountRequest) *structs.ProfileConfig {
	collectionName := shashankMongo.DatabaseName.Collection("profileConfig")
	filter := bson.M{"plan": account.BusinessPlan}
	var document *structs.ProfileConfig

	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&document)
	if err != nil {
		log.Error("setProfileConfig ERROR:")
		log.Error(err)
	}
	return document
}	

func setProfileConfig (document *structs.ProfileConfig, docID string) int64 {
	//update businessAccount
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	update := bson.M{"$set": bson.M{"profileConfig": document}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("UpdateDeliveryInfo ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}

func setDeliveryStats (docID string) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	update := bson.M{"$set":bson.M{"deliveryPending": "0", "deliveryDelivered":"0","deliveryCancelled":"0","deliveryTransit":"0"}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("setDeliveryStats ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}

func getBusinessAccount (docID string) *BusinessAccountResponse {
	var businessAccount *BusinessAccountResponse
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
    err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&businessAccount)
	if err != nil {
		log.Error("getBusinessAccount ERROR:")
		log.Error(err)
	}

	return businessAccount
}

func updateBusinessAccount(account *UpdateBusinessAccountRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(account.BybID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.D{{Key: "$set", Value: account}})
	if err != nil {
		log.Error("updateBusinessAccount ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func updatePassword(account *UpdatePasswordRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(account.BybID)
	update := bson.M{"$set":bson.M{"password": account.NewPassword}}
	filter := bson.M{"_id": id}
	res,err := collectionName.UpdateOne(shashankMongo.CtxForDB,filter, update)
	if err!=nil{
		log.Error("updateBusinessAccount ERROR:")
		log.Error(err)
		}	
	
	return res.ModifiedCount
}

func getPassword (docID string) string {
	type password struct {
		Password string `json:"password"`
	}
	var pass *password

	collectionName := shashankMongo.DatabaseName.Collection("businessAccounts")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
    err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&pass)
	if err != nil {
		log.Error("getPassword ERROR:")
		log.Error(err)
	}
	return pass.Password
}

func GetID(d *PasswordAndUsername) string {
	var accountID string

	//db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	//only used with heroku
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Error("Postgre Database connection open() ERROR:")
		log.Error(err)
	}
	defer db.Close()

	//db query
	userSQL := "SELECT bybID FROM bybIDLookup WHERE username= '"+d.UserName+"' and passHash = '"+d.Password+"'"

	err = db.QueryRow(userSQL).Scan(&accountID)

	if err != nil {
		log.Error("Postgre failed to execute query : ")
		log.Error(err)
	}

	return accountID
}

func AddBybIDToPostgre (id string,d *BusinessAccountRequest) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	//only used with heroku
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	
	if err != nil {
		log.Error("Postgre Database connection open() ERROR:")
		log.Error(err)
	}
	defer db.Close()

	//sql query
	sqlStatement := `
		INSERT INTO bybIDLookup (username, passHash, bybID)
		VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, d.UserName, d.Password, id)
	if err != nil {
		log.Error("Postgre failed to execute query : ")
		log.Error(err)
	}

}

func IsUserPresent(username string) bool {

	var isPresent bool
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	
	//only used with heroku
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	
	if err != nil {
		log.Error("Postgre Database connection open() ERROR:")
		log.Error(err)
	}
	defer db.Close()	

	//sql query
	sqlStatement := `select exists(select 1 from bybIDLookup where username='`+username+`')`
	err = db.QueryRow(sqlStatement).Scan(&isPresent)
	if err != nil {
		log.Error("Postgre failed to execute query : ")
		log.Error(err)
	}

	return isPresent
}

func CreateClusterDocument(docID string) string{
	var stringArr ClusterIDArray
	stringArr=ClusterIDArray{
		BybID:docID,
	}
	collectionName := shashankMongo.DatabaseName.Collection("cluster")
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, stringArr)
	if insertErr != nil {
		log.Error("CreateClusterDocument ERROR:")
		log.Error(insertErr)
	} else {
		fmt.Println("CreateClusterDocument() API result:", result)

		newID := result.InsertedID
		fmt.Println("CreateClusterDocument() newID:", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}