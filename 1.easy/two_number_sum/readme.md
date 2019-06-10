two number sum

- Problem:

    Tulis sebuah function yang menerima input :
    
    1. array integer 
    2. integer hasil penjumlahan yang akan dicapai

    Jika ada 2 value integer dalam array yang dijumlahkan sama dengan nilai target, akan mengembalikan array dengan 2 value integer tersebut dalam urutan dari kecil ke besar.

    Jika tidak ada yang dapat mencapai target, kembalikan empty array

- Solution:

**Summary**

Input : 
- array integer
- target value

Output :
- jika 2 value integer dijumlahkan mencapai target, kembalikan array dengan 2 value tersebut, urutan dari kecil ke besar
- jika tidak ada yang mencapai target, kembalikan empty array

**Solution 1** Brute force O(n ^ 2)

```
// twoNumberSum
result = []
lenArr = len(arr)
ulangi dari i = 0, i < lenArr, i ++
    ulangi dari j = i + 1, j < lenArr, j++
        total = arr[i] + arr[j]
        jika total = target
            minValue = min(arr[i],arr[j])
            maxValue = max(arr[i], arr[j])
            result = append(result, minValue, maxValue)
return result

// max function
max(x,y int) int

jika x > y 
    return x
return y

// min function
min(x,y int) int

jika x < y
    return x
return y

```

**Solution 2** Pakai map O(n)

```
map = {}
result = []
ulangi dari i = 0, i < len(array), i++
    num = array[i]
    n2 = target - num

    jika map[n2] ada
        minValue = min(n2, num)
        maxValue = max(n2, num)
        result = append(result, minValue, maxValue)
        return result
    
    map[num] = n2

return result
    
// max function
max(x,y int) int

jika x > y 
    return x
return y

// min function
min(x,y int) int

jika x < y
    return x
return y

```