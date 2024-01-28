package test

import (
	"testing"

	"github.com/adibhauzan/azahri_mart_be/app"
	"github.com/stretchr/testify/assert"
)

func TestDbConnetion(t *testing.T) {
	db := app.NewDbConnection()

	assert.NotNil(t, db)
}
