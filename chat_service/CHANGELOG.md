### 2022-05-22
* Switched the plain string communication to JSON for extendability. 
* Delete the member from the room after quit (but if the client just shutdown the connection without using quit command, the server can not close the connection gracefully)
* Add `/rooms -v` method for more detailed information on each rooms.
### 2022-05-20
* Add the top nav bar ui using quasar. 
* Fix the bug of stucking upon inputting empty string.
### 2022-05-18
* Integrated vue-terminal-ui. 
* Registered custom commands. 
* Implement the output supress and echo method for the messages triggered the server-side events. 
### 2022-05-16
* Switched from net to http server. 
* Upgraded http requests to websockets.
### 2022-05-15: 
* Initialized. yeahhh