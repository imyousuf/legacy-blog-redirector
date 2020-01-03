package main

import (
	"context"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// http://imyousuf-tech.blogs.smartitengineering.com/2010/06/using-gnome-blog-i-was-looking-for.html
	blogURLPattern = regexp.MustCompile("[/]?[0-9]{4}/[0-9]{2}/(.*)\\.html")
	// http://imyousuf-tech.blogs.smartitengineering.com/search/label/java
	tagURLPattern = regexp.MustCompile("[/]?search/label/(.+)")
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defaultLocation := "https://imytech.net/"
	location := defaultLocation
	if blogURLPattern.MatchString(request.Path) {
		location += "legacy-posts/" + blogURLPattern.FindStringSubmatch(request.Path)[1] + "/"
	} else if tagURLPattern.MatchString(request.Path) {
		location += "tags/" + tagURLPattern.FindStringSubmatch(request.Path)[1]
	}
	responseHeaders := make(map[string]string)
	responseHeaders["Location"] = location
	return events.APIGatewayProxyResponse{Headers: responseHeaders, StatusCode: 301}, nil
}

func main() {
	lambda.Start(handleRequest)
}
