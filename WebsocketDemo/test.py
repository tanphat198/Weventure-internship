import json

from marshmallow import Schema, fields
from websocket import create_connection
import ast

class Test:
    def __init__(self, id, name):
        self._id = id
        self.name = name

class TestSchema(Schema):
    id = fields.String()
    name = fields.String()


ws = create_connection("ws://localhost:8080/python")
res = ws.recv()
a = res.decode('utf-8')
d = json.loads(a)
schema = TestSchema()
result = schema.load(d)
testobject = Test(**result)
print(testobject._id + " " + testobject.name)
