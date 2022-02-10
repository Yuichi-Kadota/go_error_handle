# go_error_handle
goのエラー設計、ハンドリングあれこれ（備忘）
## 前提
```
以下のような層になっている場合に、各層のErrorをWrapしていく

usecase
    ↓
service
    ↓
repository
```