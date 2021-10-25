package euvies

type Request struct {
	CountryCode string `xml:"countryCode"`
	VATNumber   string `xml:"vat_number"`
}

const requestTemplate = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
   <soapenv:Header/>
   <soapenv:Body>
      <urn:checkVat>
         <urn:countryCode>{{.CountryCode}}</urn:countryCode>
         <urn:vatNumber>{{.VATNumber}}</urn:vatNumber>
      </urn:checkVat>
   </soapenv:Body>
</soapenv:Envelope>
`
