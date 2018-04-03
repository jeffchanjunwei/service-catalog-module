package main

import (
	"fmt"
	pkg "github.com/jeffchanjunwei/service-catalog-api/pkg"
	"log"
	"net/http"
)

func main() {
	var apiServer = pkg.Prepare()
	apiServer.Get("clusterservicebroker", "ClusterServiceBrokerController@Get")
	apiServer.Post("clusterservicebroker", "ClusterServiceBrokerController@Post")
	apiServer.Put("clusterservicebroker", "ClusterServiceBrokerController@Put")
	apiServer.Patch("clusterservicebroker", "ClusterServiceBrokerController@Patch")
	apiServer.Delete("clusterservicebroker", "ClusterServiceBrokerController@Delete")

}

// func Init() {
//
// }
