# Task

## Implement a function which receives two arguments:

1. the sum of digits (sum)
2. the number of digits (count)

## This function should return three values:

1. the total number of values which have count digits that add up to sum and are in increasing order
2. the lowest such value
3. the greatest such value

Note: if there're no values which satisfy these constaints, you should return an empty value.

# Examples

```
find_all(10, 3)  =>  [8, 118, 334]
find_all(27, 3)  =>  [1, 999, 999]
find_all(84, 4)  =>  []
```
