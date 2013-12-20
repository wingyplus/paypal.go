package paypalrestsdk

type Api struct {
    endpoint string
}

func NewApi(configure Configure) *Api{
    var endpoint string

    switch(configure.Mode) {
    case "live":
        endpoint = "https://api.paypal.com"
    case "sandbox":
        endpoint = "https://api.sandbox.paypal.com"
    default:
        endpoint = configure.Endpoint
    }

    return &Api{endpoint: endpoint}
}

func (api *Api) Endpoint() string {
    return api.endpoint
}
