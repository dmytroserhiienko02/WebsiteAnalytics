package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Record struct {
	UserID    int
	ProductID int
	Timestamp time.Time
}

func GetRecords(filePath string) map[int]map[int]bool {
	result := make(map[int]map[int]bool)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	csvReader := csv.NewReader(file)

	if _, err := csvReader.Read(); err != nil {
		fmt.Println("Error reading header:", err)
		return nil
	}

	for {
		rec, err := csvReader.Read()
		if err != nil {
			break
		}
		record, err := parseRec(rec)

		_, ok := result[record.UserID]
		if !ok {
			product := make(map[int]bool)
			product[record.ProductID] = true
			result[record.UserID] = product
		} else {
			result[record.UserID][record.ProductID] = true
		}

		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			continue
		}

		fmt.Printf("Parsed Record: %+v\n", record)
	}
	return result
}

func parseRec(rec []string) (Record, error) {
	var record Record
	var err error
	record.UserID = parseInt(rec[0])
	record.ProductID = parseInt(rec[1])
	record.Timestamp, err = time.Parse("2006-01-02T15:04:05", rec[2])
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return record, err
	}
	return record, err
}

func parseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func DoubleVisited(firstDay map[int]map[int]bool, secondDay map[int]map[int]bool) []int {
	fmt.Println("Visited some pages on both days:")
	visitedClients := make([]int, 0, len(firstDay))
	for i := range firstDay {
		_, ok := secondDay[i]
		if ok {
			fmt.Printf("UserID %d\n", i)
			visitedClients = append(visitedClients, i)
		}
	}
	return visitedClients
}

func VisitedSecondDay(firstDay map[int]map[int]bool, secondDay map[int]map[int]bool, visited []int) {
	fmt.Println("On the second day visited the page that hadn’t been visited by this user on the first day:")
	for _, i := range visited {
		for j := range secondDay[i] {
			_, ok := firstDay[i][j]
			if !ok {
				fmt.Printf("UserID %d\n", i)
			}
		}
	}
}

func main() {
	strDay1 := "days/day1.csv"
	strDay2 := "days/day2.csv"

	firstDay := GetRecords(strDay1)
	secondDay := GetRecords(strDay2)

	visited := DoubleVisited(firstDay, secondDay)
	VisitedSecondDay(firstDay, secondDay, visited)
}
