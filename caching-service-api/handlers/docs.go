// Package classification Emplooyee API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta

package handlers

import "caching-service/data"

// A list of employee
// swagger:response employeesResponse
type employeesResponseWrapper struct {
	// All current products
	// in:body
	Body []data.Employee
}

// Employee info
// swagger:response employeeResponse
type employeeResponseWrapper struct {
	// Single employee obj
	// in:body
	Body data.Employee
}

//BadReqeuest something is wrong with request
// swagger:response badRequestResponse
type badRequestResponse struct {
	// in:body
	Body GenericError
}

//InternalServerError something bad happended in server
// swagger:response internalServerErrorResponse
type internalServerErrorResponse struct {
	// in:body
	Body GenericError
}
