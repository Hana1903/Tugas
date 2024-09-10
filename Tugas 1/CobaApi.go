package main

import (
	"encoding/json"  
	"fmt"           
	"net/http"    
)

type TaxPayment struct {
	PaymentID string 
	TaxPayerID string      
	Amount float64
	Method string
}

var data = []TaxPayment{
	TaxPayment{"A001","001A",70000,"Cash"},
	TaxPayment{"A003","004A",55000,"Cash"},
	TaxPayment{"A005","007A",70900,"Cash"},  
	TaxPayment{"A007","010A",89000,"Cash"},
	TaxPayment{"A009","013A",90900,"Cash"},
}

func pajaks(w http.ResponseWriter, r *http.Request)  {  
	w.Header().Set("Content-Type", "application/json")  

	if r.Method == "GET" {  
		var result, err = json.Marshal(data) 

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return     
		}

		w.Write(result)
		return        
	}
	http.Error(w, "", http.StatusBadRequest)  
}

func pajak(w http.ResponseWriter, r *http.Request)  {   
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("PaymentID") 
		var result []byte
		var err error

		for _, each := range data {   
			if each.PaymentID == id {   
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)  
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main()  {
	http.HandleFunc("/pajaks", pajaks)
	http.HandleFunc("/pajak", pajak)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}