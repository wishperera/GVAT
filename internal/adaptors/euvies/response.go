package euvies

import "encoding/xml"

type Response struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPBodyResponse
}

type SOAPBodyResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBody
	FaultDetails *Fault
}

type Fault struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBody struct {
	XMLName     xml.Name `xml:"checkVatResponse"`
	CountryCode string   `xml:"countryCode"`
	VATNumber   string   `xml:"vatNumber"`
	RequestDate string   `xml:"requestDate"`
	Valid       bool     `xml:"valid"`
	Name        string   `xml:"name"`
	Address     string   `xml:"address"`
}
