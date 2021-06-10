#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}

void solution()
{
   int n; cin >> n;
   int mx = 1;
   set<int> s;
   map<int,int> mp;
   for(int i = 0; i < n ; ++i){
        int a;
       cin >> a;
       s.insert(a);
       mp[a]++;
       mx = max(mp[a],mx);
   }
   int m = s.size();
   if(n == 1) cout << 0;
   else if(n - m == 0) cout << 1;
   else if(mx >= m) cout << min(mx-1,m);
   else cout << min(mx,m);
   cout << '\n';
}
int main() {
    fast;
    int t = 1;
    cin >> t;
    while(t--) solution();
}
