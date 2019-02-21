package pojo

import "gitlab.com/z547743799/irissearch/models"

type SearchResult struct {
	RecordCount int64
	TotalPages  int
	ItemList    []models.SearchItem
}
