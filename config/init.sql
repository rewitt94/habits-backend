CREATE TABLE habits (
    id SERIAL,
    habit VARCHAR(100),
    category VARCHAR (25),
    condition VARCHAR(100),
    score INT,
    single_action BOOLEAN,
    weekday BOOLEAN,
    weekend BOOLEAN,
    private BOOLEAN,
    pass_count INT,
    total_count INT,
    best_streak INT,
    current_streak INT
)

CREATE TABLE results (
    id SERIAL,
    habit_id INT,
    success BOOLEAN,
    date DATE
)

CREATE INDEX date_index
ON results (date)

CREATE INDEX habit_index
ON results (habit_id)