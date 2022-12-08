package main

import (
	"context"
	"flag"
	"fmt"
	sdk_wrapper "github.com/fforchino/vector-go-sdk/pkg/sdk-wrapper"
	"strings"
	"vectorx/pkg/intents"
)

var Debug = false

func main() {
	var serial = flag.String("serial", "", "Vector's Serial Number")
	var locale = flag.String("locale", "", "STT Locale in use")
	var speechText = flag.String("speechText", "", "Speech text")
	flag.Parse()

	if Debug {
		println("SERIAL: " + *serial)
		println("LOCALE: " + *locale)
		println("SPEECH TEXT: " + *speechText)
	}

	if len(*speechText) > 0 {
		// Remove "" if any
		if strings.HasPrefix(*speechText, "\"") && strings.HasSuffix(*speechText, "\"") {
			*speechText = (*speechText)[1 : len(*speechText)-1]
		}

		// Register vectorx intents
		intents.RegisterIntents()

		// Make sure "locale" is just the language name
		if strings.Contains(*locale, "-") {
			*locale = strings.Split(*locale, "-")[0]
		}
		// Find out whether the speech text matches any registered intent
		xIntent, err := intents.IntentMatch(*speechText, *locale)

		if err == nil {
			// Ok, we have a match. Then extract the parameters (if any) from the intent...
			params := intents.ParseParams(*speechText, xIntent, *locale)

			// And now run the handler function (SDK code)
			sdk_wrapper.InitSDKForWirepod(*serial)

			ctx := context.Background()
			start := make(chan bool)
			stop := make(chan bool)

			go func() {
				_ = sdk_wrapper.Robot.BehaviorControl(ctx, start, stop)
			}()

			for {
				select {
				case <-start:
					returnIntent := xIntent.Handler(xIntent, params)
					stop <- true
					// Ok, intent handled. Return the intent that Wirepod has to send to the robot
					fmt.Println("{\"status\": \"ok\", \"returnIntent\": \"" + returnIntent + "\"}")
					return
				}
			}
		} else {
			// Intent cannot be handled by VectorX. Wirepod may continue its intent parsing chain
			fmt.Println("{\"status\": \"ko\", \"returnIntent\": \"\"}")
		}
	} else {
		// Intent cannot be handled by VectorX. Wirepod may continue its intent parsing chain
		fmt.Println("{\"status\": \"ko\", \"returnIntent\": \"\"}")
	}
}