# yakstack

Whether I am programming, investigating an incident, or living my day-to-day live, my pea-sized, ADD brain gets distracted and I end up shaving the yak multiple times a day. I don't really have a problem with this, since its just how my brain works, but I *do* have a problem with forgetting the tasks I was doing earlier in the day. 

To deal with this, I have made a little bin util that now works as my stack-based todo list. As I go about an investigation and discover more and more leads I need to look into before making sense of the last, I can now push them on to the stack as I go along until I hit the piece of info I need to start popping tasks off of the stack.

yakstack only has 4 commands: 

**push \<task\>**:
pushes a task on to the stack

**peek**:
prints your current task

**pop**:
removes your current task from the stack, essentially saying "done!"

**list:**
lists your entire todo list from last-in to first-in
