package meego

var languages map[string]string

func init() {
	languages = make(map[string]string)
	languages["cs"] = "C #"
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["go"] = "Golang"
}
func Get(key string) string {
	return languages[key]
}
func Add(key, value string) {
	languages[key] = value
}
func GetAll() map[string]string {
	return languages
}

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// var finallist string //variable to store final list of meetups

// //-------structs to parse json data from Meetup API
// type venue struct {
// 	Venue string `json:"name"`
// }
// type group struct {
// 	GroupName string `json:"name"`
// }
// type meetuplist struct {
// 	Name  string `json:"name"`
// 	Venue venue  `json: "venue"`
// 	Date  string `json:"local_date"`
// 	Time  string `json:"local_time"`
// 	Group group  `json: "group"`
// 	Link  string `json:"link"`
// }

// func init() {

// }

// //---------
// //fetching details of meetup of the group's url
// func getmeetups(url string) string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var meetups []meetuplist
// 	json.Unmarshal([]byte(body), &meetups)
// 	if (len(meetups)) > 0 {
// 		finallist = "Title -" + "\t" + meetups[0].Name + "\n" + "Date -" + "\t" + meetups[0].Date + "\n" + "Link -" + "\t" + meetups[0].Link + "\n\n"
// 		return finallist
// 	}
// 	return ""
// }
