package sdk

import (
	"context"
	"encoding/json"
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
	//todo currently struct input param cannot be passed
	//req := graphql.NewRequest(GetOscoreReq)
	//reqjson ,err:= json.Marshal(req)
	//if err != nil{
	//	return  -1,err
	//}
	//req.Var("data",roreq)
	//var resp int64 = -1
	//err := sdk.sendRequest(req, resp)
	//return resp, err

	tmps := getRequestOscoreReqStr(roreq)
	fmt.Printf("%s\n", tmps)
	req := graphql.NewRequest(tmps)

	req.Var("key", roreq.Key)
	req.Var("did", roreq.Did)
	req.Var("apdid", roreq.Apdid)
	req.Var("apmethod", roreq.Apmethod)
	req.Var("dpdid", roreq.Dpdid)
	req.Var("dpmethod", roreq.Dpmethod)
	req.Var("overwriteOld", roreq.overwriteOld)
	for i, wallet := range roreq.Wallets {
		req.Var(fmt.Sprintf("chain-%d", i), wallet.Chain)
		req.Var(fmt.Sprintf("address-%d", i), wallet.Address)
		req.Var(fmt.Sprintf("pubkey-%d", i), wallet.Pubkey)
		req.Var(fmt.Sprintf("sig-%d", i), wallet.Sig)
	}
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
	s := "mutation{requestOscore(input:{key:\"%s\",did:\"%s\",apdid:\"%s\",apmethod:\"%s\",dpdid:\"%s\",dpmethod:\"%s\",overwriteOld:%v,wallets:[$walletsinfo$]})}"
	str := ""
	for _, w := range req.Wallets {
		if len(str) == 0 {
			str = str + fmt.Sprintf("{chain:\"%s\",address:\"%s\",pubkey:\"%s\",sig:\"%s\"}", w.Chain, w.Address, w.Pubkey, w.Sig)
		} else {
			str = str + "," + fmt.Sprintf("{chain:\"%s\",address:\"%s\",pubkey:\"%s\",sig:\"%s\"}", w.Chain, w.Address, w.Pubkey, w.Sig)
		}
	}
	s = strings.ReplaceAll(s, "$walletsinfo$", str)
	return fmt.Sprintf(s, req.Key, req.Did, req.Apdid, req.Apmethod, req.Dpdid, req.Dpmethod, req.overwriteOld)
}

func (sdk *OscoreSDK) sendRequest(req *graphql.Request, resp interface{}) error {
	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	// define a Context for the request
	ctx := context.Background()
	return sdk.client.Run(ctx, req, &resp)
}
