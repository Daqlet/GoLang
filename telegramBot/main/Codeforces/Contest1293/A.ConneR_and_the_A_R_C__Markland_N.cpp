#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{
    int n,s,k; cin >> n >> s >> k;
    map<int,bool> mp;
    for(int i = 0; i < k; ++i){
        int x;
        cin >> x;
        mp[x] = 1;
    }
    int ans = INT_MAX;
    for(int i = s-1; i >= 1; --i){
        if(!mp[i]) {ans = min(ans,s-i);break;}
    }
    for(int i = s; i <= n; ++i)
        if(!mp[i]) {ans = min(ans,i-s); break;}
    cout << ans << '\n';
}
int main() {
    fast;
    int t = 1;
    cin >> t;
    while(t--) solution();
}
