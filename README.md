# hello-app
Simple hello-app API to be used with the hello-app-operator
- At “/hello“ returns  “Hello <value of query parameter name>“ - e.g. “/hello?name=World!” returns “Hello World!“
- At “/healthz” returns HTTP Status 200
- At “/break”. after a HTTP Post request, requests to (“/healthz”) endpoint starts to return HTTP Status 500
