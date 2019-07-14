import asyncio
import websockets


async def response(websocket, path):
    message = await websocket.recv()
    print("We got the message from the client: ",message)
    await websocket.send("server got your message!")

start_server = websockets.serve(response, 'localhost', 8000)
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()

