package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"strings"
	"strconv"
	"math/rand"
)

type Meeting struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Participants *Participants `json:"participants"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	CreatedAt time.Time `json:"createdAt"`
}
type Participants struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Rsvp string `json:"rsvp"`
}
var Meetings []Meeting




// Home Route

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Server Running at Port 8080")
}

/*
	/meetings => get all meetings
	/meetings/{id} => get meeting by ID
	/meetings/?participant={email} => get Meeting fro that mail
*/	

func getMeetings(w http.ResponseWriter, r *http.Request){
	found := 0
	par := strings.TrimPrefix(r.URL.Path, "/meetings/")
	log.Println(par)
    if len(par)>0 { 
		for _, value := range Meetings {
			if value.ID == par {
				data, err := json.MarshalIndent(value, "", "    ")
				if err != nil {
					log.Println(err)
				}
				log.Println(string(data))
				fmt.Fprintf(w, string(data))
				found = 1
				break
			}
		}
		if found==0{
			fmt.Fprintf(w, "Meeting does not exist. ")
		}
    }else{
	    	data, err := json.MarshalIndent(Meetings, "", "    ")
			if err != nil {
				log.Println(err)
			}
			log.Println(string(data))
			if len(Meetings) == 0 {
				fmt.Fprintf(w, "No Meetings")
			} else{
				  fmt.Fprintf(w, string(data))
				} 
    }
	keys, ok := r.URL.Query()["participant"]
    
    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'participant' is missing")
        return
    }
    key := string(keys[0])
	if len(key)>0 { 
		for _, value := range Meetings {
			if value.Participants.Email == key {
				data, err := json.MarshalIndent(value, "", "    ")
				if err != nil {
					log.Println(err)
				}
				log.Println(string(data))
				fmt.Fprintf(w, string(data))
				found = 1
				break
			}
		}
    }
}

/*
	/meetings/create/ => Create a Meeting 
*/
func postMeetings(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var mh Meeting
	_ = json.NewDecoder(r.Body).Decode(&mh)
	mh.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID -Generate any random No. for ID
	Meetings = append(Meetings, mh)
	json.NewEncoder(w).Encode(mh)
}


func main(){
	
	/*Meetings = append(Meetings,Meeting{ID:"1",Title:"Test",Participants:&Participants{Name:"shubham",Email:"shubhamaniket6@gmail.com",Rsvp:"Yes"},StartTime:"9",EndTime:"10",CreatedAt:time.Now()})
	Meetings = append(Meetings,Meeting{ID:"2",Title:"Test1",Participants:&Participants{Name:"shubham",Email:"shubhamaniket6@gmail.com",Rsvp:"Yes"},StartTime:"9",EndTime:"10",CreatedAt:time.Now()})
	*/
	
	http.HandleFunc("/meetings/", getMeetings)			//getAll Meetings + /meetings/{id} => get Meeting By Id
	http.HandleFunc("/meetings/create/",postMeetings)	//Create a meeting
	http.HandleFunc("/",home)							//Home Route

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}