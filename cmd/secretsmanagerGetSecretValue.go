package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/davecgh/go-spew/spew"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type cmdSecretManagerGetSecretValue struct{}

var secretsmanagerGetSecretValueCmd *cobra.Command

func init() {
	secretsManagerCmd.AddCommand(newCmdSecretsmanagerGetSecretValue())

}

func newCmdSecretsmanagerGetSecretValue() *cobra.Command {

	c := &cmdSecretManagerGetSecretValue{}

	secretsmanagerGetSecretValueCmd = &cobra.Command{
		Use: "get-secret-value",
		RunE: func(cmd *cobra.Command, args []string) error {

			return c.Run()

		},
	}

	return secretsmanagerGetSecretValueCmd
}

func (c *cmdSecretManagerGetSecretValue) Run() error {

	log.Info().Msg("hola")

	secretName := "arn:aws:secretsmanager:ap-southeast-2:177318690904:secret:DdApiKeySecret-BkF48vRfAZL4-JLeyol"

	ctx := context.Background()

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("error loading configuration")
	}
	client := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	o, err := client.GetSecretValue(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("error getting secret value")
		return nil
	}

	spew.Dump(o)

	return nil

}
