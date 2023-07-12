package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPaginate(t *testing.T) {
	mockDB, _ := gorm.Open(nil, &gorm.Config{})
	mockDB2, _ := gorm.Open(nil, &gorm.Config{})
	paginatedFunc := Paginate(1, 10)
	paginatedDB := paginatedFunc(mockDB)
	mockDB2.Offset(0).Limit(10)
	assert.Equal(t, mockDB2.Statement.SQL, paginatedDB.Statement.SQL)
}