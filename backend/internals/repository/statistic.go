package repository

import "database/sql"

func GetStudentClassesNum(db *sql.DB, userID int) (int, error) {
	query := "SELECT COUNT(class_id) from students_of_classes where student_id = ?"

	row := db.QueryRow(query, userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetTeachingClassesNum(db *sql.DB, userID int) (int, error) {
	query := "SELECT COUNT(*) from classes where teacher_id = ?"

	row := db.QueryRow(query, userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func GetAllStudentGrades(db *sql.DB, userID int) ([]int, error) {
	query := "SELECT SUM(mark) FROM marks where student_id = ? GROUP BY class_id"

	rows, err := db.Query(query, userID)
	if err != nil {
		return []int{}, err
	}

	defer rows.Close()

	var mark int
	var marks []int

	for rows.Next() {
		err := rows.Scan(&mark)
		if err != nil {
			return []int{}, err
		}
		marks = append(marks, mark)
	}

	return marks, nil
}

func GetAllStudentGradesSum(db *sql.DB, userID int) (float32, error) {
	query := "SELECT mark FROM marks WHERE student_id = ?"

	rows, err := db.Query(query, userID)
	if err != nil {

		return 0, err
	}

	defer rows.Close()

	var mark float32
	var sum float32 = 0.
	for rows.Next() {
		err := rows.Scan(&mark)
		if err != nil {
			return 0, err
		}
		sum += mark
	}

	return sum, nil
}
