
int testFunction(){
  const y = 20;
  print('$x, $y'); 
  return x + y;
}

main() {
  const x = 10;
  print('$y'); 
  int testFunction(){
    const y = 20;
    print('$x, $y');
     return x + y;
  }
  var result = testFunction();
  print('$result');

 
  int oneline(a,b) => a + b;
 
  int oneline(a,b){ return a + b }
  print(oneline(1,2));

  
  void enableFlags({bool bold, bool hidden}) { print('$bold $hidden'); }
  enableFlags(hidden: true);

  
  String say(String from, String msg, [String device = 'unknown', String mood]) {
   
    return '$from says $msg platform: ${device} mood: ${mood ?? 'unknown'}';
  }
}
