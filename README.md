# Learn_Go

A Tour Of Go Exercise を実施した記録です。

Exercise をフォルダごとに格納しています。

## Exercise: Loops and Functions

### Link

[Exercise: Loops and Functions](https://go-tour-jp.appspot.com/flowcontrol/8)

### Usage

```
cd learn_Go/ex_loop_and_functions/
```

引数なしのデフォルトは 2 としている.

```
go run loop_and_functions.go -num=3
```

### OutPut

```
0回目：2
1回目：1.75
2回目：1.7321428571428572
3回目：1.7320508100147276
4回目：1.7320508075688772
own sqrt: 1.7320508075688772
math pkg sqrt: 1.7320508075688772
```

## Exercise: Slices

### Link

[Exercise: Slices](https://go-tour-jp.appspot.com/moretypes/18)

### Usage

```
cd learn_Go/ex_slices
```

```
go run slices.go
```

### OutPut

<img src="images/exercise_slices.png" width="200px">

## Exercise: Maps

### Link

[Exercise: Maps](https://go-tour-jp.appspot.com/moretypes/23)

### Usage

```
cd learn_Go/ex_maps/
```

```
go run maps.go
```

### OutPut

```
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}
```

## Exercise: Fibonacci closure

### Link

[Exercise: Fibonacci closure](https://go-tour-jp.appspot.com/moretypes/26)

### Usage

```
cd learn_Go/ex_fibonacci_closure/
```

```
go run fibonacci_closure.go
```

### OutPut

```
1
1
2
3
5
8
13
21
34
55
```

## Exercise: Stringers

### Link

[Exercise: Stringers](https://go-tour-jp.appspot.com/methods/18)

### Usage

```
cd learn_Go/ex_stringers/
```

```
go run stringers.go
```

### OutPut

```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

## Exercise: Errors

### Link

[Exercise: Errors](https://go-tour-jp.appspot.com/methods/20)

### Usage

```
cd learn_Go/ex_errors/
```

```
go run errors.go -num=3
```

### OutPut

```
0回目：2
1回目：1.75
2回目：1.7321428571428572
3回目：1.7320508100147276
4回目：1.7320508075688772
1.7320508075688772 <nil>
----
-3 cannot Sqrt negative number: -3
```
