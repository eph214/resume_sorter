package main

import "testing"


func TestCalculateInvoice(t *testing.T) {
        cases := []Independent {
                {
	             systemId:     1,
		     subscription: 5000,
		     license: 10,
		     errMessage: "",
                },
                {
	             systemId:     2,
		     subscription: 0,
		     license: 0,
		     errMessage: "Error: invalid cost",
                },
        }

        for _, c := range cases {
                _, errStr := c.CalculateInvoice()
                if errStr != c.errMessage {
                        t.Errorf("Expected err to be %s but it was %s", c.errMessage, errStr)
                }

        }
}


       
