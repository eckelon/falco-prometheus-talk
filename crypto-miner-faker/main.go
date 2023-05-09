package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {

	for {
		randomNumber := rand.Intn(20) + 1
		ctx, cancelPi := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelPi()
		if randomNumber%10 == 0 {
			go doPi(ctx)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// start a goroutine that does some work with timeout support
		go doRequest(ctx)
		// wait for the work to complete
		time.Sleep(10 * time.Second)
	}

}

func doPi(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			time.Sleep(100 * time.Millisecond)
		default:
			fmt.Println("Starting pi decimals calculations...")
			var wg sync.WaitGroup
			numCPU := runtime.NumCPU()

			wg.Add(numCPU)

			for i := 0; i < numCPU; i++ {
				go func() {
					defer wg.Done()
					pi := calcPiLeibniz(1000000000)
					fmt.Printf("Calculated Pi: %.15f\n", pi)
				}()
			}

			wg.Wait()
		}
	}
}

func doRequest(ctx context.Context) {
	fmt.Println("")
	for {
		select {
		case <-ctx.Done():
			time.Sleep(100 * time.Millisecond)
		default:
			endpoint := selectRandomEndpoint()
			fmt.Println("requesting to " + endpoint + "...")
			_, _ = http.Get(endpoint)
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func calcPiLeibniz(iterations int64) float64 {
	pi := 0.0
	sign := 1.0

	for i := int64(0); i < iterations; i++ {
		pi += sign / (2*float64(i) + 1)
		sign = -sign
	}

	return pi * 4
}

func selectRandomEndpoint() string {

	cryptoDomains := []string{"eth-eu2.nanopool.org", "eth-hk.dwarfpool.com",
		"eth-jp1.nanopool.org", "eth-ru.dwarfpool.com",
		"eth-ru2.dwarfpool.com", "eth-sg.dwarfpool.com",
		"eth-us-east1.nanopool.org", "eth-us-west1.nanopool.org",
		"eth-us.dwarfpool.com", "eth-us2.dwarfpool.com",
		"eu.stratum.slushpool.com", "eu1.ethermine.org",
		"eu1.ethpool.org", "fr.minexmr.com",
		"mine.moneropool.com", "mine.xmrpool.net",
		"pool.minexmr.com", "pool.monero.hashvault.pro",
		"pool.supportxmr.com", "sg.minexmr.com",
		"sg.stratum.slushpool.com", "stratum-eth.antpool.com",
		"stratum-ltc.antpool.com", "stratum-zec.antpool.com",
		"stratum.antpool.com", "us-east.stratum.slushpool.com",
		"us1.ethermine.org", "us1.ethpool.org",
		"us2.ethermine.org", "us2.ethpool.org",
		"xmr-asia1.nanopool.org", "xmr-au1.nanopool.org",
		"xmr-eu1.nanopool.org", "xmr-eu2.nanopool.org",
		"xmr-jp1.nanopool.org", "xmr-us-east1.nanopool.org",
		"xmr-us-west1.nanopool.org", "xmr.crypto-pool.fr",
		"xmr.pool.minergate.com", "rx.unmineable.com",
		"ss.antpool.com", "dash.antpool.com",
		"eth.antpool.com", "zec.antpool.com",
		"xmc.antpool.com", "btm.antpool.com",
		"stratum-dash.antpool.com", "stratum-xmc.antpool.com",
		"stratum-btm.antpool.com"}

	cryptoPorts := []int{25, 3333, 3334, 3335, 3336, 3357, 4444,
		5555, 5556, 5588, 5730, 6099, 6666, 7777,
		7778, 8000, 8001, 8008, 8080, 8118, 8333,
		8888, 8899, 9332, 9999, 14433, 14444,
		45560, 45700}
	randDomainIndex := rand.Intn(len(cryptoDomains))
	randPortIndex := rand.Intn(len(cryptoPorts))
	return "http://" + cryptoDomains[randDomainIndex] + ":" + strconv.Itoa(cryptoPorts[randPortIndex])
}
