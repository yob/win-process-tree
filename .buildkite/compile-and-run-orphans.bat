go build -o win-process-tree.exe .

echo This will create some orphan processes and then exit successfully.
echo Desired behaviour: the agent should record the build as sucessful and temrinate the orphan processes

win-process-tree.exe orphans
