### simple sqlite example
First execution do a 
```
 ./dbexample -flush
```

After it won't try to create again the file 

```
➜  sqlite git:(master) ✗ ./dbexample -flush
2022/03/02 21:04:45 Create users
[{1 bot pass} {2 droid pass} {3 trolls pass}]
➜  sqlite git:(master) ✗ ./dbexample
[{1 bot pass} {4 bot pass} {2 droid pass} {5 droid pass} {3 trolls pass} {6 trolls pass}]
➜  sqlite git:(master) ✗ ./dbexample
[{1 bot pass} {4 bot pass} {7 bot pass} {2 droid pass} {5 droid pass} {8 droid pass} {3 trolls pass} {6 trolls pass} {9 trolls pass}]
➜  sqlite git:(master) ✗ ./dbexample
[{1 bot pass} {4 bot pass} {7 bot pass} {10 bot pass} {2 droid pass} {5 droid pass} {8 droid pass} {11 droid pass} {3 trolls pass} {6 trolls pass} {9 trolls pass} {12 trolls pass}]
➜  sqlite git:(master) ✗ ./dbexample
[{1 bot pass} {4 bot pass} {7 bot pass} {10 bot pass} {13 bot pass} {2 droid pass} {5 droid pass} {8 droid pass} {11 droid pass} {14 droid pass} {3 trolls pass} {6 trolls pass} {9 trolls pass} {12 trolls pass} {15 trolls pass}]
➜  sqlite git:(master) ✗ ./dbexample
[{1 bot pass} {4 bot pass} {7 bot pass} {10 bot pass} {13 bot pass} {16 bot pass} {2 droid pass} {5 droid pass} {8 droid pass} {11 droid pass} {14 droid pass} {17 droid pass} {3 trolls pass} {6 trolls pass} {9 trolls pass} {12 trolls pass} {15 trolls pass} {18 trolls pass}]
```