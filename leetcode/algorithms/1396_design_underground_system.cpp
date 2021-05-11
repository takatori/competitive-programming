using namespace std;

class UndergroundSystem
{

private:
    unordered_map<string, pair<int, int>> checkoutMap;
    unordered_map<int, pair<string, int>> checkinMap;

public:
    UndergroundSystem()
    {
    }

    void checkIn(int id, string stationName, int t)
    {
        this->checkinMap[id] = {stationName, t};
    }

    void checkOut(int id, string stationName, int t)
    {
        auto& checkIn = checkinMap[id];
        string route = checkIn.first + "_" + stationName;
        checkoutMap[route].first += t - checkIn.second;
        checkoutMap[route].second += 1;
    }

    double getAverageTime(string startStation, string endStation)
    {
        string route = startStation + "_" + endStation;
        auto& checkout = checkoutMap[route];
        return (double) checkout.first / checkout.second;
    }
};
