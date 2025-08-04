# k8s-testApplicationA

## 动态版本获取

### 方案1: 从GitHub Tags获取 (推荐)
![Helm Chart](https://img.shields.io/github/v/tag/sander-1105/k8s-testApplicationA?label=Helm%20Chart&color=blue&style=flat-square)

### 方案2: 从GitHub Releases获取
![Helm Chart](https://img.shields.io/github/v/release/sander-1105/k8s-testApplicationA?label=Helm%20Chart&color=blue&style=flat-square)

### 方案3: 使用shields.io动态获取
![Helm Chart](https://img.shields.io/badge/dynamic/json?url=https://api.github.com/repos/sander-1105/k8s-testApplicationA/releases/latest&query=$.tag_name&label=Helm%20Chart&color=blue&style=flat-square)

### 方案4: 从Chart.yaml获取 (静态)
![Helm Chart](https://img.shields.io/badge/Helm%20Chart-v0.45.0-blue.svg?style=flat-square)

## 使用方法

1. 创建新的tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. 版本会自动更新在README中显示

## 测试

当前仓库的tag: ![GitHub tag](https://img.shields.io/github/v/tag/sander-1105/k8s-testApplicationA)
