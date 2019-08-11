package main

import (
	"context"
	"golang.org/x/oauth2"
	pb "google.golang.org/genproto/googleapis/assistant/embedded/v1alpha2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	OAuth struct {
		ClientID     string   `yaml:"client_id"`
		ClientSecret string   `yaml:"client_secret"`
		Scopes       []string `yaml:"scopes"`
		AuthURL      string   `yaml:"auth_url"`
		TokenURL     string   `yaml:"token_url"`
		RefreshToken string   `yaml:"refresh_token"`
	}

	Device struct {
		Endpoint      string `yaml:"endpoint"`
		DeviceId      string `yaml:"device_id"`
		DeviceModelId string `yaml:"device_model_id"`
		LanguageCode  string `yaml:"language_code"`
	}
}

func main() {
	// Read Configuration
	data, config_file_err := ioutil.ReadFile("./config.yaml")
	if config_file_err != nil {
		panic(config_file_err)
	}

	// Parse Configuration
	config := Config{}
	yaml_err := yaml.Unmarshal(data, &config)
	if yaml_err != nil {
		panic(yaml_err)
	}

	// Setup Oauth
	oauth_conf := &oauth2.Config{
		ClientID:     config.OAuth.ClientID,
		ClientSecret: config.OAuth.ClientSecret,
		Scopes:       config.OAuth.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.OAuth.AuthURL,
			TokenURL: config.OAuth.TokenURL,
		},
	}

	// Refresh Token
	token_source := oauth_conf.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: config.OAuth.RefreshToken,
	})

	token, token_err := token_source.Token()
	if token_err != nil {
		panic(token_err)
	}

	// Connect to gRPC
	conn, conn_err := grpc.Dial(
		config.Device.Endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(token)),
	)

	if conn_err != nil {
		panic(conn_err)
	}

	// Create new Google Assistant Client
	client := pb.NewEmbeddedAssistantClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start a Sessions
	g, client_err := client.Assist(ctx)
	if client_err != nil {
		panic(client_err)
	}

	// Get text query from console
	var text string
	for _, piece := range os.Args[1:] {
		text += " " + piece
	}

	// Build the request and send it to Google Assistant Service
	g.Send(&pb.AssistRequest{
		Type: &pb.AssistRequest_Config{
			Config: &pb.AssistConfig{
				Type: &pb.AssistConfig_TextQuery{
					TextQuery: text,
				},
				AudioOutConfig: &pb.AudioOutConfig{
					Encoding:         pb.AudioOutConfig_LINEAR16,
					SampleRateHertz:  16000,
					VolumePercentage: 0,
				},
				DeviceConfig: &pb.DeviceConfig{
					DeviceId:      config.Device.DeviceId,
					DeviceModelId: config.Device.DeviceModelId,
				},
				DialogStateIn: &pb.DialogStateIn{
					LanguageCode:      config.Device.LanguageCode,
					IsNewConversation: true,
				},
			},
		},
	})

	// Wait for response and print it to the console
	for {
		res, res_err := g.Recv()
		if res_err == io.EOF {
			break
		}
		if res_err != nil {
			panic(res_err)
		}
		if res.DialogStateOut != nil {
			println(res.DialogStateOut.SupplementalDisplayText)
		}
	}
}
