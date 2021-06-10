#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{
    int n; cin >> n;
    vector<int> b;
    int a[n];
    for(int i = 0; i < n; ++i) cin >> a[i];
    sort(a,a+n);
    int c = 0,l = 0, r = n-1;
    while(l < r){
        b.pb(a[r--]);
        b.pb(a[l++]);
        ++c;
    }
    if(l != r) c--;
    cout << c << '\n';
    for(int i = 0; i < b.size(); ++i) cout << b[i] << ' ';
    if(n%2 != 0) cout << a[l];
}
int main() {
    fast;
    int t = 1;
    //cin >> t;
    while(t--) solution();
}
