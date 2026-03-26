[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/heapx/release.yml?branch=main&label=BUILD)](https://github.com/yylego/heapx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/heapx)](https://pkg.go.dev/github.com/yylego/heapx)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/heapx/main.svg)](https://coveralls.io/github/yylego/heapx?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/heapx.svg)](https://github.com/yylego/heapx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/heapx)](https://goreportcard.com/report/github.com/yylego/heapx)

# heapx

Generic indexed min-heap with Fix/Remove support — no `heap.Interface` boilerplate needed.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Features

- **Generics**: Type-safe heap with custom comparison function, no `interface{}` casting
- **Indexed Nodes**: Each node tracks its heap position, enabling O(log n) Fix and Remove
- **Clean API**: Push returns a node handle — use it to Fix or Remove without searching
- **Zero Boilerplate**: No need to implement Len/Less/Swap/Push/Pop on custom types

## Installation

```bash
go get github.com/yylego/heapx
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/yylego/heapx"
)

func main() {
    // create a min-heap of ints
    h := heapx.New[int](func(a, b int) bool { return a < b })

    h.Push(30)
    n := h.Push(10) // keep the node handle
    h.Push(20)

    fmt.Println(h.Peek().Value) // 10

    // update value and fix heap position
    n.Value = 50
    h.Fix(n)

    fmt.Println(h.Pop().Value) // 20 (now smallest)
    fmt.Println(h.Pop().Value) // 30
    fmt.Println(h.Pop().Value) // 50
}
```

### Struct Example

```go
type Task struct {
    Name     string
    Deadline int64
}

h := heapx.New[Task](func(a, b Task) bool { return a.Deadline < b.Deadline })

n1 := h.Push(Task{Name: "task-a", Deadline: 100})
h.Push(Task{Name: "task-b", Deadline: 200})

// reschedule task-a
n1.Value.Deadline = 999
h.Fix(n1)

fmt.Println(h.Peek().Value.Name) // "task-b" (now earliest)
```

## API

| Method         | Returns    | Description                          |
| -------------- | ---------- | ------------------------------------ |
| `New[V](less)` | `*Heap[V]` | Create heap with comparison function |
| `Push(v)`      | `*Node[V]` | Add value, get node handle           |
| `Pop()`        | `*Node[V]` | Remove and return min node           |
| `Peek()`       | `*Node[V]` | Return min node without removing     |
| `Fix(node)`    | —          | Re-sort after changing `node.Value`  |
| `Remove(node)` | —          | Remove specific node                 |
| `Len()`        | `int`      | Number of nodes                      |
| `node.Index()` | `int`      | Node position (-1 if removed)        |

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->
