go build -o win-process-tree.exe .

echo This will create some orphan processes and then exit successfully.
echo Desired behaviour: the agent should record the build as sucessful and temrinate the orphan processes within 2 minutes
echo If the build takes longer than that, it'll be cancelled. The child processes have tricked the agent into thinking the job is still running

win-process-tree.exe orphans
