# Hangman

* [The base](#the-base)
    * [Description](#description)
    * [Game rules](#game-rules)
    * [Tips](#tips)
    * [Game features](#game-features)
    * [Game objective](#game-objective)
* [Required](#required)
    * [Golang](#golang)
* [Run Locally](#run-locally)

## The base

### Description:
Le Jeu du Pendu is a classic and entertaining game that will put your spelling and deduction skills to the test. You'll have to guess the secret word by guessing one letter at a time, or all at once, while avoiding making too many mistakes.

### Game rules:
1. A secret word is chosen at random by the computer. This word is displayed in dashes, one dash per letter, to indicate its length.

2. The player can guess one letter at a time by typing on the keyboard.

3. If the guessed letter is in the secret word, it is revealed in its place. If the letter appears more than once in the word, all occurrences are revealed.

4. If the guessed letter is not in the secret word, an element of the hangman is drawn and you lose 1 life. The game ends when the hangman is complete (within 10 errors).

5. The player can guess the whole word by typing on the keyboard.

6. If the guessed word is the secret word, it is revealed. The game is won.

7. If the guessed word is not the secret word, you lose 2 lives instead of one.

8. The player wins if he correctly guesses the secret word before the hangman is finished. He loses if he reaches the maximum number of errors.

### Tips:
- Start by guessing common vowels, as these are more likely to appear in the secret word.
- Use the frequency of letters in your language to guide your choices.
- If you're not sure about a letter, try to guess letters that might form syllables or common words.

### Game features:
- Choose between word list and display type in the command that calls the function.
- You can stop a game and resume it later by storing it locally.
- Keep an eye on the dashboard to track the number of errors and letters already guessed and tried.
- Play solo or with friends, taking turns to guess the secret word.

### Game objective:
The aim of the game is to guess the secret word correctly before the hangman is completely drawn. You can set your own personal goals, such as trying to guess the word with as few mistakes as possible.

**Have fun playing Hangman and test your spelling and deduction skills!


## Required

### Golang

Required version : `go1.21.0`

[Documentation of GOLANG](https://go.dev/doc/)

## Run Locally

**Clone the project**

```bash
git clone https://ytrack.learn.ynov.com/git/gbenjamin/hangman.git
```

**Go to the project directory**

```bash
cd hangman
```

**Start game without ascii display**

```bash
go run main/main.go [word list file name] 
```

**Start game with ascii display**

```bash
go run main/main.go [word list file name] [ascii display file name]
```

**To stop the game** 

When you can enter a letter or a word, 
enter ``` STOP ```
then follow the instructions

**To take over a part**

```bash
go run main/main.go --startWith [save file path]
```

**You can also consult help with**
```bash
go run main/main.go --help
```
or 
```bash
go run main/main.go
```