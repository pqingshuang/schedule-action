# schedule-action

## To do

- [ ] init config file record worker situation is active or not
- [ ] when exit go routine, change config file from active to deactive
- [ ] mqtt when condition is active, recieve running signal by chann 
- [ ] when condition is not active ,exit by chann
- [ ] setting.initial check file every minute, status shows on/off status in setting
- [ ] activate shows current situation is run or not
- [ ] when it is on but not active, send to chann to invoke new gorounte
- [ ] when goroutine exit because of error, change file from active/2, 0 is deactivate by user, 1 is active and 
run, 2 is active but exit because of error
- [ ] not long connection, if based on interval, run and reocrd error and exit