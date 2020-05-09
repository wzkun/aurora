package decode

// ToStringArray function.
func ToStringArray(objs ...interface{}) []string {
	records := make([]string, 0, len(objs))
	for _, obj := range objs {
		js, _ := JSON.Marshal(obj)
		record := string(js)
		records = append(records, record)
	}
	return records
}

// ToString function.
func ToString(obj interface{}) string {
	js, _ := JSON.Marshal(obj)
	record := string(js)
	return record
}
