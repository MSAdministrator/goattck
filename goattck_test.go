package goattck_test

import (
	"testing"

	"github.com/msadministrator/goattck"
	"github.com/stretchr/testify/assert"
)

const attackURL = "https://raw.githubusercontent.com/mitre/cti/master/enterprise-attack/enterprise-attack.json"

func TestEnterprise_New(t *testing.T) {
	enterprise, err := goattck.Enterprise{}.New(attackURL)
	if err != nil {
		t.Errorf("Error, could not load Enterprise: %v", err)
	}
	assert.IsType(t, goattck.Enterprise{}, enterprise)
	enterprise, err = enterprise.Load(false)
	assert.Nil(t, err)
	assert.Greater(t, len(enterprise.Actors), 20)
	assert.Greater(t, len(enterprise.Campaigns), 5)
	assert.Greater(t, len(enterprise.DataComponents), 5)
	assert.Greater(t, len(enterprise.DataSources), 5)
	assert.Equal(t, len(enterprise.Defintions), 1)
	assert.Greater(t, len(enterprise.Malwares), 20)
	assert.Equal(t, len(enterprise.Matrices), 1)
	assert.Greater(t, len(enterprise.Mitigations), 5)
	assert.Greater(t, len(enterprise.Relationships), 5)
	assert.Equal(t, len(enterprise.Tactics), 14)
	assert.Greater(t, len(enterprise.Techniques), 200)
	assert.Greater(t, len(enterprise.Tools), 20)

	fakeURL := "hxxps://test.test.test/enterprise-legacy/enterprise-legacy.json"
	e, err := goattck.Enterprise{}.New(fakeURL)
	assert.NotNil(t, e)
	assert.Nil(t, err)
}
