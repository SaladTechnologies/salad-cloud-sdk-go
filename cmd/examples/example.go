package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdk"
	"github.com/saladtechnologies/salad-cloud-sdk-go/pkg/saladcloudsdkconfig"
)

func main() {
	loadEnv()

	config := saladcloudsdkconfig.NewConfig()
	client := saladcloudsdk.NewSaladCloudSdk(config)

	response, err := client.ContainerGroups.ListContainerGroups(context.Background(), "organizationName", "projectName")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
