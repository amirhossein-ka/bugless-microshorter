# bugless-microshorter
A simple micro service written in go, to make shorter urls !
<small><small><small>or maybe bugfull xD</small></small></small>

---


## all i know about this:

- Comunication between services is done with gRPC. *
- No redis or external library for caching
- Basically there is no authentication
- Maybe (or certainly) i should have one binary (or executably file) for each service
- Api gateway must behave like this:
    - Wait 50ms to send all requests recieved to url shortener service
    - **Or** when there is total 100 requests, send them to other sevice
    - It must send all those recieved requests with single request to url shortener service



\* i want to use mostly golangs internal libraries, so i probably use net/rpc package


#### projects structure
- `cmd/`
    - `cmd/gateway`
    - `cmd/shortner`
- `pkg/`
    -  maybe for logging or swagger ?
- `internal/`
    - this directory may contain common packages for `gateway` and `shortener`
    - cache, rpc stuff, etc...
    - `config/`
        - to config both gateway and shortener, ofc

- `internal/gateway`
    - clearly, http handlers and things to connect to shortener service
- `internal/shortener`
    - database, cache, cr~~u~~d operations,etc...

### URL Shortner


### API gateway

based on what i know for this part, there must be some cache to keep requested urls, and generate single request for all of them.

maybe i can check for duplicate requests to lower the size of duplicates or smth like that.
in that case, it would be good to give requsets coming from users some id, even for duplicate requests, and send a single request for that, with multiple ids :D

using channels looks very promising even tho i haven't started coding yet.

what i can think of about using channels are smth like this for now (kinda pseudo code):

```go
// in http handler that receiving user http request:
ch,err := getFullURL(url)
fullURL <- ch
// redirect user to fullURL


func getFullURL(shortURL string) (chan string, error) {
    // in some select/case statment, i should check for
    // ignore this comment for now hehe // some timer/ticker that do the event every 50ms (maybe read this from configs...)
    // send shortURL to some other function that checks for duplicates,
    // tags them, and wait to receive result from other channel that function reutrns 
    // or maybe that function just register the id for this request in some cache/map
    // and ... i forgot what i wanted to do with that


    // next approach is maybe dont do any select/case here
    // in a map[id]chan, i generate id here, pass short url to some function to queue it for send 
    // (check duplicate and stuff), return id that is added on map,
    // and get chan from that map and wait for result 

    //TODO: i write more here
}

// the thing i noticed now that i forgot about most of things i wanted to do. DONT CHECK YOUR PHONE.
```


anyway, these are rest endpoints this app should have:

- `/new` get url to redirect, probably `POST` request
- `/delete/{id thingy here}`  obviously, `DELETE`
- `/{id thingy here}` to redirect to url `GET`
- ~~`/update/{id thingy here}`~~ what you thought i was gonna do this ? huh you wish


#### Requirements or things i should do

- [ ] sql database
- [ ] in-memory cache
- [ ] grpc between services
- [ ] queue system for api gateway
- [ ] logging ???
- [ ] text
- [ ] dockerize services






