import asyncio
import websockets
import json

listfruits = '{"orange":"Qua cam", "strawberry":"Day tay", '\
             '"grape":"Nho", "durian":"Sau rieng"}'

async def message():
    async with websockets.connect("ws://localhost:8000") as socket:
        msg = listfruits
        await socket.send(msg)
        print(await socket.recv())


asyncio.get_event_loop().run_until_complete(message())

