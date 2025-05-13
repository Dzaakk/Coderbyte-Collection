package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	currRow := 0
	goingDown := false

	for _, c := range s {
		rows[currRow] += string(c)

		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
