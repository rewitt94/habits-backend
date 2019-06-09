package server

import (
	"local/backend/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

type Completed struct {
	Habit_id int64
}

const POST = "POST"

func complete_habit(w http.ResponseWriter, r *http.Request) {

	if (r.Method != POST) {
		http.Error(w, "Expected post method", http.StatusBadRequest)
	}

	var completed Completed;
	
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&completed)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dao.Complete_habit(completed.Habit_id)

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