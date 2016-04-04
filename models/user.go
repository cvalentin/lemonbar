package models

import (
    "encoding/json"
)

type User struct {
    //Public fields
    Name string

    //Private fields
    //Facebook API fields
    facebookId int
}



func (u User) CreateUser() {
    //JSON self

    //Push to postgres
}

func FacebookSync(){

}
