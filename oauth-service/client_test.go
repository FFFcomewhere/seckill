package main

import (
	"context"
	"fmt"
	"github.com/FFFcomewhere/sk_object/pb"
	"github.com/FFFcomewhere/sk_object/pkg/client"
	"testing"
)

func TestOauthClientImpl_Check(t *testing.T) {
	client, _ := client.NewOAuthClient("user", nil, nil)

	if response, err := client.CheckToken(context.Background(), nil, &pb.CheckTokenRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyRGV0YWlscyI6eyJVc2VySWQiOjEsIlVzZXJuYW1lIjoidXNlcm5hbWUiLCJQYXNzd29yZCI6IiIsIkF1dGhvcml0aWVzIjpbIkFkbWluIiwiU3VwZXIiXX0sIkNsaWVudERldGFpbHMiOnsiQ2xpZW50SWQiOiJjbGllbnRJZCIsIkNsaWVudFNlY3JldCI6IiIsIkFjY2Vzc1Rva2VuVmFsaWRpdHlTZWNvbmRzIjoxODAwLCJSZWZyZXNoVG9rZW5WYWxpZGl0eVNlY29uZHMiOjE4MDAwLCJSZWdpc3RlcmVkUmVkaXJlY3RVcmkiOiJodHRwOi8vMTI3LjAuMC4xIiwiQXV0aG9yaXplZEdyYW50VHlwZXMiOlsicGFzc3dvcmQiLCJyZWZyZXNoX3Rva2VuIl19LCJSZWZyZXNoVG9rZW4iOnsiUmVmcmVzaFRva2VuIjpudWxsLCJUb2tlblR5cGUiOiJqd3QiLCJUb2tlblZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SlZjMlZ5UkdWMFlXbHNjeUk2ZXlKVmMyVnlTV1FpT2pFc0lsVnpaWEp1WVcxbElqb2lkWE5sY201aGJXVWlMQ0pRWVhOemQyOXlaQ0k2SWlJc0lrRjFkR2h2Y21sMGFXVnpJanBiSWtGa2JXbHVJaXdpVTNWd1pYSWlYWDBzSWtOc2FXVnVkRVJsZEdGcGJITWlPbnNpUTJ4cFpXNTBTV1FpT2lKamJHbGxiblJKWkNJc0lrTnNhV1Z1ZEZObFkzSmxkQ0k2SWlJc0lrRmpZMlZ6YzFSdmEyVnVWbUZzYVdScGRIbFRaV052Ym1Seklqb3hPREF3TENKU1pXWnlaWE5vVkc5clpXNVdZV3hwWkdsMGVWTmxZMjl1WkhNaU9qRTRNREF3TENKU1pXZHBjM1JsY21Wa1VtVmthWEpsWTNSVmNta2lPaUpvZEhSd09pOHZNVEkzTGpBdU1DNHhJaXdpUVhWMGFHOXlhWHBsWkVkeVlXNTBWSGx3WlhNaU9sc2ljR0Z6YzNkdmNtUWlMQ0p5WldaeVpYTm9YM1J2YTJWdUlsMTlMQ0pTWldaeVpYTm9WRzlyWlc0aU9uc2lVbVZtY21WemFGUnZhMlZ1SWpwdWRXeHNMQ0pVYjJ0bGJsUjVjR1VpT2lJaUxDSlViMnRsYmxaaGJIVmxJam9pSWl3aVJYaHdhWEpsYzFScGJXVWlPbTUxYkd4OUxDSmxlSEFpT2pFMU56TTFOemN6TURRc0ltbHpjeUk2SWxONWMzUmxiU0o5LmJneEdtaGJyVmVVQWljMk1ZRWtmeFJFVFVvVHhPYUZmQXFTa2xoQk50N2ciLCJFeHBpcmVzVGltZSI6IjIwMTktMTEtMTNUMDA6NDg6MjQuMzU2OTY4KzA4OjAwIn0sImV4cCI6MTU3MzU2MTEwNCwiaXNzIjoiU3lzdGVtIn0.CRXA_vadztUpOuKPth7HSb-E4l0otZEr6YNU5ZQcEcc",
	}); err == nil {
		fmt.Println("Check error", err.Error())
	} else {
		fmt.Println("result=", response.ClientDetails.ClientId)
	}
}
