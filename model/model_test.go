package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/stock1218/KAPPAS/model/data"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stock1218/KAPPAS/model/database"
	"github.com/stock1218/KAPPAS/model/key"
)

// TestNewModel tests if a Model variable can be created.
func TestNewModel(t *testing.T) {
	var _ Model = nil
}

func TestNewAWSModel(t *testing.T) {
	var server Model = NewAWSModel()
	if server == nil {
		t.Log("Failed to create NewAWSModel")
		t.Fail()
	}
}

// Declare the model object that will be used for the follwing tests.
var server Model = NewAWSModel()

// TestPutAndGetPAN tests if PutPAN will take a string, return an id, and GetPAN() will retrieve it.
func TestPutAndGetPAN(t *testing.T) {
	myPAN := data.NewPAN()
	myPAN.SetBillingAddress("Flynn's arcade")
	myPAN.SetCardHolder("Alan Bradly")
	myPAN.SetCardNumber("123456789")
	myPAN.SetExperationDate(time.Now())
	ok := server.PutPAN(myPAN)

	if ok != nil {
		t.Log("Failed to call PutPAN: ", ok)
		t.Fail()
	} else if myPAN.GetID() == "" {
		t.Log("PutPAN returned an empty string")
		t.Fail()
	}

	getData, ok := server.GetPAN(myPAN.GetID())
	if ok != nil {
		t.Log("Error retrieving PAN: ", ok)
		t.Fail()
	} else if reflect.DeepEqual(myPAN, getData) {
		t.Log("Retrieved data was not correct")
		t.Fail()
	}
}

// TestSetAndGetDatabase will test to see if SetDatabase will set the database the AWSModel is uing, and GetDatabse will return it.
func TestSetAndGetDatabase(t *testing.T) {
	testDB := database.NewAmazonRDS("", "", "")
	server.(*AWSModel).SetDatabase(testDB)
	database := server.(*AWSModel).GetDatabase() // Cast server to a pointer to an AWSModel and call GetDatabase

	if database == nil {
		t.Log("GetDatabase() returned nil")
		t.Fail()
	} else if testDB != database {
		t.Log("GetDatabase returned incorrect value, or SetDatabase set incorrect value")
		t.Fail()
	}
}

// TestSetAndGetKey will test if SetKey will set the key used by the AWSModel, and GetKey will return it back.
func TestSetAndGetKey(t *testing.T) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)
	myKey := key.NewKMS(*sess)

	server.(*AWSModel).SetKey(myKey)
	getKey := server.(*AWSModel).GetKey()
	if getKey == nil {
		t.Log("GetKey returned nil")
		t.Fail()
	} else if myKey != getKey {
		t.Log("GetKey returned incorrect value, or SetKey set incorrect value")
	}
}
