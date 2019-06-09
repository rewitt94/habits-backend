package server

import (
	"local/backend/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

const GET = "GET"

func send_all_remaining_habits(w http.ResponseWriter, r *http.Request) {

	if (r.Method != GET) {
		http.Error(w, "Expected get method", http.StatusBadRequest)
	}

	habits := dao.Get_remaining_simplified_habits()
	fmt.Println(habits)

	js, err := json.Marshal(habits)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}