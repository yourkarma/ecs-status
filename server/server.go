package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/ecs"
	"github.com/yosssi/ace"
	"github.com/yourkarma/ecs-status/config"
	"github.com/yourkarma/ecs-status/templates"
	"github.com/yourkarma/mqttparty/log"
)

var client *ecs.ECS

type Data struct {
	ContainerInstances []ecs.ContainerInstance
	TaskDefinitions    []ecs.TaskDefinition
}

func Initialize(c *ecs.ECS) {
	client = c
}

func Start() {
	log.Infof("http/Start: server listening on port %s", config.Get().HTTP.Port)

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", config.Get().HTTP.Port), nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define helper functions
	functionMap := template.FuncMap{
		"join":     Join,
		"downcase": Downcase,
		"boolean":  Boolean,
	}

	tpl, err := ace.Load("views/layout", "views/main", &ace.Options{
		FuncMap: functionMap,
		Asset:   templates.Asset,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getData() (*Data, error) {
	return &Data{
		ContainerInstances: getContainerInstances(),
		TaskDefinitions:    getTaskDefinitions(),
	}, nil
}

func getContainerInstances() []ecs.ContainerInstance {
	response, err := client.DescribeContainerInstances(&ecs.DescribeContainerInstancesRequest{
		Cluster:            aws.String(config.Get().ECS.ClusterName),
		ContainerInstances: getContainerInstanceARNs(),
	})
	if err != nil {
		log.Errorf("main/getContainerInstances: %s", err)
		return nil
	}

	return response.ContainerInstances
}

func getContainerInstanceARNs() []string {
	response, err := client.ListContainerInstances(&ecs.ListContainerInstancesRequest{
		Cluster: aws.String(config.Get().ECS.ClusterName),
	})
	if err != nil {
		log.Errorf("main/getContainerInstanceARNs: %s", err)
		return nil
	}

	return response.ContainerInstanceARNs
}

func getTaskDefinitions() []ecs.TaskDefinition {
	var taskDefinitions []ecs.TaskDefinition

	for _, arn := range getTaskDefinitionARNs() {
		response, err := client.DescribeTaskDefinition(&ecs.DescribeTaskDefinitionRequest{
			TaskDefinition: aws.String(arn),
		})
		if err != nil {
			log.Errorf("main/getTaskDefinitions: %s", err)
			return nil
		}

		taskDefinitions = append(taskDefinitions, *response.TaskDefinition)
	}

	return taskDefinitions
}

func getTaskDefinitionARNs() []string {
	response, err := client.ListTaskDefinitions(&ecs.ListTaskDefinitionsRequest{})
	if err != nil {
		log.Errorf("main/getTaskDefinitionARNs: %s", err)
		return nil
	}

	return response.TaskDefinitionARNs
}

func Join(parts []string) string {
	return strings.Join(parts, " ")
}

func Downcase(text string) string {
	return strings.ToLower(text)
}

func Boolean(value bool) string {
	if value == true {
		return "Yes"
	}

	return "No"
}
