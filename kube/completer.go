package kube

import (
	"strings"

	"github.com/c-bata/go-prompt-toolkit"
)

func Completer(s string) []prompt.Completion {
	if s == "" {
		return []prompt.Completion{}
	}
	args := strings.Split(s, " ")
	l := len(args)

	if strings.HasPrefix(args[l-1], "-") {
		return optionCompleter(args, strings.HasPrefix(args[l-1], "--"))
	}

	if len(args) == 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	if len(args) == 2 {
		return secondArgsCompleter(args[0], args[1])
	}

	if len(args) == 3 {
		return thirdArgsCompleter(args[0], args[1], args[2])
	}

	return []prompt.Completion{}
}

func strToCompletionList(x []string) []prompt.Completion {
	l := len(x)
	y := make([]prompt.Completion, l)
	for i := 0; i < l; i++ {
		y[i] = prompt.Completion{Text: x[i]}
	}
	return y
}

var commands = []prompt.Completion{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "replace", Description: "Replace a resource by filename or stdin."},
	{Text: "patch", Description: "Update field(s) of a resource using strategic merge patch."},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "namespace", Description: "SUPERSEDED: Set and view the current Kubernetes namespace"},
	{Text: "logs", Description: "Print the logs for a container in a pod."},
	{Text: "rolling-update", Description: "Perform a rolling update of the given ReplicationController."},
	{Text: "scale", Description: "Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job."},
	{Text: "cordon", Description: "Mark node as unschedulable"},
	{Text: "drain", Description: "Drain node in preparation for maintenance"},
	{Text: "uncordon", Description: "Mark node as schedulable"},
	// {Text: "attach", Description: "Attach to a running container."},  // still not supported
	// {Text: "exec", Description: "Execute a command in a container."}, // still not supported
	// {Text: "port-forward", Description: "Forward one or more local ports to a pod."}, // still not supported
	{Text: "proxy", Description: "Run a proxy to the Kubernetes API server"},
	{Text: "run", Description: "Run a particular image on the cluster."},
	{Text: "expose", Description: "Take a replication controller, service, or pod and expose it as a new Kubernetes Service"},
	{Text: "autoscale", Description: "Auto-scale a Deployment, ReplicaSet, or ReplicationController"},
	{Text: "rollout", Description: "rollout manages a deployment"},
	{Text: "label", Description: "Update the labels on a resource"},
	{Text: "annotate", Description: "Update the annotations on a resource"},
	{Text: "config", Description: "config modifies kubeconfig files"},
	{Text: "cluster-info", Description: "Display cluster info"},
	{Text: "api-versions", Description: "Print the supported API versions on the server, in the form of 'group/version'."},
	{Text: "version", Description: "Print the client and server version information."},
	{Text: "explain", Description: "Documentation of resources."},
	{Text: "convert", Description: "Convert config files between different API versions"},
}

func secondArgsCompleter(first, second string) []prompt.Completion {
	switch first {
	case "get":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), second, true)
	case "describe":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), second, true)
	case "create":
		subcommands := []prompt.Completion{
			{Text: "configmap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "deployment", Description: "Create a deployment with the specified name."},
			{Text: "namespace", Description: "Create a namespace with the specified name"},
			{Text: "quota", Description: "Create a quota with the specified name."},
			{Text: "secret", Description: "Create a secret using specified subcommand"},
			{Text: "service", Description: "Create a service using specified subcommand."},
			{Text: "serviceaccount", Description: "Create a service account with the specified name"},
		}
		return prompt.FilterHasPrefix(subcommands, second, true)
	case "replace":
	case "patch":
	case "delete":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), second, true)
	case "edit":
	case "apply":
	case "namespace":
	case "logs":
	case "rolling-update":
	case "scale":
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	case "uncordon":
		return prompt.FilterHasPrefix(getNodeCompletions(), second, true)
	//case "attach": // still not supported
	//case "exec":   // still not supported
	case "port-forward":
	case "proxy":
	case "run":
	case "expose":
	case "autoscale":
	case "rollout":
	case "label":
	case "annotate":
	case "config":
	case "cluster-info":
		subCommands := []prompt.Completion{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		return prompt.FilterHasPrefix(subCommands, second, true)
	case "api-versions":
	case "version":
	case "explain":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), second, true)
	case "convert":
	default:
		return []prompt.Completion{}
	}
	return []prompt.Completion{}
}

func thirdArgsCompleter(first, second, third string) []prompt.Completion {
	switch first {
	case "describe":
		switch second {
		case "po":
			fallthrough
		case "pods":
			return prompt.FilterContains(getPodCompletions(), third, true)
		case "deploy":
			fallthrough
		case "deployments":
			return prompt.FilterContains(strToCompletionList(getDeploymentNames()), third, true)
		case "no":
			fallthrough
		case "nodes":
			return prompt.FilterContains(getNodeCompletions(), third, true)
		}
	}
	return []prompt.Completion{}
}
