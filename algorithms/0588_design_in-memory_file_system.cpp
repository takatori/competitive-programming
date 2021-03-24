using namespace std;

class Node
{
    string name;
    bool is_file;
    string content;
    vector<reference_wrapper<Node> > nodes;

public:
    Node()
    {
        this->name = "";
        this->is_file = false;
        this->nodes = vector<reference_wrapper<Node> >();
    }

    Node(string name, bool is_file, string content)
    {
        this->name = name;
        this->is_file = is_file;
        this->content = content;
        this->nodes = vector<reference_wrapper<Node> >();
    }

    Node &find(string name)
    {
        for (Node &node : nodes)
        {
            if (node.name == name)
                return node;
        }
        return &Node("*", false, "");
    }
}

class FileSystem
{
    &Node root;

public:
    FileSystem()
    {
        root = &Node();
    }

    vector<string> ls(string path)
    {
        vector<string> paths = split(path, '/');
        &Node target = root;
        for (string p : paths)
        {
            target = target.find(p);
        }
        vector<string> results;
        for (Node &node : target.nodes)
        {
            results.push_back(node.name);
        }
        return results;
    }

    void mkdir(string path)
    {
        vector<string> paths = split(path, '/');
        &Node target = root;
        for (string p : paths)
        {
            auto t = target.find(p);
            if (t.name == "*")
            {
                Node &newNode = &Node(p, false, "")
                                     target.nodes.push_back(newNode);
                target = newNode;
            }
            else
            {
                target = t;
            }
        }
    }

    void addContentToFile(string filePath, string content)
    {
        vector<string> paths = split(path, '/');
        &Node target = root;
        string fileName = paths.back();
        paths.pop_back();
        for (string p : paths)
        {
            target = target.find(p);
        }
        auto n = target.find(fileName);
        if (n.name == "*")
        {
            &Node file = &Node(fileName, true, content);
            target.nodes.push_back(file);
        }
        else
        {
            n.content += content;
        }
    }

    string readContentFromFile(string filePath)
    {
        vector<string> paths = split(path, '/');
        &Node target = root;
        for (string p : paths)
        {
            target = target.find(p);
        }
        return target.content;
    }

    vector<string> split(const string &str, char sep)
    {
        vector<string> v;
        stringstream ss(str);
        string buffer;
        while (getline(ss, buffer, sep))
        {
            v.push_back(buffer);
        }
        return v;
    }
};