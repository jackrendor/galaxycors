# GalaxyCors
![](icon.png)

GalaxyCors is a tool to determinate if http endpoints have basic CORS Misconfiguration.

## Usage
```code
  -cookie string
        Cookies to use during requests. Example: co1=val;id=000.
  -data string
        Data to be sent in the body.
  -header string
        Headers to use during requests. head1: data; head2: data2. (default "User-Agent: GalaxyCors 0.1")
  -method string
        Method to use during requests. (default "GET")
  -timeout int
        Timeout for connection. (default 10)
  -url string
        Url to the target.
```

## Thanks
A sincere thank to my girlfriend and partner in crime.
She was telling me about Alexa and some of the vulnerable subdomains owned by Amazon, and she mentioned CORS. She understood the vulnerability, and allowed me to explain this techy and boring stuff. :')