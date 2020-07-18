# Edge Auth System.


Define an edge authorization and authentication system.  I started this project while working on my master's in Cybersecurity at Syracuse University.
The Company that I worked for at the time was specialized in Point of Sale for Convenience stores. A common problem was how to remove 
the dependency on the Network/Internet but still take advantage of modern Cloud technology. Most Convenience stores operate 24 hours a day; 
any downtime at the POS can cause significant monetary loss, so I wanted to perform an OAuth like authorization but in a distributed manner at the site.

The backend is coded in GoLang. As far as unit test goes, I think I have pretty good code coverage, probably over 90%. 
The frontend I was a lot less careful. I have used this project as a tool to learn Flutter. I like that I can design the UI programmatically,
and I can get good results quickly. I think Google has done an outstanding job creating the Flutter framework. 
I got the same code base to run on Chrome, MacOS, Android, and even Linux and Windows (Even though those latter platforms are in the early stages of support).

The project is currently dependent on MongoDB. I may came back at a later time and add support for other databases. The port the GoLang Web Server listens on is also hardcoded (9119) and I need to add functionality to rotate the API Keys for machine to machine remote logins.  
