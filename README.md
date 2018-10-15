![Go-Emoji-Utils](http://up.tmdvs.me/52074bebc945/d)

# Go Emoji Utils
A collection of useful functions for working with emoji. For example: look up the definition of a specific emoji, or search for all occurrences of emojis in a string.

## Features
 - [x] Return a list of all emojis in a string
 - [ ] Find the location of a specific emoji in a string
 - [ ] Return the definition of a single emoji
 - [ ] Return the definitions of an array of emoji characters

## Examples
### Return a list of all emojis in a string
You can search a string for all occurrences of emoji. You will be returned an array of results specifying which emojis were found, and how many times each occurred.

```go
str := "This is a string 😄 🐷 with some 👍🏻🙈 emoji! 🐷 🏃🏿‍♂️"
emoji.DetectEmoji(str)

// Result: SearchResults[ SearchResult{ Match: Emoji{…}, Occurences: 1 }, …]
```
