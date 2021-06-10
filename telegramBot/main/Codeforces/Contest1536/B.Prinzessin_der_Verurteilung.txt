#include <bits/stdc++.h>
using namespace std;
#define ll long long

bool good(string &s, string &t) {
    for(int i = 0, j = 0; i < s.length(); ++i) {
        if(s[i] == t[j]) {
            bool gd = 1;
            int k = 0;
            for(int l = i; l < s.length() && k < t.length(); ++k, ++l) {
                if(s[l] != t[k]) {
                    gd = 0;
                    break;
                } 
            }
            if(gd && k == t.length()) return 0;
        }
    }
    return 1;
}

string ans = "";

void rec(string &s, string &t) {
    for(char a = 'a'; a <= 'z'; ++a) {
        t += a;
        if(good(s, t)) {
            if(ans == "" || ans.length() > t.length()) ans = t;
            else if(ans.length() == t.length()) ans = min(ans, t);
            t.pop_back();
            return;
        }
        t.pop_back();
    }
    if(t.length() < 20) {
        for(char a = 'a'; a <= 'z'; ++a) {
            t += a;
            rec(s, t);
            t.pop_back();
        }
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(0); cout.tie(0);
    int t = 1; cin >> t;
    while(t--) {
        int n; cin >> n;
        string s; cin >> s;
        string b = "";
        ans = "";
        rec(s, b);
        cout << ans << endl;
    }
    return 0;
}