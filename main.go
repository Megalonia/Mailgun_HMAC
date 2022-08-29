package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"strconv"
	"encoding/hex"
	// mailgun "github.com/mailgun/mailgun-go/v3"
)


func main() {
	var timestamp int64 = 1661549497
	var token string = "23bfc3d51cc8ec4c214d7b53b3d7a93ea8bbada4d253fc68d0"

	timeStr := strconv.FormatInt(timestamp,10)
	// concat1 := token + timeStr
	concat2 := timeStr + token


	// asBytes1 := []byte(concat1)
	asBytes2 := []byte(concat2)

	var secret_string string = "69510e8a8f5190adfd8f960d2a08cf06"
	signature := "e2120fc62e932c8ce2456a1a5207616d8e737a33eb5701cd9a929df5e2253faf"
	// var domain string = "http://bin.mailgun.net/79dc8ac0"

	// mg := mailgun.NewMailgun(domain,secret_string)
	
	// b,err := mg.VerifyWebhookSignature(mailgun.Signature{
	// 	TimeStamp: timeStr,
	// 	Token: token,
	// 	Signature: signature,
	// })

	// if err != nil {
	// 	fmt.Println("err",err)
	// }

	// fmt.Println("b",b)



	 secret := []byte(secret_string)

	 // toHash1 := append(secret,asBytes1...)
	 toHash2 := append(secret,asBytes2...)
	 signature2 := hmac.New(sha256.New, toHash2)

	// encode1 := hex.EncodeToString(signature1.Sum(nil))
	 encode2 := hex.EncodeToString(signature2.Sum(nil))

	// actual := "e2120fc62e932c8ce2456a1a5207616d8e737a33eb5701cd9a929df5e2253faf"

	// fmt.Println("V1: signing key + tokenime  = ",encode1)
	 fmt.Println("V2: signing key + timetoken = ",encode2)
	 fmt.Println("Actual = ",signature) 


	// if encode1 == actual { 
	// 	fmt.Println("V1 is TRUE")
	// } else {
	// 	fmt.Println("V1 is FALSE")
	// }
	 if encode2 == signature { 
	 	fmt.Println("V2 is TRUE")
	 } else {
	 	fmt.Println("V2 is FALSE")
	 }
}
