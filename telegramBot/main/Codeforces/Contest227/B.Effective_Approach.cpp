#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{

}
int main() {
    fast;
    int n, m, x;
    cin >> n;
    vector<int> a(n+1);
    for (int i = 0; i < n; i++) {
        cin >> x;
        a[x] = i;
    }
    cin >> m;
    ll v = 0, p = 0;
    for (int i = 0; i < m; i++) {
        cin >> x;
        v += a[x] + 1;
        p += n - a[x];
    }
    cout << v << ' ' << p;
}
