#include <bits/stdc++.h>
using namespace std;
#define ll long long

void solve() {
    int n; cin >> n;
    int a[n];
    int mx = 0, mn = 0;
    for(int i = 0; i < n; ++i) {
        cin >> a[i];
        if(a[i] > a[mx]) mx = i;
        if(a[i] < a[mn]) mn = i;
    }
    mx++; ++mn; ++n;
    int f = min(max(mn, mx), max(n-mn, n-mx));
    int s = min(mx + n - mn, mn + n - mx);
    cout << min(f, s) << endl;
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