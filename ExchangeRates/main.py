from Exchange import Exchange
from ExchangeRates import exchangeService 
from Client import client

newResponse = exchangeService("https://api.exchangeratesapi.io").get_by_history_with_base('11','12','2011','11','12','2013','USD')


print(newResponse)
