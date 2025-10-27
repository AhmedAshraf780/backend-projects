# HTTP
- HTTP is a protocol that is used to transfer data over the internet
- It's just a couple lines of code listening on port 80 for requests
- It takes a request and parsing it with the protocol they use and anaylizing it
- It's a stateless protocol, each request doesn't know anything about the previous one

# HTTP Use Cases
- You open the browser and type youtube.com
- The browser takes this url and sends it to DNS Server which is another machine that has http code that takes requests
- The DNS sends the ip of this url back to the browsers
- the browser now has the following ( your ip, what you want , the distanation ip, etc) and it combines all of this info into the request
- The browser sends this request to the distanation on port 80
- The distanation has http code that listens on port 80 so he recieves this request from there
- The distanation has the code that parses the request and sends back the response
- The browser recieves the response and shows it to you

# HTTP Cons
- The Main issue is the data is not encrypted, it's literally just text
- hence, your request can be exposed to the public at the same network
- You are vulnerable to man in the middle attacks
- There's no way to verify the request is from the real source

## HTTPS
- HTTPS is a protocol that is used to transfer data over the internet like HTTP
- It's listening on port 443 for requests
- It requires a certificate to be signed by a trusted authority (TLS/SSL handshake)


