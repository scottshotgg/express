<!--
Try using this later if you wanna do something
<span style="color:#6ba1f9">
hey *its* me
</span> 
-->

# Express-rearch
An overhaul of Express, from the ground up
<br>
<br>

## Example Program

Below is the most advanced program that can be currently written
in Express. You can find the **full** *uncommented* version under the `samples` directory labeled `advanced.expr`.
<br>
You will find examples of the many flexibilities and 
optional verbosity that allow the language to be so _Expressive_
<br>
<br>
```
Let's Begin:
```
____

<br>
<br>

> Start off by declaring some variables

```c
int powerlevel = 9000;
bool over9k = false;
float pi = 3.14159265359;
string arizona = "iced out boys";
```
<br>

> Both of these are `int` variables, with the latter showing a *type inference* based on evaluation of the `rvalue` expression


```c
int zero = 0
one := 1
```
<br>

> You may also use the `set` operator even outside of `object`s


```c
anotherOne: "anotherOne"
```
<br>

> `Optional Verbosity`: <br>
You will observe that the usage of commas and semicolons as statement delimiters is acceptable if you prefer, but are in no way required if you'd rather not deal with that. <br>
In the underlying parser architecture, they serve a semantic purpose by marking the end of a statement parse. By default, the ending of the statement will be semantically inferred if not *expressly* specified, however, there are compiler `flags` and [ECMA-335](https://www.ecma-international.org/publications/standards/Ecma-335.htm) `attributes` to modify the default action as granularly (or entirely) as you prefer to enforce strict punctuation.<br>
In this regard, having the flexibility to allow the compiler to semantically infer the end of the statement, while also retaining the ability to manually signal when a statement should end, can be very relaxing.<br>
This allowed flexibility is known as *`optional verbosity`* in Express and is one of the key motivators in the development.
 
<br>

> Below shows a type inferred `object` where most properties are also type inferred:

```c
testerinoObject := {
    id: "ba3d4793-cfae-48d1-ad51-47cbfd70f98a"`
```
<br>

> Reference one of the above variables

```c
    time: timestamp
```
<br>

> The `assignment` operator with a type can also be used within objects

```c
    float price = 55.3592,
```
<br>

> Also the `inference` operator can also be used

```c
    dank_meme := true
```
<br>

> **Any** unicode character is supported

```c
    🔥420🔥: "blaze it" 
    deeJayKhalid: anotherOne
```
<br>

> Array composition (`type` and `length` inference) using the above definitions of `zero` and `one`.

```c
    ten: [ one, zero, one, zero ]
```
<br>

> A few nested objects

```c
     throw_more_dots: { 
         throw_more_dots: { 
             more_dots: { 
                 more_dots: { 

                     ok: "stop dots", 
                     
                 }, 
             }, 
         }, 
     }
```
<br>

> Ending of `testerinoObject`

```c
}
```

<br>

> Arrays can be declared as well; this is a `string` type `array` using composition to infer a `static` length.

```c
string[] stringArray = [ "hi", "my", "name", "is", "scott" ]
```
<br>

> Expanding on the above example:
<br>
Delineatiation of elements from one another using commas isn't required but should be used at your own descretion of verbosity.
<br>
The spacing also doesn't matter, but *readable* code does. 
<br>
Again - its all *semantics* `¯\_(ツ)_/¯`


```c
string ay = "ayy" 
string waddup = "waddup" 
string[] here_comes = 
[
  ayy 
      waddup 
             "its", 
                    "dat" 
                          "boi" 
                                ]
```
<br>

> Quick power level check

```c
if powerlevel < 9001 { 
   powerlevel = 9001 
   over9k = true 
}
```
<br>

> A simple for loop

```c
for j := one, j < 10, j++ { 
   zero = 1 
}
```
