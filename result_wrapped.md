#ExtraAPI Routes: 2
   Denco: 864 Bytes
   Echo: 3616 Bytes
   Gin: 1192 Bytes
   GorillaMux: 6648 Bytes
   HttpTreeMux: 1704 Bytes
   ServeMux: 1816 Bytes

#GithubAPI Routes: 203
   Chi: 94904 Bytes
   Denco: 37640 Bytes
   Echo: 123400 Bytes
   Gin: 58808 Bytes
   GorillaMux: 1319712 Bytes
   HttpRouter: 37104 Bytes
   HttpTreeMux: 78816 Bytes
   ServeMux: 121216 Bytes

#GPlusAPI Routes: 13
   Chi: 8024 Bytes
   Denco: 3240 Bytes
   Echo: 11344 Bytes
   Gin: 4560 Bytes
   GorillaMux: 68016 Bytes
   HttpRouter: 2776 Bytes
   HttpTreeMux: 7456 Bytes
   ServeMux: 7408 Bytes

#ParseAPI Routes: 26
   Chi: 9672 Bytes
   Denco: 4096 Bytes
   Echo: 14416 Bytes
   Gin: 7880 Bytes
   GorillaMux: 105400 Bytes
   HttpRouter: 5040 Bytes
   HttpTreeMux: 7864 Bytes
   ServeMux: 13648 Bytes

#Static Routes: 157
   HttpServeMux: 74776 Bytes
   Chi: 83176 Bytes
   Denco: 10216 Bytes
   Echo: 97368 Bytes
   Gin: 34360 Bytes
   GorillaMux: 599512 Bytes
   HttpRouter: 21696 Bytes
   HttpTreeMux: 73464 Bytes
   ServeMux: 76632 Bytes

