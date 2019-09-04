package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

//LcDataEvent type
//input event: basic data Lc
type LcDataEvent struct {
	Name             string `json:"Nombre"`
	Description      string `json:"Descripcion"`
	EnterpriseCode   int    `json:"CodigoOrganismo"`
	EnterpriseName   string `json:"NombreOrganismo"`
	OrganizationRut  string `json:"RutUnidad"`
	OrganizationCode int    `json:"CodigoUnidad"`
	OrganizationName string `json:"NombreUnidad"`
	TypeLCCode       int    `json:"CodigoTipo"`
	SubTypeLC        string `json:"Tipo"`
	GoodServiceCode  int    `json:"CodigoProducto"`
	GoodServiceName  string `json:"NombreProducto"`
	CategoryCode     int    `json:"CodigoCategoria"`
	CategoryName     string `json:"NombreCategoria"`
	ItemDescription  string `json:"DescripcionItem"`
}

//OrganizationBasicInformation type
//output event: basic data of organization
type OrganizationBasicInformation struct {
	Rut       string  `json:"Rutproveedor"`
	Relevance float32 `json:"Relevancia"`
}

//HandleEvent func
func HandleEvent(ctx context.Context, event LcDataEvent) ([]OrganizationBasicInformation, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	return []OrganizationBasicInformation{
		{Rut: "76.753.070-6", Relevance: 90.0},
		{Rut: "13.887.231-9", Relevance: 50.0},
		{Rut: "7.299.547-3", Relevance: 50.0},
	}, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
