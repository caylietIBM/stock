package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	person "github.ibm.com/Caylie-Taylor/geaux-go/http_api_gateway/var_global"
	pb "github.ibm.com/Caylie-Taylor/geaux-go/stock/proto"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello Caylie")
}

func CompanyInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	symbol := params["symbol"]
	companyInfoRequest := pb.CompanyInfoRequest{Symbol: symbol}
	fmt.Fprintln(w, symbol)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stockCompanyResponse, err := person.New_client.GetCompanyInfo(ctx, &companyInfoRequest)
	fmt.Println("GRPC response: ", stockCompanyResponse)
	//fmt.Fprint(w, stockCompanyResponse.CompanyName)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if ctx.Err() != nil {
		fmt.Fprint(w, "request cancelled, likely due to time out")
		return
	}
	fmt.Println("GRPC response: ", stockCompanyResponse)
	fmt.Fprint(w, stockCompanyResponse.CompanyName)
	fmt.Fprint(w, stockCompanyResponse.Ceo)
}
func StockPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticker := params["ticker"]
	stockPriceRequest := pb.StockPriceRequest{Ticker: ticker}
	fmt.Fprintln(w, ticker)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stockPriceResponse, err := person.New_client.GetStockPrice(ctx, &stockPriceRequest)
	fmt.Println("GRPC response: ", stockPriceResponse)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if ctx.Err() != nil {
		fmt.Fprint(w, "request cancelled, likely due to time out")
		return
	}
	fmt.Println("GRPC response: ", stockPriceResponse)
	fmt.Fprint(w, stockPriceResponse.Price)
}
