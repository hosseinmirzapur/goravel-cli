package artisan

import (
	"os"

	"github.com/hosseinmirzapur/goravel-cli/utils"
	"github.com/spf13/cobra"
)

var decryptKey string = ""

// env group
var envGroup = &cobra.Group{
	ID:    "env",
	Title: "env",
}

// env encrypt command
var encryptEnvCmd = &cobra.Command{
	Use:     "env:encrypt",
	Short:   "encrypt .env file content with AES algorithm",
	Example: "goravel-cli artisan key:encrypt --key CUSTOM_KEY",
	GroupID: envGroup.ID,
	Run:     encryptEnv,
}

// env decrypt command
var decryptEnvCmd = &cobra.Command{
	Use:     "env:decrypt",
	Short:   "decrypt .env file content with AES algorithm",
	Example: "goravel-cli artisan key:decrypt --key CUSTOM_KEY",
	GroupID: envGroup.ID,
	Run:     decryptEnv,
}

func encryptEnv(cmd *cobra.Command, args []string) {
	// read .env file
	data, err := os.ReadFile(".env")
	if err != nil {
		utils.Error("env", "error encrypting .env file", err)
	}

	// encrypt .env content
	resData, err := utils.Encrypt(data)
	if err != nil {
		utils.Error("env", "error encrypting .env file", err)
	}

	// write encrypted data into .env.encrypted
	err = utils.WriteFile(".env.encrypted", resData.EncryptedContent)
	if err != nil {
		utils.Error("env", "error writing .env file", err)
	}

	utils.Success(".env file encrypted successfully", false)
	utils.Info(string(resData.Key), false)

}

func decryptEnv(cmd *cobra.Command, args []string) {
	// read content from .env.encrypted
	data, err := os.ReadFile(".env.encrypted")
	if err != nil {
		utils.Error("env", "error decrypting .env file", err)
	}
	// decrypt .env.encrypted content
	decrypted, err := utils.Decrypt(data, []byte(decryptKey))
	if err != nil {
		utils.Error("env", "error decrypting .env file", err)
	}

	// write decrypted data into .env
	err = utils.WriteFile(".env", decrypted)
	if err != nil {
		utils.Error("env", "error writing .env file", err)
	}
	utils.Success("decryption successful", true)
	utils.Info("please remove .env.encrypted file", false)
	utils.Alert("if .env file is still encrypted, you have provided the wrong key", false)
}

func init() {
	decryptEnvCmd.Flags().StringVarP(&decryptKey, "key", "k", "", "decryption key")
	err := decryptEnvCmd.MarkFlagRequired("key")
	utils.Error("env", "unable to recognize flag", err)

}
