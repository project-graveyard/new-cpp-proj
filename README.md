# new-cpp-proj: Simplify Your C++ Project Creation

Generate a C++ project with Makefile

## Usage

- [new-cpp-proj](/docs/new-cpp-proj.md)
- [new-cpp-proj init](/docs/new-cpp-proj.md)

## Template Information

### C++ Template

```cpp
include <iostream>
using namespace std;

int main() {
    cout << "Hello, World!" << endl;
    return 0;
}
```

### Makefile Template

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
