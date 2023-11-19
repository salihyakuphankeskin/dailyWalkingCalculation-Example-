package main

import (
	"fmt"
)

func main() {
	ExerciseMap := map[int]int{}
	var totalExercise int = 0
	var totalExerciseDay int = 0

	for day := 0; ; day++ {
		var walkMeasure int;
		fmt.Println("How many meter did you walk today?")	
		_, err :=fmt.Scanln(&walkMeasure)
		if err != nil {
			break
		}
		ExerciseMap[day]= walkMeasure
		
		totalExercise += walkMeasure
		totalExerciseDay += 1
	}	
	fmt.Println("")
	for i := 0; i < len(ExerciseMap); i++ {
		firstText := fmt.Sprintf("The %d. day \t",i)
		fmt.Print(firstText)
		ExerciseMapText := 	fmt.Sprintf("%d meter", ExerciseMap[i])		
		fmt.Println(ExerciseMapText)		
	}
	endText := fmt.Sprintf("In %d days %d meters has been done.",totalExerciseDay, totalExercise)
	fmt.Println("")
	fmt.Println(endText)

}