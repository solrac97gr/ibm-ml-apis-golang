package natural

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageclassifierv1"
)

func Natural() {

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("NLC_API_KEY"),
	}

	options := &naturallanguageclassifierv1.NaturalLanguageClassifierV1Options{
		Authenticator: authenticator,
	}

	naturalLanguageClassifier, naturalLanguageClassifierErr := naturallanguageclassifierv1.NewNaturalLanguageClassifierV1(options)

	if naturalLanguageClassifierErr != nil {
		panic(naturalLanguageClassifierErr)
	}

	naturalLanguageClassifier.SetServiceURL(os.Getenv("NLC_URL"))

	result, _, responseErr := naturalLanguageClassifier.Classify(
		&naturallanguageclassifierv1.ClassifyOptions{
			ClassifierID: core.StringPtr(os.Getenv("NLC_MODEL_ID")),
			Text:         core.StringPtr("Como estas?"),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
