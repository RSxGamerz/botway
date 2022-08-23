package deno

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/deno.dockerfile", "botway", botName)
}

func Resources() string {
	return templates.Content("telegram/deno.md", "resources", "")
}

func MainTsContent() string {
	return templates.Content("main.ts", "telegram-deno", "")
}
