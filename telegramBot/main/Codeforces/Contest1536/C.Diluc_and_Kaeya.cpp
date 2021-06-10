#include <bits/stdc++.h>
#include <stdio.h>
using namespace std;

#define all(v) v.begin(), v.end()
#define eb emplace_back
#define ll long long

#define fr first
#define sc second

const int N = 5e5 + 5;

int lp[N];
int pref[2][N];

vector<int> divs[N];

void precalc() {
   /* for(int i = 2; i < N; i += 1) {
        if (!lp[i]) {
            lp[i] = i;
        }
        if (lp[i] != i) continue;
        for(int j = i + i; j < N; j += i) {
            if (!lp[j]) {
                lp[j] = lp[i];
            }
        }
    }*/
    for(int i = 1; i < N; i += 1) {
        for(int j = 2; j * j <= i; j += 1) {
            if (i % j == 0) {
                divs[i].push_back(j);
                if (j != i / j) {
                    divs[i].push_back(i / j);
                }
            }
        }
        divs[i].push_back(1);
    }
}

void solve() {
    int n;
    cin >> n;
    string s;
    cin >> s;
    pref[0][0] = pref[1][0] = 0;
    for(int i = 0; i < n; i += 1) {
        int x = (s[i] == 'K');
        if (i > 0) {
            pref[0][i] = pref[0][i - 1];
            pref[1][i] = pref[1][i - 1];
        }
        pref[x][i] += 1;
    }
    auto get_pair = [&] (int i, int j) {
        return make_pair(pref[0][j] - (i == 0 ? 0 : pref[0][i - 1]),
                pref[1][j] - (i == 0 ? 0 : pref[1][i - 1]));
    };

    map<pair<int, int>, int> pos;
    for(int i = 0; i < n; i += 1) {
        int len = i + 1;
        int res = 1;
        pair<ll, ll> coef = get_pair(0, i);
        ll g = gcd(coef.fr, coef.sc);
        coef.fr /= g;
        coef.sc /= g;
        pos[coef] += 1;
        cout << pos[coef] << ' ';
    }
    cout << endl;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);
    //precalc();
//    freopen("taskA.in", "r", stdin);
//    freopen("taskA.out", "w", stdout);
    int t = 1;
    cin >> t;
    while(t--) {
        solve();
    }
    return 0;
}
