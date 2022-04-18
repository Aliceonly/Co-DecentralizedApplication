package utils
func Selectalldate() int {
	var sqlStr=` SELECT COUNT(transationNumber) FROM alldate`
	//alldate,err :=db.Prepare(sqlStr)
	rows,err :=db.Query(sqlStr)
	if err!=nil {
		panic(err)
	}
	var date int
	for rows.Next() {
		err :=rows.Scan(&date)
		if err!=nil {
			panic(err)
		}
		fmt.Println(date)
	}
	return date
}