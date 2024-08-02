package main

import (
	"fmt"
	"strconv"
)
//Student Grade Calculator

// Student struct with name, grades, and average grade
type Student struct {
	Name       string
	Grades     map[string]float64
	AverageGrade float64
}

func main() {
	var student Student

	//Accepting student user inputs with appropriate error handling
	fmt.Println("Student Grade Calculator, Calculate Your Grades Here!")

	for {
		fmt.Print("Enter your name: ")
		_, err := fmt.Scanln(&student.Name)
		if err == nil && student.Name != "" {
			break
		}
		fmt.Println("Invalid input. Please enter a valid name.")
	}

	fmt.Print("Enter the number of subjects you have taken: ")
	var numSubjectsStr string
	fmt.Scanln(&numSubjectsStr)
	numSubjects, err := strconv.Atoi(numSubjectsStr)
	if err != nil || numSubjects <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}
	// a map to store grades
	student.Grades = make(map[string]float64)

	// Loop for each subject
	for i := 1; i <= numSubjects; i++ {
		var subject string

		fmt.Printf("Enter the name of subject %d: ", i)
		fmt.Scanln(&subject)

		// Check if the subject name is empty
		if subject == "" {
			fmt.Println("Invalid input. Please enter a valid subject name.")
			continue
		}

		for {
			fmt.Printf("Enter the numeric grade for %s: ", subject)
			var gradeInput string
			fmt.Scanln(&gradeInput)

			// Check if the grade input is empty
			if gradeInput == "" {
				fmt.Println("Invalid input. Please enter a valid numeric grade.")
				continue
			}

			// Check if the grade input is a valid number
			grade, err := strconv.ParseFloat(gradeInput, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid numeric grade.")
				continue
			}

			// Check if the grade is within the valid range
			if grade < 0 || grade > 100 {
				fmt.Println("Invalid grade entered. Please enter a value between 0 and 100.")
				continue
			}

			// Adding the grade to the student's grades map
			student.Grades[subject] = grade
			break
		}
	}

	// Calculate the average grade
	student.AverageGrade = calculateAverageGrade(student.Grades)

	// Printing the student's grade summary
	fmt.Println("\n--- Student Grade Summary ---")
	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Println("Subject Grades:")
	for subject, grade := range student.Grades {
		fmt.Printf("- %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", student.AverageGrade)
}

// calculate the average grade of a student
func calculateAverageGrade(grades map[string]float64) float64 {
	var total float64
	count := len(grades)
	for _, grade := range grades {
		total += grade
	}
	return total / float64(count)
}
