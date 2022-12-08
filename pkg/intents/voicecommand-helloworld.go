package intents

import (
	sdk_wrapper "github.com/fforchino/vector-go-sdk/pkg/sdk-wrapper"
)

/**********************************************************************************************************************/
/*                                                HELLO WORLD                                                         */
/**********************************************************************************************************************/

/*
To implement an extended voice command there are just two functions you need to implement in your command file.
The first function you have to implement is "COMMAND_Register". It registers your custom intent.
A call to HelloWorld_Register() must be added in RegisterIntents() in the intents.go file.
*/

func HelloWorld_Register(intentList *[]IntentDef) error {
	utterances := make(map[string][]string)
	utterances[LOCALE_ENGLISH] = []string{"hello world"}
	utterances[LOCALE_ITALIAN] = []string{"ciao mondo"}

	var intent = IntentDef{
		IntentName: "extended_intent_hello_world",
		Utterances: utterances,
		Parameters: []string{},
		Handler:    helloWorld,
	}
	*intentList = append(*intentList, intent)
	return nil
}

/*
The second function you have to implement is the handler of your intent, e.g. the code that will be executed when the intent is
recognized. The engine passes this fuction the matched intent and its parameters, and expects that it returns the wirepod intent
to be sent back to the robot. That's all!
*/

func helloWorld(intent IntentDef, params IntentParams) string {
	returnIntent := STANDARD_INTENT_GREETING_HELLO
	sdk_wrapper.SayText("Hello world!")
	return returnIntent
}