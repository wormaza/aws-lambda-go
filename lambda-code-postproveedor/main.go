package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

//recomendaciones type
//input event: Proveedor a guardar
type recomendaciones struct {
	Code         string  `json:"Code"`
	ExternalCode string  `json:"ExternalCode"`
	Rut          string  `json:"Rut"`
	Rubros       string  `json:"Rubros"`
	Relevancia   float32 `json:"Relevancia"`
	Activo       bool    `json:"Activo"`
}

//PostRecomendaciones func
//Ingresa nuevas recomendaciones
func PostRecomendaciones(re recomendaciones) (bool, error) {
	return postItem(re)
}

//PostResponse type
//output event: Response de la operacion
type PostResponse struct {
	Message string `json:"Message"`
}

//HandleEvent func
func HandleEvent(ctx context.Context, event recomendaciones) (PostResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	resultado, err := PostRecomendaciones(event)

	if err != nil {
		fmt.Println("Dynamo call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	if resultado == true {
		return PostResponse{Message: "Ok"}, nil
	}

	return PostResponse{Message: "No es posible almacenar el registro"}, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
