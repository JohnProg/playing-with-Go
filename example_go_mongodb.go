package main

import (
    "fmt"
    "time"
)

import (
    "labix.org/v2/mgo"
)

func connectToMongo() bool {
    ret := false
    fmt.Println("enter main - connecting to mongo")

    // tried doing this - doesn't work as intended
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Detected panic")
            var ok bool
            err, ok := r.(error)
            if !ok {
                fmt.Printf("pkg:  %v,  error: %s", r, err)
            }
        }
    }()

    maxWait := time.Duration(5 * time.Second)
    session, sessionErr := mgo.DialWithTimeout("localhost:27017", maxWait)
    if sessionErr == nil {
        session.SetMode(mgo.Monotonic, true)
        coll := session.DB("MyDB").C("MyCollection")
        if ( coll != nil ) {
            fmt.Println("Got a collection object")
            ret = true
        }
    } else { // never gets here
        fmt.Println("Unable to connect to local mongo instance!")
    }
    return ret
}

func main() {
    if ( connectToMongo() ) {
        fmt.Println("Connected")
    } else {
        fmt.Println("Not Connected")
    }
}
