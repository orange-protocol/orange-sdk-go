package sdk

type AlgorithmProvider struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Introduction    string `json:"introduction"`
	Did             string `json:"did"`
	CreateTime      int64  `json:"createTime"`
	Title           string `json:"title"`
	Provider        string `json:"provider"`
	InvokeFrequency int64  `json:"invokeFrequency"`
	APIState        int64  `json:"apiState"`
	Author          string `json:"author"`
	Popularity      int64  `json:"popularity"`
	Delay           int64  `json:"delay"`
	Icon            string `json:"icon"`
}

type DataProvider struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Introduction    string `json:"introduction"`
	Did             string `json:"did"`
	CreateTime      int64  `json:"createTime"`
	Title           string `json:"title"`
	Provider        string `json:"provider"`
	InvokeFrequency int64  `json:"invokeFrequency"`
	APIState        int64  `json:"apiState"`
	Author          string `json:"author"`
	Popularity      int64  `json:"popularity"`
	Delay           int64  `json:"delay"`
	Icon            string `json:"icon"`
}

type ProviderMethod struct {
	Name         string `json:"name"`
	ParamSchema  string `json:"paramSchema"`
	ResultSchema string `json:"resultSchema"`
}

type RequestOscoreData struct {
	Userdid      string        `json:"userDid"`
	Apdid        string        `json:"apDid"`
	Apmethod     string        `json:"apMethod"`
	Dpdid        string        `json:"dpDid"`
	Dpmethod     string        `json:"dpMethod"`
	OverwriteOld bool          `json:"overwriteOld"`
	Wallets      []*UserWallet `json:"wallets"`
}

type RequestOscoreReq struct {
	AppDid string            `json:"appDid"`
	Data   RequestOscoreData `json:"data"`
	Sig    string            `json:"sig"`
}

type UserWallet struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
	Sig     string `json:"sig"`
}

type UserTasks struct {
	TaskID      string  `json:"taskId"`
	UserDid     string  `json:"userDID"`
	ApDid       string  `json:"apDID"`
	ApName      string  `json:"apName"`
	ApMethod    string  `json:"apMethod"`
	DpDid       string  `json:"dpDID"`
	DpName      string  `json:"dpName"`
	DpMethod    string  `json:"dpMethod"`
	CreateTime  string  `json:"createTime"`
	UpdateTime  string  `json:"updateTime"`
	TaskStatus  string  `json:"taskStatus"`
	TaskResult  *string `json:"taskResult"`
	ResultFile  *string `json:"resultFile"`
	IssueTxhash *string `json:"issueTxhash"`
}

type GetAlgorithmProviderMethodResp struct {
	GetAlgorithmProviderMethods []*ProviderMethod `json:"getAlgorithmMethods"`
}

type GetDataProviderMethodResp struct {
	GetDataProviderMethods []*ProviderMethod `json:"getDataMethods"`
}

type GetAlgorithmProvidersResp struct {
	GetAllAlgorithmProviders []*AlgorithmProvider `json:"getAllAlgorithmProviders"`
}

type GetDataProvidersResp struct {
	GetAllDataProviders []*DataProvider `json:"getAllDataProviders"`
}

type GetUserTaskResp struct {
	GetUserTask *UserTasks `json:"getUserTask"`
}

type GetOscoreResp struct {
	Oscore int64 `json:"requestOscore"`
}
