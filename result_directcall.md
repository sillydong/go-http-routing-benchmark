#ExtraAPI Routes: 2
   Denco: 848 Bytes
   Echo: 3600 Bytes
   Gin: 1176 Bytes
   GorillaMux: 6632 Bytes
   HttpTreeMux: 1688 Bytes
   ServeMux: 1800 Bytes

#GithubAPI Routes: 203
   Chi: 94888 Bytes
   Denco: 37912 Bytes
   Echo: 123384 Bytes
   Gin: 58792 Bytes
   GorillaMux: 1319696 Bytes
   HttpRouter: 37088 Bytes
   HttpTreeMux: 78800 Bytes
   ServeMux: 121408 Bytes

#GPlusAPI Routes: 13
   Chi: 8008 Bytes
   Denco: 3224 Bytes
   Echo: 11328 Bytes
   Gin: 4544 Bytes
   GorillaMux: 68000 Bytes
   HttpRouter: 2760 Bytes
   HttpTreeMux: 7440 Bytes
   ServeMux: 7392 Bytes

#ParseAPI Routes: 26
   Chi: 9656 Bytes
   Denco: 4080 Bytes
   Echo: 14192 Bytes
   Gin: 7864 Bytes
   GorillaMux: 105384 Bytes
   HttpRouter: 5024 Bytes
   HttpTreeMux: 7848 Bytes
   ServeMux: 13632 Bytes

#Static Routes: 157
   HttpServeMux: 72056 Bytes
   Chi: 83160 Bytes
   Denco: 9912 Bytes
   Echo: 97352 Bytes
   Gin: 34344 Bytes
   GorillaMux: 599496 Bytes
   HttpRouter: 21680 Bytes
   HttpTreeMux: 73448 Bytes
   ServeMux: 76824 Bytes

