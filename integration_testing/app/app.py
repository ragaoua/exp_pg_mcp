import os, sys
import asyncio

from openai import AsyncOpenAI

from agents import Agent, Runner, OpenAIChatCompletionsModel
from agents.mcp import MCPServerStreamableHttp
from agents.run_context import RunContextWrapper


async def main():
    for var in ["OPENAI_BASE_URL", "MCP_HOST", "MODEL"]:
        if var not in os.environ:
            print(f"Variable {var} not found")
            sys.exit(1)

    openaiBaseUrl = os.getenv("OPENAI_BASE_URL")
    mcpHost = os.getenv("MCP_HOST")
    model = os.getenv("MODEL")

    client = AsyncOpenAI(
        base_url=openaiBaseUrl,
        api_key="",
    )

    async with MCPServerStreamableHttp(params={"url": mcpHost}) as mcpServer:
        agent = Agent(
            name="Assistant",
            instructions="use the tools to answer the questions",
            model=OpenAIChatCompletionsModel(
                model=model,
                openai_client=client,
            ),
            mcp_servers=[mcpServer],
        )

        result = await Runner.run(
            starting_agent=agent,
            input="Can you tell me how many roles there are in the cluster ?",
        )
        print(result)


if __name__ == "__main__":
    asyncio.run(main())