goos: darwin
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkChi_Param                       2402067               482.2 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_Param                     7684461               145.6 ns/op            56 B/op          2 allocs/op
BenchmarkEcho_Param                      5840106               187.2 ns/op            40 B/op          2 allocs/op
BenchmarkGin_Param                       7987465               133.1 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_Param                 887809              1325 ns/op            1112 B/op          9 allocs/op
BenchmarkHttpRouter_Param                9149572               125.8 ns/op            56 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param               3081914               395.3 ns/op           376 B/op          4 allocs/op
BenchmarkServeMux_Param                  4116349               287.6 ns/op            40 B/op          2 allocs/op
BenchmarkChi_Param5                      1755902               687.9 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_Param5                    4357380               260.7 ns/op           184 B/op          2 allocs/op
BenchmarkEcho_Param5                     4699460               254.0 ns/op            40 B/op          2 allocs/op
BenchmarkGin_Param5                      5975685               199.8 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_Param5                594471              1989 ns/op            1176 B/op          9 allocs/op
BenchmarkHttpRouter_Param5               4887799               232.2 ns/op           184 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param5              1444081               842.3 ns/op           600 B/op          7 allocs/op
BenchmarkServeMux_Param5                 1756706               686.4 ns/op           264 B/op          5 allocs/op
BenchmarkChi_Param20                     1000000              1340 ns/op             360 B/op          3 allocs/op
BenchmarkDenco_Param20                   1668542               726.4 ns/op           728 B/op          2 allocs/op
BenchmarkEcho_Param20                    2450503               476.3 ns/op            40 B/op          2 allocs/op
BenchmarkGin_Param20                     3316458               333.4 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_Param20               235434              4513 ns/op            3284 B/op         11 allocs/op
BenchmarkHttpRouter_Param20              1930668               603.4 ns/op           728 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param20              325648              3687 ns/op            3219 B/op         11 allocs/op
BenchmarkServeMux_Param20                 728596              1661 ns/op            1032 B/op          7 allocs/op
BenchmarkChi_ParamWrite                  2653207               438.9 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_ParamWrite                6438277               171.4 ns/op            64 B/op          3 allocs/op
BenchmarkEcho_ParamWrite                 5484009               212.6 ns/op            48 B/op          3 allocs/op
BenchmarkGin_ParamWrite                  7128556               159.2 ns/op            32 B/op          2 allocs/op
BenchmarkGorillaMux_ParamWrite            898590              1328 ns/op            1120 B/op         10 allocs/op
BenchmarkHttpRouter_ParamWrite           7750078               147.5 ns/op            64 B/op          3 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2916175               407.1 ns/op           384 B/op          5 allocs/op
BenchmarkServeMux_ParamWrite             5198062               222.7 ns/op            16 B/op          1 allocs/op
BenchmarkDenco_ExtraStatic               1451322               820.9 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_ExtraStatic                 921679              1204 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraStatic                 3831018               332.0 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraStatic          1346230               880.0 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraStatic         1265390               958.2 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraStatic             608238              1979 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_ExtraParam                1321983               814.9 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_ExtraParam                  866134              1204 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraParam                  3764395               310.3 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraParam           1240334               960.5 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraParam          1419510               843.5 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraParam              692074              1680 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_Extra2Params              1436050               805.1 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_Extra2Params               1016449              1177 ns/op             896 B/op         10 allocs/op
BenchmarkGin_Extra2Params                3831333               316.9 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_Extra2Params         1220664               963.1 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_Extra2Params        1404056               835.9 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_Extra2Params            604580              1772 ns/op            1024 B/op         17 allocs/op
BenchmarkDenco_ExtraAll                  4793409               239.1 ns/op            80 B/op          3 allocs/op
BenchmarkEcho_ExtraAll                   3036296               390.1 ns/op            80 B/op          4 allocs/op
BenchmarkGin_ExtraAll                    2579630               434.7 ns/op           176 B/op          4 allocs/op
BenchmarkHttpTreeMux_ExtraAll            2238794               528.8 ns/op           400 B/op          5 allocs/op
BenchmarkServeMux_ExtraAll               2033703               592.0 ns/op            64 B/op          3 allocs/op
BenchmarkChi_GithubStatic                2515555               481.3 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_GithubStatic             13599382                81.58 ns/op           24 B/op          1 allocs/op
BenchmarkEcho_GithubStatic               5484403               208.3 ns/op            40 B/op          2 allocs/op
BenchmarkGin_GithubStatic                8362200               136.8 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_GithubStatic          334430              3206 ns/op             808 B/op          8 allocs/op
BenchmarkHttpRouter_GithubStatic        12307555                91.30 ns/op           24 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubStatic       11252208                98.27 ns/op           24 B/op          1 allocs/op
BenchmarkServeMux_GithubStatic           4824884               236.8 ns/op            24 B/op          1 allocs/op
BenchmarkChi_GithubParam                 1908456               630.0 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_GithubParam               4683799               241.5 ns/op           152 B/op          2 allocs/op
BenchmarkEcho_GithubParam                4158852               280.1 ns/op            40 B/op          2 allocs/op
BenchmarkGin_GithubParam                 5719396               203.4 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_GithubParam           219951              4747 ns/op            1128 B/op          9 allocs/op
BenchmarkHttpRouter_GithubParam          5365462               213.0 ns/op           120 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubParam         2108346               554.2 ns/op           408 B/op          5 allocs/op
BenchmarkServeMux_GithubParam            2086089               557.2 ns/op            72 B/op          3 allocs/op
BenchmarkChi_GithubAll                      8985            131793 ns/op           73081 B/op        609 allocs/op
BenchmarkDenco_GithubAll                   23391             54945 ns/op           25096 B/op        370 allocs/op
BenchmarkEcho_GithubAll                    17857             64048 ns/op            8120 B/op        406 allocs/op
BenchmarkGin_GithubAll                     27409             48783 ns/op            4872 B/op        203 allocs/op
BenchmarkGorillaMux_GithubAll                430           2843577 ns/op          217549 B/op       1791 allocs/op
BenchmarkHttpRouter_GithubAll              25970             45883 ns/op           18664 B/op        370 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            126638 ns/op           70728 B/op        874 allocs/op
BenchmarkServeMux_GithubAll                10000            124539 ns/op           14616 B/op        540 allocs/op
BenchmarkChi_GPlusStatic                 2469238               476.9 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_GPlusStatic              16080384                74.60 ns/op           24 B/op          1 allocs/op
BenchmarkEcho_GPlusStatic                6225674               189.0 ns/op            40 B/op          2 allocs/op
BenchmarkGin_GPlusStatic                 9650610               127.6 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_GPlusStatic          1268251               991.8 ns/op           808 B/op          8 allocs/op
BenchmarkHttpRouter_GPlusStatic         13866670                77.37 ns/op           24 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        13361671                87.65 ns/op           24 B/op          1 allocs/op
BenchmarkServeMux_GPlusStatic            6368907               184.0 ns/op            24 B/op          1 allocs/op
BenchmarkChi_GPlusParam                  2199004               541.4 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_GPlusParam                6617103               177.3 ns/op            88 B/op          2 allocs/op
BenchmarkEcho_GPlusParam                 5234647               220.6 ns/op            40 B/op          2 allocs/op
BenchmarkGin_GPlusParam                  6988406               173.8 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_GPlusParam            567523              1978 ns/op            1112 B/op          9 allocs/op
BenchmarkHttpRouter_GPlusParam           6576224               166.0 ns/op            88 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2832715               433.5 ns/op           376 B/op          4 allocs/op
BenchmarkServeMux_GPlusParam             3151465               384.0 ns/op            40 B/op          2 allocs/op
BenchmarkChi_GPlus2Params                1915664               611.7 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_GPlus2Params              5191447               213.1 ns/op            88 B/op          2 allocs/op
BenchmarkEcho_GPlus2Params               4755480               247.8 ns/op            40 B/op          2 allocs/op
BenchmarkGin_GPlus2Params                6197532               192.5 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_GPlus2Params          289845              3878 ns/op            1128 B/op          9 allocs/op
BenchmarkHttpRouter_GPlus2Params         6111842               183.1 ns/op            88 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        2043349               570.7 ns/op           408 B/op          5 allocs/op
BenchmarkServeMux_GPlus2Params           1992519               593.6 ns/op            72 B/op          3 allocs/op
BenchmarkChi_GPlusAll                     149733              7779 ns/op            4680 B/op         39 allocs/op
BenchmarkDenco_GPlusAll                   413450              2857 ns/op             984 B/op         24 allocs/op
BenchmarkEcho_GPlusAll                    312202              3451 ns/op             520 B/op         26 allocs/op
BenchmarkGin_GPlusAll                     634213              2475 ns/op             312 B/op         13 allocs/op
BenchmarkGorillaMux_GPlusAll               39432             32039 ns/op           13928 B/op        115 allocs/op
BenchmarkHttpRouter_GPlusAll              490444              2377 ns/op             952 B/op         24 allocs/op
BenchmarkHttpTreeMux_GPlusAll             179402              6105 ns/op            4344 B/op         51 allocs/op
BenchmarkServeMux_GPlusAll                195734              5878 ns/op             648 B/op         29 allocs/op
BenchmarkChi_ParseStatic                 2448097               485.8 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_ParseStatic              13748854                84.41 ns/op           24 B/op          1 allocs/op
BenchmarkEcho_ParseStatic                5921918               198.9 ns/op            40 B/op          2 allocs/op
BenchmarkGin_ParseStatic                 8445585               123.9 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_ParseStatic           848359              1198 ns/op             808 B/op          8 allocs/op
BenchmarkHttpRouter_ParseStatic         14255942                82.43 ns/op           24 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseStatic        10855117               112.6 ns/op            24 B/op          1 allocs/op
BenchmarkServeMux_ParseStatic            5165193               218.4 ns/op            24 B/op          1 allocs/op
BenchmarkChi_ParseParam                  2275634               534.1 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_ParseParam                6452120               186.4 ns/op            88 B/op          2 allocs/op
BenchmarkEcho_ParseParam                 5481141               210.8 ns/op            40 B/op          2 allocs/op
BenchmarkGin_ParseParam                  8367566               148.3 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_ParseParam            736794              1452 ns/op            1112 B/op          9 allocs/op
BenchmarkHttpRouter_ParseParam           7846512               160.9 ns/op            88 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseParam          2636179               422.6 ns/op           376 B/op          4 allocs/op
BenchmarkServeMux_ParseParam             2998899               368.6 ns/op            40 B/op          2 allocs/op
BenchmarkChi_Parse2Params                2081517               546.0 ns/op           360 B/op          3 allocs/op
BenchmarkDenco_Parse2Params              5762701               198.8 ns/op            88 B/op          2 allocs/op
BenchmarkEcho_Parse2Params               5430128               217.5 ns/op            40 B/op          2 allocs/op
BenchmarkGin_Parse2Params                6633669               161.4 ns/op            24 B/op          1 allocs/op
BenchmarkGorillaMux_Parse2Params          730944              1642 ns/op            1128 B/op          9 allocs/op
BenchmarkHttpRouter_Parse2Params         7051012               159.5 ns/op            88 B/op          2 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2302215               516.8 ns/op           408 B/op          5 allocs/op
BenchmarkServeMux_Parse2Params           2609415               446.3 ns/op            72 B/op          3 allocs/op
BenchmarkChi_ParseAll                      81310             14156 ns/op            9360 B/op         78 allocs/op
BenchmarkDenco_ParseAll                   249819              4319 ns/op            1552 B/op         42 allocs/op
BenchmarkEcho_ParseAll                    177633              5888 ns/op            1040 B/op         52 allocs/op
BenchmarkGin_ParseAll                     253527              4497 ns/op             624 B/op         26 allocs/op
BenchmarkGorillaMux_ParseAll               21876             53405 ns/op           25920 B/op        224 allocs/op
BenchmarkHttpRouter_ParseAll              292665              3661 ns/op            1264 B/op         42 allocs/op
BenchmarkHttpTreeMux_ParseAll             124321              8625 ns/op            6352 B/op         77 allocs/op
BenchmarkServeMux_ParseAll                123384              9259 ns/op             976 B/op         45 allocs/op
BenchmarkHttpServeMux_StaticAll            23180             50596 ns/op            3768 B/op        157 allocs/op
BenchmarkChi_StaticAll                     13984             85218 ns/op           56521 B/op        471 allocs/op
BenchmarkDenco_StaticAll                   75424             15374 ns/op            3768 B/op        157 allocs/op
BenchmarkEcho_StaticAll                    29372             40253 ns/op            6280 B/op        314 allocs/op
BenchmarkGin_StaticAll                     41338             28117 ns/op            3768 B/op        157 allocs/op
BenchmarkGorillaMux_StaticAll               1796            574358 ns/op          126859 B/op       1256 allocs/op
BenchmarkHttpRouter_StaticAll              57958             19436 ns/op            3768 B/op        157 allocs/op
BenchmarkHttpTreeMux_StaticAll             59157             19613 ns/op            3768 B/op        157 allocs/op
BenchmarkServeMux_StaticAll                22916             51446 ns/op            3768 B/op        157 allocs/op
PASS
ok      github.com/julienschmidt/go-http-routing-benchmark      235.152s