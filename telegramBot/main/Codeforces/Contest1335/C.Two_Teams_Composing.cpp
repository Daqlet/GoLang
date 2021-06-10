#include <bits/stdc++.h>
using namespace std;
#define ll long long
#define fast {ios_base::sync_with_stdio(0), cin.tie(0), cout.tie(0);}
#define pb push_back

ll binPow(ll a,ll n){ if(n == 0) return 1; if(n%2 == 0) return binPow(a*a,n/2); else return a*binPow(a,n-1);}
bool check(vector<int> &a, int x) {
  map<int, int> m;
  for(int i = 0; i < a.size(); ++i) {
    m[a[i]]++;
  }
  int mx = 0;
  for(auto i = m.begin(); i != m.end(); ++i) {
    mx = max(mx, i->second);
  }
  if((mx > x && m.size() >= x) || (mx == x && m.size() > x)) return 1;
  else return 0;
}

void solution(){
    int n; cin >> n;
    vector<int> a(n);
    for(int i = 0; i < n; ++i) cin >> a[i];
    int l = 0, r = n;
    while(l < r-1) {
      int m = l + ((r - l)>>1);
      if(check(a, m)) l = m;
      else r = m;
    }
    cout << l << '\n';
}
int main() {
    fast;
    int t = 1;
    cin >> t;
    while(t--) solution();
}
