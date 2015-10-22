# Playing Cards

[GoDoc](https://godoc.org/github.com/stensonb/playingcards)

# Usage

Import the package (note, I'm importing math/rand and time in order to supply Shuffle() with a seed):
```
import (
  "github.com/stensonb/playingcards/standarddeck"
  "math/rand"
  "time"
)
```

Now, we can use the library:
```
// get a new standard deck (all values in all four suites)
d := standarddeck.New()

// shuffle them, providing the rand object
d.Shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))

// deal 5 cards
for i := 0; i < 5; i++ {
  c, _ := d.NextCard()
  fmt.Println(c)
}
```
