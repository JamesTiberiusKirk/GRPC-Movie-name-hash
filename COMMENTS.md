# Comments

- Unit testing the Echo API handler was annoying & messy, shoud figure out a different
unit testing strategy
- I'm was not certain how I should have initialised the grpcClinet
  - If I should be doing it in the handler or outside the handler and share it across
  the entire app like I have done
- Since these microservices share a go.mod file, I needed to pass a different context to 
the docker build (reffer to the docker-compose) and I was not sure weather this is a
common convention or weather I should have done something else.
- Something else which I know is common in the industry but I personally dislike it
(hence why I did not follow that convention here) is to separate the entire logic of the 
controller out from the handler function.
  - I can understand the reason why this is usually done.
  - However, in my experience it creates to much *bloat*.
  - IMO, this should be abstracted on a case to case basis 
    - E.G. where the logic might actually be reused somewhere else) and not by default
    on every single handler.

P.S. I know convention is that we do not commit `.env`, but theres no actual secrets in
in there, and for running locally, its needed.  
P.S.S Sorry, on projects where I work alone I dont typically follow propper git etiquate,
So no git history.  
