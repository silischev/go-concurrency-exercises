//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream) <-chan *Tweet {
	tweetStream := make(chan *Tweet)

	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				close(tweetStream)
				return
			} else {
				tweetStream <- tweet
			}
		}
	}()

	return tweetStream
}

func consumer(in <-chan *Tweet) {
	wg := &sync.WaitGroup{}

	for t := range in {
		tweet := t
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets about golang")
			} else {
				fmt.Println(tweet.Username, "\tdoes not tweet about golang")
			}
		}()
	}

	wg.Wait()
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweetStream := producer(stream)
	consumer(tweetStream)

	fmt.Printf("Process took %s\n", time.Since(start))
}
