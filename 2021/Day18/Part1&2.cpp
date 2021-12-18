#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <math.h>
using namespace std;

class Node {
    Node* left = nullptr; Node* right = nullptr; Node* parent = nullptr;
    int d = 0, l, r;

    public:
    static Node* fromString(string s, int depth){
        string aux; int open = 0;
        Node* n = new Node();
        n->d = depth;
        for(int i = 0; i < s.size(); i++){
            if(i != 0)
                aux.push_back(s[i]);

            if(s[i] == '['){
                open++;
            }else if (s[i] == ',' && open == 1){
                aux.pop_back();
                if(aux.size() == 1){
                    n->l = aux[0] - '0';
                }else{
                    n->left = Node::fromString(aux, depth + 1);
                    n->left->parent = n;
                }
                aux = string(s.begin() + i + 1, s.end());
                aux.pop_back();
                if(aux.size() == 1){
                    n->r = aux[0] - '0';
                }else{
                    n->right = Node::fromString(aux, depth + 1);
                    n->right->parent = n;
                }

            }else if(s[i] == ']')
                open--;
        }
        return n;
    }

    string toString(){
        string aux = "[";
        if(left == nullptr)
            aux+= to_string(l);
        else
            aux += left->toString();

        aux += ",";
        if(right == nullptr)
            aux+= to_string(r);
        else
            aux += right->toString();

        return aux + "]";
    }

    void reduce(){
        explode();
        while(split())
            explode();
    }

    bool split(){
        if(l > 9){
            left = new Node();
            left->l = l / 2;
            left->r = ceil(1.0*l/2);
            left->d = d + 1;
            left->parent = this;
            l = 0;
            return true;
        }
        if(left != nullptr){
            if(left->split())
                return true;
        }
        if(r > 9){
            right = new Node();
            right->l = r / 2;
            right->r = ceil(1.0*r/2);
            right->d = d + 1;
            right->parent = this;
            r = 0;
            return true;
        }
        if(right != nullptr){
            if(right->split())
                return true;
        }
        return false;
    }
    
    void explode(){
        if(this->left != nullptr)
            this->left->explode();

        if(this->right != nullptr)
            this->right->explode();

        if(d >= 4){
            auto p = this->parent;
            char c = ' ';
            if(this->parent != nullptr && this->parent->left == this){
                c = 'l';
            }else if(this->parent != nullptr && this->parent->right == this){
                c = 'r';
            }
            while(p != nullptr){
                if(p->left == nullptr){
                    p->l += this->l;
                    break;
                }else if(p->left != nullptr && c == 'r'){
                    auto nn = p->left; bool f = false;
                    while (nn != nullptr){
                        if(nn->right == nullptr){
                            nn->r += this->l;
                            f = true;
                            break;
                        }
                        nn = nn->right;
                    }
                    if(f) break;
                }
                if(p->parent != nullptr && p->parent->left == p){
                    c = 'l';
                }else if(p->parent != nullptr && p->parent->right == p){
                    c = 'r';
                }
                p = p->parent;
            }

            p = this->parent;
            if(this->parent != nullptr && this->parent->left == this){
                c = 'l';
            }else if(this->parent != nullptr && this->parent->right == this){
                c = 'r';
            }
            while(p != nullptr){
                if(p->right == nullptr){
                    p->r += this->r;
                    break;
                }else if(p->right != nullptr && c == 'l'){
                    auto nn = p->right; bool f = false;
                    while (nn != nullptr){
                        if(nn->left == nullptr){
                            nn->l += this->r;
                            f = true;
                            break;
                        }
                        nn = nn->left;
                    }
                    if(f) break;
                }
                if(p->parent != nullptr && p->parent->left == p){
                    c = 'l';
                }else if(p->parent != nullptr && p->parent->right == p){
                    c = 'r';
                }
                p = p->parent;
            }
            if(this->parent->left == this){
                this->parent->left = nullptr;
                this->parent->l = 0;
            }else if(this->parent->right == this){
                this->parent->right = nullptr;
                this->parent->r = 0;
            }
        }
    }

    long long magnitude(){
        long long ans = 0;
        if(left == nullptr)
            ans += 1LL *l * 3;
        else
            ans += left->magnitude() * 3;

        if(right == nullptr)
            ans += 1LL * r * 2;
        else
            ans += right->magnitude() * 2;

        return ans;
    }
};

int main(){
    string s1, s2;
    vector<string> ss;
    cin >> s1;
    ss.push_back(s1);
    Node* total;
    while(cin >> s2){
        auto res = "[" + s1 + "," + s2 + "]";
        ss.push_back(s2);
        total = Node::fromString(res, 0);
        total->reduce();
        s1 = total->toString();
    }
    cout << "Part 1: " << total->magnitude() << endl;

    long long maxE = 0;
    for(int i = 0; i < ss.size(); i++){
        for(int j = i + 1; j < ss.size(); j++){
            auto res = "[" + ss[i] + "," + ss[j] + "]";
            total = Node::fromString(res, 0);
            total->reduce();
            maxE = max(maxE, total->magnitude());
            res = "[" + ss[j] + "," + ss[i] + "]";
            total = Node::fromString(res, 0);
            total->reduce();
            maxE = max(maxE, total->magnitude());
        }
    }
    cout << "Part 2: " << maxE;

    return 0;
}