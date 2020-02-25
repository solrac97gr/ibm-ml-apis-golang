package naturalu

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

func Naturalu(urls string) {
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("NLU_API_KEY"),
	}

	options := &naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
		Version:       "2019-07-12",
		Authenticator: authenticator,
	}

	naturalLanguageUnderstanding, naturalLanguageUnderstandingErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(options)

	if naturalLanguageUnderstandingErr != nil {
		panic(naturalLanguageUnderstandingErr)
	}

	naturalLanguageUnderstanding.SetServiceURL(os.Getenv("NLU_URL"))

	url := urls

	result, _, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			// You can pass URL,TEXT,HTML
			URL: &url,
			Features: &naturallanguageunderstandingv1.Features{
				// You can edit all freatures for more customization or you can delete some Features if you dont need
				Categories:    &naturallanguageunderstandingv1.CategoriesOptions{},
				Concepts:      &naturallanguageunderstandingv1.ConceptsOptions{},
				Emotion:       &naturallanguageunderstandingv1.EmotionOptions{},
				Entities:      &naturallanguageunderstandingv1.EntitiesOptions{},
				Keywords:      &naturallanguageunderstandingv1.KeywordsOptions{},
				Metadata:      &naturallanguageunderstandingv1.MetadataOptions{},
				Relations:     &naturallanguageunderstandingv1.RelationsOptions{},
				SemanticRoles: &naturallanguageunderstandingv1.SemanticRolesOptions{},
				Sentiment:     &naturallanguageunderstandingv1.SentimentOptions{},
				Syntax:        &naturallanguageunderstandingv1.SyntaxOptions{},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "   ")
	fmt.Println(string(b))
}
