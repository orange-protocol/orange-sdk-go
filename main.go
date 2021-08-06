package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oscore/oscore-sdk-go/graphql"
	"github.com/oscore/oscore-sdk-go/sdk"
)

func main() {
	fmt.Println("sdk start!")
	client := graphql.NewClient("http://localhost:8080/query")

	// make a request
	req := graphql.NewRequest(`
		query getAlgorithmMethods($apdid:String!){
			getAlgorithmMethods(did:$apdid){
				name,
				paramSchema,
				resultSchema
			}
		}
	`)
	//req := graphql.NewRequest(sdk.GetAllAlgorithmProvidersReq)

	// set any variables
	req.Var("$apdid", "did:ont:testap")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData sdk.GetAlgorithmProviderMethodResp
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", respData)
	for _, al := range respData.GetAlgorithmProviderMethods {
		fmt.Printf("name:%s\n", al.Name)
		fmt.Printf("did:%s\n", al.ResultSchema)
	}
}
