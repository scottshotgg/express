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

Below is currently the most advanced program that can be written in Express. You can find the **full** *uncommented* version labeled `advanced.expr` under the `samples` directory
<br>
You will find a few examples of the allowed flexibilities and optional verbosity that allow the language to be so _Expressive_
<br>
<br>
```
Let's Begin:
```
____

<br>
<br>

> Start off by declaring some variables:

```csharp
int powerlevel = 9000;
bool over9k = false;
float pi = 3.14159265359;
string arizona = "iced out boys";
```
<br>

> Comments also take on a familiar syntax:
```csharp
// Inline comment

/*
  Multiline
  Comment
*/
```
<br>

> Both of these are `int` variables, with the latter showing a *type inference* based on evaluation of the `rvalue` expression:

```csharp
int zero = 0
    one := 1    // tabbed for visibility
```
<br>

> `Optional Verbosity`: <br>
Before moving on, let's explain a primary motivator in the Express language development. Above, you will observe that the usage of commas and semicolons as statement delimiters seems to be optional. This is because they are! Statement delimiters are acceptable if you prefer, but are in no way required if you'd rather not deal with that. <br>
In the underlying parser architecture, they serve a semantic purpose by marking the end of a statement parse. By default, the ending of the statement will be semantically inferred if not *expressly* specified, however, there are compiler `flags` and [ECMA-335](https://www.ecma-international.org/publications/standards/Ecma-335.htm) `attributes` to modify the default action to enforce strict punctuation as granularly (or entirely) as you prefer.<br>
In this regard, having the flexibility to allow the compiler to semantically infer the end of the statement, while also retaining the ability to manually signal when a statement should end, can be very relaxing.<br>
This allowed flexibility is known as *`optional verbosity`* in Express and is one of the key motivators in it's development.

<br>

> Now, I wouldn't do this, but as a testament towards the semantic reasoning within the parser; you can even write statements on the same line as each other:
```csharp
string ay = "ayy" int timestamp = 1527799745 string waddup = "waddup"
```
<br>

> Usage of the `set` operator is permitted even outside of `object` declarations. 
<!--The difference from a standard declaration or assignment is that the `set` operator will enforce a *`non-destructive`* local declaration of a variable; meaning it will not crawl the scope tree to perform reassignment and will enforce that it is the  
crawl _down_ the scope tree instead of _up_? -->

```csharp
anotherOne: "anotherOne"
```
<br>

> Below shows a type inferred `object` where most of its properties are also type inferred.<br>
*Note:* The following statements until the ending brace are all within the `testerino` variable

```csharp
testerino := {
    id: "ba3d4793-cfae-48d1-ad51-47cbfd70f98a"
```
<br>

> Reference one of the above variables in a new property declaration

```csharp
    time: timestamp
```
<br>

> You can also use the `assignment` operator along with a type to crimp the type of the variable instead of leaving it up for interpretation by the compiler

```csharp
    float price = 55.3592,
```
<br>

> The `inference` operator can also be used within objects. Although currently it doesn't do anything different than the `set` operator within an object, it may have a more impactful use later

```csharp
    dank_meme := true
```
<br>

> **Any** unicode character is supported

```csharp
    🔥420🔥    : "🅱️laze it" 
    ⒹⒿKhalid : anotherOne
```
<br>

> An `array` composition using the above definitions of `zero` and `one` to derive a `type` and `length` inference

```csharp
    ten: [ one, zero, one, zero ]
```
<br>

> A few nested objects

```csharp
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

> Ending of `testerino`

```csharp
}
```

<br>

> Arrays can be declared as well; below is a `static` `string` type `array` using composition to infer a `length`.

```csharp
string[] stringArray = [ "hi", "my", "name", "is", "scott" ]
```
<br>

> Expanding on the above example:<br>
Delineatiation of elements from one another using commas isn't required but should be used at your own descretion of verbosity.<br>
The spacing also doesn't matter, but *readable* code does. <br>
Again - its all *semantics* `¯\_(ツ)_/¯`

```csharp
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

> Quick power level check before blasting off

```vb
if powerlevel < 9001 { 
   powerlevel = 9001 
   over9k = true 
}
```
<br>

> A simple for loop

```csharp
percent: 0
for progress := 0, progress < 100, progress++ { 
   percent = progress
}
```
