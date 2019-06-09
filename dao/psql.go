package dao

import (
	"fmt"
	"database/sql"
	"io/ioutil"
	"bytes"
	"strings"
	"sync"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "life_stats"
)

var (
	lock sync.Mutex
	db *sql.DB
)

func Get_database_pool() *sql.DB {

	lock.Lock()
	defer lock.Unlock()

	if db == nil {

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
		db, err := sql.Open("postgres", psqlInfo)
		
		if err != nil {
			panic(err)
		}
	
		return db

	} else {

		return db

	}


}

func Close_database_pool() {

	db := Get_database_pool()
	db.Close()

}

func Initialise_database(relative_path_to_database_init string, habits []Habit) {

	create_and_index_tables(relative_path_to_database_init)
	save_habits(habits)

}

func execute_sql_script(sql_script string) {

	db := Get_database_pool()
	commands := strings.Split(sql_script, "\n\n")

	for command_index := range commands {
		command := commands[command_index]
		fmt.Printf("%v\n\n", command)
		_, err := db.Exec(command)
		if err != nil {
			panic(err)
		}
	}

}

func create_and_index_tables(relative_path_to_database_init string) {

	init_script, err := ioutil.ReadFile(relative_path_to_database_init)
	if err != nil {
		panic(err)
	}

	execute_sql_script(string(init_script))

}

func save_habits(habits []Habit) {

    var sql_script_buffer bytes.Buffer

	for habit_index := range habits {
		habit := habits[habit_index]
		sql_script_buffer.WriteString(
			"INSERT INTO habits(habit, category, condition, score, single_action, weekday, weekend, private, pass_count, total_count, best_streak, current_streak)" +
			fmt.Sprintf("VALUES ('%v','%v','%v',%v,%v,%v,%v,%v,0,0,0,0);\n\n", habit.Habit, habit.Category, habit.Condition, habit.Score, habit.SingleAction, habit.Weekday, habit.Weekend, habit.Private))
	}

	execute_sql_script(sql_script_buffer.String())

}

func Get_remaining_simplified_habits() []SimplifiedHabit {

	all_habits := get_simplified_habits()
	db := Get_database_pool()
	var completed_habit_ids []int64
	var remaining_habits []SimplifiedHabit

	rows, err := db.Query("SELECT habit_id FROM results where DATE(date) = DATE(NOW())")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var habit_id int64
		err = rows.Scan(&habit_id)
		if err != nil {
			panic(err)
		}
		completed_habit_ids = append(completed_habit_ids, habit_id)
	}

	for _,habit := range all_habits {
		var completed bool;
		for _, habit_id := range completed_habit_ids {
			if habit.Id == habit_id {
				completed = true
			}
		}
		if (!completed) {
			remaining_habits = append(remaining_habits, habit)
		}
	}

	return remaining_habits

}

func get_simplified_habits() []SimplifiedHabit {

	db := Get_database_pool()
	var habits []SimplifiedHabit

	rows, err := db.Query("SELECT id, habit, category, condition, score FROM habits")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var habit_name string
		var category string
		var condition string
		var score int64

		err = rows.Scan(&id, &habit_name, &category, &condition, &score)
		if err != nil {
			panic(err)
		}
		negative := score < 0
		habit := SimplifiedHabit{ 
			Id: id,
			Habit: habit_name,
			Category: category,
			Condition: condition,
			Negative: negative,
		}
		habits = append(habits, habit)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return habits

}

func Complete_habit(habit_id int64) {

	execute_sql_script(fmt.Sprintf("INSERT INTO results(habit_id, success, date) VALUES (%v, true, DATE(NOW()))", habit_id))

}