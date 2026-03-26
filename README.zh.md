[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/heapx/release.yml?branch=main&label=BUILD)](https://github.com/yylego/heapx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/heapx)](https://pkg.go.dev/github.com/yylego/heapx)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/heapx/main.svg)](https://coveralls.io/github/yylego/heapx?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/heapx.svg)](https://github.com/yylego/heapx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/heapx)](https://goreportcard.com/report/github.com/yylego/heapx)

# heapx

泛型索引最小堆，支持 Fix/Remove 操作 — 无需手写 `heap.Interface` 样板代码。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 功能特性

- **泛型支持**：类型安全的堆，自定义比较函数，无需 `interface{}` 类型断言
- **索引节点**：每个节点维护堆中位置，支持 O(log n) 的 Fix 和 Remove
- **简洁 API**：Push 返回节点句柄 — 用它来 Fix 或 Remove，无需搜索
- **零样板代码**：无需在自定义类型上实现 Len/Less/Swap/Push/Pop

## 安装

```bash
go get github.com/yylego/heapx
```

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/yylego/heapx"
)

func main() {
    // 创建 int 最小堆
    h := heapx.New[int](func(a, b int) bool { return a < b })

    h.Push(30)
    n := h.Push(10) // 保存节点句柄
    h.Push(20)

    fmt.Println(h.Peek().Value) // 10

    // 更新值并修复堆位置
    n.Value = 50
    h.Fix(n)

    fmt.Println(h.Pop().Value) // 20（现在最小）
    fmt.Println(h.Pop().Value) // 30
    fmt.Println(h.Pop().Value) // 50
}
```

### 结构体示例

```go
type Task struct {
    Name     string
    Deadline int64
}

h := heapx.New[Task](func(a, b Task) bool { return a.Deadline < b.Deadline })

n1 := h.Push(Task{Name: "task-a", Deadline: 100})
h.Push(Task{Name: "task-b", Deadline: 200})

// 重新调度 task-a
n1.Value.Deadline = 999
h.Fix(n1)

fmt.Println(h.Peek().Value.Name) // "task-b"（现在最早）
```

## 接口

| 方法           | 返回值     | 说明                         |
| -------------- | ---------- | ---------------------------- |
| `New[V](less)` | `*Heap[V]` | 使用比较函数创建堆           |
| `Push(v)`      | `*Node[V]` | 添加值，获取节点句柄         |
| `Pop()`        | `*Node[V]` | 移除并返回最小节点           |
| `Peek()`       | `*Node[V]` | 返回最小节点但不移除         |
| `Fix(node)`    | —          | 修改 `node.Value` 后重新排序 |
| `Remove(node)` | —          | 移除指定节点                 |
| `Len()`        | `int`      | 节点数量                     |
| `node.Index()` | `int`      | 节点位置（-1 表示已移除）    |

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
