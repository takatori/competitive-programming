using namespace std;

class Solution {
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
      int L = beginWord.length();
      unordered_map<string, vector<string>> allComboDict;
      
      for (string word : wordList) {
        for (int i = 0; i < L; i++) {
          string newWord = word.substr(0, i) + '*' + word.substr(i+1, L);
          vector<string> transformations = (allComboDict.find(newWord) != allComboDict.end()) ? allComboDict[newWord] : vector<string>(); 
          transformations.push_back(word);
          allComboDict[newWord] = transformations;
        }
      }
       queue<pair<string, int>> q;
       q.push(make_pair(beginWord, 1));
       unordered_map<string, bool> visited;
        visited[beginWord] = true;
        
        while (!q.empty()) {
          auto node = q.front();
          q.pop();
          string word = node.first;
          int level = node.second;
          for (int i = 0; i < L; i++) {
            string newWord = word.substr(0, i) + '*' + word.substr(i+1, L);
            for (string adjacentWord : allComboDict[newWord]) {
              if (adjacentWord == endWord) {
                return level + 1;
              }
              if (visited.find(adjacentWord) == visited.end()) {
                visited[adjacentWord] = true;
                q.push(make_pair(adjacentWord, level+1));
              }
            }
          }
        }
      return 0;
    }
};