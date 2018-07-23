for i := 0, i < 10, i ++ {
  
}

for i in [ 0, 2, 4, 6, 8 ] {
  i = key
}

for i of [ 0, 2, 4, 6, 8 ] {
  i = value
}

for i over [ 0, 2, 4, 6, 8 ] {
  i = { key, value }
}

for k,v over [ 0, 2, 4, 6, 8 ] {
  k = key
  v = value
}

// async
over [ 0, 2, 4, 6, 8 ] {
  key
  value
}

// async
over 1..10 {
  value
}

from 1 to 10 {
  value
}

for 1..10 {
  value
}

for i := 1..10 {
  i = value
}

for k,v := 1..10 {
  k = key
  v = value
}

for k,v over 1..10 {
  key
  value
}

while(i < 10) {

}

do {

} while(i < 10)

// All numbers from 1 to 100
i := [ 1, 2, ..100 ]
i := 1..100
i := from 1 to 100
i := numbers until 100
// Don't think I'm going to allow this
// i := until 1000
i := 1 until 100

i := [ 2, 4, ..100 ]
i := 2,4, ..100
i := 2,4 to 100
i := evens until 100
i := 2,4 until 100