import os, sys
import json
import traceback

from fastmcp import Client
import asyncio

from openai import OpenAI

async def main():
    for var in ["OPENAI_BASE_URL", "MCP_HOST", "MODEL"]:
        if var not in os.environ:
            print(f"Variable {var} not found")
            sys.exit(1)

    model = os.getenv("MODEL")
    client = OpenAI(api_key="")
    messages = [
        {
            "role": "system",
            "content": "You're a helpful assistant for a database system administrator"
        },
        {
            "role": "user",
            "content": "Can you tell me how many roles there are in the cluster ?"
        }
    ]

    try:
        async with Client(os.getenv("MCP_HOST")) as mcpClient:
            tools = await mcpClient.list_tools()

            while True:
                response = client.chat.completions.create(
                    model=model,
                    messages=messages,
                    tools=[
                        {
                            "type": "function",
                            "function": {
                                "name": tool.name,
                                "description": tool.description,
                            },
                            "parameters": tool.inputSchema
                        } for tool in tools
                    ],
                ).choices[0].message

                if not response.tool_calls:
                    messages.append({
                        "role": "assistant",
                        "content": response.content
                    })
                    break

                messages.append({
                    "role": "assistant",
                    "tool_calls": [
                        {
                            "id": tc.id,
                            "function": tc.function,
                            "type": tc.type
                        } for tc in response.tool_calls
                    ]
                })

                for tc in response.tool_calls:
                    args = json.loads(tc.function.arguments)

                    tc_result = await mcpClient.call_tool(tc.function.name, **args)
                    tc_result_content = tc_result.content[0].text
                    messages.append({
                        "role": "tool",
                        "tool_call_id": tc.id,
                        "name": tc.function.name,
                        "content": tc_result_content
                    })
    except Exception:
        print(traceback.format_exc())

    print("####################################")
    print(messages)
    print("####################################")

if __name__ == "__main__":
    asyncio.run(main())
