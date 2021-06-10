#include <bits/stdc++.h>
using namespace std;
#define ll long long

int count(int n) {
    int res = 0;
    for(int i = 2; i <= sqrt(n); ++i) {
        while(n%i == 0) {
            n /= i;
            res++;
        }
    }
    if(n > 1) ++res;
    return res;
}

void solve() {
    int a, b, k; cin >> a >> b >> k;
    if(k > 60) {
        cout << "No" << endl;
        return;
    }
    int c1 = count(a);
    int c2 = count(b);
    if(k == 1 && a == b) cout << "No" << endl;
    else if(k == 1 && (a%b == 0 || b%a == 0)) cout << "Yes" << endl;
    else if(k == 1) cout << "No" << endl;
    else if(k <= c1+c2) cout << "Yes" << endl;
    else cout << "No" << endl;
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