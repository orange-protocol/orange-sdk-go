package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOscoreSDK_GetAlgorithmProviders(t *testing.T) {
	sdk,err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t,err)
	aps,err := sdk.GetAlgorithmProviders()
	assert.Nil(t,err)
	assert.NotNil(t,aps)
	assert.Greater(t,len(aps),0)

}
