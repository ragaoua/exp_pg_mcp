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
podman run --rm --name tmp_pg -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 postgres
~~~

Next, run the MCP server :

~~~bash
go run .
~~~

Finally, use the python script to test access to the MCP server and test the execution of the list_all_roles function :

~~~bash
python -m venv venv
source venv/bin/activate
pip install fastmcp
python test.py
~~~
