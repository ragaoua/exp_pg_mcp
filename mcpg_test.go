package main

import (
	"log"
	"testing"
	"context"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func TestMcpg(t *testing.T) {
	log.Println("##################################################################")
	log.Println("################## Connecting to the MCP server ##################")
	log.Println("##################################################################")

	c, err := client.NewStreamableHttpClient("http://localhost:8080/mcp")
	if err != nil {
		t.Logf("Error instantiating client : %v", err)
		t.Fail()
		return
	}
	defer func() {
		err = c.Close()
		if err != nil {
			t.Logf("Error closing client : %v", err)
			t.Fail()
		}
	}()

	ctx := context.Background()

	initRequest := mcp.InitializeRequest{}
	_, err = c.Initialize(ctx, initRequest)
	if err != nil {
		t.Logf("Error initializing client : %v", err)
		t.Fail()
		return
	}

        log.Println("Connection successful")
        log.Println()
        log.Println()


        log.Println("#############################################################")
        log.Println("################## Listing available tools ##################")
        log.Println("#############################################################")
	toolsRequest := mcp.ListToolsRequest{}
	tools, err := c.ListTools(ctx, toolsRequest)
	if err != nil {
		t.Logf("Error listing tools : %v", err)
		t.Fail()
		return
	}

	log.Println("Available tools:")
	for _, tool := range tools.Tools {
		log.Printf("%v : %v\n", tool.Name, tool.Description)
	}
        log.Println()
        log.Println()


        log.Println("##############################################################")
        log.Println("################## Executing list_all_roles ##################")
        log.Println("##############################################################")
	result, err := c.CallTool(
		ctx,
		mcp.CallToolRequest{
			Params: mcp.CallToolParams{ Name: "list_all_roles" },
		},
	)
	if err != nil {
		t.Logf("Error executing tool : %v", err)
		t.Fail()
		return
	}
	log.Println("Result:")
	log.Printf("%v", result.Content)
}
