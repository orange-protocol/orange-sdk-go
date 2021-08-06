package sdk

var (
	GetAllAlgorithmProvidersReq = `
		query{
			getAllAlgorithmProviders(){
				name,
				type,
				introduction,
				did,
				createTime,
				title,
				provider,
				invokeFrequency,
				apiState,
				author,
				popularity,
				delay,icon
			}
		}
	`

	GetAllDataProvidersReq = `
		query{
			getAllDataProviders(){
				name,
				type,
				introduction,
				did,
				createTime,
				title,
				provider,
				invokeFrequency,
				apiState,
				author,
				popularity,
				delay,icon
			}
		}
	`

	GetAlgorithmProviderMethodsReq = `
		query getAlgorithmMethods($apdid:String!){
			getAlgorithmMethods(did:$apdid){
				name,
				paramSchema,
				resultSchema
			}
		}
	`

	GetDataProviderMethodsReq = `
		query getDataMethods($dpdid:String!){
			getDataMethods(did:$dpdid){
				name,
				paramSchema,
				resultSchema
			}
		}
	`

	//GetOscoreReq = `
	//	mutation requestOscore($data:RequestOscoreReq!){
	//			requestOscore(input:$data)
	//	}
	//`
	GetOscoreReq = `
		mutation{
				requestOscore(input:{
					key:$key,
					did:$did,
					apdid:$apdid,
					apmethod:$apmethod,
					dpdid:$dpdid,
					dpmethod:$dpmethod,
					overwriteOld:$overwriteOld,
					wallets:[$walletsinfo$]
				})
		}
	`

	GetUserTask = `
		query getUserTask($key:String!,$taskId:Int!){
			getUserTask(key:$key,taskId:$taskId){
				taskId,
				userDID,
				apDID,
				apName,
				apMethod,
				dpDID,
				dpName,
				dpMethod,
				createTime,
				updateTime,
				taskStatus,
				taskResult,
				resultFile,
				issueTxhash
			}
		}
		`
)
