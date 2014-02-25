package main

import (
	"net/url"
	"github.com/dougblack/sleepy"
)

type Market struct{}
type Fixture struct{}

func (market Market) Get(values url.Values) (int, interface{}) {
	markets := []string{"1 x 2", "First Goal"}
	data := map[string][]string	{"markets": markets}
	return 200, data
}

func (fixture Fixture) Get(value url.Values) (int, interface{}) {
	fixtures := []string{"Man Utd v Chelsea","Arsenal v Liverpool", "WBA v Norwich"}
	data := map[string][]string{"fixtures": fixtures}
	return 200, data
}


func main() {
	market:=new(Market)
	fixture:=new(Fixture)
	
	go func() {
		var marketsApi = sleepy.NewAPI()
		marketsApi.AddResource(market,"/markets")
		marketsApi.Start(3000)
		
	}()
	
	
	var fixturesApi = sleepy.NewAPI()
	fixturesApi.AddResource(fixture,"/fixtures")
	fixturesApi.Start(3001)
}