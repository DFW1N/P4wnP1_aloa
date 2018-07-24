package main

import (
	pb "../proto/gopherjs"
	"honnef.co/go/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"github.com/HuckRidgeSW/hvue"
)

var (
	document   = dom.GetWindow().Document().(dom.HTMLDocument)
	serverAddr = GetBaseURL()
	Client     = NewRpcClient(serverAddr + ":80")
	GS *pb.GadgetSettings
)

func GetBaseURL() string {
	document := js.Global.Get("window").Get("document")
	location := document.Get("location")
	port := location.Get("port").String()
	url := location.Get("protocol").String() + "//" + location.Get("hostname").String()
	if len(port) > 0 {
		url = url + ":" + port
	}
	return url
}

type appController struct {
	*js.Object
}

func main() {
	println(GetBaseURL())

	/*
	println("Listening for RPC events ...")
	err := Client.StartListenEvents(common.EVT_ANY)
	if err != nil {println(err)}

	time.Sleep(time.Second * 5)

	Client.StopEventListening()
	println("... done listening for RPC events")

	time.Sleep(time.Second)

	println("Listening for RPC events ...")
	err := Client.StartListenEvents(common.EVT_LOG)
	if err != nil {println(err)}
	*/

	/*
	fmt.Printf("Address %v\n", strings.TrimSuffix(document.BaseURI(), "/"))
	fmt.Printf("Client %v\n", Client)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	gs, err := Client.GetDeployedGadgetSetting(ctx, &pb.Empty{})
	if err == nil {
		//export Gadget setting
		js.Global.Set("gs", gs)
		GS = gs
	} else {
		fmt.Printf("Error rpc call: %v\n", err)
	}
	*/

	InitCompEthernetAddresses2()
	InitCompToggleSwitch()
	InitCompUSBSettings()
	InitCompTab()
	InitCompTabs()
	InitCompCodeEditor()
	InitCompHIDScript()
	InitCompLogger()
	vm := hvue.NewVM(
		hvue.El("#app"),
		//add "testString" to data
		hvue.DataFunc(func(vm *hvue.VM) interface{} {
			data := struct{
				*js.Object
				TestString string `js:"testString"`
			}{Object: O()}
			data.TestString = "type('hello');"
			return &data
		}),
		//add console to app as computed property, to allow debug output on vue events
		hvue.Computed(
			"console",
			func(vm *hvue.VM) interface{} {
			return js.Global.Get("console")
		}))
	js.Global.Set("vm",vm)

}
