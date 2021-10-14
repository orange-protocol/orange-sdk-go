package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/orange-protocol/orange-sdk-go/graphql"
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
	req.Var("apdid", apdid)
	resp := &GetAlgorithmProviderMethodResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetAlgorithmProviderMethods, err
}
func (sdk *OscoreSDK) GetDataMethods(dpdid string) ([]*ProviderMethod, error) {
	req := graphql.NewRequest(GetDataProviderMethodsReq)
	req.Var("dpdid", dpdid)
	resp := &GetDataProviderMethodResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetDataProviderMethods, err
}

func (sdk *OscoreSDK) RequestOscore(roreq *RequestOscoreReq) (int64, error) {
	tmps := getRequestOscoreReqStr(roreq)
	fmt.Printf("%s\n", tmps)
	req := graphql.NewRequest(tmps)

	tmp, _ := json.Marshal(req.Vars())
	fmt.Printf("vars:%s\n", tmp)

	resp := &GetOscoreResp{}
	err := sdk.sendRequest(req, resp)
	return resp.Oscore, err
}

func (sdk *OscoreSDK) GetUserTask(key string, taskId int64) (*UserTasks, error) {
	req := graphql.NewRequest(GetUserTask)
	req.Var("key", key)
	req.Var("taskId", taskId)

	resp := &GetUserTaskResp{}
	err := sdk.sendRequest(req, resp)

	return resp.GetUserTask, err
}

func getRequestOscoreReqStr(req *RequestOscoreReq) string {
	s := "mutation{requestOscore(input:{appdid:\"%s\",data:{userdid:\"%s\",apdid:\"%s\",apmethod:\"%s\",dpdid:\"%s\",dpmethod:\"%s\",overwriteOld:%v,wallets:[$walletsinfo$]},sig:\"%s\"})}"
	str := ""
	for _, w := range req.Data.Wallets {
		if len(str) == 0 {
			str = str + fmt.Sprintf("{chain:\"%s\",address:\"%s\",pubkey:\"%s\",sig:\"%s\"}", w.Chain, w.Address, w.Pubkey, w.Sig)
		} else {
			str = str + "," + fmt.Sprintf("{chain:\"%s\",address:\"%s\",pubkey:\"%s\",sig:\"%s\"}", w.Chain, w.Address, w.Pubkey, w.Sig)
		}
	}
	s = strings.ReplaceAll(s, "$walletsinfo$", str)
	return fmt.Sprintf(s, req.AppDid, req.Data.Userdid, req.Data.Apdid, req.Data.Apmethod, req.Data.Dpdid, req.Data.Dpmethod, req.Data.OverwriteOld, req.Sig)
}

func (sdk *OscoreSDK) sendRequest(req *graphql.Request, resp interface{}) error {
	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	// define a Context for the request
	ctx := context.Background()
	return sdk.client.Run(ctx, req, &resp)
}
