// MIT License
//
// Copyright (c) 2020 Alex W. Baulé
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package main

import (
	"log"

	"github.com/alexwbaule/give-help/v2/apihandler"
	"github.com/alexwbaule/give-help/v2/generated/models"
	"github.com/alexwbaule/give-help/v2/generated/restapi"
	"github.com/alexwbaule/give-help/v2/generated/restapi/operations"
	runtimeApp "github.com/alexwbaule/give-help/v2/runtime"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

var httpAdapter *httpadapter.HandlerAdapter

func init() {

	rt, err := runtimeApp.NewRuntime()
	if err != nil {
		log.Fatal(err.Error())
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	api := operations.NewGiveHelpAPI(swaggerSpec)
	server := restapi.NewServer(api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "give-help-service"
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	/*
	 * App Handlers
	 */

	// Applies when the "x-api-token" header is set
	api.APIKeyHeaderAuth = func(token string, roles []string) (*models.LoggedUser, error) {
		return apihandler.CheckAPIKeyAuth(rt, token, roles)
	}

	/* API Proposal */
	api.ProposalAddProposalHandler = apihandler.AddProposalHandler(rt)
	api.ProposalAddProposalImagesHandler = apihandler.AddProposalImagesHandler(rt)
	api.ProposalChangeProposalImagesHandler = apihandler.ChangeProposalImagesHandler(rt)
	api.ProposalAddProposalTagsHandler = apihandler.AddProposalTagsHandler(rt)
	api.ProposalChangeProposalStateHandler = apihandler.ChangeProposalStateHandler(rt)
	api.ProposalChangeProposalTextHandler = apihandler.ChangeProposalTextHandler(rt)
	api.ProposalChangeProposalValidateHandler = apihandler.ChangeProposalValidateHandler(rt)
	api.ProposalSearchProposalsHandler = apihandler.SearchProposalsHandler(rt)
	api.ProposalGetProposalByIDHandler = apihandler.GetProposalByIDHandler(rt)
	api.ProposalGetProposalByUserIDHandler = apihandler.GetProposalByUserIDHandler(rt)
	api.ProposalGetProposalShareDataIDHandler = apihandler.GetProposalShareDataIDHandler(rt)
	api.ProposalAddProposalComplaintHandler = apihandler.AddProposalComplaintHandler(rt)

	/* API User */
	api.UserAddUserHandler = apihandler.AddUserHandler(rt)
	api.UserUpdateUserByIDHandler = apihandler.UpdateUserByIDHandler(rt)
	api.UserGetUserByIDHandler = apihandler.GetUserByIDHandler(rt)

	/* API Transaction */
	api.TransactionAddTransactionHandler = apihandler.AddTransactionHandler(rt)
	api.TransactionChangeTransactionStatusHandler = apihandler.ChangeTransactionStatusHandler(rt)
	api.TransactionGetTransactionByIDHandler = apihandler.GetTransactionByIDHandler(rt)
	api.TransactionGetTransactionByProposalIDHandler = apihandler.GetTransactionByProposalIDHandler(rt)
	api.TransactionGetTransactionByUserIDHandler = apihandler.GetTransactionByUserIDHandler(rt)
	api.TransactionTransactionGiverReviewHandler = apihandler.TransactionGiverReviewHandler(rt)
	api.TransactionTransactionTakerReviewHandler = apihandler.TransactionTakerReviewHandler(rt)

	api.TransactionAcceptTransactionHandler = apihandler.TransactionAcceptTransactionHandler(rt)
	api.TransactionFinishTransactionHandler = apihandler.TransactionFinishTransactionHandler(rt)
	api.TransactionCancelTransactionHandler = apihandler.TransactionCancelTransactionHandler(rt)

	/* API Tags */
	api.TagsGetTagsHandler = apihandler.GetTagsHandler(rt)

	/* API Banks */
	api.BanksGetBankListHandler = apihandler.GetBankListHandler(rt)

	/* API Etc */
	api.EtcGetEtcListHandler = apihandler.GetEtcListHandler(rt)

	/* API Terms */
	api.TermsPutUserAcceptHandler = apihandler.TermsPutUserAcceptHandler(rt)
	api.TermsGetTermsHandler = apihandler.TermsGetTermsHandler(rt)
	api.TermsGetUserAcceptedHandler = apihandler.TermsGetUserAcceptedHandler(rt)

	server.ConfigureAPI()

	httpAdapter = httpadapter.New(server.GetHandler())

}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return httpAdapter.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
