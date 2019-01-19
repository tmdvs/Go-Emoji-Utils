![Go-Emoji-Utils](http://up.tmdvs.me/52074bebc945/d)

# Go Emoji Utils
A collection of useful functions for working with emoji. For example: look up the definition of a specific emoji, or search for all occurrences of emojis in a string.

## Features
 - [x] Find all occurrences of emoji in a string
 - [x] Search a string for the presence specific emoji
 - [x] Look up the definition of a single emoji
 - [x] Look up the definitions for a list of emojis
 - [x] Remove all emoji from a string
 - [x] Import tool to update Emoji data with [Emojipedia](http://emojipedia.org/) specs
 - [x] Find the location of and occurrences of a specific emoji in a string

## Examples
### Find all occurrences of emoji in a string
You can search a string for all occurrences of emoji. You will be returned an array of results specifying which emojis were found, and how many times each occurred. The `Locations` property comprises of an array containing the start and end locations of each occurance of the emoji within the string you're searching.

```go
input := "This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️"
result := emoji.FindAll(input)

// result: SearchResults{ SearchResult{ Match: Emoji{…}, Occurrences: 1, Locations: […] }, …}
```

### Search a string for the presence specific emoji
You can search a string for the presence of a specific emoji. You will be returned a `SearchResult` struct with the definition of the matching emoji, how many times it occurred in the string, and its location within the string. 

```go
input := "This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️"
result := emoji.Find("🐷", input)

// result: SearchResult{ Match: Emoji{ Key:"1F437", Value:"🐷", Descriptor: "pig" }, Occurrences: 2, Locations: [[19 19] [42 42]  } }
```

### Checking search results for the occurrence of a specific emoji
The `SearchResults` struct has an `IndexOf` method for conveniently checking whether or not a specific emoji appears within the results set. If the emoji is not found a position of `-1` is returned.

```go
input := "This is a string 😄 🐷"
results := emoji.FindAll(input)

pigEmoji := emoji.Emojis["1F437"]
pigIndex := results.IndexOf(pigEmoji)
// pigIndex: 1
```

### Look up the definition of a single emoji
You can lookup the definition of a single emoji character using the `LookupEmoji` function. If a matching emoji can't be found an error will be returned.

```go
result, err := emoji.LookupEmoji("🐷")

// result: Emoji{ Key:"1F437", Value:"🐷", Descriptor: "pig" }
```

### Look up the definitions for a list of emojis
You can lookup the definition of a for each emoji character in an array of emojis using the `LookupEmojis` function. This function is a convenience function that loops through the input slice and passes each emoji to the `LookupEmoji` function.

Results are returned in the same order that the input strings were provided in. If a matching emoji can't be found an error will be appear in for that position in the result slice.

```go
result := emoji.LookupEmojis([]string{"🐷", "🙈"})

// result: []interface{}{ Emoji{ Key:"1F437", Value:"🐷", Descriptor: "pig" }, …}
```

### Remove all emoji from a string
You can remove all the emoji characters from a string. This function finds all occurences of emoji using the `FindAll` function, and uses the `Location` field to remove runes at those indexes.

```go
input := "This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️ 🥰"
output := emoji.RemoveAll(input)

// output: "This is a string with some emoji!"
```

### Updating the Emoji definitions data file
Definitions of each emoji can be found in the `data/emoji.json` file. This JSON file is unmarshalled into a map at runtime with the emoji's hex representation as it's key. You can update the definitions file using definitions from [Emojipedia](http://emojipedia.org/) with the `import` program bundled with this package. You can find it at `import/main.go`.