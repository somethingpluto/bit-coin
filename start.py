from os import system,popen


for i in range(1,100001):
    cmdstr = f"bitcoin.ext addblock -data '{i}'"
    print(cmdstr)
    popen(cmdstr)