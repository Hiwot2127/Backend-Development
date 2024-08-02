package models

import "time"

// returning a slice of initial mock tasks.
func MockData() []Task {
	return []Task{
		{
			ID:          1,
			Title:       "Complete Homework",
			Description: "Finish math and science homework",
			DueDate:     time.Now().AddDate(0, 0, 3), // Due in 3 days
			Status:      "Pending",
		},
		{
			ID:          2,
			Title:       "Grocery Shopping",
			Description: "Buy milk, eggs, and bread",
			DueDate:     time.Now().AddDate(0, 0, 1), // Due tomorrow
			Status:      "Pending",
		},
		{
			ID:          3,
			Title:       "Clean the House",
			Description: "Vacuum and dust all rooms",
			DueDate:     time.Now().AddDate(0, 0, 5), // Due in 5 days
			Status:      "Pending",
		},
		{
			ID:          4,
			Title:       "Finish Project Report",
			Description: "Complete the final project report for school",
			DueDate:     time.Now().AddDate(0, 0, 2), // Due in 2 days
			Status:      "In Progress",
		},
		{
			ID:          5,
			Title:       "Exercise",
			Description: "Go for a 30-minute run",
			DueDate:     time.Now().AddDate(0, 0, 1), // Due tomorrow
			Status:      "Completed",
		},
	}
}
