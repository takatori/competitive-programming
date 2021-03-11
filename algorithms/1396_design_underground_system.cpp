using namespace std;

class UndergroundSystem
{

private:
    unordered_map<string, unordered_map<string, tuple<long, long>>> travels;
    unordered_map<int, tuple<string, int>> checkInHistory;

public:
    UndergroundSystem()
    {
    }

    void checkIn(int id, string stationName, int t)
    {
        this->checkInHistory[id] = make_tuple(stationName, t);
    }

    void checkOut(int id, string stationName, int t)
    {
        auto history = this->checkInHistory[id];
        string checkInStation = get<0>(history);
        int checkInTime = get<1>(history);

        auto current = this->travels[checkInStation][stationName];
        long totalTime = get<0>(current);
        long count = get<1>(current);
        totalTime = totalTime + (t - checkInTime);
        this->travels[checkInStation][stationName] = make_tuple(totalTime, count++);
    }

    double getAverageTime(string startStation, string endStation)
    {
        auto station = this->travels[startStation][endStation];
        return double(get<0>(station)) / get<1>(station);
    }
};

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * UndergroundSystem* obj = new UndergroundSystem();
 * obj->checkIn(id,stationName,t);
 * obj->checkOut(id,stationName,t);
 * double param_3 = obj->getAverageTime(startStation,endStation);
 */