package render

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botwaygo"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func AddTokens(serviceId, apiToken string) {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s/env-vars", serviceId)

	botType := botwaygo.GetBotInfo("bot.type")

	bot_token := ""
	app_token := ""
	payload_content := ""

	if botType == "discord" {
		bot_token = "DISCORD_TOKEN"
		app_token = "DISCORD_CLIENT_ID"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken(), app_token, botwaygo.GetAppId())
	} else if botType == "slack" {
		bot_token = "SLACK_TOKEN"
		app_token = "SLACK_APP_TOKEN"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"%s\",\"value\":\"%s\"},{\"key\":\"SIGNING_SECRET\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken(), app_token, botwaygo.GetAppId(), botwaygo.GetSigningSecret())
	} else if botType == "telegram" {
		bot_token = "TELEGRAM_TOKEN"
		payload_content = fmt.Sprintf("[{\"key\":\"%s\",\"value\":\"%s\"}]", bot_token, botwaygo.GetToken())
	}

	payload := strings.NewReader(payload_content)
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode == 200 {
		fmt.Println(constants.HEADING + constants.BOLD.Render("Tokens added successfuly"))
	} else {
		fmt.Println(string(body))
	}
}

func ConnectService() {
	id := gjson.Get(string(constants.BotwayConfig), "render.user.id").String()
	apiToken := gjson.Get(string(constants.BotwayConfig), "render.user.api_token").String()

	serviceName := strings.ReplaceAll(botwaygo.GetBotInfo("bot.name"), " ", "%20")

	url := fmt.Sprintf("https://api.render.com/v1/services?name=%s&type=web_service&env=docker&ownerId=%s&limit=20", serviceName, id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, serr := http.DefaultClient.Do(req)

	if serr != nil {
		panic(serr)
	}

	body, _ := ioutil.ReadAll(res.Body)

	serviceId := gjson.Get(string(body), "0.service.id").String()
	serviceSlug := gjson.Get(string(body), "0.service.slug").String()
	serviceRepo := gjson.Get(string(body), "0.service.repo").String()

	renderPath := "render.projects." + serviceSlug

	service, _ := sjson.Set(string(constants.BotwayConfig), renderPath+".id", serviceId)
	addRepoToservice, _ := sjson.Set(service, renderPath+".repo", serviceRepo)

	remove := os.Remove(constants.BotwayConfigFile)

	if remove != nil {
		log.Fatal(remove)
	}

	newBotConfig := os.WriteFile(constants.BotwayConfigFile, []byte(addRepoToservice), 0644)

	if newBotConfig != nil {
		panic(newBotConfig)
	}

	AddTokens(serviceId, apiToken)
}
