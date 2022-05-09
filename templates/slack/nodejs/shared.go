package nodejs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/botway/templates"
	"github.com/abdfnx/looker"
	"github.com/tidwall/sjson"
)

var Packages = "slackbots botway.js"

func IndexJSContent() string {
	return templates.Content("slack", "nodejs", "src/index.js", "")
}

func Resources() string {
	return `# Botway Slack (Node.js) Resources

> Here is some useful links and resources to help you to build your own bot

## Setup

- [Setup Slack bot](https://github.com/abdfnx/botway/discussions/6)

## API

- [Simple way to control your Slack Bot](https://github.com/mishk0/slack-bot-api)

## Examples

[Examples](https://github.com/mishk0/slack-bot-api/tree/master/test)

big thanks to [**@mishk0**](https://github.com/mishk0)`
}

func SlackNodejs(botName, pm string) {
	npmPath, nerr := looker.LookPath("npm")
	pmPath, err := looker.LookPath(pm)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" " + pm + " is not installed"))
	} else {
		if nerr != nil {
			fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
			fmt.Println(constants.FAIL_FOREGROUND.Render(" npm is not installed"))
		} else {
			npmInit := npmPath + " init -y"

			cmd := exec.Command("bash", "-c", npmInit)

			if runtime.GOOS == "windows" {
				cmd = exec.Command("powershell.exe", npmInit)
			}

			cmd.Dir = botName
			err = cmd.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}

			packageJson, err := ioutil.ReadFile(filepath.Join(botName, "package.json"))

			if err != nil {
				log.Printf("error: %v\n", err)
			}

			version, _ := sjson.Set(string(packageJson), "version", "0.1.0")
			description, _ := sjson.Delete(version, "description")
			keywords, _ := sjson.Delete(description, "keywords")
			license, _ := sjson.Delete(keywords, "license")
			main, _ := sjson.Set(string(license), "main", "src/index.js")
			author, _ := sjson.Delete(string(main), "author")
			final, _ := sjson.Delete(author, "scripts")

			newPackageJson := ioutil.WriteFile(filepath.Join(botName, "package.json"), []byte(final), 0644)

			if newPackageJson != nil {
				log.Printf("error: %v\n", newPackageJson)
			}

			DockerfileContent := templates.Content("slack", "nodejs", pm + "/Dockerfile", botName)

			indexFile := os.WriteFile(filepath.Join(botName, "src", "index.js"), []byte(IndexJSContent()), 0644)
			dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent), 0644)
			procFile := os.WriteFile(filepath.Join(botName, "Procfile"), []byte("process: node ./src/index.js"), 0644)
			resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources()), 0644)

			if resourcesFile != nil {
				log.Fatal(resourcesFile)
			}

			if indexFile != nil {
				log.Fatal(indexFile)
			}

			if dockerFile != nil {
				log.Fatal(dockerFile)
			}

			if procFile != nil {
				log.Fatal(procFile)
			}

			icmd := func () string {
				if pm == "npm" {
					return " i " + Packages
				} else {
					return " add " + Packages
				}
			}

			pmInstall := pmPath + icmd()
			installCmd := exec.Command("bash", "-c", pmInstall)

			if runtime.GOOS == "windows" {
				installCmd = exec.Command("powershell.exe", pmInstall)
			}

			installCmd.Dir = botName
			installCmd.Stdin = os.Stdin
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			err = installCmd.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}

			templates.CheckProject(botName, "slack")
		}
	}
}
