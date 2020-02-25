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
	limit := int64(3)
	targets := []string{"growht", "hacking"}
	targetss := []string{"stocks"}
	sentiment := true
	emotion := true
	sentences := true
	lemma := true
	partOfSpeech := true

	result, _, responseErr := naturalLanguageUnderstanding.Analyze(
		&naturallanguageunderstandingv1.AnalyzeOptions{
			URL: &url,
			Features: &naturallanguageunderstandingv1.Features{
				Categories: &naturallanguageunderstandingv1.CategoriesOptions{
					Limit: &limit,
				},
				Concepts: &naturallanguageunderstandingv1.ConceptsOptions{
					Limit: &limit,
				},
				Emotion: &naturallanguageunderstandingv1.EmotionOptions{
					Targets: targets,
				},
				Entities: &naturallanguageunderstandingv1.EntitiesOptions{
					Sentiment: &sentiment,
					Limit:     &limit,
				},
				Keywords: &naturallanguageunderstandingv1.KeywordsOptions{
					Sentiment: &sentiment,
					Emotion:   &emotion,
					Limit:     &limit,
				},
				Metadata:      &naturallanguageunderstandingv1.MetadataOptions{},
				Relations:     &naturallanguageunderstandingv1.RelationsOptions{},
				SemanticRoles: &naturallanguageunderstandingv1.SemanticRolesOptions{},
				Sentiment: &naturallanguageunderstandingv1.SentimentOptions{
					Targets: targetss,
				},
				Syntax: &naturallanguageunderstandingv1.SyntaxOptions{
					Sentences: &sentences,
					Tokens: &naturallanguageunderstandingv1.SyntaxOptionsTokens{
						Lemma:        &lemma,
						PartOfSpeech: &partOfSpeech,
					},
				},
			},
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "   ")
	fmt.Println(string(b))
}
