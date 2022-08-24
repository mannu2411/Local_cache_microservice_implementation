package main

import (
	"github.com/razorpay/razorpay-go"
	"html/template"
	"net/http"
)

type PageVars struct {
	OrderId string
}

func Payment(w http.ResponseWriter, r *http.Request) {
	client := razorpay.NewClient("rzp_test_XkDWTHT1Xa5vGz", "BR7TjcTcMH0rxdPNK19UCzAV")

	data := map[string]interface{}{
		"amount":   50000,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := body["id"]
	str := value.(string)
	HomePageVars := PageVars{OrderId: str}
	t, err := template.ParseFiles("app.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
