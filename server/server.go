package server

import (
	"net/http"
	"log"
	"fmt"
)

func App() {

	http.HandleFunc("/remaining", send_all_remaining_habits)
	http.HandleFunc("/complete", complete_habit)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
