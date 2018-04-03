package pkg

import (
	"log"
	"net/http"
	"net/url"
)

type ClusterServiceBrokerController struct {
}

func (c *ClusterServiceBrokerController) GetCSB(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {

}

func (c *ClusterServiceBrokerController) PostCSB(request map[string]url.Values, headers http.Header) {

}

func (c *ClusterServiceBrokerController) PutCSB() {

}

func (c *ClusterServiceBrokerController) PatchCSB() {

}

func (c *ClusterServiceBrokerController) DeleteCSB() {

}

//type ClusterServiceBrokerListController struct {
//}
