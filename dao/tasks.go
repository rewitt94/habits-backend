package dao

type SimplifiedHabit struct {
	Id				int64
	Habit     		string
	Category 		string
	Condition 		string
	Negative	 	bool
}

type Habit struct {
	Id				int64
	Habit     		string
	Category 		string
	Condition 		string
	Score    		int64
	SingleAction 	bool
	Weekday  		bool
	Weekend  		bool
	Private 		bool
	PassCount 		int64
    TotalCount		int64
    BestStreak		int64
    CurrentStreak 	int64
}