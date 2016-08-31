#!/bin/bash
VMs=$(VBoxManage list runningvms | awk '{print $2}' | sed 's/[{,}]//g')
for uuid in $VMs 
do
VBoxManage controlvm $uuid poweroff
VBoxManage unregistervm $uuid
done
