package server

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Start(db_url string) error {
	mcp_server := server.NewMCPServer(
		"MCPG",
		"0.1",
		server.WithToolCapabilities(true),
	)

	mcp_server.AddTool(
		mcp.NewTool(
			"list_all_roles",
			mcp.WithDescription("list all roles in the cluster"),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return listAllRolesHandler(ctx, request, db_url)
		},
	)

	log.Println("Starting StreamableHTTP server on :8080")
	httpServer := server.NewStreamableHTTPServer(mcp_server)
	err := httpServer.Start(":8080")
	return err
}

func listAllRolesHandler(ctx context.Context, request mcp.CallToolRequest, db_url string) (*mcp.CallToolResult, error) {
	roles, err := listAllRoles(db_url)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", roles)), nil
}

func listAllRoles(db_url string) ([]string, error) {
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT rolname FROM pg_roles")
	if err != nil {
		return nil, fmt.Errorf("Query error : %v\n", err)
	}

	var roles []string
	for rows.Next() {
		var role string

		err := rows.Scan(&role)
		if err != nil {
			return nil, fmt.Errorf("Fetching error : %v\n", err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}
