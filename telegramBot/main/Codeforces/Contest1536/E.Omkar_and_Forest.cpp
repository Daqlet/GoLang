#include <bits/stdc++.h>
#include <stdio.h>
using namespace std;

#define all(v) v.begin(), v.end()
#define eb emplace_back
#define ll long long

namespace MATH_CAL {
    const int mod = 1e9 + 7;
    inline ll add(ll a, ll b) { return a + b >= mod ? a + b - mod : a + b; }
    inline ll sub(ll a, ll b) { return a - b < 0 ? a - b + mod : a - b; }
    inline ll mul(ll a, ll b) { return (ll) a * b % mod; }
    inline void Add(ll &a, ll b) { (a += b) >= mod ? a -= mod : 0; }
    inline void Sub(ll &a, ll b) { (a -= b) < 0 ? a += mod : 0; }
    inline ll qpow(ll x, ll n) { ll r = 1; for ( ; n; n /= 2, x = mul(x, x)) if (n & 1) r = mul(r, x); return r; }
    inline ll mod_inv(ll x) { return qpow(x, mod - 2); }
    inline ll gcd(ll a, ll b) {while(b) {a %= b; swap(a, b);} return a;}
    inline ll lcm(ll a, ll b) {return a * b / gcd(a, b);}
} using namespace MATH_CAL;


const int N = 2e3 + 5;

int dx[] = {-1, 0, 1, 0};
int dy[] = {0, 1, 0, -1};

int a[N][N];
bool used[N][N];
int cnt;
int n, m;

void dfs(int x, int y) {
    used[x][y] = 1;
    cnt += 1;
    for(int k = 0; k < 4; k += 1) {
        int to_x = dx[k] + x;
        int to_y = dy[k] + y;
        if (to_x >= 0 && to_x < n && to_y >= 0 && to_y < m && !used[to_x][to_y] && a[to_x][to_y]) {
            dfs(to_x, to_y);
        }
    }
}

ll calc(int x) {
    return qpow(2, x);
}

void solve() {
    cin >> n >> m;
    bool any = false;
    for(int i = 0; i < n; i += 1) {
        for(int j = 0; j < m; j += 1) {
            char c;
            cin >> c;
            if (c == '#') {
                a[i][j] = 1;
            } else {
                a[i][j] = 0;
            }
            if (!a[i][j]) {
                any = true;
            }
            used[i][j] = 0;
        }
    }
    vector<int> components;
    for(int i = 0; i < n; i += 1) {
        for(int j = 0; j < m; j += 1) {
            if (a[i][j] && !used[i][j]) {
                dfs(i, j);
                components.push_back(cnt);
                cnt = 0;
            }
        }
    }
    ll res = 1;
    for(int x : components) {
        res = mul(res, calc(x));
    }
    if (components.size() == 1 && !any) {
        res -= 1;
    }
    cout << res << endl;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(0);
    cout.tie(0);
//    freopen("taskA.in", "r", stdin);
//    freopen("taskA.out", "w", stdout);
    int t = 1;
    cin >> t;
    while(t--) {
        solve();
    }
    return 0;
}
