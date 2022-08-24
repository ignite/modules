package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/simulation"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"mint/InflationRateChange", "InflationRateChange", "\"0.230000000000000000\"", "mint"},
		{"mint/InflationMax", "InflationMax", "\"0.200000000000000000\"", "mint"},
		{"mint/InflationMin", "InflationMin", "\"0.070000000000000000\"", "mint"},
		{"mint/GoalBonded", "GoalBonded", "\"0.670000000000000000\"", "mint"},
		{"mint/DistributionProportions", "DistributionProportions", "{\"staking\":\"0.250000000000000000\",\"funded_addresses\":\"0.210000000000000000\",\"community_pool\":\"0.540000000000000000\"}", "mint"},
		{"mint/FundedAddresses", "FundedAddresses", "[{\"address\":\"cosmos1vt6cdw69vhrm59v7tfkqm03rd4d6rlnqfwkq8w\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos16n6wfzxxexnn96jzh9narl6y8dtmwejr96p9p5\",\"weight\":\"0.010346444681782616\"},{\"address\":\"cosmos1dcg2d7kt092xdc60me5zkz33uj4rs9ukmkafvq\",\"weight\":\"0.012317549730149727\"},{\"address\":\"cosmos1repxmyy9mx4xq4fajgjxaahaw0yjlmh5uk64m6\",\"weight\":\"0.019897379547884085\"},{\"address\":\"cosmos1n6wnkglm8m3sxr2f7g9rmv0u6ekjc7e4t5ta2f\",\"weight\":\"0.006141422857328581\"},{\"address\":\"cosmos17q32k8jlkac2yapq7mxf5lak0huh06p0g6sqmy\",\"weight\":\"0.007683902997839140\"},{\"address\":\"cosmos1ls6shhp8swr5ards3fp2lst5ftysse25dqmk2x\",\"weight\":\"0.006856440371419632\"},{\"address\":\"cosmos1xq7lhaqxvkkljqp73xt5h3uqwe0f3gj2pwc0xv\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos14g2upxq6hu49lt544a0jrs6j376r334h9w6tus\",\"weight\":\"0.016662840271533599\"},{\"address\":\"cosmos1kcu2rlxffgpycjn0e2fzs9de3xvdw8gf5n8em0\",\"weight\":\"0.016406308110665711\"},{\"address\":\"cosmos14lh203x67r7a7vd9as2vqk7uc08uumlpkvuu5g\",\"weight\":\"0.008878619785299237\"},{\"address\":\"cosmos1dec5rwvvcygyx6eg0hjv35vxfr58hawx7xvekw\",\"weight\":\"0.016267509677487911\"},{\"address\":\"cosmos1tlxztuxxs8rxmxv97247vsn5cafkv9sgplhmag\",\"weight\":\"0.017180571755071666\"},{\"address\":\"cosmos17ktwhac3c75k7vr62lsvmtx8u9cs6t8g0t8ts6\",\"weight\":\"0.011536944938043694\"},{\"address\":\"cosmos1mtjrpensxply0kstkgfdjwnjakvvwdnhrg2at0\",\"weight\":\"0.010610596013385205\"},{\"address\":\"cosmos1uqv7xkqey6sv8c3ta8aqn08pd97qg5dyqrf6a4\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1hc08vrfk2fyhe2wkdep4uswcqtepy73rxpe5me\",\"weight\":\"0.010676596986934073\"},{\"address\":\"cosmos1rl5exulw99m0mx43c09m0k4ydmzdurxu8c4mpm\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos12yeysqnpx43uddpuve72xs4w9ghnhkf2mggrhz\",\"weight\":\"0.005312909508254276\"},{\"address\":\"cosmos1r557dv93c3qtzl4ws5ne9rgk6pp3wqchycem0q\",\"weight\":\"0.010491849933835756\"},{\"address\":\"cosmos1uhp7u3jflawdn64ehe0g0gvlnrtz9uvzzc9uvk\",\"weight\":\"0.006390850270813198\"},{\"address\":\"cosmos13uh4804tdnrffhvxajlce7ru9wn84lgl94cs6v\",\"weight\":\"0.019172715206289138\"},{\"address\":\"cosmos1nulylqzrde5utkjrz0twh6e2kp2csat2f8n9je\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1906mhe8x2xge0gm4neczdv7wvs0fzxnfh05j7r\",\"weight\":\"0.015322247994609446\"},{\"address\":\"cosmos1vxnfc3v6fx6vwxn0gees08n02cuaf8jrjq7hxv\",\"weight\":\"0.014120472015057072\"},{\"address\":\"cosmos1j2uxc248vlfw32ws0auyrjatgkh58c6zepkjku\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1fqd5p03lv502gy24r5fxyexpq6zyccgc56yagm\",\"weight\":\"0.006481158671130950\"},{\"address\":\"cosmos1wl62xxnkeftjksvmx73wj2fa3vpreg5fyczurn\",\"weight\":\"0.011313084918839887\"},{\"address\":\"cosmos1pmuazlfeqjpv0fdktkuh20epw276zvg58advj8\",\"weight\":\"0.014478223295274740\"},{\"address\":\"cosmos1nnzlwke77snwzdq54ryaau9gzw47a7sg8mrng4\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1xzyprgtm3hjrqnr75qclqk7ypqunmfl7k26dza\",\"weight\":\"0.011906575459237459\"},{\"address\":\"cosmos1hyt7g5533ypamrkhm29jqdzu9w3mws0dgpwzzd\",\"weight\":\"0.013018175136719084\"},{\"address\":\"cosmos1ct3n6rn6qvps6j2f5pfhsyszjw9r3z6ps89x4m\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1rvp5k0g7lk5sgul6prt3hs6z55ddd6k2jxqxtp\",\"weight\":\"0.012310359216781740\"},{\"address\":\"cosmos1aye9n83a4eulzntgq25zkyq4faljg225x5hm7y\",\"weight\":\"0.013831594424707402\"},{\"address\":\"cosmos1q9puau53l07ys8en7k5ejjlu49plr607jqku2m\",\"weight\":\"0.019075909099606813\"},{\"address\":\"cosmos186js4cmqzqeesltj9sjspsj6xthehg5zuk8ryq\",\"weight\":\"0.010506634053162084\"},{\"address\":\"cosmos1h3ytxxhq5kerq080uv985jmrmpwmhjlq4cjgk5\",\"weight\":\"0.011408884176143559\"},{\"address\":\"cosmos1fdufk5dtcnf5r5ktfk9l997c7tlhcfln620t6f\",\"weight\":\"0.014516652780079007\"},{\"address\":\"cosmos1mh4yza5rszduv6jtpkhucggp2gqwm2k42h8nz8\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1c7xwgn6x4pp0zj2c33m6j87jl9t02wkn3yt7s7\",\"weight\":\"0.005000000000000000\"},{\"address\":\"cosmos1h2qndghy9tvl7qz9y5z3433kvhuxcwt5sk43j7\",\"weight\":\"0.011152746629363641\"},{\"address\":\"cosmos19amx0kcpnrdx4g9pttcm2v04m76atlxztcq8sd\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1g9rzaedphtthffy9nrh0ukjdehau4p0aq72k5k\",\"weight\":\"0.015477169807850692\"},{\"address\":\"cosmos1gf73lfe9tzn27xpzz6xksku0c08rvk3gekka5h\",\"weight\":\"0.007218936958952564\"},{\"address\":\"cosmos10l5my6s0z3j74jz5fn9a6hfldv0jt2aawdug5n\",\"weight\":\"0.013692731693070881\"},{\"address\":\"cosmos1zj8sm2fsazhcn2jg2h24084f2l3eeu20d396dg\",\"weight\":\"0.014945145118083484\"},{\"address\":\"cosmos1w0ln5gz9tldh5dun9qx8ed3ww82zfsp4vemh79\",\"weight\":\"0.006580834237555860\"},{\"address\":\"cosmos1s2kdm985deuj52yw9mvmrsafhe95qxaghr5ng2\",\"weight\":\"0.020000000000000000\"},{\"address\":\"cosmos1ryaxl3wjyqf793yx0upz0c5se8p3uzj98cev93\",\"weight\":\"0.354811011669756390\"}]", "mint"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 6)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
