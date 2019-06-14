from Client import client


class Exchange(client):
    def __init__(self , url):
        super().__init__(url)
    
    def get(self , params):
        return self.doRequest("GET" , params)