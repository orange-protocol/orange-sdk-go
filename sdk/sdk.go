package sdk

import (
	"context"
	"fmt"
	"strings"

	"github.com/oscore/oscore-sdk-go/graphql"
)

type OscoreSDK struct {
	client *graphql.Client
}

func NewOscoreSDK(url string) (*OscoreSDK, error) {
	return &OscoreSDK{client: graphql.NewClient(url)}, nil
}

func (sdk *OscoreSDK) GetAlgorithmProviders() ([]*AlgorithmProvider, error) {
	req := graphql.NewRequest(GetAllAlgorithmProvidersReq)
	resp := &GetAlgorithmProvidersResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetAllAlgorithmProviders, err
}

func (sdk *OscoreSDK) GetDataProviders() ([]*DataProvider, error) {
	req := graphql.NewRequest(GetAllDataProvidersReq)
	resp := &GetDataProvidersResp{}
	err := sdk.sendRequest(req, resp)
	if err != nil {
		return nil, err
	}
	return resp.GetAllDataProviders, nil
}

func (sdk *OscoreSDK) GetAlgorithmMethods(apdid string) ([]*ProviderMethod, error) {
	req := graphql.NewRequest(GetAlgorithmProviderMethodsReq)
	req.Var("did", apdid)
	resp := &GetAlgorithmProviderMethodResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetAlgorithmProviderMethods, err
}
func (sdk *OscoreSDK) GetDataMethods(dpdid string) ([]*ProviderMethod, error) {
	req := graphql.NewRequest(GetDataProviderMethodsReq)
	req.Var("did", dpdid)
	resp := &GetDataProviderMethodResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetDataProviderMethods, err
}

func (sdk *OscoreSDK) RequestOscore(roreq *RequestOscoreReq) (int64, error) {
	req := graphql.NewRequest(getRequestOscoreReqStr(roreq.Wallets))

	req.Var("key", roreq.Key)
	req.Var("did", roreq.Did)
	req.Var("apdid", roreq.Apdid)
	req.Var("apmethod", roreq.Apmethod)
	req.Var("dpdid", roreq.Dpdid)
	req.Var("dpmethod", roreq.Dpmethod)
	req.Var("overwriteOld",roreq.overwriteOld)
	for i, wallet := range roreq.Wallets {
		req.Var(fmt.Sprintf("chain-%d", i), wallet.Chain)
		req.Var(fmt.Sprintf("address-%d", i), wallet.Address)
		req.Var(fmt.Sprintf("pubkey-%d", i), wallet.Pubkey)
		req.Var(fmt.Sprintf("sig-%d", i), wallet.Sig)
	}
	var resp int64 = -1
	err := sdk.sendRequest(req, resp)
	return resp, err
}

func (sdk *OscoreSDK) GetUserTask(key string, taskId int64) (*UserTasks, error) {
	req := graphql.NewRequest(GetUserTask)
	req.Var("key", key)
	req.Var("taskId", taskId)

	resp := &GetUserTaskResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetUserTask, err
}

func getRequestOscoreReqStr(wallets []*UserWallet) string {

	str := ""
	for i, _ := range wallets {
		if len(str) == 0 {
			str = str + fmt.Sprintf("{chain:$chain-%d,address:$address-%d,pubkey:$pubkey-%d,sig:$sig-%d}", i, i, i, i)
		} else {
			str = str + "," + fmt.Sprintf("{chain:$chain-%d,address:$address-%d,pubkey:$pubkey-%d,sig:$sig-%d}", i, i, i, i)
		}
	}
	return strings.Replace(GetOscoreReq, "%walletsinfo%", str, 0)
}

func (sdk *OscoreSDK) sendRequest(req *graphql.Request, resp interface{}) error {
	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	// define a Context for the request
	ctx := context.Background()
	return sdk.client.Run(ctx, req, &resp)
}
