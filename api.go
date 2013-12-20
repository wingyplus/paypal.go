package paypalrestsdk

type Api struct {
    endpoint string
}

func NewApi(options map[string]string) *Api{
    var endpoint string

    switch(options["mode"]) {
    case "live":
        endpoint = "https://api.paypal.com"
    case "sandbox":
        endpoint = "https://api.sandbox.paypal.com"
    default:
        endpoint = options["endpoint"]
    }

    return &Api{endpoint: endpoint}
}

func (api *Api) Endpoint() string {
    return api.endpoint
}
