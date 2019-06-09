package dao

import (
	"io/ioutil"
	"bytes"
	"strconv"
	"fmt"
)

func Load_habits_from_csv(relative_path_to_csv string) ([]Habit, error) {
	var habits []Habit
	filename := relative_path_to_csv
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(content, []byte("\n"))
	for line_index := range lines {
		/*
			skip title row in config file
		*/
		if line_index > 0 {
			line := lines[line_index];
			habit, err := create_habit_from_comma_separated_line(line)
			if (err != nil) {
				fmt.Printf("Could not format line: %v\nERROR : %v\nINFO : Habit will be ommitted", line, err)
			} else {
				habits = append(habits, habit)
			}
		}
	}
	return habits, nil
}

func create_habit_from_comma_separated_line(raw []byte) (Habit, error) {
	data := bytes.Split(raw, []byte(","))
	score, score_err := strconv.ParseInt(string(data[3]), 10, 64)
	if score_err != nil {
		return Habit{}, score_err
	}
	single_action, single_action_err := strconv.ParseBool(string(data[4]))
	if single_action_err != nil {
		return Habit{}, single_action_err
	}
	weekday, weekday_err := strconv.ParseBool(string(data[5]))
	if weekday_err != nil {
		return Habit{}, weekday_err
	}
	weekend, weekend_err := strconv.ParseBool(string(data[6]))
	if weekend_err != nil {
		return Habit{}, weekend_err
	}
	private, private_err := strconv.ParseBool(string(data[7]))
	if private_err != nil {
		return Habit{}, private_err
	}
	return Habit{
		Habit:     		string(data[0]),
		Category: 		string(data[1]),
		Condition: 		string(data[2]),
		Score: 			score,
		SingleAction: 	single_action,
		Weekday:  		weekday,
		Weekend:  		weekend,
		Private:  		private,
	}, nil
	
}