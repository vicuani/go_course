package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

func readUsers(reader io.Reader) ([]User, error) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var users []User
	for i, record := range records {
		if i == 0 {
			continue
		}

		age, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		tripsCount, err := strconv.Atoi(record[6])
		if err != nil {
			return nil, err
		}

		user := User{
			FirstName:  record[0],
			LastName:   record[1],
			Email:      record[2],
			Age:        age,
			Gender:     record[4],
			City:       record[5],
			TripsCount: tripsCount,
			Profession: record[7],
		}
		users = append(users, user)
	}
	return users, nil
}

func readStatistics(reader io.Reader) ([]Statistics, error) {
	var stats []Statistics

	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] {
		averageTrips, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		stats = append(stats, Statistics{
			City:         record[0],
			AgeRange:     record[1],
			AverageTrips: averageTrips,
		})
	}

	return stats, nil
}
