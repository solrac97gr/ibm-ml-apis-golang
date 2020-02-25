package personality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
)

func Personality() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	authenticator := &core.IamAuthenticator{

		ApiKey: os.Getenv("PI_API_KEY"),
	}

	options := &personalityinsightsv3.PersonalityInsightsV3Options{
		Version:       "2017-10-13",
		Authenticator: authenticator,
	}

	personalityInsights, personalityInsightsErr := personalityinsightsv3.NewPersonalityInsightsV3(options)

	if personalityInsightsErr != nil {
		panic(personalityInsightsErr)
	}

	personalityInsights.SetServiceURL(os.Getenv("PI_URL"))

	profile, profileErr := ioutil.ReadFile("./personalityInsights/profile.json")
	if profileErr != nil {
		panic(profileErr)
	}

	content := new(personalityinsightsv3.Content)
	json.Unmarshal(profile, content)

	result, _, responseErr := personalityInsights.Profile(
		&personalityinsightsv3.ProfileOptions{
			Content:                content,
			ContentType:            core.StringPtr("application/json"),
			RawScores:              core.BoolPtr(true),
			ConsumptionPreferences: core.BoolPtr(true),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
