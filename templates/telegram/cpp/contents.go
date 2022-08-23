package cpp

import "github.com/abdfnx/botway/templates"

func DockerfileContent(botName string) string {
	return templates.Content("dockerfiles/cmake-telegram.dockerfile", "botway", botName)
}

func Resources() string {
	return templates.Content("telegram/cpp.md", "resources", "")
}

func BWCPPFileContent(botName string) string {
	return templates.Content("packages/bwpp/main.hpp", "botway", botName)
}

func MainCppContent(botName string) string {
	return templates.Content("src/main.cpp", "telegram-cpp", botName)
}

func DotDockerIgnoreContent() string {
	return templates.Content(".dockerignore", "telegram-cpp", "")
}

func CmakeListsContent(botName string) string {
	return templates.Content("CMakeLists.txt", "telegram-cpp", botName)
}

func RunPsFileContent() string {
	return templates.Content("run.ps1", "telegram-cpp", "")
}
