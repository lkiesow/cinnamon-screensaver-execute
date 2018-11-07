Cinnamon Screensaver Execute
============================


A simple tool to execute commands when Cinnamon's screensaver engages and locks
the session. You can use this for example to adjust the lighting of peripheral
devices (e.g. turn your keyboard LEDs red) when you system is locked.


Build the Software
------------------

```sh
go build cinnamon-screensaver-execute.go
```


Configuration
-------------

For the configuration, create a JSON file in your home directory and put the
command to execute in as array for the keys `locked` and `unlocked`. Take a
look at the example configuration provided in this repository.

```sh
cp example-dot-cinnamon-screensaver-execute.json ~/.cinnamon-screensaver-execute.json
```


Run Service
-----------

To run and test it simply execute the built executable:

```sh
./cinnamon-screensaver-execute
```

Add this to Cinnamon's startup applications if you want it to be run
automatically whenever you use Cinnamon.
