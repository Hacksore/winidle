# winidle

Detect if windows is idle for a given amount of seconds and run a command


### Example

```
winidle.exe -t 300 -c "rundll32.exe powrprof.dll,SetSuspendState 0,1,0"
```