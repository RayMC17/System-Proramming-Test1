// filename: main.go
//Raynisha Cornelio
//Test Number 1
//System Programming
// Mr.Dalwin Lewis

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// create a new Mux instance to route HTTP requests
	mux := http.NewServeMux()

	// register the root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Mr. Lewis, my name is Raynisha Michelle Cornelio, but you already know that... Welcome to my website!\n\n")
		fmt.Fprintf(w, "I am pursuing an Associate Degree and planning to continue with a Bachelor's in IT. After that, I am not sure what I want to do, Is that normal? I have so many plans that I don't know what to choose...\n\n")
		fmt.Fprintf(w, "I always dreamed of becoming a teacher but then I want to work as a Intelligence Analyst,which work along with the Police Department and with Crime etc.\n")
		fmt.Fprintf(w, "Well I'll see what's gonna happen after I finish my studies, hoping all goes well.\n")
		fmt.Fprintf(w, "I just want to make my parents proud and I want to help them in every way I can and be the Role model for my other two sisters\n\n")
		fmt.Fprintf(w, "Some of my hobbies include playing musical instruments like guitar, bass and piano, drawing, reading, cooking, and baking.\n\n")
		fmt.Fprintf(w, "Something I don't like hmm...I don't like spaghetti and I don't like washing dishes I do wash it but it's not something I enjoy, I prefer to wash clothes.\n\n")
		fmt.Fprintf(w, "Thank you for visiting my website. My email address is 2021154121@ub.edu.bz Have a blessed and fantastic day!\n")
	})

	// register the greeting handler
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		dayOfWeek := currentTime.Weekday()

		switch dayOfWeek {
		case time.Monday:
			fmt.Fprintf(w, "Happy Monday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Tuesday:
			fmt.Fprintf(w, "Happy Tuesday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Wednesday:
			fmt.Fprintf(w, "Happy Wednesday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Thursday:
			fmt.Fprintf(w, "Happy Thursday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Friday:
			fmt.Fprintf(w, "Happy Friday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Saturday:
			fmt.Fprintf(w, "Happy Saturday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		case time.Sunday:
			fmt.Fprintf(w, "Blessed Sunday! %s\n", currentTime.Format("2006-01-02 15:04:05"))
		} 
	})

	// register the random quote handler
	mux.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		quotes := []string{
			"Life is short. Smile while you still have teeth.",
			"Did you know DIET stands for 'Did I eat that?'",
			"Life has no remote... get up and change it yourself. -Mark A. Cooper",
			"Life teaches you a new lesson every day, if you are attentive enough in the class of life. -Invajy",
			"It always seems impossible until it's done. -Nelson Mandela",
			"The future belongs to those who believe in the beauty of their dreams.",
		}

		rand.Seed(time.Now().UnixNano())
		randomQuote := quotes[rand.Intn(len(quotes))]

		fmt.Fprintf(w, "Random Quote: %s\n", randomQuote)
	})

	// start the server
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux) 
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	} 
}
