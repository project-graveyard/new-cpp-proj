# new-cpp-proj

Generate a C++ project with Makefile

### main.cpp template

```cpp
inclide <iostream>
using namespace std;

int main() {
    cout << "Hello, World!" << endl;
    return 0;
}
```

### Makefile template

```makefile
all: run clean

build:
  @g++ main.cpp -o main.out

run: build
  @./main.out

clean:
  @rm *.out

debug:
  g++ -g main.cpp -o main.out
  gdb main.out
```
