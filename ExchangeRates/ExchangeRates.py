from Exchange import Exchange

class exchangeService(Exchange):
    def __init__(self , url):
        super().__init__(url)

    def get_latest(self):
        querystring = '/latest'
        return self.get(params = querystring)

    def get_by_time(self , day , month , year):
        querystring = "/{}-{}-{}".format(year , month , day)
        return self.get(params = querystring)

    def get_by_base(self , moneyunit):
        querystring = "/latest?base={}".format(moneyunit)
        return self.get(params = querystring)

    def get_by_symbols(self , moneyunit , type):
        querystring =  "/latest?symbols={},{}".format(moneyunit , type)
        return self.get(params = querystring)

    def get_by_history(self , day1 , month1 , year1 , day2 , month2 , year2):
        querystring =  "/history?start_at={}-{}-{}&end_at={}-{}-{}".format(year1 , month1 , day1, year2 , month2 , day2)
        return self.get(params = querystring)

    def get_by_history_with_symbols(self , day1 , month1 , year1 , day2 , month2 , year2 , moneyunit , type):
        querystring =  "/history?start_at={}-{}-{}&end_at={}-{}-{}&symbols={},{}".format(year1 , month1 , day1 , year2 , month2 , day2 , moneyunit , type)
        return self.get(params = querystring)

    def get_by_history_with_base(self , day1 , month1 , year1 , day2 , month2 , year2 , moneyunit):
        querystring =  "/history?start_at={}-{}-{}&end_at={}-{}-{}&base={}".format(year1 , month1 , day1 , year2 , month2 , day2 , moneyunit)
        return self.get(params = querystring)
     

