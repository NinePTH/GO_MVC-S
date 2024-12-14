package services

import (
	"learn-Go/src/database"
	"fmt"
	"strings"
)

func SelectData(table string, fields []string, where bool, whereCon string) ([]map[string]interface{}, error) {
	var query string = "SELECT "

	for _, field := range fields {
		query += field + ", "
	}

	query = strings.TrimRight(query, ", ")

	query += " FROM " + table

	if where {
		query += " WHERE " + whereCon
	}

	fmt.Println("Executing query:", query)

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	// Hard part again krub pom
	// Iterate over the rows
	for rows.Next() {
		// Create a values for store value that we want and valuePointers for using with scan method 
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))

		// Assign value address to value pointer
		for i := range values {
			valuePointers[i] = &values[i]
		}

		// Scan the row value into the value pointers to assign the value into values variable
		err := rows.Scan(valuePointers...)
		if err != nil {
			return nil, err
		}

		// Map the column names to the values
		rowMap := make(map[string]interface{})
		for i, column := range columns {
			// Add each column and its corresponding value to the map
			rowMap[column] = values[i]
		}

		// Append the row map to the results slice
		results = append(results, rowMap)
		// fmt.Println("rowMap =",rowMap)
	}
	fmt.Printf("results = %v\n\n",results)

	// Check if there were any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}