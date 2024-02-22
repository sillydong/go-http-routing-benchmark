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
   Denco: 9928 Bytes
   Echo: 97368 Bytes
   Gin: 34360 Bytes
   GorillaMux: 599512 Bytes
   HttpRouter: 21696 Bytes
   HttpTreeMux: 73464 Bytes
   ServeMux: 76840 Bytes

goos: darwin
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkChi_Param                       2087209               599.9 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_Param                     5399943               206.6 ns/op            96 B/op          3 allocs/op
BenchmarkEcho_Param                      4813693               249.5 ns/op            72 B/op          3 allocs/op
BenchmarkGin_Param                       6306952               188.0 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_Param                 732076              1434 ns/op            1136 B/op         10 allocs/op
BenchmarkHttpRouter_Param                6761620               173.0 ns/op            96 B/op          3 allocs/op
BenchmarkHttpTreeMux_Param               2745614               449.6 ns/op           400 B/op          5 allocs/op
BenchmarkServeMux_Param                  3406209               383.3 ns/op            64 B/op          3 allocs/op
BenchmarkChi_Param5                      1515321               751.0 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_Param5                    3824071               302.8 ns/op           224 B/op          3 allocs/op
BenchmarkEcho_Param5                     3883113               296.1 ns/op            72 B/op          3 allocs/op
BenchmarkGin_Param5                      4815877               247.7 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_Param5                494395              2225 ns/op            1200 B/op         10 allocs/op
BenchmarkHttpRouter_Param5               3957816               320.9 ns/op           224 B/op          3 allocs/op
BenchmarkHttpTreeMux_Param5              1226887               918.6 ns/op           624 B/op          8 allocs/op
BenchmarkServeMux_Param5                 1530224               789.2 ns/op           288 B/op          6 allocs/op
BenchmarkChi_Param20                      987591              1483 ns/op             384 B/op          4 allocs/op
BenchmarkDenco_Param20                   1402417               841.6 ns/op           768 B/op          3 allocs/op
BenchmarkEcho_Param20                    2031148               572.1 ns/op            72 B/op          3 allocs/op
BenchmarkGin_Param20                     2747575               423.5 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_Param20               232614              4953 ns/op            3307 B/op         12 allocs/op
BenchmarkHttpRouter_Param20              1499593               836.7 ns/op           768 B/op          3 allocs/op
BenchmarkHttpTreeMux_Param20              277723              4307 ns/op            3244 B/op         12 allocs/op
BenchmarkServeMux_Param20                 588363              1866 ns/op            1056 B/op          8 allocs/op
BenchmarkChi_ParamWrite                  2098473               578.2 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_ParamWrite                4077717               287.3 ns/op           120 B/op          5 allocs/op
BenchmarkEcho_ParamWrite                 3753741               330.8 ns/op            96 B/op          5 allocs/op
BenchmarkGin_ParamWrite                  4429651               267.6 ns/op            72 B/op          4 allocs/op
BenchmarkGorillaMux_ParamWrite            806547              1475 ns/op            1160 B/op         12 allocs/op
BenchmarkHttpRouter_ParamWrite           4777334               256.1 ns/op           120 B/op          5 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2266411               566.7 ns/op           424 B/op          7 allocs/op
BenchmarkServeMux_ParamWrite             2324829               493.2 ns/op            88 B/op          5 allocs/op
BenchmarkDenco_ExtraStatic               1256125               865.5 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_ExtraStatic                 845119              1269 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraStatic                 3692772               325.7 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraStatic          1349493               892.4 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraStatic         1402136               847.8 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraStatic             688623              1745 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_ExtraParam                1431872               829.9 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_ExtraParam                  849496              1221 ns/op             896 B/op         10 allocs/op
BenchmarkGin_ExtraParam                  3656671               314.4 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_ExtraParam           1213384               996.0 ns/op           896 B/op          8 allocs/op
BenchmarkHttpTreeMux_ExtraParam          1389002               839.9 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_ExtraParam              689226              1769 ns/op            1008 B/op         17 allocs/op
BenchmarkDenco_Extra2Params              1454143               818.3 ns/op           880 B/op          8 allocs/op
BenchmarkEcho_Extra2Params                880195              1225 ns/op             896 B/op         10 allocs/op
BenchmarkGin_Extra2Params                3764235               312.3 ns/op           400 B/op          2 allocs/op
BenchmarkGorillaMux_Extra2Params         1200806              1028 ns/op             896 B/op          8 allocs/op
BenchmarkHttpTreeMux_Extra2Params        1394052               850.3 ns/op           848 B/op          7 allocs/op
BenchmarkServeMux_Extra2Params            556333              1828 ns/op            1024 B/op         17 allocs/op
BenchmarkDenco_ExtraAll                  3385212               347.2 ns/op           160 B/op          5 allocs/op
BenchmarkEcho_ExtraAll                   2348588               506.0 ns/op           144 B/op          6 allocs/op
BenchmarkGin_ExtraAll                    2245033               527.0 ns/op           224 B/op          6 allocs/op
BenchmarkHttpTreeMux_ExtraAll            1840358               639.9 ns/op           448 B/op          7 allocs/op
BenchmarkServeMux_ExtraAll               1625546               716.5 ns/op           112 B/op          5 allocs/op
BenchmarkChi_GithubStatic                2272874               518.9 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_GithubStatic              9078430               124.2 ns/op            64 B/op          2 allocs/op
BenchmarkEcho_GithubStatic               4503680               258.9 ns/op            72 B/op          3 allocs/op
BenchmarkGin_GithubStatic                6295491               184.9 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_GithubStatic          339248              3225 ns/op             832 B/op          9 allocs/op
BenchmarkHttpRouter_GithubStatic         8178879               137.6 ns/op            64 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubStatic        8064549               141.6 ns/op            48 B/op          2 allocs/op
BenchmarkServeMux_GithubStatic           3964291               299.6 ns/op            48 B/op          2 allocs/op
BenchmarkChi_GithubParam                 1762029               669.5 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_GithubParam               4013496               292.4 ns/op           192 B/op          3 allocs/op
BenchmarkEcho_GithubParam                3711766               311.5 ns/op            72 B/op          3 allocs/op
BenchmarkGin_GithubParam                 4829560               239.0 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_GithubParam           213649              5012 ns/op            1152 B/op         10 allocs/op
BenchmarkHttpRouter_GithubParam          4495662               260.2 ns/op           160 B/op          3 allocs/op
BenchmarkHttpTreeMux_GithubParam         1947478               613.3 ns/op           432 B/op          6 allocs/op
BenchmarkServeMux_GithubParam            1856818               644.5 ns/op            96 B/op          4 allocs/op
BenchmarkChi_GithubAll                      7186            145752 ns/op           77953 B/op        812 allocs/op
BenchmarkDenco_GithubAll                   19516             61068 ns/op           33216 B/op        573 allocs/op
BenchmarkEcho_GithubAll                    17742             67006 ns/op           14616 B/op        609 allocs/op
BenchmarkGin_GithubAll                     24144             49040 ns/op            9744 B/op        406 allocs/op
BenchmarkGorillaMux_GithubAll                464           2561909 ns/op          222422 B/op       1994 allocs/op
BenchmarkHttpRouter_GithubAll              22196             52796 ns/op           26784 B/op        573 allocs/op
BenchmarkHttpTreeMux_GithubAll              9096            116541 ns/op           75600 B/op       1077 allocs/op
BenchmarkServeMux_GithubAll                 9366            131737 ns/op           19488 B/op        743 allocs/op
BenchmarkChi_GPlusStatic                 2370885               491.2 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_GPlusStatic               9875689               114.4 ns/op            64 B/op          2 allocs/op
BenchmarkEcho_GPlusStatic                4825885               235.0 ns/op            72 B/op          3 allocs/op
BenchmarkGin_GPlusStatic                 6975657               169.8 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_GPlusStatic          1250350               968.4 ns/op           832 B/op          9 allocs/op
BenchmarkHttpRouter_GPlusStatic          9133286               121.7 ns/op            64 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusStatic         8896178               130.7 ns/op            48 B/op          2 allocs/op
BenchmarkServeMux_GPlusStatic            5045293               234.2 ns/op            48 B/op          2 allocs/op
BenchmarkChi_GPlusParam                  2095414               566.9 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_GPlusParam                5155281               223.7 ns/op           128 B/op          3 allocs/op
BenchmarkEcho_GPlusParam                 4406540               263.0 ns/op            72 B/op          3 allocs/op
BenchmarkGin_GPlusParam                  5857400               199.7 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_GPlusParam            642459              1860 ns/op            1136 B/op         10 allocs/op
BenchmarkHttpRouter_GPlusParam           5634120               202.9 ns/op           128 B/op          3 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2598680               448.6 ns/op           400 B/op          5 allocs/op
BenchmarkServeMux_GPlusParam             2820890               423.2 ns/op            64 B/op          3 allocs/op
BenchmarkChi_GPlus2Params                1888552               626.2 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_GPlus2Params              4418935               264.9 ns/op           128 B/op          3 allocs/op
BenchmarkEcho_GPlus2Params               3985641               296.3 ns/op            72 B/op          3 allocs/op
BenchmarkGin_GPlus2Params                5215785               219.2 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_GPlus2Params          266562              3773 ns/op            1152 B/op         10 allocs/op
BenchmarkHttpRouter_GPlus2Params         5081314               229.6 ns/op           128 B/op          3 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        2013682               608.5 ns/op           432 B/op          6 allocs/op
BenchmarkServeMux_GPlus2Params           1858538               643.9 ns/op            96 B/op          4 allocs/op
BenchmarkChi_GPlusAll                     140286              7964 ns/op            4992 B/op         52 allocs/op
BenchmarkDenco_GPlusAll                   306394              3614 ns/op            1504 B/op         37 allocs/op
BenchmarkEcho_GPlusAll                    257892              4116 ns/op             936 B/op         39 allocs/op
BenchmarkGin_GPlusAll                     391784              2990 ns/op             624 B/op         26 allocs/op
BenchmarkGorillaMux_GPlusAll               37710             30140 ns/op           14240 B/op        128 allocs/op
BenchmarkHttpRouter_GPlusAll              391363              2950 ns/op            1472 B/op         37 allocs/op
BenchmarkHttpTreeMux_GPlusAll             181052              6186 ns/op            4656 B/op         64 allocs/op
BenchmarkServeMux_GPlusAll                179115              6467 ns/op             960 B/op         42 allocs/op
BenchmarkChi_ParseStatic                 2135661               538.2 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_ParseStatic               9547832               127.9 ns/op            64 B/op          2 allocs/op
BenchmarkEcho_ParseStatic                4696572               253.5 ns/op            72 B/op          3 allocs/op
BenchmarkGin_ParseStatic                 6792006               182.0 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_ParseStatic           947336              1370 ns/op             832 B/op          9 allocs/op
BenchmarkHttpRouter_ParseStatic          8730080               135.2 ns/op            64 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseStatic         7743080               142.7 ns/op            48 B/op          2 allocs/op
BenchmarkServeMux_ParseStatic            4363630               280.7 ns/op            48 B/op          2 allocs/op
BenchmarkChi_ParseParam                  1888156               636.4 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_ParseParam                5060760               234.2 ns/op           128 B/op          3 allocs/op
BenchmarkEcho_ParseParam                 4526246               257.1 ns/op            72 B/op          3 allocs/op
BenchmarkGin_ParseParam                  6206336               188.8 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_ParseParam            700513              1466 ns/op            1136 B/op         10 allocs/op
BenchmarkHttpRouter_ParseParam           5569603               215.2 ns/op           128 B/op          3 allocs/op
BenchmarkHttpTreeMux_ParseParam          2131436               532.1 ns/op           400 B/op          5 allocs/op
BenchmarkServeMux_ParseParam             2878530               398.2 ns/op            64 B/op          3 allocs/op
BenchmarkChi_Parse2Params                1919667               617.8 ns/op           384 B/op          4 allocs/op
BenchmarkDenco_Parse2Params              4662492               250.7 ns/op           128 B/op          3 allocs/op
BenchmarkEcho_Parse2Params               4319011               281.3 ns/op            72 B/op          3 allocs/op
BenchmarkGin_Parse2Params                5572830               228.2 ns/op            48 B/op          2 allocs/op
BenchmarkGorillaMux_Parse2Params          585163              2022 ns/op            1152 B/op         10 allocs/op
BenchmarkHttpRouter_Parse2Params         5040607               235.6 ns/op           128 B/op          3 allocs/op
BenchmarkHttpTreeMux_Parse2Params        1814935               719.6 ns/op           432 B/op          6 allocs/op
BenchmarkServeMux_Parse2Params           2159425               570.2 ns/op            96 B/op          4 allocs/op
BenchmarkChi_ParseAll                      56500             18674 ns/op            9984 B/op        104 allocs/op
BenchmarkDenco_ParseAll                   169123              6418 ns/op            2592 B/op         68 allocs/op
BenchmarkEcho_ParseAll                    131928              8426 ns/op            1872 B/op         78 allocs/op
BenchmarkGin_ParseAll                     204438              5929 ns/op            1248 B/op         52 allocs/op
BenchmarkGorillaMux_ParseAll               16794             64803 ns/op           26544 B/op        250 allocs/op
BenchmarkHttpRouter_ParseAll              192469              5910 ns/op            2304 B/op         68 allocs/op
BenchmarkHttpTreeMux_ParseAll              95234             12593 ns/op            6976 B/op        103 allocs/op
BenchmarkServeMux_ParseAll                 78985             13274 ns/op            1600 B/op         71 allocs/op
BenchmarkHttpServeMux_StaticAll            15924             67588 ns/op            7536 B/op        314 allocs/op
BenchmarkChi_StaticAll                     10000            115379 ns/op           60289 B/op        628 allocs/op
BenchmarkDenco_StaticAll                   46022             26337 ns/op           10048 B/op        314 allocs/op
BenchmarkEcho_StaticAll                    23462             55234 ns/op           11304 B/op        471 allocs/op
BenchmarkGin_StaticAll                     30092             37418 ns/op            7536 B/op        314 allocs/op
BenchmarkGorillaMux_StaticAll               1598            656923 ns/op          130626 B/op       1413 allocs/op
BenchmarkHttpRouter_StaticAll              39633             29319 ns/op           10048 B/op        314 allocs/op
BenchmarkHttpTreeMux_StaticAll             34359             32223 ns/op           10048 B/op        314 allocs/op
BenchmarkServeMux_StaticAll                18027             66214 ns/op            7536 B/op        314 allocs/op
PASS
ok      github.com/julienschmidt/go-http-routing-benchmark      242.038s