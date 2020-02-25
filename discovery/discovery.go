package discovery

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

func Discovery() {
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("D_API_KEY"),
	}

	options := &discoveryv1.DiscoveryV1Options{
		Version:       "2019-07-12",
		Authenticator: authenticator,
	}

	discovery, discoveryErr := discoveryv1.NewDiscoveryV1(options)

	if discoveryErr != nil {
		panic(discoveryErr)
	}

	discovery.SetServiceURL(os.Getenv("D_URL"))

	result, _, err := discovery.Query(discovery.NewQueryOptions("system", "news-es"))

	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
