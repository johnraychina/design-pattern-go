package main

import "fmt"

//Why would you want to control access to an object?
//Here is an example: you have a massive object that consumes a vast amount of system resources.
//You need it from time to time, but not always.

// You could implement lazy initialization: create this object only when it’s actually needed.
//All of the object’s clients would need to execute some deferred initialization code.
//Unfortunately, this would probably cause a lot of code duplication.
//
//In an ideal world, we’d want to put this code directly into our object’s class, but that isn’t always possible.
//For instance, the class may be part of a closed 3rd-party library

//The Proxy pattern suggests that you create a new proxy class with the same interface as an original service object.
//Then you update your app so that it passes the proxy object to all the original object’s clients.
//Upon receiving a request from a client,
//the proxy creates a real service object and delegates all the work to it.
func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
