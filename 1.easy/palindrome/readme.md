palindrome

Problem:
- Tulis suatu function yang dapat mengecek palindrome
- palindrome merupakan 

// Solution 1
```
tmpStr = ""

Ulangi dari i = 0, i < len(str), i++
    tmpStr = str[i] + tmpStr

if str != tmpStr
    return false

return true
```
// Solution 2

```
middle = len(str) / 2

ulangi dari i = 0, i < middle, i++
    left = str[i]
    right = str[len(str) - 1 - i]

    jika left != right
        return false

return true
```
