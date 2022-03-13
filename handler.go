package main

import (
	"encoding/json"
	"net/http"
	pb "onlineboutiqueapi/genproto"
)

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {

	products, err := fe.getProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type productView struct {
		Item  *pb.Product
		Price *pb.Money
	}

	ps := make([]productView, len(products))
	for i, p := range products {
		price, err := fe.convertCurrency(r.Context(), p.GetPriceUsd(), "USD")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ps[i] = productView{p, price}
	}

	responseJSON, _ := json.Marshal(ps)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
