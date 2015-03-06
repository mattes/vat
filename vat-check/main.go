package main

import (
	"fmt"
	"github.com/mattes/vat"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %v <vat-number>\n", os.Args[0])
		os.Exit(2)
	}
	r, err := vat.CheckVAT(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	fmt.Println("Request date:", r.RequestDate)
	fmt.Println("  VAT number:", r.VATnumber)
	fmt.Println("     Country:", r.CountryCode)
	fmt.Println("        Name:", r.Name)
	fmt.Println("     Address:", r.Address)

	if r.Valid {
		fmt.Println("\nSuccess: The VAT number is valid!")
		os.Exit(0)
	} else {
		fmt.Println("\nError: The VAT number is not valid!")
		os.Exit(1)
	}
}
