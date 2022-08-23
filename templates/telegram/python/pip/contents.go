package pip

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/pip.dockerfile", "botway", botName)
}

func RequirementsContent() string {
	return templates.Content("requirements.txt", "telegram-python", "")
}
