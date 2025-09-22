package main

import "ecommerce/cmd"

func main() {

	cmd.Server()
	// var s string
	// s = "a"
	// byteArr := []byte(s)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// base64Str := enc.EncodeToString(byteArr)
	// decodeStr, err := enc.DecodeString(base64Str)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println(byteArr)
	// fmt.Println(base64Str)
	// fmt.Println(decodeStr)
	// var s string
	// s = "Hello world"
	// byteStr := []byte(s)
	// hashStr := sha256.Sum256(byteStr)
	// fmt.Println(hashStr)

	// secret := []byte("abcdef")
	// message := []byte("Print Message")

	// hash := hmac.New(sha256.New, []byte(secret))
	// hash.Write(message)
	// text := hash.Sum(nil)
	// fmt.Println(text)
}
