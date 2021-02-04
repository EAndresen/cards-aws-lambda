# Cards - In AWS Lambda

Ready to go app and to be deployed in AWS Lambda.

To Run:

| For PC  |
| ------------- |
| set GOOS=linux  |
| go build -o main main.go deck.go game.go player.go |
| %USERPROFILE%\Go\bin\build-lambda-zip.exe -output main.zip main |



## Battle of the Beasts!
In a time before ages there were two giants colliding in two different parts of the world.
But one day they met and can't stand the others' greatness!

So the only solution for this is to get a deck of card each, shuffle them around and deal a hand of 7!
Calculate the value of the hand and se who, once and for all, is the number one.

Unless there is a tie then no one wins and both go home.
