An MCP server providing tools for interacting with a PostgreSQL database.

This is a small project for me to learn Golang.

# Tools

- list all roles
- list all databases
- list all schemas in a database
- list all objects in a schema or a database
- execute a SQL query

# Test

First, spawn a PostgreSQL database cluster :

~~~bash
podman run --rm --detach --name mcpg_tmp_pg -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgres
~~~

Next, run tests :

~~~bash
cd server
go test
~~~

This will start the MCP server and create a client to access it and test the execution of the list_all_roles function
