Screen command

Refer : https://linuxize.com/post/how-to-use-linux-screen/

Why ? : Situation where you are performing a long running task on remote machine and it somehow gets disconnected and work is lost.

GNU screen is a terminal multiplexer, it means you can start a screen session any then open any number of windows inside that session.

Check version
$ screen --version

$ screen : This will open a screen session, create a new window, and start a shell in that window.
           
To get a list of all commands : Ctrl + a + ?

Starting a named session
$ screen -S session_name

Working with Linux screen windows
You can have multiple screens inside one session.

Ctrl + a c        : Creates a new window with shell.
Ctrl + a "        : List all widows
Ctrl + a 0        : Switch to window 0(by number)
Ctrl + a A        : Rename the current window
Ctrl + a S        : Split the current window region horizontally
Ctrl + a |        : Split veritcally
Ctrl + a tab      : Swtich the input focus to next region
Ctrl + a Ctrl + a : Toggle between current and previous window
Ctrl + a Q        : Close all regions from the current one
Ctrl + a X        : Close the current region

Detach from the screen session
You can detach from the screen session at any time by : Ctrl + a d
The program running in the screen session will continue to run after you detach from the session.

To resume your screen session : $ screen -r

Incase you have multiple screen sessions running on your machine, you will need to append the screen session ID
after r switch

To list the screen ID's
$ screen -ls


Customize linux screen
When the screen is started, it read conf from /etc/screenrc and ~/.screenrc if file is present.
We can modify the default screen settings according to our preferences by modifying the screenrc file.

Basic linux screen usage :
1. screen
2. Run the desired program
3. Ctrl a + Ctrl d to detach from the screen session
4. Retach the screen session by screen -r
