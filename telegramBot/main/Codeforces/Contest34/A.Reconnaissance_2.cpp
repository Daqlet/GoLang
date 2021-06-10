#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{
    int n,x,y; cin >> n;
    int a[n];
    for(int i = 0; i < n ; ++i) cin >> a[i];
    int mn = abs(a[0]-a[n-1]);
    x = 1; y = n;
    for(int i = 1; i < n; ++i){
        if(mn > abs(a[i]-a[i-1])){
            mn = abs(a[i]-a[i-1]);
            x = i+1;
            y = i;
        }
    }
    cout << x << ' ' << y << '\n';
}
int main() {
    fast;
    int t = 1;
    //cin >> t;
    while(t--) solution();
}
