import requests

class client():
    def __init__(self , url):
        self.url = url

    def get():
        pass 

    def doRequest(self ,method , params):
        link = "".join([self.url , params])
        res =  requests.request(method , url = link)
        return res.text