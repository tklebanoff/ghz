/*
Package ghz can be used to perform benchmarking and load testing against gRPC services.

	report, err := ghz.Run(
		"helloworld.Greeter.SayHello",
		"localhost:50051",
		WithProtoFile("greeter.proto", []string{}),
		WithDataFromFile("data.json"),
		WithInsecure(true),
	)
*/
package ghz