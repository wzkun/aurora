package mysql

// MakeMysqlPage
func MakeMysqlPage(pageSize, pageIndex int32) (limit int, off int) {
	offset := int32(0)
	if pageSize <= 0 {
		pageSize = 100
	}

	if pageIndex > 0 {
		offset = (pageIndex - 1) * pageSize
	}

	return int(pageSize), int(offset)
}
