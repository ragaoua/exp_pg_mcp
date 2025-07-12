import asyncio
from fastmcp import Client

async def main():
    print("##################################################################")
    print("################## Connecting to the MCP server ##################")
    print("##################################################################")

    async with Client("http://localhost:8080/mcp") as client:
        print("Connection successful")
        print()
        print()

        print("#############################################################")
        print("################## Listing available tools ##################")
        print("#############################################################")
        tools = await client.list_tools()
        print("Available tools:")
        print(f"{tools}")
        print()
        print()


        print("##############################################################")
        print("################## Executing list_all_roles ##################")
        print("##############################################################")
        result = await client.call_tool("list_all_roles")
        print("Result:")
        print(f"{result}")
        print()
        print()

asyncio.run(main())
