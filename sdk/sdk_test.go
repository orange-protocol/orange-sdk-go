package sdk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOscoreSDK_GetAlgorithmProviders(t *testing.T) {
	sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	aps, err := sdk.GetAlgorithmProviders()
	assert.Nil(t, err)
	assert.NotNil(t, aps)
	assert.Greater(t, len(aps), 0)

}

func TestOscoreSDK_GetDataProviders(t *testing.T) {
	sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	aps, err := sdk.GetDataProviders()
	assert.Nil(t, err)
	assert.NotNil(t, aps)
	assert.Greater(t, len(aps), 0)
}

func TestOscoreSDK_GetAlgorithmMethods(t *testing.T) {
	sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	aps, err := sdk.GetAlgorithmProviders()
	assert.Nil(t, err)
	assert.NotNil(t, aps)
	assert.Greater(t, len(aps), 0)
	apdid := aps[0].Did
	fmt.Printf("apdid:%s\n", apdid)

	methods, err := sdk.GetAlgorithmMethods(apdid)
	assert.Nil(t, err)
	assert.NotNil(t, methods)
	assert.Greater(t, len(methods), 0)
	fmt.Printf("method:%s\n", methods[0].Name)
}

func TestOscoreSDK_GetDataMethods(t *testing.T) {
	sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	dps, err := sdk.GetDataProviders()
	assert.Nil(t, err)
	assert.NotNil(t, dps)
	assert.Greater(t, len(dps), 0)
	dpdid := dps[0].Did
	fmt.Printf("dpdid:%s\n", dpdid)

	methods, err := sdk.GetDataMethods(dpdid)
	assert.Nil(t, err)
	assert.NotNil(t, methods)
	assert.Greater(t, len(methods), 0)
	for _, m := range methods {
		fmt.Printf("method:%s\n", m.Name)
	}
}

func TestOscoreSDK_RequestOscore(t *testing.T) {
	//sdk, err := NewOscoreSDK("http://172.168.3.38:8080/query")
	sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	req := &RequestOscoreReq{
		AppDid: "did:ont:ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev",
		Data: RequestOscoreData{
			Userdid:      "did:ont:AGAMr5P2Ngi7SGvhKd3s5vWTWpid5uGywL",
			Apdid:        "did:ont:testap",
			Apmethod:     "calc30xWithDefi",
			Dpdid:        "did:ont:abcdefg",
			Dpmethod:     "queryXdaysSumWithDefi",
			OverwriteOld: true,
			Wallets: []*UserWallet{{
				Chain:   "eth",
				Address: "0x45929D79A6DDdaA3C8154D4F245d17d1D80DbBcc",
				Pubkey:  "HN6l5UfFXVd4GHcB3HDeO13Iu6N7uDjG62kQmD2zbUw=",
				Sig:     "0xc0cd6419d10fc3dcf1483b20f69c9b20c7ee44208868399dda50184305370be00fa5b8aacdf51fd56027efac09ac0d997b71f745d7383bb6a49c4c7d0d05d8371c",
			},
			},
		},
		Sig: "43319a9c43d777b32b8320dbdfce32bb61c8314f81b10c838ad57bde57306ddf38cfabd64c33212d60f73e051534967015cd3222675892332a6ba5c521f59bba",
	}

	//tmp ,_:= json.Marshal(req.Data)
	//tmphex := hex.EncodeToString(tmp)
	//fmt.Printf("%s",tmphex)

	taskid, err := sdk.RequestOscore(req)
	assert.Nil(t, err)
	fmt.Printf("taskid:%d\n", taskid)
}

func TestNewOscoreSDK_GetUserTask(t *testing.T) {
	sdk, err := NewOscoreSDK("http://172.168.3.38:8080/query")
	//sdk, err := NewOscoreSDK("http://localhost:8080/query")
	assert.Nil(t, err)

	task, err := sdk.GetUserTask("key", 123)
	assert.Nil(t, err)
	assert.NotNil(t, task)
	fmt.Printf("%v\n", task)
	fmt.Printf("%s\n", *task.ResultFile)
	fmt.Printf("%s\n", *task.TaskResult)

}
