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
		query{
			getAlgorithmMethods(did:$did){
				name,
				paramSchema,
				resultSchema
			}
		}
	`

	GetDataProviderMethodsReq = `
		query{
			getDataMethods(did:$did){
				name,
				paramSchema,
				resultSchema
			}
		}
	`

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
					wallets:[%walletsinfo%]
				})
		}
	`

	GetUserTask = `
		query{
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