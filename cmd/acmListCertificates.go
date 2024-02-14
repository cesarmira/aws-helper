package cmd

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type cmdAcm struct{}

func init() {
	acmCmd.AddCommand(newCmdAcmListCertificates())
}

func newCmdAcmListCertificates() *cobra.Command {

	c := cmdAcm{}

	secretsManagerCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		RunE: func(cmd *cobra.Command, args []string) error {

			return c.Run()

		},
	}

	return secretsManagerCmd
}

func (c *cmdAcm) Run() error {

	//log.Info().Msg("hola")
	ctx := context.Background()

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("error loading configuration")
	}
	client := acm.NewFromConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating acm client")
	}

	secrets, err := client.ListCertificates(ctx, &acm.ListCertificatesInput{})
	if err != nil {
		log.Error().Err(err).Msg("error listing certificates")
	}

	var i int
	for _, v := range secrets.CertificateSummaryList {
		log.Info().Msgf("iterating %d", i)
		i++

		cert, err := client.GetCertificate(ctx, &acm.GetCertificateInput{
			CertificateArn: v.CertificateArn,
		})
		if err != nil {
			log.Error().Err(err).Msg("error getting certificate")
			continue
		}
		spew.Dump(cert.Certificate)
	}

	return nil

}
