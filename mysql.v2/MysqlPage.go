package mysql

func MakeMysqlPage(pageSize, pageIndex int32) (int, int) {
	offset := int32(0)
	if pageSize <= 0 {
		pageSize = 100
	}

	if pageIndex > 0 {
		offset = (pageIndex - 1) * pageSize
	}

	return int(pageSize), int(offset)
}

type PageMySql struct {
	PageSize int
	Offset   int
}

// MysqlMap 构建分页信息
func MysqlMap(count int64) ([][]*PageMySql, error) {
	offset := 0

	s := int(count / 10000)
	yu := int(count % 10000)

	pagelist := make([][]*PageMySql, 0)
	if s < 1 {
		pages := make([]*PageMySql, 0)
		for {
			if count > 0 {
				page := &PageMySql{}
				page.PageSize = 1000
				page.Offset = offset

				pages = append(pages, page)
			}

			if count <= 1000 {
				break
			}

			count = count - 1000
			offset = offset + 1000
		}

		pagelist = append(pagelist, pages)
	} else {
		for i := 1; i <= s; i++ {
			pages := make([]*PageMySql, 0)
			for j := 1; j <= 10; j++ {
				page := &PageMySql{}
				page.PageSize = 1000
				page.Offset = offset

				pages = append(pages, page)
				count = count - 1000
				offset = offset + 1000
			}
			pagelist = append(pagelist, pages)
		}
	}

	if s > 1 && yu > 0 {
		pages := make([]*PageMySql, 0)
		for {
			if count > 0 {
				page := &PageMySql{}
				page.PageSize = 1000
				page.Offset = offset

				pages = append(pages, page)
			}

			if count <= 1000 {
				break
			}

			count = count - 1000
			offset = offset + 1000
		}

		pagelist = append(pagelist, pages)
	}

	return pagelist, nil
}
