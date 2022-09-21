package pip

import (
	"fmt"

	"github.com/abdfnx/botway/templates"
)

func DockerfileContent(botName, hostService string) string {
	return templates.Content(fmt.Sprintf("dockerfiles/%s/pip.dockerfile", hostService), "botway", botName, "telegram")
}

func RequirementsContent() string {
	return templates.Content("requirements.txt", "telegram-python", "", "")
}
