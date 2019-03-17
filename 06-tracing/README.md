## Step 1 :: Start Zipkin server
```
docker container run --rm -p 9411:9411 openzipkin/zipkin
```

Open url=http://localhost:9411 in browser

## Step 2 :: Start service 02

```
$cd service02
$go build
$./service02
```

Open url=http://localhost:9002/hello/{state} in browser
* state = success|fail

## Step 3 :: Start service 01

```
$cd service01
$go build
$./service01
```

Open url=http://localhost:9001/call{state} in browser
* state = success|fail