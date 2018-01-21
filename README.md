# Guess and Guesser

This is a sample set of Go scripts for handling command line prompts.

The first script is `guess`.
It is a minimal "guess the number" game.
The initial implementation doesn't even randomize the number.

The second script is `guesser`.
It takes the path to `guess` as its argument.
It runs `guess` and reads what `guess` prints.
It reads the input, adjusts its guesses appropriately, and "types" them into guess.
