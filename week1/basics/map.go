package main

import "fmt"

func main() {

	var user map[string]string = map[string]string{
		"name": "Spira",
		"lastName": "404",
	}

	// set capacity
	profile := make(map[string]string, 10)

	fmt.Printf("%+v\n", profile)

	// if there is no key, you will get default value
	mName := user["middlename"]
	fmt.Println("mName", mName)

	// check if key exists
	mName, mNameExist := user["middlename"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// delete key
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
