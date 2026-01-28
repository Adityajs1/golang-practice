package main

import "fmt"

type Student struct{
    Name  string
    Age   int
    Grade []int
}

func addGrade(student *Student, grade int) {
    student.Grade = append(student.Grade, grade)
}

func averageGrade(student Student) float64{
    if len(student.Grade) == 0{
        return 0.0
    }
    sum := 0
    for _, grade := range student.Grade {
        sum += grade
    }
    average := float64(sum) / float64(len(student.Grade))
    return average
}

func findTopStudent(students []Student) Student {
    topStudent := students[0]
    highestAverage := averageGrade(topStudent)
    
    for i := 1; i < len(students); i++ {
        currentAverage := averageGrade(students[i])
        
        if currentAverage > highestAverage {
            highestAverage = currentAverage
            topStudent = students[i]
        }
    }
    
    return topStudent
}

func studentGrade(){
    s := Student{
        Name: "Aditya",
        Age: 18,
        Grade: []int{70, 42, 56},
    }

    s1 := Student{
        Name: "Chitranshi",
        Age: 18,
        Grade: []int{72, 90, 78},
    }

    s2 := Student{
        Name: "Arnav",
        Age: 19,
        Grade: []int{78, 48, 50},
    }

    addGrade(&s, 78)
    addGrade(&s1, 70)
    addGrade(&s2, 65)

    // Create slice of students
    students := []Student{s, s1, s2}
    
    // Print each student's average
    for _, student := range students {
        avg := averageGrade(student)
        fmt.Printf("%s's average: %.2f\n", student.Name, avg)
    }
    
    // Find and print top student
    topStudent := findTopStudent(students)
    topAvg := averageGrade(topStudent)
    fmt.Printf("Top student: %s with average %.2f\n", topStudent.Name, topAvg)
}