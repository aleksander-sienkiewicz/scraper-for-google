package main //always here

import (
	"fmt" //print format
	"math/rand"
	"net/http" //make http requests in this project
	"net/url"
	"strings"
	"time" //selecting random user agent

	"github.com/PuerkitoBio/goquery" //go query, help us work with scraped items
	//CLI cmd for goquery import: [go get "github.com/PuerkitoBio/goquery"]
)

var googleDomains = map[string]string{ //contain list of all google domains out there for that country code
	"com": "https://www.google.com/search?q=", //?[parameter]
	"ac":  "https://www.google.ac/search?q=",  //take link for that country
	"ad":  "https://www.google.ad/search?q=",  //then append a querry to the end of it
	"ae":  "https://www.google.ae/search?q=",  //and boom, program know which search to use
	"af":  "https://www.google.com.af/search?q=",
	"ag":  "https://www.google.com.ag/search?q=",
	"ai":  "https://www.google.com.ai/search?q=",
	"al":  "https://www.google.al/search?q=",
	"am":  "https://www.google.am/search?q=",
	"ao":  "https://www.google.co.ao/search?q=",
	"ar":  "https://www.google.com.ar/search?q=",
	"as":  "https://www.google.as/search?q=",
	"at":  "https://www.google.at/search?q=",
	"au":  "https://www.google.com.au/search?q=",
	"az":  "https://www.google.az/search?q=",
	"ba":  "https://www.google.ba/search?q=",
	"bd":  "https://www.google.com.bd/search?q=",
	"be":  "https://www.google.be/search?q=",
	"bf":  "https://www.google.bf/search?q=",
	"bg":  "https://www.google.bg/search?q=",
	"bh":  "https://www.google.com.bh/search?q=",
	"bi":  "https://www.google.bi/search?q=",
	"bj":  "https://www.google.bj/search?q=",
	"bn":  "https://www.google.com.bn/search?q=",
	"bo":  "https://www.google.com.bo/search?q=",
	"br":  "https://www.google.com.br/search?q=",
	"bs":  "https://www.google.bs/search?q=",
	"bt":  "https://www.google.bt/search?q=",
	"bw":  "https://www.google.co.bw/search?q=",
	"by":  "https://www.google.by/search?q=",
	"bz":  "https://www.google.com.bz/search?q=",
	"ca":  "https://www.google.ca/search?q=",
	"kh":  "https://www.google.com.kh/search?q=",
	"cc":  "https://www.google.cc/search?q=",
	"cd":  "https://www.google.cd/search?q=",
	"cf":  "https://www.google.cf/search?q=",
	"cat": "https://www.google.cat/search?q=",
	"cg":  "https://www.google.cg/search?q=",
	"ch":  "https://www.google.ch/search?q=",
	"ci":  "https://www.google.ci/search?q=",
	"ck":  "https://www.google.co.ck/search?q=",
	"cl":  "https://www.google.cl/search?q=",
	"cm":  "https://www.google.cm/search?q=",
	"co":  "https://www.google.com.co/search?q=",
	"cr":  "https://www.google.co.cr/search?q=",
	"cu":  "https://www.google.com.cu/search?q=",
	"cv":  "https://www.google.cv/search?q=",
	"cy":  "https://www.google.com.cy/search?q=",
	"cz":  "https://www.google.cz/search?q=",
	"de":  "https://www.google.de/search?q=",
	"dj":  "https://www.google.dj/search?q=",
	"dk":  "https://www.google.dk/search?q=",
	"dm":  "https://www.google.dm/search?q=",
	"do":  "https://www.google.com.do/search?q=",
	"dz":  "https://www.google.dz/search?q=",
	"ec":  "https://www.google.com.ec/search?q=",
	"ee":  "https://www.google.ee/search?q=",
	"eg":  "https://www.google.com.eg/search?q=",
	"es":  "https://www.google.es/search?q=",
	"et":  "https://www.google.com.et/search?q=",
	"fi":  "https://www.google.fi/search?q=",
	"fj":  "https://www.google.com.fj/search?q=",
	"fm":  "https://www.google.fm/search?q=",
	"fr":  "https://www.google.fr/search?q=",
	"ga":  "https://www.google.ga/search?q=",
	"ge":  "https://www.google.ge/search?q=",
	"gf":  "https://www.google.gf/search?q=",
	"gg":  "https://www.google.gg/search?q=",
	"gh":  "https://www.google.com.gh/search?q=",
	"gi":  "https://www.google.com.gi/search?q=",
	"gl":  "https://www.google.gl/search?q=",
	"gm":  "https://www.google.gm/search?q=",
	"gp":  "https://www.google.gp/search?q=",
	"gr":  "https://www.google.gr/search?q=",
	"gt":  "https://www.google.com.gt/search?q=",
	"gy":  "https://www.google.gy/search?q=",
	"hk":  "https://www.google.com.hk/search?q=",
	"hn":  "https://www.google.hn/search?q=",
	"hr":  "https://www.google.hr/search?q=",
	"ht":  "https://www.google.ht/search?q=",
	"hu":  "https://www.google.hu/search?q=",
	"id":  "https://www.google.co.id/search?q=",
	"iq":  "https://www.google.iq/search?q=",
	"ie":  "https://www.google.ie/search?q=",
	"il":  "https://www.google.co.il/search?q=",
	"im":  "https://www.google.im/search?q=",
	"in":  "https://www.google.co.in/search?q=",
	"io":  "https://www.google.io/search?q=",
	"is":  "https://www.google.is/search?q=",
	"it":  "https://www.google.it/search?q=",
	"je":  "https://www.google.je/search?q=",
	"jm":  "https://www.google.com.jm/search?q=",
	"jo":  "https://www.google.jo/search?q=",
	"jp":  "https://www.google.co.jp/search?q=",
	"ke":  "https://www.google.co.ke/search?q=",
	"ki":  "https://www.google.ki/search?q=",
	"kg":  "https://www.google.kg/search?q=",
	"kr":  "https://www.google.co.kr/search?q=",
	"kw":  "https://www.google.com.kw/search?q=",
	"kz":  "https://www.google.kz/search?q=",
	"la":  "https://www.google.la/search?q=",
	"lb":  "https://www.google.com.lb/search?q=",
	"lc":  "https://www.google.com.lc/search?q=",
	"li":  "https://www.google.li/search?q=",
	"lk":  "https://www.google.lk/search?q=",
	"ls":  "https://www.google.co.ls/search?q=",
	"lt":  "https://www.google.lt/search?q=",
	"lu":  "https://www.google.lu/search?q=",
	"lv":  "https://www.google.lv/search?q=",
	"ly":  "https://www.google.com.ly/search?q=",
	"ma":  "https://www.google.co.ma/search?q=",
	"md":  "https://www.google.md/search?q=",
	"me":  "https://www.google.me/search?q=",
	"mg":  "https://www.google.mg/search?q=",
	"mk":  "https://www.google.mk/search?q=",
	"ml":  "https://www.google.ml/search?q=",
	"mm":  "https://www.google.com.mm/search?q=",
	"mn":  "https://www.google.mn/search?q=",
	"ms":  "https://www.google.ms/search?q=",
	"mt":  "https://www.google.com.mt/search?q=",
	"mu":  "https://www.google.mu/search?q=",
	"mv":  "https://www.google.mv/search?q=",
	"mw":  "https://www.google.mw/search?q=",
	"mx":  "https://www.google.com.mx/search?q=",
	"my":  "https://www.google.com.my/search?q=",
	"mz":  "https://www.google.co.mz/search?q=",
	"na":  "https://www.google.com.na/search?q=",
	"ne":  "https://www.google.ne/search?q=",
	"nf":  "https://www.google.com.nf/search?q=",
	"ng":  "https://www.google.com.ng/search?q=",
	"ni":  "https://www.google.com.ni/search?q=",
	"nl":  "https://www.google.nl/search?q=",
	"no":  "https://www.google.no/search?q=",
	"np":  "https://www.google.com.np/search?q=",
	"nr":  "https://www.google.nr/search?q=",
	"nu":  "https://www.google.nu/search?q=",
	"nz":  "https://www.google.co.nz/search?q=",
	"om":  "https://www.google.com.om/search?q=",
	"pk":  "https://www.google.com.pk/search?q=",
	"pa":  "https://www.google.com.pa/search?q=",
	"pe":  "https://www.google.com.pe/search?q=",
	"ph":  "https://www.google.com.ph/search?q=",
	"pl":  "https://www.google.pl/search?q=",
	"pg":  "https://www.google.com.pg/search?q=",
	"pn":  "https://www.google.pn/search?q=",
	"pr":  "https://www.google.com.pr/search?q=",
	"ps":  "https://www.google.ps/search?q=",
	"pt":  "https://www.google.pt/search?q=",
	"py":  "https://www.google.com.py/search?q=",
	"qa":  "https://www.google.com.qa/search?q=",
	"ro":  "https://www.google.ro/search?q=",
	"rs":  "https://www.google.rs/search?q=",
	"ru":  "https://www.google.ru/search?q=",
	"rw":  "https://www.google.rw/search?q=",
	"sa":  "https://www.google.com.sa/search?q=",
	"sb":  "https://www.google.com.sb/search?q=",
	"sc":  "https://www.google.sc/search?q=",
	"se":  "https://www.google.se/search?q=",
	"sg":  "https://www.google.com.sg/search?q=",
	"sh":  "https://www.google.sh/search?q=",
	"si":  "https://www.google.si/search?q=",
	"sk":  "https://www.google.sk/search?q=",
	"sl":  "https://www.google.com.sl/search?q=",
	"sn":  "https://www.google.sn/search?q=",
	"sm":  "https://www.google.sm/search?q=",
	"so":  "https://www.google.so/search?q=",
	"st":  "https://www.google.st/search?q=",
	"sr":  "https://www.google.sr/search?q=",
	"sv":  "https://www.google.com.sv/search?q=",
	"td":  "https://www.google.td/search?q=",
	"tg":  "https://www.google.tg/search?q=",
	"th":  "https://www.google.co.th/search?q=",
	"tj":  "https://www.google.com.tj/search?q=",
	"tk":  "https://www.google.tk/search?q=",
	"tl":  "https://www.google.tl/search?q=",
	"tm":  "https://www.google.tm/search?q=",
	"to":  "https://www.google.to/search?q=",
	"tn":  "https://www.google.tn/search?q=",
	"tr":  "https://www.google.com.tr/search?q=",
	"tt":  "https://www.google.tt/search?q=",
	"tw":  "https://www.google.com.tw/search?q=",
	"tz":  "https://www.google.co.tz/search?q=",
	"ua":  "https://www.google.com.ua/search?q=",
	"ug":  "https://www.google.co.ug/search?q=",
	"uk":  "https://www.google.co.uk/search?q=",
	"us":  "https://www.google.com/search?q=",
	"uy":  "https://www.google.com.uy/search?q=",
	"uz":  "https://www.google.co.uz/search?q=",
	"vc":  "https://www.google.com.vc/search?q=",
	"ve":  "https://www.google.co.ve/search?q=",
	"vg":  "https://www.google.vg/search?q=",
	"vi":  "https://www.google.co.vi/search?q=",
	"vn":  "https://www.google.com.vn/search?q=",
	"vu":  "https://www.google.vu/search?q=",
	"ws":  "https://www.google.ws/search?q=",
	"za":  "https://www.google.co.za/search?q=",
	"zm":  "https://www.google.co.zm/search?q=",
	"zw":  "https://www.google.co.zw/search?q=",
}

