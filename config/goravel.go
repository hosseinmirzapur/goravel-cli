package config

type goravelConfig struct {
	GithubRepo string
}

func GetGoravelConfig() *goravelConfig {
	return &goravelConfig{
		GithubRepo: "https://github.com/goravel/goravel.git",
	}
}
