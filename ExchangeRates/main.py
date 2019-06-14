from Exchange import Exchange
from ExchangeRates import exchangeService 
from Client import client

newResponse = exchangeService("https://api.exchangeratesapi.io").get_by_history('11','12','2011','11','12','2013')
print(newResponse)
# anotherResponse = client.doRequest(method="GET",params="")
# a = Exchange(anotherResponse)
# b = exchangeService(a).get_by_base
# print(b)