type SearchResult struct { //define struct to store domain in accessible and printable way
	ResultRank  int    //when we pass results from google we will only hold these 4 values
	ResultURL   string //easy storage to data base
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{ //list of user agents, to access google site, always need these for the classic scraper projects
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string { //to select a random user agent - uses the math import
	rand.Seed(time.Now().Unix())            //initialize with seed func, can read abt it on randint library page on golang docs
	randNum := rand.Int() % len(userAgents) //from len of user agents, select a random number and select it, maybe should have done len-1
	return userAgents[randNum]              //selected for each different request
}

func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	//accept search term country code, language code, number of pages, and num results expected
	toScrape := []string{}                                 //slice of type string
	searchTerm = strings.Trim(searchTerm, " ")             //trim all space in search term
	searchTerm = strings.Replace(searchTerm, " ", "+", -1) //where ever  there is a space, put a plus. llike how google URL has + in there

	if googleBase, found := googleDomains[countryCode]; found { //index and find result for that country code from map-
		//so if its found, range over pages [number of pages u choose to scrape from google]
		for i := 0; i < pages; i++ { //for every page
			start := i * count                                                                                                  //count is number of results returned
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start) //the format google is expecting
			//scrapeURL is the actuale url we want to build, so it will Sprint from google link based of country code, users search,
			//count(results expected), start, and language code - all things we need to build the URL
			toScrape = append(toScrape, scrapeURL) //append scrapeURL to toScrape
		}
	} else { //the country code that user passed is not supported
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err //return nil for the slice, and error
	}
	return toScrape, nil //return the slice string to scrape and nil for the error

}

