package key

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// TestNewKey tests if a Key type variable can be created.
//
func TestNewKey(t *testing.T) {
	var _ Key = nil
}

// TestNewKMS tests if a KMS struct can be created.
//
func TestNewKMS(t *testing.T) {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)

	var key Key = NewKMS(*sess)
	if key == nil {
		t.Log("Key: failed to create new KMS")
		t.Fail()
	}
}

// KeySetUp is used as a setup function for creating a testable KMS.
//
func KeySetUp() Key {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	var key Key = NewKMS(*sess)
	return key
}

// TestSetAndGetID tests if SetID will set the ID of a key, and GetID will return it.
//
func TestSetAndGetID(t *testing.T) {
	key := KeySetUp()

	initialID := "123"
	key.SetID("123")
	returnedID := key.GetID()

	if initialID != returnedID {
		t.Log("SetID set the wrong value, or GetID returned the wrong value.")
		t.Fail()
	}

}

// TestSetAndGetClient tests if SetClient will set the client for a KMS struct, and GetClient will return it.
//
func TestSetAndGetClient(t *testing.T) {
	key := KeySetUp()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)

	initialClient := kms.New(sess)

	key.(*KMS).SetClient(*initialClient)

	returnedClient := key.(*KMS).GetClient()
	if initialClient.Config.Region != returnedClient.Config.Region {
		t.Log("SetSession set the incorrect value, or GetSession returned the wrong value")
		t.Fail()
	}
}

// TestEncryptAndDecrypt tests if encrypt and decrypt will encrypt and decrypt a message.
//
func TestEncryptAndDecrypt(t *testing.T) {
	key := KeySetUp()

	plaintext := "this is my secret"

	key.SetID("arn:aws:kms:us-east-1:188809012751:key/e33ea72b-0ffd-4ba3-aa12-bebced4bd23d")
	ciphertext := key.Encrypt(plaintext)
	if ciphertext == plaintext {
		t.Log("Encrypt didn't encrypt the string")
		t.Fail()
	}

	decryptedText := key.Decrypt(ciphertext)
	if decryptedText != plaintext {
		t.Log("Decrypt failed to decrypt the string")
		t.Fail()
	}

}
