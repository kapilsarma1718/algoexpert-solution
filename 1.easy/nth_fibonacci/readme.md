nth fibonacci

- Problem : 

    tulis sebuah function yang menerima input:

    1. n integer

    dan mengembalikan bilangan fibonacci sesuai urutan n

- Solution :

**Summary**

Input :
- n integer, urutan n fibonacci

Output:
- bilangan fibonacci sesuai urutan n

Contoh:
- input : 6
- output : 5 (0,1,1,2,3,5)
- Urutan fibonacci berbeda dari yang biasanya:

[0, 1, 1, 2, 3, 5]
1   2  3  4  5  6

0 merupakan urutan pertama

**Solution 1** loop O(n)

```sh

nums = [0,1]
result = 0


if n <= 2 
    return nums[n-1]

ulangi i = 2, i < n, i++
    result = [i - 2] + [i - 1]
    data = append(data, result)

return result
```

**Solution 2** recusive O(2 ^ n)

if n == 1 
    return 0
if n == 2
    return 1
ret fib(n - 2) + fib(n - 1)