goos: darwin
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkChi_Param                       2681167               441.4 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_Param                    12943513                88.59 ns/op           32 B/op          1 allocs/op
BenchmarkEcho_Param                     25348876                47.28 ns/op            0 B/op          0 allocs/op
BenchmarkGin_Param                      22900412                51.12 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_Param                 807553              1309 ns/op            1088 B/op          8 allocs/op
BenchmarkHttpRouter_Param               16721287                68.10 ns/op           32 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param               3440352               340.8 ns/op           352 B/op          3 allocs/op
BenchmarkServeMux_Param                  5384269               222.7 ns/op            16 B/op          1 allocs/op
BenchmarkChi_Param5                      1841041               637.0 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_Param5                    5730218               206.4 ns/op           160 B/op          1 allocs/op
BenchmarkEcho_Param5                    12156172                97.03 ns/op            0 B/op          0 allocs/op
BenchmarkGin_Param5                     12408232                93.75 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_Param5                604698              1998 ns/op            1152 B/op          8 allocs/op
BenchmarkHttpRouter_Param5               6554326               177.5 ns/op           160 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5              1521284               782.0 ns/op           576 B/op          6 allocs/op
BenchmarkServeMux_Param5                 1858088               622.4 ns/op           240 B/op          4 allocs/op
BenchmarkChi_Param20                     1000000              1248 ns/op             336 B/op          2 allocs/op
BenchmarkDenco_Param20                   1798250               653.2 ns/op           704 B/op          1 allocs/op
BenchmarkEcho_Param20                    4156710               276.5 ns/op             0 B/op          0 allocs/op
BenchmarkGin_Param20                     4891606               235.3 ns/op             0 B/op          0 allocs/op
BenchmarkGorillaMux_Param20               245829              4499 ns/op            3259 B/op         10 allocs/op
BenchmarkHttpRouter_Param20              2097354               574.8 ns/op           704 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20              281034              3763 ns/op            3195 B/op         10 allocs/op
BenchmarkServeMux_Param20                 662598              1652 ns/op            1008 B/op          6 allocs/op
BenchmarkChi_ParamWrite                  2544122               454.6 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_ParamWrite               11226268                97.91 ns/op           32 B/op          1 allocs/op
BenchmarkEcho_ParamWrite                 9916639               115.1 ns/op             8 B/op          1 allocs/op
BenchmarkGin_ParamWrite                 16503542                61.98 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_ParamWrite            832291              1435 ns/op            1088 B/op          8 allocs/op
BenchmarkHttpRouter_ParamWrite          14011472                81.29 ns/op           32 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite          3128733               371.6 ns/op           352 B/op          3 allocs/op
BenchmarkServeMux_ParamWrite             5119350               227.9 ns/op            16 B/op          1 allocs/op
BenchmarkDenco_ExtraStatic               1398222               839.2 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_ExtraStatic                 914376              1284 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraStatic                 3217419               349.4 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraStatic          1326595               959.1 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraStatic         1328217               911.5 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraStatic             583723              1819 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_ExtraParam                1000000              1008 ns/op             880 B/op          8 allocs/op
BenchmarkEcho_ExtraParam                  751464              1450 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraParam                  3159781               414.1 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraParam           1000000              1012 ns/op             896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraParam          1327606               898.8 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraParam              625386              1794 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_Extra2Params              1307618               884.5 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_Extra2Params                843258              1314 ns/op             896 B/op         10 allocs/op
BenchmarkGin_Extra2Params                3536451               336.3 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_Extra2Params         1000000              1077 ns/op             896 B/op          8 allocs/op
BenchmarkHttpTreeMux_Extra2Params        1290243               940.8 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_Extra2Params            555848              1974 ns/op            1024 B/op         17 allocs/op
BenchmarkDenco_ExtraAll                  8323833               149.0 ns/op            32 B/op          1 allocs/op
BenchmarkEcho_ExtraAll                  10308734               113.1 ns/op             0 B/op          0 allocs/op
BenchmarkGin_ExtraAll                    3623834               336.4 ns/op           128 B/op          2 allocs/op
BenchmarkHttpTreeMux_ExtraAll            2692083               426.8 ns/op           352 B/op          3 allocs/op
BenchmarkServeMux_ExtraAll               2707315               442.9 ns/op            16 B/op          1 allocs/op
BenchmarkChi_GithubStatic                2533878               435.1 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_GithubStatic             54933049                24.07 ns/op            0 B/op          0 allocs/op
BenchmarkEcho_GithubStatic              17613999                67.51 ns/op            0 B/op          0 allocs/op
BenchmarkGin_GithubStatic               20252904                56.75 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_GithubStatic          309649              3257 ns/op             784 B/op          7 allocs/op
BenchmarkHttpRouter_GithubStatic        36573325                32.17 ns/op            0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       25696156                41.93 ns/op            0 B/op          0 allocs/op
BenchmarkServeMux_GithubStatic           6504540               175.6 ns/op             0 B/op          0 allocs/op
BenchmarkChi_GithubParam                 1960807               596.0 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_GithubParam               5667382               212.2 ns/op           128 B/op          1 allocs/op
BenchmarkEcho_GithubParam               10488637               116.0 ns/op             0 B/op          0 allocs/op
BenchmarkGin_GithubParam                10637775               105.1 ns/op             0 B/op          0 allocs/op
BenchmarkGorillaMux_GithubParam           225250              4832 ns/op            1104 B/op          8 allocs/op
BenchmarkHttpRouter_GithubParam          7497891               156.6 ns/op            96 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam         2193717               526.6 ns/op           384 B/op          4 allocs/op
BenchmarkServeMux_GithubParam            2280853               509.0 ns/op            48 B/op          2 allocs/op
BenchmarkChi_GithubAll                      8726            132056 ns/op           68209 B/op        406 allocs/op
BenchmarkDenco_GithubAll                   26930             41546 ns/op           20224 B/op        167 allocs/op
BenchmarkEcho_GithubAll                    46532             25268 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GithubAll                     59317             20494 ns/op               0 B/op          0 allocs/op
BenchmarkGorillaMux_GithubAll                436           2522831 ns/op          212677 B/op       1588 allocs/op
BenchmarkHttpRouter_GithubAll              39760             30832 ns/op           13792 B/op        167 allocs/op
BenchmarkHttpTreeMux_GithubAll             12453             98240 ns/op           65856 B/op        671 allocs/op
BenchmarkServeMux_GithubAll                10000            112502 ns/op            9744 B/op        337 allocs/op
BenchmarkChi_GPlusStatic                 2965288               436.8 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_GPlusStatic              78656700                15.24 ns/op            0 B/op          0 allocs/op
BenchmarkEcho_GPlusStatic               22236816                48.13 ns/op            0 B/op          0 allocs/op
BenchmarkGin_GPlusStatic                25424128                45.40 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_GPlusStatic          1251261               922.9 ns/op           784 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusStatic         68929677                15.97 ns/op            0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        48785482                24.24 ns/op            0 B/op          0 allocs/op
BenchmarkServeMux_GPlusStatic           11517031               101.8 ns/op             0 B/op          0 allocs/op
BenchmarkChi_GPlusParam                  2429389               501.2 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_GPlusParam                9143235               122.4 ns/op            64 B/op          1 allocs/op
BenchmarkEcho_GPlusParam                17181405                67.40 ns/op            0 B/op          0 allocs/op
BenchmarkGin_GPlusParam                 18251988                65.98 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_GPlusParam            586773              1900 ns/op            1088 B/op          8 allocs/op
BenchmarkHttpRouter_GPlusParam          11441790               100.5 ns/op            64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam          3254619               380.0 ns/op           352 B/op          3 allocs/op
BenchmarkServeMux_GPlusParam             3900388               298.6 ns/op            16 B/op          1 allocs/op
BenchmarkChi_GPlus2Params                2135352               578.4 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_GPlus2Params              6805082               169.1 ns/op            64 B/op          1 allocs/op
BenchmarkEcho_GPlus2Params              12572076                95.19 ns/op            0 B/op          0 allocs/op
BenchmarkGin_GPlus2Params               13718084                83.36 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_GPlus2Params          295299              3702 ns/op            1104 B/op          8 allocs/op
BenchmarkHttpRouter_GPlus2Params         9435702               124.4 ns/op            64 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        2372356               503.5 ns/op           384 B/op          4 allocs/op
BenchmarkServeMux_GPlus2Params           2220436               530.9 ns/op            48 B/op          2 allocs/op
BenchmarkChi_GPlusAll                     157998              7068 ns/op            4368 B/op         26 allocs/op
BenchmarkDenco_GPlusAll                   639408              1782 ns/op             672 B/op         11 allocs/op
BenchmarkEcho_GPlusAll                   1000000              1085 ns/op               0 B/op          0 allocs/op
BenchmarkGin_GPlusAll                    1341753               872.2 ns/op             0 B/op          0 allocs/op
BenchmarkGorillaMux_GPlusAll               40370             28459 ns/op           13616 B/op        102 allocs/op
BenchmarkHttpRouter_GPlusAll              887737              1323 ns/op             640 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll             223813              4818 ns/op            4032 B/op         38 allocs/op
BenchmarkServeMux_GPlusAll                257348              4678 ns/op             336 B/op         16 allocs/op
BenchmarkChi_ParseStatic                 2986364               396.8 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_ParseStatic              65839329                17.45 ns/op            0 B/op          0 allocs/op
BenchmarkEcho_ParseStatic               22403420                48.50 ns/op            0 B/op          0 allocs/op
BenchmarkGin_ParseStatic                26425470                44.97 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_ParseStatic           964310              1141 ns/op             784 B/op          7 allocs/op
BenchmarkHttpRouter_ParseStatic         66596787                16.98 ns/op            0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        31576119                37.30 ns/op            0 B/op          0 allocs/op
BenchmarkServeMux_ParseStatic            8802049               129.5 ns/op             0 B/op          0 allocs/op
BenchmarkChi_ParseParam                  2635974               451.9 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_ParseParam                9808893               115.7 ns/op            64 B/op          1 allocs/op
BenchmarkEcho_ParseParam                19008793                60.95 ns/op            0 B/op          0 allocs/op
BenchmarkGin_ParseParam                 21802174                55.27 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_ParseParam            739580              1380 ns/op            1088 B/op          8 allocs/op
BenchmarkHttpRouter_ParseParam          12395030                91.84 ns/op           64 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam          3220826               428.8 ns/op           352 B/op          3 allocs/op
BenchmarkServeMux_ParseParam             4756806               242.5 ns/op            16 B/op          1 allocs/op
BenchmarkChi_Parse2Params                2223050               603.7 ns/op           336 B/op          2 allocs/op
BenchmarkDenco_Parse2Params              7675392               150.2 ns/op            64 B/op          1 allocs/op
BenchmarkEcho_Parse2Params              16150070                73.60 ns/op            0 B/op          0 allocs/op
BenchmarkGin_Parse2Params               18628184                69.07 ns/op            0 B/op          0 allocs/op
BenchmarkGorillaMux_Parse2Params          587480              1760 ns/op            1104 B/op          8 allocs/op
BenchmarkHttpRouter_Parse2Params        10692589               109.5 ns/op            64 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2392962               503.2 ns/op           384 B/op          4 allocs/op
BenchmarkServeMux_Parse2Params           3025788               407.2 ns/op            48 B/op          2 allocs/op
BenchmarkChi_ParseAll                      88462             13837 ns/op            8736 B/op         52 allocs/op
BenchmarkDenco_ParseAll                   442122              2870 ns/op             928 B/op         16 allocs/op
BenchmarkEcho_ParseAll                    590954              2012 ns/op               0 B/op          0 allocs/op
BenchmarkGin_ParseAll                     661280              1736 ns/op               0 B/op          0 allocs/op
BenchmarkGorillaMux_ParseAll               21408             55583 ns/op           25296 B/op        198 allocs/op
BenchmarkHttpRouter_ParseAll              555274              1991 ns/op             640 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll             156620              7579 ns/op            5728 B/op         51 allocs/op
BenchmarkServeMux_ParseAll                162368              7586 ns/op             352 B/op         19 allocs/op
BenchmarkHttpServeMux_StaticAll            30514             40103 ns/op               0 B/op          0 allocs/op
BenchmarkChi_StaticAll                     15034             94904 ns/op           52752 B/op        314 allocs/op
BenchmarkDenco_StaticAll                  318667              4232 ns/op               0 B/op          0 allocs/op
BenchmarkEcho_StaticAll                    68386             17845 ns/op               0 B/op          0 allocs/op
BenchmarkGin_StaticAll                     80202             15238 ns/op               0 B/op          0 allocs/op
BenchmarkGorillaMux_StaticAll               1749            630743 ns/op          123090 B/op       1099 allocs/op
BenchmarkHttpRouter_StaticAll             144763              8327 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll            147956              8127 ns/op               0 B/op          0 allocs/op
BenchmarkServeMux_StaticAll                28986             39718 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/julienschmidt/go-http-routing-benchmark      314.139s