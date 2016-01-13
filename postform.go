package main

import (
	"fmt"
	"github.com/eaciit/toolkit"
)

func main() {
	url := "http://www.dce.com.cn/PublicWeb/MainServlet"

	formvalues := toolkit.M{}.Set("Pu00231_Input.trade_date", "20151214").
		Set("Pu00231_Input.variety", "i").
		Set("Pu00231_Input.trade_type", "0").
		Set("Submit", "Go").
		Set("action", "Pu00231_result")

	config := toolkit.M{}.Set("formvalues", formvalues)

	r, e := toolkit.HttpCall(url, "POST", nil, config)
	if e != nil {
		fmt.Printf("ERROR:\n%s\n", e.Error())
		return
	}

	fmt.Printf("Result:\n%s\n", toolkit.HttpContentString(r))
}
