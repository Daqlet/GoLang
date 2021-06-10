#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n, L, R; cin >> n >> L >> R;
    int a[n];
    for(int i = 0; i < n; ++i) {
        cin >> a[i];
    }
    sort(a, a+n);
    ll ans = 0;
    for(int i = 0; i < n; ++i) {
        int l = i, r = n;
        int mn = max(0, L-a[i]);
        int mx = max(0, R-a[i]);
        while(l < r-1) {
            int m = (l + r)/2;
            if(a[m] < mn) l = m;
            else r = m;
        }
        if(r == n) continue;
        mn = r;
        l = i; r = n;
        while(l < r-1) {
            int m = (l + r)/2;
            if(a[m] > mx) r = m;
            else l = m;
        }
        if(l == i) continue;
        mx = l;
        //cout << a[i] << " " << mn << ' ' << mx << endl;
        ans += (mx - mn + 1);
    }
    cout << ans << endl;
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