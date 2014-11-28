package main

// START BASE
import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type MeetupPing struct {
    Owner   string      `bson:"Owner"`
    When    time.Time   `bson:"When"`
    Secret  string      `bson:"-"`
}
// END BASE

// START SAVE
func (mp *MeetupPing) save() {
    sess, err := mgo.Dial("localhost")
    if err != nil {
        fmt.Printf("Erreur de connexion a Mongodb : %v", err)
    }
    defer sess.Close()
    sess.SetSafe(&mgo.Safe{})

    collection := sess.DB("meetup").C("pings")
    err = collection.Insert(mp)
    if err != nil {
        fmt.Printf("Erreur a la sauvegarde du ping : %v", err)
    }
}
// END SAVE

// START GET
func getLastPingForMeetup(meetupName string) (MeetupPing, error) {
    var lastPing MeetupPing
    sess, err := mgo.Dial("localhost")
    if err != nil {
        return lastPing, err
    }
    defer sess.Close()
    sess.SetSafe(&mgo.Safe{})
    err = sess.DB("meetup").
        C("pings").
        Find(bson.M{"Owner": meetupName}).
        Sort("-When").
        One(&lastPing)
    if err != nil {
        return lastPing, err
    }

    return lastPing, nil
}
// END GET

// START MAIN
func main() {
    pingNow := MeetupPing {
        Owner: "golang-paris",
        When: time.Now(),
        Secret: "I love Php",
    }
    pingNow.save()
    lastPing, error := getLastPingForMeetup("golang-paris")
    if error != nil {
        fmt.Println(error)
        return
    }
    fmt.Printf("Quel Meetup : %v\n", lastPing.Owner)
    fmt.Printf("Quand : %v\n", lastPing.When)
    fmt.Printf("Le secret: %v", lastPing.Secret)
}
//END MAIN
