#include <iostream>
using namespace std;

int ** a;
int n = 51;
int * was;

bool dfs(int key){
  was[key] = 1;
  for(int i = 0; i < n; i++){
    if(a[key][i]){
      if(was[i] == 0) {if(dfs(i)) return true;}
      else if(was[i] == 1) { return true;}
    }
  }
  was[key] = 2;
  return false;
}

int main() {
  cin >> n; a = new int *[n]; was = new int [n];
  for(int i = 0; i < n; i++){
    a[i] = new int [n]; was[i] = 0;
    for(int j = 0; j < n; j++) cin >> a[i][j];
  }
  for(int i = 0; i < n; i++){
    if(!was[i]&&dfs(i)) {cout << "1\n"; return 0;}
  }
  cout << "0\n";
  return 0;
}
