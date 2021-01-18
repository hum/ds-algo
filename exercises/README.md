# Exercises

This is a list of short exercises for reference and a bit of problem solving. Feel free to check the solutions as well if you get stuck.

<details>
<summary>Exercise 1 (reversing a string)</summary>
Reverse a given string.

In:
```go
s = "Hello, my name is..."
```
Out:
```go
result = "...si eman ym ,olleH"
```
</details>
<details>
<summary>Exercise 2 (array merge)</summary>
Given 2 sorted arrays (slices), merge them into eachother while keeping the final array sorted.

In:
```go
array1 = [0, 3, 4, 31]
array2 = [4, 6, 30]
```
Out:
```go
result = [0, 3, 4, 4, 6, 30, 31]
```
</details>
<details>
<summary>Exercise 3 (first recurring number)</summary>
Given a non-sorted array, find the first recurring number. If no such pair exists, return an error.

In:
```go
array = [2, 5, 1, 2, 3, 5, 1, 2, 4]
```
Out:
```go
2
```
</details>
<details>
<summary>Exercise 4 (reverse linked list)</summary>
Given a Singly Linked List. Reverse the order of the nodes.

In:
```go
linkedList = [1, 10, 16, 88]
```
Out:
```go
linkedList = [88, 16, 10, 1]
```
[Solution](https://github.com/hum/ds-algo/blob/27338caee1a9fcd45cc63f6d2f5b5e429c94ccd0/ds/linkedlist.go#L91)
</details>
<details>
<summary>Exercise 5 (queue implementation with stack)</summary>
Implement a Queue using Stack.

In:
```go
MyQueue queue = new MyQueue()

queue.push(1)
queue.push(2)
print(queue.peek())
print(queue.pop())
print(queue.empty())
```
Out:
```go
1
1
false
```
</details>
