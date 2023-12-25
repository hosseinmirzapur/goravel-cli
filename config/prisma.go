package config

type prismaConfig struct {
	GithubRepo string
}

func GetPrismaConfig() *prismaConfig {
	return &prismaConfig{
		GithubRepo: "github.com/steebchen/prisma-client-go",
	}
}
