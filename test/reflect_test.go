package test

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	ID      uint64 `json:"id"`
	Address string `js:"address"`
}

func Test_Run1(t *testing.T) {

fmt.Println(fmt.Errorf("wallet asset(%d)",uint64(22)))

	user := User{ID: 1, Address: "123"}
	//user := A()
	//userType := reflect.TypeOf(*user)
	//userValue := reflect.ValueOf(*user)
	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)
	fmt.Println(userType.Name())
	fmt.Println(userType.Kind())
	//var update = make(map[string]interface{})
	for i := 0; i < userType.NumField(); i++ {

		fmt.Println(userType.Field(i).Name)
		fmt.Println(userValue.Field(i).Interface())
		//update[userType.Field(i).Name] = userValue.Field(i).Interface()
	}
}
func A() *User {
	user := User{ID: 1, Address: "123"}
	return &user
}

func B() *User {
	user := User{ID: 1, Address: "123"}
	return &user
}
