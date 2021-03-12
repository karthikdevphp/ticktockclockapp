package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)


const sec = 1
const min = 60
const hour = 3600
const stoptime = 10800


type server struct {
	doneChan chan bool
	tickerA  ticker
	tickerB  ticker
	tickerC  ticker
}

type ticker struct {
	period time.Duration
	ticker time.Ticker
}

var usersInput = make(map[string]string)
var out io.Writer = os.Stdout //for testing

func main() {
	clockStart(sec,min,hour,stoptime)
}

/*
Func to start the clock
@accepts sec time.Duration
@accepts min time.Duration
@accepts hour time.Duration
@accepts stopTime time.Duration
 */
func clockStart(sec time.Duration, min time.Duration, hour time.Duration,stopTime time.Duration) error {
	doneChan := make(chan bool)
	tickerA := createTicker(sec * time.Second)
	tickerB := createTicker(min * time.Second)
	tickerC := createTicker(hour * time.Second)

	s := &server{doneChan, *tickerA, *tickerB, *tickerC}
	start := time.Now()
	go s.listener(start,stopTime)
	go startServer()
	msg := <-doneChan

	if msg {
		tickerA.ticker.Stop()
		tickerB.ticker.Stop()
		tickerC.ticker.Stop()
	}
	return nil
}

/*
Func to start http server

*/
func startServer() {
	http.HandleFunc("/userConcole", changeValue)
	err := http.ListenAndServe(":9093", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
Func to start the HTTP SERVER

*/
func createTicker(period time.Duration) *ticker {
	return &ticker{period, *time.NewTicker(period)}
}

/*
Func to start the listener
@accepts start time.Time
@accepts stopTime time.Duration
*/
func (s *server) listener(start time.Time,stopTime time.Duration) {

	for {
		select {
		case <-s.tickerC.ticker.C:
			elapsed := time.Since(start)
			//fmt.Println(time.Duration(elapsed.Seconds()))
			if usersInput["Hour"] != "" {
				fmt.Fprint(out, "Elapsed: ", math.Round(elapsed.Seconds()), usersInput["Hour"]+"\n")
			} else {
				fmt.Fprint(out, "Elapsed: ", math.Round(elapsed.Seconds()), " Bong\n")
			}

			if time.Duration(elapsed.Seconds()) == stopTime {
				s.doneChan <- true
				return
			}
		case <-s.tickerB.ticker.C:
			elapsed := time.Since(start)
			m := time.Duration(elapsed.Seconds())
			if m%hour != 0 {
				if usersInput["Minute"] != "" {
					fmt.Fprint(out,"Elapsed: ", math.Round(elapsed.Seconds()), usersInput["Minute"]+"\n")
				} else {
					fmt.Fprint(out, "Elapsed: ", math.Round(elapsed.Seconds()), " Tock\n")
				}
			}
		case <-s.tickerA.ticker.C:
			elapsed := time.Since(start)
			s := time.Duration(elapsed.Seconds())
			if s%min != 0 && s%hour != 0 {
				if usersInput["Second"] != "" {
					fmt.Fprint(out,"Elapsed: ", math.Round(elapsed.Seconds()), usersInput["Second"]+"\n")
				} else {
					fmt.Fprint(out, "Elapsed: ", math.Round(elapsed.Seconds()), " Tick\n")
				}
			}
		case <-s.doneChan:
			return
		}
	}

}

/*
Func for http handler

*/
func changeValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["Second"][0] != "" {
			usersInput["Second"] = r.Form["Second"][0]
		}
		if r.Form["Minute"][0] != "" {
			usersInput["Minute"] = r.Form["Minute"][0]
		}
		if r.Form["Hour"][0] != "" {
			usersInput["Hour"] = r.Form["Hour"][0]
		}
	}
	t, _ := template.ParseFiles("tmpl/template_file.gohtml")
	t.Execute(w, nil)
}

