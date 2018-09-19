// Copyright © 2018 Paul Carleton
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/nmrshll/oauth2-noserver"
	pb "github.com/pcarleton/cc-grpc/proto/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

type localClient struct {
	token     string
	apiClient pb.ApiClient
}

type googleConfig struct {
	Web struct {
		ClientID                string   `json:"client_id"`
		ProjectID               string   `json:"project_id"`
		AuthUri                 string   `json:"auth_uri"`
		TokenUri                string   `json:"token_uri"`
		AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
		JavascriptOrigins       []string `json:"javascript_origins"`
	} `json:"web"`
}

func getNewToken() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tokenPath := path.Join(home, ".cc-token")
	info, err := os.Stat(tokenPath)

	if err == nil {
		log.Printf("Found file")

		duration := time.Now().Sub(info.ModTime()).Minutes()
		log.Printf("Duration: %s", duration)
		if duration < 30 {
			contents, err := ioutil.ReadFile(tokenPath)
			if err == nil {
				return string(contents)
			} else {
				log.Printf("Failed to read: %s", err)
			}
		}
	}

	secretsPath := viper.GetString("CLIENT_SECRETS_PATH")
	if secretsPath == "" {
		log.Fatalf("No client secrets specified under CLIENT_SECRETS_PATH")
	}

	b, err := ioutil.ReadFile(secretsPath)

	if err != nil {
		panic(err)
	}

	var configData googleConfig

	json.Unmarshal(b, &configData)

	conf := &oauth2.Config{
		ClientID:     configData.Web.ClientID,     // also known as slient key sometimes
		ClientSecret: configData.Web.ClientSecret, // also known as secret key
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  configData.Web.AuthUri,
			TokenURL: configData.Web.TokenUri,
		},
	}

	client, err := oauth2ns.AuthenticateUser(conf)
	if err != nil {
		panic(err)
	}
	token := client.Token.Extra("id_token").(string)

	ioutil.WriteFile(tokenPath, []byte(token), 0644)

	return token
}

func getClient() *localClient {
	var conn *grpc.ClientConn
	var err error

	if insecure {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			panic(err)
		}

		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
	}

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := pb.NewApiClient(conn)

	return &localClient{
		token:     getNewToken(),
		apiClient: c,
	}
}

func (l *localClient) GetContext() context.Context {
	md := metadata.Pairs("token", l.token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx
}

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report [month_num] [bank_label] [account] [spreadsheet_id]",
	Short: "Create and share a report based on the specified config",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		spreadsheetId := args[3]
		plaidTokenLabel := args[1]
		accountNick := args[2]
		month, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		client := getClient()
		ctx := client.GetContext()

		r2, err := client.apiClient.CreateReport(ctx, &pb.CreateReportRequest{
			Month:         int32(month),
			Namespace:     plaidTokenLabel,
			AccountId:     accountNick,
			SpreadsheetId: spreadsheetId,
		})
		if err != nil {
			log.Fatalf("could not create report: %s", err)
		}
		log.Printf("Create Report Response: %s", r2.Result)

	},
}

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Print the health of the server",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		ctx := client.GetContext()

		resp, err := client.apiClient.GetHealth(ctx, &pb.GetHealthRequest{})
		if err != nil {
			log.Fatalf("could not fetch health: %s", err)
		}
		log.Printf("Version: %s", resp.Version)
		for _, check := range resp.Statuses {
			log.Printf("Health Response: %s [%s] %s", check.Label, check.Status, check.Result)
		}
	},
}

func init() {
	RootCmd.AddCommand(reportCmd)
	RootCmd.AddCommand(healthCmd)
}