/*
were sending the response that we recieve after we use scrapeClientRequest, and result counter
sending a response of type httpresponse
result counter called rank here
return searchResult, and error
*/
func googleResultParsing(response *http.Response, rank int) ([]SearchResult, error) {
	// will use goQuerry lib here
	// newdocfromresponse, takes response. returns document, and error is error
	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil { //when ever there is an error u ALWAYS check for an error :D
		return nil, err
	}

	results := []SearchResult{} //define results equal to search results
	sel := doc.Find("div.g")    //sel will find div.g in the google document
	//when we make that request. we recieve the response in a document format that we can work with
	//and querry, and run find functions and string and trimming functions, why we convert to a document
	//we want to find div.g[this ID] in the doc cuz itll contain every request

	rank++                     //rank +1
	for i := range sel.Nodes { //sel will have dif Nodes, they have the individual results were after
		item := sel.Eq(i) //start assigning vars and assignign values
		//we finding all our vars and assinging them from the document
		linkTag := item.Find("a")       //in html 'a' always has the linktag
		link, _ := linkTag.Attr("href") //link is always in linktag.attribute 'href'
		titleTag := item.Find("h3.r")   //h3.r is the element that has title tag will convert to text to get title
		descTag := item.Find("span.st") //description tags found using span.st - these all r used by google[the stuff in the ' ']
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ") //trim empty spaces

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") { //if link isnt empty and link isnt a null link
			//a nul link is represented by the # there and strings.hasprifix like '/'
			result := SearchResult{ //then result = searchresult
				rank, //create one search result with these attributes
				link,
				/*we defined searchresult as a struct earlier as a struct, it has these atribtues. thats why we found them
				in the doc in the code just abovvein the for loop for sel.Nodes, for the items in text
				*/
				title,
				desc,
			}
			results = append(results, result)
			//now we have the things we need we take results which is type slice searchresults, which we can append
			//so thats what we do in the line above
			rank++
		}
	}
	return results, err //return results, and error as nil if all goes well

}

