#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{
    int n,s = 0,mx = 0,k = 0; cin >> n;
    vector<int> a(n);
    map<int,int> mp;
    for(int i = 0; i < n; ++i){
        cin >> a[i];
        mp[a[i]]++;
        s += a[i];
        mx = max(a[i],mx);
    }
    if(mp[a[0]] == n || n == 1) cout << 0 << '\n';
    else if(s%n != 0) cout << -1 << '\n';
    else {
        int ave = s/n;
        for(int i = 0; i < n; ++i) if(a[i] > ave) k++;
        cout << k << '\n';
    }
}
int main() {
    fast;
    int t = 1;
    cin >> t;
    while(t--) solution();
}
