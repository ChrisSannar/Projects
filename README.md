# Projects
A collection of various things that I've learned, in various languages, all made by yours truly.

## Easy Poll

Goal: Make a simple website that people can make a quick poll on a few topics

### Features

Poll Creator:
* Create a poll - Either by a name (case sensetive?), or picks a random name
  * Check if the name already exists.
* Set up questions (set built-in limits  like 10 questions, 10 answers) - Triggers session creation: Users can now join
  * Question title
  * Choices
  * List number of users?
  * Load questions via file?
  * Button to create
* Start
* Alternate between the two:
  * Awaiting responses - Option to close
  * Display results (Graphs? Values? etc.)
* At the end is the option to make another
  * Save responses via CSV
  * Save Poll via file

Poll User:
* Join a poll - By a name
  * Check if poll exists
  * Include a name? Optional?
* Once joined, It will wait till the quiz starts
* Question opens:
  * They are given options
  * Once selected, they can change till they submit (disable everything)
* When the Creator closes the poll, display the results to everyone
  * Opens next question when Creator changes
* Once poll is over, asks if you would like to try another one

Engineering:
* Interactivity with the server - Use socket.io to get values from user to admin.
  * React socket io library?
* Use `rsuite` for components
* Rxjs to handle state managemenet
* No database needed, unless we want to store statistics (which we don't really care about)

Display:
* Allow to hide or show results

### Concerns

How is this different from google forms?
* Realtime feedback

How about security, how do we know people won't just try an join?
* Pick more obscure word for poll name
