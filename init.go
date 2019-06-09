package main

import (
	"local/backend/dao"
)

const relative_path_to_habits_csv = "./config/habits.csv"
const relative_path_to_database_init = "./config/init.sql"

func main() {

	habits, err := dao.Load_habits_from_csv(relative_path_to_habits_csv)
	if (err != nil) {
		panic(err)
	}
	dao.Initialise_database(relative_path_to_database_init, habits)
	dao.Close_database_pool()
	
}