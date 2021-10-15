# Orange Server SDK 

## 1. 简介

orange server 使用 [graphql](https://graphql.org/) 作为对外的API 协议，本SDK 是对面向开发者和第三方的接口的封装



## 2. API KEY 的申请

orange sdk api需要注册DID

申请入口 ：[TBD](https://app.orangeprotocol.io/business)



## 3. 接口介绍

graphql schema

```

type AlgorithmProvider{
    name:String!
    type:String!
    introduction:String!
    did:String!
    createTime:Int!
    title:String!
    provider:String!
    invokeFrequency:Int!
    apiState:Int!
    author:String!
    popularity:Int!
    delay:Int!
    icon:String!
}

type DataProvider {
    name:String!
    type:String!
    introduction:String!
    did:String!
    createTime:Int!
    title:String!
    provider:String!
    invokeFrequency:Int!
    apiState:Int!
    author:String!
    popularity:Int!
    delay:Int!
    icon:String!
}

type ProviderMethod {
    name:String!
    paramSchema:String!
    resultSchema:String!
}

type UserTasks {
    taskId:String!
    userDID:String!
    apDID:String!
    apName:String!
    apMethod:String!
    dpDID:String!
    dpName:String!
    dpMethod:String!
    createTime:String!
    updateTime:String!
    taskStatus:String!
    taskResult:String
    resultFile:String
    issueTxhash:String
}

input RequestOrangeScoreReq{
    appdid:String!
    data:RequestOrangeScoreData!
    sig:String!
}
input RequestOrangeScoreData {
    userdid:String!
    apdid:String!
    apmethod:String!
    dpdid:String!
    dpmethod:String!
    overwriteOld:Boolean!
    wallets:[UserWallet!]!
}
input UserWallet{
    chain:String!
    address:String!
    pubkey:String!
    sig:String!
}

type Query {
  getAllAlgorithmProviders:[AlgorithmProvider!]!
  getAllDataProviders:[DataProvider!]!
  getAlgorithmMethods(did:String!):[ProviderMethod!]!
  getDataMethods(did:String!):[ProviderMethod!]!
  getUserTask(key:String!,taskId:Int!):UserTasks
}

type Mutation {
  requestOrangeScore(input:RequestOrangeScoreReq):Int!
}

```

### 3.1 取得所有支持的Algorithm provider

取得系统所有注册的算法提供方的信息

#### golang

```GetAlgorithmProviders```

parameters: 

returns: []*AlgorithmProvider

```
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
```



### 3.2 取得Algorithm methods

取得算法提供方支持的所有方法

#### golang

```GetAlgorithmMethods```

parameters: apdid string               did of algorithm provider

returns:         []*ProviderMethod

```
type ProviderMethod struct {
	Name         string `json:"name"`
	ParamSchema  string `json:"paramSchema"`
	ResultSchema string `json:"resultSchema"`
}
```



### 3.3 取得所有支持的Data provider

取得系统所有注册的数据提供方的信息

#### golang

```GetDataProviders```

parameters: 

returns: []*DataProvider

```
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
```

### 3.4 取得Data methods

取得数据提供方支持的所有方法

#### golang

```GetDataMethods```

parameters: dpdid string               did of data provider

returns:         []*ProviderMethod

```
type ProviderMethod struct {
	Name         string `json:"name"`
	ParamSchema  string `json:"paramSchema"`
	ResultSchema string `json:"resultSchema"`
}
```



### 3.5 申请计算Orange score

申请计算OrangeScore, 计算是一个异步过程，本次只返回申请的task id

### golang

```RequestOrangeScore```

parameters: *RequestOrangeScoreReq

return int64

```
type RequestOrangeScoreReq struct {
	AppDid string            `json:"appDid"`                    //app did
	Data   RequestOrangeScoreData `json:"data"`                      //request data
	Sig    string            `json:"sig"`                       //signature for data by did  
}

type RequestOrangeScoreData struct {
	Userdid      string        `json:"userDid"`                 //用户的DID
    Apdid        string        `json:"apdid"`                   //算法提供方的DID
    Apmethod     string        `json:"apmethod"`                //算法提供方的接口名称 
    Dpdid        string        `json:"dpdid"`                   //数据提供方的DID
    Dpmethod     string        `json:"dpmethod"`                //数据提供方的接口名称    
    overwriteOld bool          `json:"overwriteOld"`            //是否覆盖之前已存在的task
    Wallets      []*UserWallet `json:"wallets"`                 //用户钱包信息 

}
type UserWallet struct {
	Chain   string `json:"chain"`                               //链的名称 “eth" ,"bsc"等
	Address string `json:"address"`                             //绑定的钱包地址
	Pubkey  string `json:"pubkey"`                              //钱包地址对应的公钥
	Sig     string `json:"sig"`                                 //钱包私钥对用户DID的签名
}

```



### 3.6 查询用户Orange score task

根据task id 查询orange score task 信息

#### golang

```GetUserTask```

parameters: key string            //apikey

​	        taskId int64          //taskid

return :  *UserTasks

```
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
	TaskResult  *string `json:"taskResult"`                             // point
	ResultFile  *string `json:"resultFile"`                             // credential file       
	IssueTxhash *string `json:"issueTxhash"`                            //transaction hash for the credential on ontology
}
```



## 4. how to use

golang

add dependency in go.mod

```
github.com/orange-protocol/orange-sdk-go latest
```

```golang
func TestOrangeSDK_GetAlgorithmProviders(t *testing.T) {
	sdk, err := NewOrangeSDK("http://localhost:8080/query")
	assert.Nil(t, err)
	aps, err := sdk.GetAlgorithmProviders()
	assert.Nil(t, err)
	assert.NotNil(t, aps)
	assert.Greater(t, len(aps), 0)
}

```

