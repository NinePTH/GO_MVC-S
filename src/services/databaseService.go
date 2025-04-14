package services

import (
	"fmt"
	"strings"

	"github.com/NinePTH/GO_MVC-S/src/utils/databaseConnector"
)

// select distinct,etc.
func SelectInnerJoin(mainTable string,joinTable string,joinCondition string,fields []string,where bool,whereCon string,whereArgs []interface{},) ([]map[string]interface{}, error) {

	var query string = "SELECT "

	// เติม field ที่ต้องการ
	for _, field := range fields {
		query += field + ", "
	}
	query = strings.TrimRight(query, ", ")
	query += " FROM " + mainTable

	// ถ้า joinCondition ว่าง แปลว่าเขียน join ทั้งหมดไว้ใน joinTable แล้ว
	if joinCondition != "" {
		query += " INNER JOIN " + joinTable + " ON " + joinCondition
	} else {
		query += " INNER JOIN " + joinTable
	}

	// ถ้ามีเงื่อนไข WHERE
	if where {
		query += " WHERE " + whereCon
	}

	// log query
	fmt.Println("Executing query:", query)

	// run query
	rows, err := databaseConnector.DB.Query(query, whereArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, column := range columns {
			rowMap[column] = values[i]
		}
		results = append(results, rowMap)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("results = %v\n\n", results)
	return results, nil
}



func SelectData(table string, fields []string, where bool, whereCon string, whereArgs []interface{}) ([]map[string]interface{}, error) {
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

	rows, err := databaseConnector.DB.Query(query, whereArgs...)
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
	fmt.Printf("results = %v\n\n", results)

	// Check if there were any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
func UpdateData(table string, data map[string]interface{}, condition string, conditionValues []interface{}) (int64, error) {
	var setClauses []string
	var values []interface{}
	// Start by appending the values for condition
	values = append(values, conditionValues...)

	// Construct the SET clause and add placeholders
	// for example: "name = $1, age = $2"
	for column, value := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, len(values)+1))
		values = append(values, value)
	}

	// Construct the full query with the correct placeholders for PostgreSQL
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(setClauses, ", "), condition)

	fmt.Println("Executing query:", query)
	fmt.Println("With values:", values)

	// Prepare the statement
	stmt, err := databaseConnector.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement
	result, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}


func InsertData(table string, data map[string]interface{}) (int64, error) {
	var columns []string
	var placeholders []string
	var values []interface{}

	for column, value := range data {
		columns = append(columns, column)
		values = append(values, value)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(placeholders)+1))
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	fmt.Println("Executing query:", query)

	// Prepare the statement
	stmt, err := databaseConnector.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement
	result, err := stmt.Exec(values...)
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteData(table string, condition string, conditionValues []interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, condition)
	fmt.Println("Executing query:", query)

	// Prepare the statement
	stmt, err := databaseConnector.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(conditionValues...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
