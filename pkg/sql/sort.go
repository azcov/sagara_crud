package sql

import (
	"fmt"
	"strings"
)

func GenerateOrderByClause(orderBy string, availableColumnForOrder map[string]string, defaultQuery string) string {
	orderByClause := ""
	stringSplitted := strings.Split(orderBy, ",")

	for idx := range stringSplitted {
		stringSplitted[idx] = strings.TrimSpace(stringSplitted[idx])
		columnName := strings.Trim(stringSplitted[idx], "-")
		columnName = strings.Trim(columnName, "+")
		if stringSplitted[idx] != "" && CheckColumnCanOrder(availableColumnForOrder, columnName) {
			if idx == 0 {
				orderByClause += "ORDER BY "
			}

			columnName = availableColumnForOrder[columnName]
			columnOrder := fmt.Sprintf("%s %s", columnName, "asc")
			if stringSplitted[idx][0:1] == "-" {
				columnOrder = fmt.Sprintf("%s %s", columnName, "desc")
			}

			orderByClause += columnOrder
			if idx != len(stringSplitted)-1 && strings.TrimSpace(stringSplitted[idx+1]) != "" {
				orderByClause += ", "
			}
		}
	}

	if len(orderByClause) == 0 {
		orderByClause += fmt.Sprintf("ORDER BY %s", defaultQuery)
	}

	return orderByClause
}

func CheckColumnCanOrder(columnList map[string]string, requestColumn string) bool {
	if _, ok := columnList[strings.ToLower(requestColumn)]; ok {
		return true
	}
	return false
}

// func CheckReplaceColumnName(columnName string, replaceColumn map[string]string) string {
// 	for idx := range replaceColumn {
// 		if strings.EqualFold(strings.ToLower(idx), strings.ToLower(columnName)) {
// 			return replaceColumn[idx]
// 		}
// 	}
// 	return columnName
// }

func GenerateOffsetCluse(page, limit int32) (offset int32) {
	return (page - 1) * limit
}
