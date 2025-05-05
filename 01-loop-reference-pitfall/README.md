# Loop Variable Reference Pitfall in Go

This example demonstrates a classic mistake when taking the address of a loop variable in Go using a `for range` loop.

## ❌ The Buggy Code

Using `&book` inside a `for _, book := range` loop causes all pointers to point to the same memory address.

## ✅ The Fix

Use `for i := range slice` and take the address of `slice[i]`.

## 📁 Files

- `main.go`: Contains both buggy and fixed versions with explanations.
