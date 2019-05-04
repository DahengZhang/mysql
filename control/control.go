package control

type Test struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func InsertDB(title, content string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO test_table (title, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	stmtIns.Exec(title, content)

	return nil
}
func SelectDB(o_id int) (Test, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM test_table WHERE id=?")
	if err != nil {
		return Test{}, err
	}
	defer stmtOut.Close()

	var id int
	var title, content, create_time, update_time string

	err = stmtOut.QueryRow(o_id).Scan(&id, &title, &content, &create_time, &update_time)
	if err != nil {
		return Test{}, err
	}

	return Test{
		Id:         id,
		Title:      title,
		Content:    content,
		CreateTime: create_time,
		UpdateTime: update_time,
	}, nil
}

func SelectAll() ([]Test, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM test_table")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	result := []Test{}

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var title, content, create_time, update_time string
		err := rows.Scan(&id, &title, &content, &create_time, &update_time)
		if err != nil {
			return nil, err
		}

		result = append(result, Test{
			Id:         id,
			Title:      title,
			Content:    content,
			CreateTime: create_time,
			UpdateTime: update_time,
		})
	}

	return result, nil
}

func UpdateDB(id int, title, content string) error {
	stmtUpd, err := dbConn.Prepare("UPDATE test_table SET title=?, content=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmtUpd.Close()

	_, err = stmtUpd.Exec(title, content, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteDB(id int) error {
	_, err = dbConn.Query("DELETE FROM test_table WHERE id=" + string(id))
	if err != nil {
		return err
	}

	return nil
}
