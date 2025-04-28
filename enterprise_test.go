package goattck_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/msadministrator/goattck"
)

func TestEnterprise_buildRelationshipMap(t *testing.T) {
	enterprise, err := goattck.Enterprise{}.New(attackURL)
	assert.Nil(t, err)
	enterprise, err = enterprise.Load(false)
	assert.Nil(t, err)
	assert.NotNil(t, enterprise)

	assert.Equal(t, 181, len(enterprise.Actors))
}
