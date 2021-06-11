#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n; cin >> n;
    int a[n];
    ll sum = 0;
    for(int i = 0; i < n; ++i) {
        cin >> a[i];
        sum += a[i];
    }
    if(sum%n != 0) cout << -1 << endl;
    else {
        ll k = 0, av = sum/n;
        for(int i = 0; i < n; ++i) {
            if(a[i] > av) ++k;
        }
        cout << k << endl;
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(0); cout.tie(0);
    int t = 1; cin >> t;
    while(t--) {
        solve();
    }
    return 0;
}