/*
this func will return the http client were gonna use to scrape, we can either say use proxy as our
http client or use default http client[which is return &http.Client{}]
*/
func getScrapeClient(proxyString interface{}) *http.Client { //passing pointer

	switch v := proxyString.(type) {

	case string: //if passed string as proxy
		proxyUrl, _ := url.Parse(v)                                                     //
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}} //this proxy doesnt need to be here!
		//

		//just if i need to see proxy code in the future but im prolly just gonna watch a vid in that case

	default:
		return &http.Client{}
	}
}

func GoogleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {

	results := []SearchResult{} //return result, slice of type searchresult
	resultCounter := 0
	//count num of results found fo rparticular querry

	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	//will create search queery for google, [just basic usage of querry-jsut focus on golang]
	//we want to list of urls we will search
	if err != nil { //if the error is there
		return nil, err //return no search result, and the error
	}

	//for, return page, equal to range, googlepages
	for _, page := range googlePages { //range over pages
		res, err := scrapeClientRequest(page, proxyString) //make request to query we built, one by one go thru all query in googlepages var, one by one make request to that url using scraprequest func
		//we will pass page, and proxystring - we will keep proxy as nil, proxy is if we want our server or computer to make a request to google, can pass if we wanted to
		if err != nil { //if there is an error
			return nil, err //return nil and error
		}

		//
		data, err := googleResultParsing(res, resultCounter) //make request to query we built, one by one go thru all query in googlepages var, one by one make request to url using scraprequest func
		/*for all URLS we need to call three dif funcs, one where we will scrape client request, we will make request to client
		and one where we parse the request that we get.
		This takes response that we get after we made request to URL and resultcounter
		*/

		if err != nil { //we will also handle the error here
			return nil, err //return nil for search result and pass error to the main func
		}
		resultCounter += len(data)    //data that we get from googleResultParsing
		for _, result := range data { //append that into our result
			results = append(results, result)
			/*results is the slice of searchResult, we will recieve search result 1 by 1 from
			googleResultParsing, we take that here and range over it and append it to our results slice
			*/
		}
		time.Sleep(time.Duration(backoff) * time.Second) //every time u loop over lets say a page, sleep for some time
		//then it will go for next time, we will pass 10 sec before scraping next page, we ALSO pass nil as proxy string
	}
	return results, nil //return results and nil for error
}

/*
takes searchURL(the page), the query we built from buildGoogleURLS, that single url we are passing here
takes proxyString

returns http response or error
*/

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {
	//to make this request we need a client, func is called getScrapeClient
	baseClient := getScrapeClient(proxyString) //will take our proxyString, if there is something it will
	//create the base clioent from that or it will return default base client if nothing is entered
	req, _ := http.NewRequest("GET", searchURL, nil) //first we will make a request
	req.Header.Set("User-Agent", randomUserAgent())  //also set header for user agent using random user agent function
	/*we have set the header and we have made a get request to the searchURL, the searchurl we got from buildurl func
	 */
	res, err := baseClient.Do(req)
	if res.StatusCode != 200 { //if there is some problem -guy in bid didnt rlly elaborate we will see
		err := fmt.Errorf("scraper received a non-200 status code suggesting a ban") //will update yall on this one
		return nil, err                                                              //obviously return nil for results and error
	}

	if err != nil { //if theres an error
		return nil, err //same same
	}
	return res, nil //otherwise simply give response and nil for error <3
}

func main() {
	//main calls to google scrape the search, country code for the website, language code, error code, number of pages you want to scrape, number of results you want,
	//nil is passed for proxy string and 10 is passed for delay time[backoff]
	res, err := GoogleScrape("aleksander sienkiewicz", "com", "en", nil, 1, 30, 10) //text to search, the name can be anything
	if err == nil {                                                                 //if no error
		for _, res := range res { //range over response
			fmt.Println(res) //print response
		}
	}
}
