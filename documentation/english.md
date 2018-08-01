# Proposal for English words within Express

Prepositions are currently used in for loops to make implicit key and value iterations much easier. A similar usage could be applied to boolean expressions and **other operations**.

For instance, the following expressions logically compose into the same reading:<br>
```cs
if i == 7 {

}
```
```cs
if i is 7 {

}
```
<br>

As well as these two:
```cs
if i < 7  {

}
```
```cs
if i is less than 7 {

}
```
<br>

For type equality comparison with `var` types:
```cs
var i = 6 // var is an int type
int j = 6

if i === 6 {

}

if i is the same as 6 {

}
```