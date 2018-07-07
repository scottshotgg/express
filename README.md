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

<code>
int powerlevel = 9000;
<br>
bool over9k = false;
<br>
float pi = 3.14159265359;
<br>
string arizona = "iced out boys";
</code>

<br>
<br>

> Both of these are `int` variables, with the latter showing a *type inference* based on evaluation of the `rvalue` expression


<code>
int zero = 0
<br>
one := 1
</code>

<br>
<br>

> You may also use the `set` operator even outside of `object`s

`anotherOne: "anotherOne"`
<br>
<br>

> `Optional Verbosity`: 
<br>
You will observe that the usage of commas and semicolons as statement delimiters is acceptable if you prefer, but are in no way required if you'd rather not deal with that. 
<br>
In the underlying parser architecture, they serve a semantic purpose by marking the end of a statement parse. By default, the ending of the statement will be semantically inferred if not *expressly* specified, however, there are compiler `flags` and [ECMA-335](https://www.ecma-international.org/publications/standards/Ecma-335.htm) `attributes` to modify the default action as granularly (or entirely) as you prefer to enforce strict punctuation.
<br>
In this regard, having the flexibility to allow the compiler to semantically infer the end of the statement, while also retaining the ability to manually signal when a statement should end, can be very relaxing.
<br>
This allowed flexibility is known as *`optional verbosity`* in Express and is one of the key motivators in the development.
 
<br>

> Below shows a type inferred `object` where most properties are also type inferred:

<code>
  testerinoObject := {
<br>
&nbsp;&nbsp;&nbsp;&nbsp;
  id: "ba3d4793-cfae-48d1-ad51-47cbfd70f98a" 
<br>
</code> 
<br>
  
> Reference one of the above variables

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  time: timestamp
</code>
  <br>
  <br>

> The `assignment` operator with a type can also be used within objects

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  float price = 55.3592,
</code>

<br>
<br>

> Also the `inference` operator can also be used

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  dank_meme := true
</code>
<br>
<br>

> **Any** unicode character is supported

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  🔥420🔥: "blaze it" 
<br>
&nbsp;&nbsp;&nbsp;&nbsp;
  deeJayKhalid: anotherOne
</code>

<br>
<br>

> Array composition (`type` and `length` inference) using the above definitions of `zero` and `one`.

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  ten: [ one, zero, one, zero ]
</code>

<br>
<br>

> A few nested objects

<code>
&nbsp;&nbsp;&nbsp;&nbsp;
  throw_more_dots: {
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    throw_more_dots: {
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      more_dots: {
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        more_dots: {
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          ok: "stop dots",
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        },
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      },
<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    },
<br>
&nbsp;&nbsp;&nbsp;&nbsp;
  }
</code>

<br>
<br>

> End of `testerinoObject`

<code>
}
</code>

<br>
<br>

> Arrays can be declared as well; this is a `string` type `array` using composition to infer a `static` length.

<code>
string[] stringArray = [ "hi", "my", "name", "is", "scott" ]
</code>

<br>
<br>

> Expanding on the above example:
<br>
Delineatiation of elements from one another using commas isn't required but should be used at your own descretion of verbosity.
<br>
The spacing also doesn't matter, but *readable* code does. 
<br>
Again - its all *semantics* `¯\_(ツ)_/¯`

<code> 
string ay = "ayy"
<br>
string waddup = "waddup" 
<br>
string[] hereHeComes = 
<br>
[
<br>
&nbsp;&nbsp;
ayy
<br>
&nbsp;
&nbsp;
&nbsp;&nbsp;
waddup
<br>
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
"its" 
<br>
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
"dat"
<br>
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
"boi" 
<br>
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;
&nbsp;&nbsp;
]
</code>

<br>
<br>

> Quick power level check

<code>
if powerlevel < 9001 {
<br>
&nbsp;&nbsp;
powerlevel = 9001
<br>
&nbsp;&nbsp;
over9k = true
<br>
}
</code>

<br>
<br>

> A simple for loop

<code>
for j := one, j < 10, j++ {
<br>
&nbsp;&nbsp;
zero = 1
<br>
}
</code>
