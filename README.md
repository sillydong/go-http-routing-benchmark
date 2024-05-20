Go HTTP Router Benchmark
========================

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

 * [Beego](http://beego.me/)
 * [go-json-rest](https://github.com/ant0ine/go-json-rest)
 * [Denco](https://github.com/naoina/denco)
 * [Gocraft Web](https://github.com/gocraft/web)
 * [Goji](https://github.com/zenazn/goji/)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)
 * [Martini](https://github.com/go-martini/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [Possum](https://github.com/mikespook/possum)
 * [R2router](https://github.com/vanng822/r2router)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)


## Motivation

Go is a great language for web applications. Since the [default *request multiplexer*](http://golang.org/pkg/net/http/#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Be aware that we are comparing apples and oranges here. We compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

Personally, I prefer slim and optimized software, which is why I implemented [HttpRouter](https://github.com/julienschmidt/httprouter), which is also tested here. In fact, this benchmark suite started as part of the packages tests, but was then extended to a generic benchmark suite.
So keep in mind, that I am not completely unbiased :relieved:


## Results

Benchmark System:
 * Intel Core i5-2500K (4x 3,30GHz + Turbo Boost), CPU-governor: performance
 * 2x 4 GiB DDR3-1333 RAM, dual-channel
 * go version go1.3rc1 linux/amd64
 * Ubuntu 14.04 amd64 (Linux Kernel 3.13.0-29), fresh installation


### Memory Consumption

Besides the micro-benchmarks, there are 3 sets of benchmarks where we play around with clones of some real-world APIs, and one benchmark with static routes only, to allow a comparison with [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).
The following table shows the memory required only for loading the routing structure for the respective API.
The best 3 values for each test are bold. I'm pretty sure you can detect a pattern :wink:

| Router           |      Static |      GitHub |    Google+ |      Parse |
|:-----------------|------------:|------------:|-----------:|-----------:|
| __HttpServeMux__ |     76824 B |    121408 B |     7392 B |    13632 B |
| Ace              |     30648 B |     48632 B |     3664 B |     6656 B |
| Aero             |     30304 B |    488272 B |    25960 B |    28368 B |
| Bear             |     30200 B |     82080 B |     7096 B |    12216 B |
| Beego            |     98008 B |    150632 B |    10256 B |    19096 B |
| Bone             |     40480 B |    100080 B |     6688 B |    11440 B |
| __Chi__          |     83160 B |     94888 B |     8008 B |     9656 B |
| CloudyKitRouter  |     30376 B |     93648 B |     6728 B |    11184 B |
| Denco            |     10776 B | __37624 B__ |     3224 B | __4080 B__ |
| __Echo__         |     97288 B |    123320 B |    11472 B |    14336 B |
| __Gin__          |     34344 B |     58792 B |     4544 B |     7864 B |
| Gocraft Web      |     55304 B |     95656 B |     7480 B |    12712 B |
| Goji             |     27216 B |     46416 B | __2928 B__ |     5248 B |
| Gojiv2           |    107904 B |    105648 B |     7264 B |    14400 B |
| Go-Json-Rest     |    137872 B |    142192 B |    11360 B |    14096 B |
| GoRestful        |    803336 B |   1254336 B |    72152 B |   119920 B |
| __Gorilla Mux__  |    599496 B |   1319696 B |    68000 B |   105384 B |
| GowwwRouter      |     24512 B |     78640 B |     5640 B |     9280 B |
| __HttpRouter__   |     21680 B | __37088 B__ | __2760 B__ | __5024 B__ |
| HttpTreeMux      |     73448 B |     78800 B |     7440 B |     7848 B |
| Kocha            |    123656 B |    785632 B |   128880 B |   181712 B |
| LARS             |     30640 B |     48592 B |     3656 B |     6632 B |
| Macaron          |     38624 B |     93120 B |     8632 B |    13624 B |
| Martini          |    314360 B |    474792 B |    24656 B |    43760 B |
| Pat              | __20200 B__ | __19400 B__ | __1848 B__ | __2552 B__ |
| Possum           |     87504 B |     79944 B |     6624 B |     8112 B |
| R2router         | __23256 B__ |     46616 B |     3864 B |     6920 B |
| Rivet            | __24608 B__ |     42840 B |     3064 B |     5680 B |
| Tango            |     28264 B |     54840 B |     5168 B |     8920 B |
| TigerTonic       |     78752 B |     96832 B |     9392 B |     9768 B |
| Traffic          |    555248 B |    916720 B |    48176 B |    78832 B |
| Vulcan           |    368744 B |    424960 B |    25648 B |    44224 B |

The first place goes to [Pat](https://github.com/bmizerany/pat), followed by [HttpRouter](https://github.com/julienschmidt/httprouter). Now, before everyone starts reading the documentation of Pat, `[SPOILER]` this low memory consumption comes at the price of relatively bad routing performance. The routing structure of Pat is simple - probably too simple. `[/SPOILER]`.

Moreover main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll         5000     706222 ns/op          96 B/op        6 allocs/op
BenchmarkBeego_StaticAll                2000    1408954 ns/op      482433 B/op    14088 allocs/op
BenchmarkDenco_StaticAll              200000      12679 ns/op           0 B/op        0 allocs/op
BenchmarkGocraftWeb_StaticAll          10000     154142 ns/op       51468 B/op      947 allocs/op
BenchmarkGoji_StaticAll                20000      80518 ns/op           0 B/op        0 allocs/op
BenchmarkGoJsonRest_StaticAll           2000     978164 ns/op      180973 B/op     3945 allocs/op
BenchmarkGorillaMux_StaticAll           1000    1763690 ns/op       71804 B/op      956 allocs/op
BenchmarkHttpRouter_StaticAll         100000      15010 ns/op           0 B/op        0 allocs/op
BenchmarkHttpTreeMux_StaticAll        100000      15123 ns/op           0 B/op        0 allocs/op
BenchmarkKocha_StaticAll              100000      23093 ns/op           0 B/op        0 allocs/op
BenchmarkMartini_StaticAll               500    3444278 ns/op      156015 B/op     2351 allocs/op
BenchmarkPat_StaticAll                  1000    1640745 ns/op      549187 B/op    11186 allocs/op
BenchmarkTigerTonic_StaticAll          50000      58264 ns/op        7714 B/op      157 allocs/op
BenchmarkTraffic_StaticAll               500    7230129 ns/op     3763731 B/op    27453 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkHttpServeMux_Param           	 5465816	       217.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkAce_Param                    	10911700	       114.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkAero_Param                   	33411831	        35.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_Param                   	 2238254	       506.8 ns/op	     456 B/op	       5 allocs/op
BenchmarkBeego_Param                  	 1528132	       786.7 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_Param                   	 1585542	       754.6 ns/op	     720 B/op	       5 allocs/op
BenchmarkChi_Param                    	 1724193	       713.2 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_Param        	56165307	        20.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_Param                  	12856096	        90.29 ns/op	      32 B/op	       1 allocs/op
BenchmarkEcho_Param                   	25199661	        46.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param                    	23224945	        48.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param             	 1605573	       728.7 ns/op	     640 B/op	       8 allocs/op
BenchmarkGoji_Param                   	 3021921	       392.1 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Param                 	 1000000	      1004 ns/op	    1040 B/op	       8 allocs/op
BenchmarkGoJsonRest_Param             	 1357318	       863.6 ns/op	     617 B/op	      13 allocs/op
BenchmarkGoRestful_Param              	  353190	      3069 ns/op	    4200 B/op	      15 allocs/op
BenchmarkGorillaMux_Param             	 1000000	      1370 ns/op	    1088 B/op	       8 allocs/op
BenchmarkGowwwRouter_Param            	 4125868	       287.3 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_Param             	17156538	        68.93 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param            	 3461013	       344.3 ns/op	     352 B/op	       3 allocs/op
BenchmarkKocha_Param                  	 6755576	       180.8 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_Param                   	27456474	        42.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param                	  991389	      1540 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_Param                	  426780	      3021 ns/op	    1096 B/op	      12 allocs/op
BenchmarkPat_Param                    	 1412851	       840.1 ns/op	     488 B/op	       9 allocs/op
BenchmarkPossum_Param                 	 4397672	       264.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_Param               	 3226189	       367.9 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_Param                  	12100791	       102.2 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_Param                  	 2218171	       540.0 ns/op	     192 B/op	       6 allocs/op
BenchmarkTigerTonic_Param             	 1000000	      1371 ns/op	     664 B/op	      12 allocs/op
BenchmarkTraffic_Param                	  532950	      2213 ns/op	    1824 B/op	      20 allocs/op
BenchmarkVulcan_Param                 	 4017600	       300.7 ns/op	      98 B/op	       3 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkHttpServeMux_Param5          	 1899558	       622.2 ns/op	     240 B/op	       4 allocs/op
BenchmarkAce_Param5                   	 5213516	       235.8 ns/op	     160 B/op	       1 allocs/op
BenchmarkAero_Param5                  	21325314	        54.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_Param5                  	 1648374	       838.9 ns/op	     501 B/op	       5 allocs/op
BenchmarkBeego_Param5                 	 1000000	      1043 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_Param5                  	 1000000	      1103 ns/op	     768 B/op	       5 allocs/op
BenchmarkChi_Param5                   	 1000000	      1037 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_Param5       	13417222	        88.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_Param5                 	 5516372	       225.8 ns/op	     160 B/op	       1 allocs/op
BenchmarkEcho_Param5                  	11434164	       117.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5                   	12216440	       101.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param5            	 1000000	      1297 ns/op	     912 B/op	      11 allocs/op
BenchmarkGoji_Param5                  	 2165875	       602.5 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Param5                	 1000000	      1501 ns/op	    1104 B/op	       8 allocs/op
BenchmarkGoJsonRest_Param5            	  914564	      1932 ns/op	    1065 B/op	      16 allocs/op
BenchmarkGoRestful_Param5             	  299751	      4197 ns/op	    4312 B/op	      15 allocs/op
BenchmarkGorillaMux_Param5            	  608173	      2531 ns/op	    1152 B/op	       8 allocs/op
BenchmarkGowwwRouter_Param5           	 2340406	       486.8 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_Param5            	 4883850	       238.4 ns/op	     160 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param5           	 1000000	      1064 ns/op	     576 B/op	       6 allocs/op
BenchmarkKocha_Param5                 	 1610242	       733.0 ns/op	     440 B/op	      10 allocs/op
BenchmarkLARS_Param5                  	15985417	        75.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param5               	  846705	      1827 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_Param5               	  296218	      3723 ns/op	    1256 B/op	      13 allocs/op
BenchmarkPat_Param5                   	  603826	      2320 ns/op	     776 B/op	      23 allocs/op
BenchmarkPossum_Param5                	 4417070	       287.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_Param5              	 2394753	       499.5 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_Param5                 	 3622023	       333.3 ns/op	     240 B/op	       1 allocs/op
BenchmarkTango_Param5                 	 1641811	       734.6 ns/op	     312 B/op	       6 allocs/op
BenchmarkTigerTonic_Param5            	  230798	      5374 ns/op	    1975 B/op	      27 allocs/op
BenchmarkTraffic_Param5               	  315151	      3718 ns/op	    2176 B/op	      26 allocs/op
BenchmarkVulcan_Param5                	 2902212	       415.9 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_Param20         	  904251	      1792 ns/op	    1008 B/op	       6 allocs/op
BenchmarkAce_Param20                  	 1714713	       688.1 ns/op	     704 B/op	       1 allocs/op
BenchmarkAero_Param20                 	51319618	        23.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_Param20                 	  453313	      2255 ns/op	    1665 B/op	       5 allocs/op
BenchmarkBeego_Param20                	  486447	      2452 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_Param20                 	  386766	      3159 ns/op	    1935 B/op	       5 allocs/op
BenchmarkChi_Param20                  	  287170	      4102 ns/op	    2524 B/op	       6 allocs/op
BenchmarkCloudyKitRouter_Param20      	 3026379	       380.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_Param20                	 1634906	       743.8 ns/op	     704 B/op	       1 allocs/op
BenchmarkEcho_Param20                 	 3685407	       322.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20                  	 4900233	       246.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Param20           	  237270	      4828 ns/op	    3787 B/op	      15 allocs/op
BenchmarkGoji_Param20                 	  747334	      1692 ns/op	    1247 B/op	       2 allocs/op
BenchmarkGojiv2_Param20               	  671704	      1724 ns/op	    1344 B/op	       8 allocs/op
BenchmarkGoJsonRest_Param20           	  196104	      6358 ns/op	    4581 B/op	      20 allocs/op
BenchmarkGoRestful_Param20            	  158557	      7094 ns/op	    6660 B/op	      17 allocs/op
BenchmarkGorillaMux_Param20           	  214329	      5394 ns/op	    3259 B/op	      10 allocs/op
BenchmarkGowwwRouter_Param20          	 2065264	       613.4 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_Param20           	 1862535	       630.7 ns/op	     704 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Param20          	  278740	      4072 ns/op	    3195 B/op	      10 allocs/op
BenchmarkKocha_Param20                	  515193	      2308 ns/op	    1872 B/op	      27 allocs/op
BenchmarkLARS_Param20                 	 6114734	       213.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Param20              	  290316	      4329 ns/op	    2916 B/op	      12 allocs/op
BenchmarkMartini_Param20              	  165742	      6670 ns/op	    3620 B/op	      15 allocs/op
BenchmarkPat_Param20                  	  125610	      9461 ns/op	    4071 B/op	      72 allocs/op
BenchmarkPossum_Param20               	 4433713	       273.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_Param20             	  405236	      2790 ns/op	    2251 B/op	       6 allocs/op
BenchmarkRivet_Param20                	 1000000	      1121 ns/op	    1024 B/op	       1 allocs/op
BenchmarkTango_Param20                	 1000000	      1338 ns/op	     872 B/op	       6 allocs/op
BenchmarkTigerTonic_Param20           	   59702	     19677 ns/op	    9297 B/op	      77 allocs/op
BenchmarkTraffic_Param20              	   93975	     12701 ns/op	    7823 B/op	      46 allocs/op
BenchmarkVulcan_Param20               	 1748520	       718.3 ns/op	      98 B/op	       3 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkHttpServeMux_ParamWrite      	 5198677	       238.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkAce_ParamWrite               	 6888921	       170.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkAero_ParamWrite              	24208358	        44.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_ParamWrite              	 2210568	       559.6 ns/op	     456 B/op	       5 allocs/op
BenchmarkBeego_ParamWrite             	 1409678	       839.1 ns/op	     360 B/op	       4 allocs/op
BenchmarkBone_ParamWrite              	 1401434	       808.6 ns/op	     720 B/op	       5 allocs/op
BenchmarkChi_ParamWrite               	 1598197	       746.2 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_ParamWrite   	50733813	        21.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_ParamWrite             	11273503	       105.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkEcho_ParamWrite              	10872991	       119.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkGin_ParamWrite               	15480328	        74.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParamWrite        	 1447512	       813.3 ns/op	     648 B/op	       9 allocs/op
BenchmarkGoji_ParamWrite              	 2840547	       434.0 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_ParamWrite            	 1000000	      1157 ns/op	    1072 B/op	      10 allocs/op
BenchmarkGoJsonRest_ParamWrite        	  966004	      1448 ns/op	    1096 B/op	      18 allocs/op
BenchmarkGoRestful_ParamWrite         	  325435	      3290 ns/op	    4208 B/op	      16 allocs/op
BenchmarkGorillaMux_ParamWrite        	 1000000	      1456 ns/op	    1088 B/op	       8 allocs/op
BenchmarkGowwwRouter_ParamWrite       	 1412277	       838.2 ns/op	     752 B/op	       6 allocs/op
BenchmarkHttpRouter_ParamWrite        	13338344	        84.50 ns/op	      32 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParamWrite       	 3159184	       391.2 ns/op	     352 B/op	       3 allocs/op
BenchmarkKocha_ParamWrite             	 5839240	       200.1 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_ParamWrite              	17950340	        62.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParamWrite           	  645147	      1831 ns/op	    1136 B/op	      14 allocs/op
BenchmarkMartini_ParamWrite           	  296262	      3535 ns/op	    1168 B/op	      16 allocs/op
BenchmarkPat_ParamWrite               	 1000000	      1342 ns/op	     912 B/op	      13 allocs/op
BenchmarkPossum_ParamWrite            	 4181599	       275.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_ParamWrite          	 2800455	       440.3 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_ParamWrite             	 5687090	       211.0 ns/op	     112 B/op	       2 allocs/op
BenchmarkTango_ParamWrite             	 4022582	       298.6 ns/op	      96 B/op	       3 allocs/op
BenchmarkTigerTonic_ParamWrite        	  654310	      2212 ns/op	    1104 B/op	      17 allocs/op
BenchmarkTraffic_ParamWrite           	  398692	      3090 ns/op	    2248 B/op	      24 allocs/op
BenchmarkVulcan_ParamWrite            	 3811351	       313.2 ns/op	      98 B/op	       3 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkHttpServeMux_ParseStatic     	 9048776	       132.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAce_ParseStatic              	26481710	        44.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseStatic             	41322638	        28.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_ParseStatic             	 4252377	       288.7 ns/op	     120 B/op	       3 allocs/op
BenchmarkBeego_ParseStatic            	 1412271	       817.5 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_ParseStatic             	 2940679	       409.4 ns/op	     144 B/op	       3 allocs/op
BenchmarkChi_ParseStatic              	 2827348	       430.6 ns/op	     336 B/op	       2 allocs/op
BenchmarkCloudyKitRouter_ParseStatic  	55619330	        20.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_ParseStatic            	64137386	        18.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_ParseStatic             	24743558	        48.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseStatic              	25828315	        46.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseStatic       	 2178440	       461.4 ns/op	     288 B/op	       5 allocs/op
BenchmarkGoji_ParseStatic             	 8245336	       156.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_ParseStatic           	 1237357	      1031 ns/op	    1024 B/op	       7 allocs/op
BenchmarkGoJsonRest_ParseStatic       	 1353042	       787.3 ns/op	     297 B/op	      11 allocs/op
BenchmarkGoRestful_ParseStatic        	  245815	      4228 ns/op	    4296 B/op	      14 allocs/op
BenchmarkGorillaMux_ParseStatic       	  847635	      1210 ns/op	     784 B/op	       7 allocs/op
BenchmarkGowwwRouter_ParseStatic      	45508498	        23.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseStatic       	57818761	        20.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_ParseStatic      	29671724	        39.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_ParseStatic            	38851652	        27.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_ParseStatic             	28013059	        42.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseStatic          	  857366	      1343 ns/op	     728 B/op	       8 allocs/op
BenchmarkMartini_ParseStatic          	  330825	      3047 ns/op	     792 B/op	      11 allocs/op
BenchmarkPat_ParseStatic              	 2308255	       533.4 ns/op	     240 B/op	       5 allocs/op
BenchmarkPossum_ParseStatic           	 2248917	       527.0 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_ParseStatic         	 5710326	       207.3 ns/op	     112 B/op	       3 allocs/op
BenchmarkRivet_ParseStatic            	27727514	        37.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_ParseStatic            	 1971240	       640.2 ns/op	     192 B/op	       6 allocs/op
BenchmarkTigerTonic_ParseStatic       	 6275888	       180.6 ns/op	      48 B/op	       1 allocs/op
BenchmarkTraffic_ParseStatic          	  548349	      1917 ns/op	    1224 B/op	      18 allocs/op
BenchmarkVulcan_ParseStatic           	 3506895	       355.4 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_ParseParam      	 4078958	       255.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkAce_ParseParam               	 6451381	       156.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkAero_ParseParam              	25746639	        43.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_ParseParam              	 1915686	       583.7 ns/op	     467 B/op	       5 allocs/op
BenchmarkBeego_ParseParam             	 1388889	       860.7 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_ParseParam              	  988784	      1083 ns/op	     800 B/op	       6 allocs/op
BenchmarkChi_ParseParam               	 1495933	       836.5 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_ParseParam   	35893054	        30.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_ParseParam             	 8650219	       131.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_ParseParam              	19458058	        58.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseParam               	21312324	        56.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseParam        	 1372714	       817.5 ns/op	     656 B/op	       8 allocs/op
BenchmarkGoji_ParseParam              	 2119477	       575.6 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_ParseParam            	  856065	      1288 ns/op	    1072 B/op	       9 allocs/op
BenchmarkGoJsonRest_ParseParam        	  989040	      1189 ns/op	     617 B/op	      13 allocs/op
BenchmarkGoRestful_ParseParam         	  253110	      4906 ns/op	    4616 B/op	      15 allocs/op
BenchmarkGorillaMux_ParseParam        	  690854	      1529 ns/op	    1088 B/op	       8 allocs/op
BenchmarkGowwwRouter_ParseParam       	 3482479	       337.4 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_ParseParam        	11444320	        99.84 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_ParseParam       	 2807654	       413.4 ns/op	     352 B/op	       3 allocs/op
BenchmarkKocha_ParseParam             	 5577115	       213.7 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_ParseParam              	25113572	        46.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseParam           	  608556	      1755 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_ParseParam           	  310753	      3337 ns/op	    1096 B/op	      12 allocs/op
BenchmarkPat_ParseParam               	  678992	      1584 ns/op	     928 B/op	      12 allocs/op
BenchmarkPossum_ParseParam            	 4062718	       293.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_ParseParam          	 2850289	       475.1 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_ParseParam             	 8555395	       125.2 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_ParseParam             	 1902853	       701.8 ns/op	     224 B/op	       6 allocs/op
BenchmarkTigerTonic_ParseParam        	  769236	      1972 ns/op	     672 B/op	      11 allocs/op
BenchmarkTraffic_ParseParam           	  352258	      2945 ns/op	    1864 B/op	      20 allocs/op
BenchmarkVulcan_ParseParam            	 2838192	       548.6 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_Parse2Params    	 2390662	       424.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkAce_Parse2Params             	 7121470	       181.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkAero_Parse2Params            	19779590	        59.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_Parse2Params            	 1633340	       653.5 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_Parse2Params           	 1235704	      1088 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_Parse2Params            	  928798	      1162 ns/op	     752 B/op	       5 allocs/op
BenchmarkChi_Parse2Params             	 1000000	      1029 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_Parse2Params 	24705878	        44.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_Parse2Params           	 8138413	       144.8 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_Parse2Params            	17260183	        69.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Parse2Params             	17627776	        68.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_Parse2Params      	 1300046	       922.1 ns/op	     704 B/op	       9 allocs/op
BenchmarkGoji_Parse2Params            	 2228156	       489.9 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_Parse2Params          	  953311	      1146 ns/op	    1056 B/op	       8 allocs/op
BenchmarkGoJsonRest_Parse2Params      	  850563	      1280 ns/op	     681 B/op	      14 allocs/op
BenchmarkGoRestful_Parse2Params       	  239770	      4861 ns/op	    5040 B/op	      15 allocs/op
BenchmarkGorillaMux_Parse2Params      	  691602	      1858 ns/op	    1104 B/op	       8 allocs/op
BenchmarkGowwwRouter_Parse2Params     	 3457646	       341.6 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_Parse2Params      	 9073614	       119.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_Parse2Params     	 2309328	       511.0 ns/op	     384 B/op	       4 allocs/op
BenchmarkKocha_Parse2Params           	 3041686	       330.1 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_Parse2Params            	21266056	        56.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_Parse2Params         	  634230	      1746 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_Parse2Params         	  315937	      3495 ns/op	    1176 B/op	      13 allocs/op
BenchmarkPat_Parse2Params             	  844086	      1430 ns/op	     688 B/op	      13 allocs/op
BenchmarkPossum_Parse2Params          	 4172289	       299.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_Parse2Params        	 2561953	       451.5 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_Parse2Params           	 6822342	       172.6 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_Parse2Params           	 1860228	       644.9 ns/op	     264 B/op	       6 allocs/op
BenchmarkTigerTonic_Parse2Params      	  555007	      2109 ns/op	    1008 B/op	      16 allocs/op
BenchmarkTraffic_Parse2Params         	  385471	      2696 ns/op	    1912 B/op	      21 allocs/op
BenchmarkVulcan_Parse2Params          	 3012182	       384.4 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_ParseAll        	  139440	      7602 ns/op	     352 B/op	      19 allocs/op
BenchmarkAce_ParseAll                 	  384374	      3264 ns/op	     640 B/op	      16 allocs/op
BenchmarkAero_ParseAll                	 1011798	      1142 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_ParseAll                	   82966	     13367 ns/op	    8920 B/op	     110 allocs/op
BenchmarkBeego_ParseAll               	   53978	     22345 ns/op	    9152 B/op	      78 allocs/op
BenchmarkBone_ParseAll                	   45613	     24798 ns/op	   14672 B/op	     131 allocs/op
BenchmarkChi_ParseAll                 	   66042	     17975 ns/op	   14112 B/op	      84 allocs/op
BenchmarkCloudyKitRouter_ParseAll     	 1502521	       787.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_ParseAll               	  404728	      2823 ns/op	     928 B/op	      16 allocs/op
BenchmarkEcho_ParseAll                	  655888	      1842 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseAll                 	  785426	      1524 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_ParseAll          	   62310	     20622 ns/op	   13520 B/op	     181 allocs/op
BenchmarkGoji_ParseAll                	  108182	     11585 ns/op	    5376 B/op	      32 allocs/op
BenchmarkGojiv2_ParseAll              	   39866	     31299 ns/op	   26960 B/op	     199 allocs/op
BenchmarkGoJsonRest_ParseAll          	   44426	     26759 ns/op	   13034 B/op	     321 allocs/op
BenchmarkGoRestful_ParseAll           	    9660	    113881 ns/op	  119408 B/op	     380 allocs/op
BenchmarkGorillaMux_ParseAll          	   20108	     60825 ns/op	   25296 B/op	     198 allocs/op
BenchmarkGowwwRouter_ParseAll         	  166765	      6509 ns/op	    5376 B/op	      32 allocs/op
BenchmarkHttpRouter_ParseAll          	  615121	      2182 ns/op	     640 B/op	      16 allocs/op
BenchmarkHttpTreeMux_ParseAll         	  155962	      7860 ns/op	    5728 B/op	      51 allocs/op
BenchmarkKocha_ParseAll               	  260070	      4908 ns/op	    1112 B/op	      54 allocs/op
BenchmarkLARS_ParseAll                	  879696	      1434 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_ParseAll             	   37680	     35537 ns/op	   18928 B/op	     208 allocs/op
BenchmarkMartini_ParseAll             	   13250	     89313 ns/op	   25696 B/op	     305 allocs/op
BenchmarkPat_ParseAll                 	   37822	     31773 ns/op	   14304 B/op	     272 allocs/op
BenchmarkPossum_ParseAll              	   86901	     13547 ns/op	   10815 B/op	      78 allocs/op
BenchmarkR2router_ParseAll            	  122778	      9746 ns/op	    7520 B/op	      94 allocs/op
BenchmarkRivet_ParseAll               	  390045	      3067 ns/op	     912 B/op	      16 allocs/op
BenchmarkTango_ParseAll               	   67813	     17637 ns/op	    5864 B/op	     156 allocs/op
BenchmarkTigerTonic_ParseAll          	   36517	     31174 ns/op	   14112 B/op	     262 allocs/op
BenchmarkTraffic_ParseAll             	   17911	     70479 ns/op	   44689 B/op	     579 allocs/op
BenchmarkVulcan_ParseAll              	  109344	     10679 ns/op	    2548 B/op	      78 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkHttpServeMux_GithubStatic    	 6761012	       178.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAce_GithubStatic             	16953830	        69.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic            	30110296	        37.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GithubStatic            	 4184908	       290.8 ns/op	     120 B/op	       3 allocs/op
BenchmarkBeego_GithubStatic           	 1482186	       817.9 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_GithubStatic            	  157602	      7767 ns/op	    2880 B/op	      60 allocs/op
BenchmarkCloudyKitRouter_GithubStatic 	28246358	        37.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkChi_GithubStatic             	 2635764	       461.6 ns/op	     336 B/op	       2 allocs/op
BenchmarkDenco_GithubStatic           	49484086	        22.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GithubStatic            	18174693	        65.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubStatic             	18930952	        63.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubStatic      	 2600448	       465.2 ns/op	     288 B/op	       5 allocs/op
BenchmarkGoji_GithubStatic            	 7852814	       159.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GithubStatic          	 1000000	      1049 ns/op	    1024 B/op	       7 allocs/op
BenchmarkGoRestful_GithubStatic       	  172465	      6915 ns/op	    4296 B/op	      14 allocs/op
BenchmarkGoJsonRest_GithubStatic      	 1699747	       704.5 ns/op	     297 B/op	      11 allocs/op
BenchmarkGorillaMux_GithubStatic      	  399555	      3379 ns/op	     784 B/op	       7 allocs/op
BenchmarkGowwwRouter_GithubStatic     	19757596	        59.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GithubStatic      	33845946	        34.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GithubStatic     	31087028	        38.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GithubStatic           	30327891	        40.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_GithubStatic            	22255046	        53.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubStatic         	 1000000	      1172 ns/op	     728 B/op	       8 allocs/op
BenchmarkMartini_GithubStatic         	  237128	      5411 ns/op	     792 B/op	      11 allocs/op
BenchmarkPat_GithubStatic             	  152574	      7435 ns/op	    3648 B/op	      76 allocs/op
BenchmarkPossum_GithubStatic          	 2425903	       478.7 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_GithubStatic        	 5641338	       209.4 ns/op	     112 B/op	       3 allocs/op
BenchmarkRivet_GithubStatic           	19223307	        62.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_GithubStatic           	 1680769	       707.2 ns/op	     192 B/op	       6 allocs/op
BenchmarkTigerTonic_GithubStatic      	 7049305	       170.7 ns/op	      48 B/op	       1 allocs/op
BenchmarkTraffic_GithubStatic         	  158122	      7687 ns/op	    4632 B/op	      89 allocs/op
BenchmarkVulcan_GithubStatic          	 2644215	       458.8 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_GithubParam     	 2398608	       508.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkAce_GithubParam              	 5612031	       213.4 ns/op	      96 B/op	       1 allocs/op
BenchmarkAero_GithubParam             	14859574	        81.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GithubParam             	 1763348	       666.1 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_GithubParam            	 1000000	      1025 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_GithubParam             	  257737	      4502 ns/op	    1792 B/op	      18 allocs/op
BenchmarkChi_GithubParam              	 1315644	       908.2 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_GithubParam  	10455927	       112.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GithubParam            	 5798680	       207.5 ns/op	     128 B/op	       1 allocs/op
BenchmarkEcho_GithubParam             	10150003	       116.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubParam              	11051881	       107.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubParam       	 1268599	       922.9 ns/op	     704 B/op	       9 allocs/op
BenchmarkGoji_GithubParam             	 1826998	       652.3 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GithubParam           	 1000000	      1499 ns/op	    1120 B/op	      10 allocs/op
BenchmarkGoJsonRest_GithubParam       	 1000000	      1256 ns/op	     681 B/op	      14 allocs/op
BenchmarkGoRestful_GithubParam        	  130336	      9261 ns/op	    4296 B/op	      15 allocs/op
BenchmarkGorillaMux_GithubParam       	  254048	      5026 ns/op	    1104 B/op	       8 allocs/op
BenchmarkGowwwRouter_GithubParam      	 3246871	       368.4 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_GithubParam       	 7116517	       166.0 ns/op	      96 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GithubParam      	 2209821	       549.8 ns/op	     384 B/op	       4 allocs/op
BenchmarkKocha_GithubParam            	 3381410	       357.0 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_GithubParam             	11035650	        91.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubParam          	  820053	      1789 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_GithubParam          	  174650	      6843 ns/op	    1176 B/op	      13 allocs/op
BenchmarkPat_GithubParam              	  198342	      5919 ns/op	    2280 B/op	      41 allocs/op
BenchmarkPossum_GithubParam           	 4379406	       280.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_GithubParam         	 2838924	       427.9 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_GithubParam            	 4747910	       258.1 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_GithubParam            	 1474998	       806.9 ns/op	     296 B/op	       6 allocs/op
BenchmarkTigerTonic_GithubParam       	  641031	      2269 ns/op	     960 B/op	      16 allocs/op
BenchmarkTraffic_GithubParam          	  175928	      6421 ns/op	    2760 B/op	      39 allocs/op
BenchmarkVulcan_GithubParam           	 1725438	       695.7 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_GithubAll       	   10000	    105116 ns/op	    9744 B/op	     337 allocs/op
BenchmarkAce_GithubAll                	   30418	     40465 ns/op	   13792 B/op	     167 allocs/op
BenchmarkAero_GithubAll               	   75484	     15703 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GithubAll               	   10000	    140081 ns/op	   86448 B/op	     943 allocs/op
BenchmarkBeego_GithubAll              	    7094	    211377 ns/op	   71456 B/op	     609 allocs/op
BenchmarkBone_GithubAll               	     650	   1827501 ns/op	  704128 B/op	    8453 allocs/op
BenchmarkChi_GithubAll                	    6446	    181966 ns/op	  124320 B/op	     740 allocs/op
BenchmarkCloudyKitRouter_GithubAll    	   66229	     17947 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GithubAll              	   29504	     40705 ns/op	   20224 B/op	     167 allocs/op
BenchmarkEcho_GithubAll               	   46300	     26654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll                	   56760	     21259 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GithubAll         	    5242	    191267 ns/op	  130032 B/op	    1686 allocs/op
BenchmarkGoji_GithubAll               	    3991	    311754 ns/op	   56112 B/op	     334 allocs/op
BenchmarkGojiv2_GithubAll             	    2384	    462300 ns/op	  294256 B/op	    3712 allocs/op
BenchmarkGoJsonRest_GithubAll         	    5113	    250919 ns/op	  127875 B/op	    2737 allocs/op
BenchmarkGoRestful_GithubAll          	     692	   1699081 ns/op	  914568 B/op	    3009 allocs/op
BenchmarkGorillaMux_GithubAll         	     472	   2551269 ns/op	  212674 B/op	    1588 allocs/op
BenchmarkGowwwRouter_GithubAll        	   16252	     74022 ns/op	   56112 B/op	     334 allocs/op
BenchmarkHttpRouter_GithubAll         	   38299	     31104 ns/op	   13792 B/op	     167 allocs/op
BenchmarkHttpTreeMux_GithubAll        	   10000	    101065 ns/op	   65856 B/op	     671 allocs/op
BenchmarkKocha_GithubAll              	   15237	     84964 ns/op	   23304 B/op	     843 allocs/op
BenchmarkLARS_GithubAll               	   67335	     18234 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GithubAll            	    4299	    239927 ns/op	  147784 B/op	    1624 allocs/op
BenchmarkMartini_GithubAll            	     493	   2533526 ns/op	  231418 B/op	    2731 allocs/op
BenchmarkPat_GithubAll                	     430	   2740790 ns/op	 1406616 B/op	   22348 allocs/op
BenchmarkPossum_GithubAll             	   10000	    112181 ns/op	   84448 B/op	     609 allocs/op
BenchmarkR2router_GithubAll           	   12776	     94096 ns/op	   70832 B/op	     776 allocs/op
BenchmarkRivet_GithubAll              	   24200	     49211 ns/op	   16272 B/op	     167 allocs/op
BenchmarkTango_GithubAll              	    9650	    169307 ns/op	   53849 B/op	    1215 allocs/op
BenchmarkTigerTonic_GithubAll         	    2665	    443264 ns/op	  166560 B/op	    3462 allocs/op
BenchmarkTraffic_GithubAll            	     474	   2281219 ns/op	  814175 B/op	   13911 allocs/op
BenchmarkVulcan_GithubAll             	   10000	    120555 ns/op	   19894 B/op	     609 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkHttpServeMux_GPlusStatic     	12028470	        99.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkAce_GPlusStatic              	24615852	        48.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusStatic             	45418159	        26.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GPlusStatic             	 5421914	       222.0 ns/op	     104 B/op	       3 allocs/op
BenchmarkBeego_GPlusStatic            	 1663098	       721.8 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_GPlusStatic             	10218438	       118.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkChi_GPlusStatic              	 2896255	       411.1 ns/op	     336 B/op	       2 allocs/op
BenchmarkCloudyKitRouter_GPlusStatic  	62745422	        19.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GPlusStatic            	83691806	        14.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_GPlusStatic             	24680076	        48.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusStatic              	27089271	        44.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusStatic       	 2929268	       408.7 ns/op	     272 B/op	       5 allocs/op
BenchmarkGoji_GPlusStatic             	10874802	       111.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGojiv2_GPlusStatic           	 1288410	       937.5 ns/op	    1024 B/op	       7 allocs/op
BenchmarkGoJsonRest_GPlusStatic       	 1990029	       602.3 ns/op	     297 B/op	      11 allocs/op
BenchmarkGoRestful_GPlusStatic        	  371722	      2996 ns/op	    3872 B/op	      14 allocs/op
BenchmarkGorillaMux_GPlusStatic       	 1267455	       931.7 ns/op	     784 B/op	       7 allocs/op
BenchmarkGowwwRouter_GPlusStatic      	57062635	        21.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusStatic       	66416431	        16.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic      	40881658	        26.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkKocha_GPlusStatic            	53502736	        22.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkLARS_GPlusStatic             	27932422	        38.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusStatic          	 1000000	      1138 ns/op	     728 B/op	       8 allocs/op
BenchmarkMartini_GPlusStatic          	  523855	      2706 ns/op	     792 B/op	      11 allocs/op
BenchmarkPat_GPlusStatic              	 5869297	       200.4 ns/op	      96 B/op	       2 allocs/op
BenchmarkPossum_GPlusStatic           	 2556429	       465.7 ns/op	     416 B/op	       3 allocs/op
BenchmarkR2router_GPlusStatic         	 6640219	       175.6 ns/op	     112 B/op	       3 allocs/op
BenchmarkRivet_GPlusStatic            	37382448	        32.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkTango_GPlusStatic            	 2311534	       515.7 ns/op	     152 B/op	       6 allocs/op
BenchmarkTigerTonic_GPlusStatic       	10638046	       108.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkTraffic_GPlusStatic          	 1000000	      1333 ns/op	    1080 B/op	      15 allocs/op
BenchmarkVulcan_GPlusStatic           	 4236765	       284.7 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_GPlusParam      	 4081885	       294.6 ns/op	      16 B/op	       1 allocs/op
BenchmarkAce_GPlusParam               	 7840766	       150.4 ns/op	      64 B/op	       1 allocs/op
BenchmarkAero_GPlusParam              	26666523	        44.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GPlusParam              	 2265758	       508.0 ns/op	     472 B/op	       5 allocs/op
BenchmarkBeego_GPlusParam             	 1401324	       858.3 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_GPlusParam              	 1420401	       834.0 ns/op	     720 B/op	       5 allocs/op
BenchmarkChi_GPlusParam               	 1618464	       736.8 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_GPlusParam   	32907574	        36.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GPlusParam             	 9917472	       119.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_GPlusParam              	15575216	        75.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusParam               	18191904	        66.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusParam        	 1409857	       793.3 ns/op	     640 B/op	       8 allocs/op
BenchmarkGoji_GPlusParam              	 2563616	       448.1 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GPlusParam            	 1000000	      1219 ns/op	    1040 B/op	       8 allocs/op
BenchmarkGoJsonRest_GPlusParam        	 1000000	      1102 ns/op	     617 B/op	      13 allocs/op
BenchmarkGoRestful_GPlusParam         	  325508	      3668 ns/op	    4216 B/op	      15 allocs/op
BenchmarkGorillaMux_GPlusParam        	  873915	      1865 ns/op	    1088 B/op	       8 allocs/op
BenchmarkGowwwRouter_GPlusParam       	 3859592	       306.0 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_GPlusParam        	11630635	       105.0 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GPlusParam       	 2787085	       423.7 ns/op	     352 B/op	       3 allocs/op
BenchmarkKocha_GPlusParam             	 5139994	       210.0 ns/op	      56 B/op	       3 allocs/op
BenchmarkLARS_GPlusParam              	18890013	        64.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusParam           	  933618	      1734 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_GPlusParam           	  391327	      3543 ns/op	    1096 B/op	      12 allocs/op
BenchmarkPat_GPlusParam               	 1000000	      1085 ns/op	     528 B/op	       9 allocs/op
BenchmarkPossum_GPlusParam            	 4359518	       286.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_GPlusParam          	 3074923	       406.3 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_GPlusParam             	 8877614	       132.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkTango_GPlusParam             	 1819788	       649.7 ns/op	     216 B/op	       6 allocs/op
BenchmarkTigerTonic_GPlusParam        	  872292	      1579 ns/op	     744 B/op	      12 allocs/op
BenchmarkTraffic_GPlusParam           	  464478	      2641 ns/op	    1840 B/op	      20 allocs/op
BenchmarkVulcan_GPlusParam            	 2934632	       413.3 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_GPlus2Params    	 2314809	       515.3 ns/op	      48 B/op	       2 allocs/op
BenchmarkAce_GPlus2Params             	 6517754	       194.6 ns/op	      64 B/op	       1 allocs/op
BenchmarkAero_GPlus2Params            	15281761	        74.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GPlus2Params            	 1794117	       658.3 ns/op	     496 B/op	       5 allocs/op
BenchmarkBeego_GPlus2Params           	 1000000	      1052 ns/op	     352 B/op	       3 allocs/op
BenchmarkBone_GPlus2Params            	  655886	      2355 ns/op	    1072 B/op	       9 allocs/op
BenchmarkChi_GPlus2Params             	 1264939	       849.0 ns/op	     672 B/op	       4 allocs/op
BenchmarkCloudyKitRouter_GPlus2Params 	19263645	        62.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GPlus2Params           	 6986532	       172.5 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho_GPlus2Params            	12743157	        93.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlus2Params             	13883689	        86.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlus2Params      	 1270357	      1002 ns/op	     704 B/op	       9 allocs/op
BenchmarkGoji_GPlus2Params            	 1735557	       688.1 ns/op	     336 B/op	       2 allocs/op
BenchmarkGojiv2_GPlus2Params          	  992594	      1647 ns/op	    1120 B/op	      11 allocs/op
BenchmarkGoJsonRest_GPlus2Params      	 1000000	      1273 ns/op	     681 B/op	      14 allocs/op
BenchmarkGoRestful_GPlus2Params       	  283909	      4000 ns/op	    4312 B/op	      15 allocs/op
BenchmarkGorillaMux_GPlus2Params      	  305373	      3953 ns/op	    1104 B/op	       8 allocs/op
BenchmarkGowwwRouter_GPlus2Params     	 3385851	       347.2 ns/op	     336 B/op	       2 allocs/op
BenchmarkHttpRouter_GPlus2Params      	 9161599	       130.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params     	 1991605	       558.3 ns/op	     384 B/op	       4 allocs/op
BenchmarkKocha_GPlus2Params           	 3373486	       369.8 ns/op	     128 B/op	       5 allocs/op
BenchmarkLARS_GPlus2Params            	14403019	        80.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlus2Params         	  734269	      1832 ns/op	    1064 B/op	      10 allocs/op
BenchmarkMartini_GPlus2Params         	  181400	      6716 ns/op	    1224 B/op	      15 allocs/op
BenchmarkPat_GPlus2Params             	  276254	      4059 ns/op	    2056 B/op	      27 allocs/op
BenchmarkPossum_GPlus2Params          	 4340661	       282.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkR2router_GPlus2Params        	 2255727	       504.2 ns/op	     400 B/op	       4 allocs/op
BenchmarkRivet_GPlus2Params           	 5327124	       194.8 ns/op	      96 B/op	       1 allocs/op
BenchmarkTango_GPlus2Params           	 1581183	       741.4 ns/op	     296 B/op	       6 allocs/op
BenchmarkTigerTonic_GPlus2Params      	  592950	      2469 ns/op	    1040 B/op	      16 allocs/op
BenchmarkTraffic_GPlus2Params         	  215583	      5330 ns/op	    2216 B/op	      27 allocs/op
BenchmarkVulcan_GPlus2Params          	 1679697	       683.5 ns/op	      98 B/op	       3 allocs/op

BenchmarkHttpServeMux_GPlusAll        	  235812	      4844 ns/op	     336 B/op	      16 allocs/op
BenchmarkAce_GPlusAll                 	  737550	      2006 ns/op	     640 B/op	      11 allocs/op
BenchmarkAero_GPlusAll                	 1654206	       727.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkBear_GPlusAll                	  151028	      8226 ns/op	    5488 B/op	      61 allocs/op
BenchmarkBeego_GPlusAll               	   98876	     13697 ns/op	    4576 B/op	      39 allocs/op
BenchmarkBone_GPlusAll                	   58953	     19716 ns/op	   10688 B/op	      98 allocs/op
BenchmarkChi_GPlusAll                 	  121375	     10783 ns/op	    8064 B/op	      48 allocs/op
BenchmarkCloudyKitRouter_GPlusAll     	 2081798	       585.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkDenco_GPlusAll               	  853987	      1934 ns/op	     672 B/op	      11 allocs/op
BenchmarkEcho_GPlusAll                	 1000000	      1048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusAll                 	 1000000	      1090 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocraftWeb_GPlusAll          	   94009	     11109 ns/op	    7936 B/op	     103 allocs/op
BenchmarkGoji_GPlusAll                	  163762	      6921 ns/op	    3696 B/op	      22 allocs/op
BenchmarkGojiv2_GPlusAll              	   67167	     18018 ns/op	   13872 B/op	     115 allocs/op
BenchmarkGoJsonRest_GPlusAll          	   83306	     15009 ns/op	    7701 B/op	     170 allocs/op
BenchmarkGoRestful_GPlusAll           	   19191	     55121 ns/op	   55328 B/op	     193 allocs/op
BenchmarkGorillaMux_GPlusAll          	   38548	     31231 ns/op	   13616 B/op	     102 allocs/op
BenchmarkGowwwRouter_GPlusAll         	  263602	      4080 ns/op	    3696 B/op	      22 allocs/op
BenchmarkHttpRouter_GPlusAll          	 1000000	      1380 ns/op	     640 B/op	      11 allocs/op
BenchmarkHttpTreeMux_GPlusAll         	  233121	      5789 ns/op	    4032 B/op	      38 allocs/op
BenchmarkKocha_GPlusAll               	  376674	      3670 ns/op	     976 B/op	      43 allocs/op
BenchmarkLARS_GPlusAll                	 1297966	       826.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMacaron_GPlusAll             	   66444	     17007 ns/op	    9464 B/op	     104 allocs/op
BenchmarkMartini_GPlusAll             	   21888	     55383 ns/op	   14328 B/op	     171 allocs/op
BenchmarkPat_GPlusAll                 	   38169	     29888 ns/op	   14432 B/op	     231 allocs/op
BenchmarkPossum_GPlusAll              	  163995	      7357 ns/op	    5408 B/op	      39 allocs/op
BenchmarkR2router_GPlusAll            	  200451	      5493 ns/op	    4624 B/op	      50 allocs/op
BenchmarkRivet_GPlusAll               	  759523	      2080 ns/op	     768 B/op	      11 allocs/op
BenchmarkTango_GPlusAll               	  137558	      8572 ns/op	    3008 B/op	      78 allocs/op
BenchmarkTigerTonic_GPlusAll          	   42288	     26871 ns/op	   10120 B/op	     188 allocs/op
BenchmarkTraffic_GPlusAll             	   24654	     47617 ns/op	   25832 B/op	     328 allocs/op
BenchmarkVulcan_GPlusAll              	  194070	      6360 ns/op	    1274 B/op	      39 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.

## Usage

If you'd like to run these benchmarks locally, you'll need to install the package first:

```bash
go get github.com/julienschmidt/go-http-routing-benchmark
```
This may take a while due to the large number of dependencies that need to be downloaded. Once that command has finished you can run the full set of benchmarks like this:

```bash
cd $GOPATH/src/github.com/julienschmidt/go-http-routing-benchmark
go test -bench=.
```

> **Note:** If you run the tests and it SIGQUIT's make the go test timeout longer (#44)
>
```
go test -timeout=2h -bench=.
```


You can bench specific frameworks only by using a regular expression as the value of the `bench` parameter:
```bash
go test -bench="Martini|Gin|HttpMux"
```
