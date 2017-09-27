package main

import (  
    "fmt"
)

type InvoiceCalculator interface {  
    CalculateInvoice() (int,string)
}

type Independent struct {  
    systemId     int
    subscription int
    license 	 int
    errMessage   string 
}

type Partner struct {  
    systemId  int
    license      int
    errMessage 	 string
}

//invoice for Independent is the cost of subscription plus license
func (s Independent) CalculateInvoice() (int, string) {  
    total := s.subscription + s.license
    if total == 0 {
        return total, "Error: invalid cost"
    } 
    return total, "" 
}

//invoice for Partner is the cost of license alone
func (p Partner) CalculateInvoice() (int, string) {  
    if p.license == 0 {
        return p.license, "Error: invalid cost"
    } 
    return p.license, "" 
}


/*
total revenue is calculated by iterating InvoiceCalculator slice and adding up all costs for partners and independents 
*/
func totalRevenue(s []InvoiceCalculator) {  
    totalInvoice := 0
    for _, t := range s {
        result,_ := t.CalculateInvoice()
        if result > 0 {
         totalInvoice = totalInvoice + result
	}
    }
    fmt.Printf("Total Invoices Per Month $%d", totalInvoice)
}

func main() {  
    largeIndependent := Independent{1, 5000, 2000,""}
    smallIndependent := Independent{2, 300, 100,""}
    bigPartner := Partner{3, 30000,""}
    invoiceCollection := []InvoiceCalculator{largeIndependent, smallIndependent, bigPartner}
    totalRevenue(invoiceCollection)

}
