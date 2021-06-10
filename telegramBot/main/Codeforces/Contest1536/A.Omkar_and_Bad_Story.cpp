#include <bits/stdc++.h>
#include <stdio.h>
using namespace std;

#define all(v) v.begin(), v.end()
#define eb emplace_back
#define ll long long

void solve() {
    int n;
    cin >> n;
    vector<int> a(n);
    bool neg = 0;
    for(int i = 0; i < n; i += 1) {
        cin >> a[i];
        if (a[i] < 0) {
            neg = 1;
        }
    }
    if (neg) {
        cout << "NO\n";
        return;
    }
    cout << "YES\n";
    cout << 200 << "\n";
    for(int i = 0; i < 200; i += 1) {
        cout << i << ' ';
    }
    cout << "\n";
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
