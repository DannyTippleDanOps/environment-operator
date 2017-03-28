package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/runtime"
	"k8s.io/client-go/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

// Client is a top level struct, wrapping all other clients
type Client struct {
	Interface kubernetes.Interface
	Namespace string
}

// Service builds Service client
func (c *Client) Service() *Service {
	return &Service{Interface: c.Interface, Namespace: c.Namespace}
}

// Deployment builds Deployment client
func (c *Client) Deployment() *Deployment {
	return &Deployment{Interface: c.Interface, Namespace: c.Namespace}
}

// PVC builds PersistentVolumeClaim client
func (c *Client) PVC() *PersistentVolumeClaim {
	return &PersistentVolumeClaim{Interface: c.Interface, Namespace: c.Namespace}
}

// Ingress builds Ingress client
func (c *Client) Ingress() *Ingress {
	return &Ingress{Interface: c.Interface, Namespace: c.Namespace}
}

// Ns builds Ingress client
func (c *Client) Ns() *Namespace {
	return &Namespace{Interface: c.Interface, Namespace: c.Namespace}
}

// ThirdPartyResource() builds TPR client
func (c *Client) ThirdPartyResource(kind string) *ThirdPartyResource {
	var restcli *rest.RESTClient
	var err error

	config, err := rest.InClusterConfig()
	if err != nil {
		return &ThirdPartyResource{}
	}

	config.GroupVersion = &unversioned.GroupVersion{
		Group:   "prsn.io",
		Version: "v1",
	}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}

	// TPR request/response debug stuff below.
	//
	// config.CAFile = ""
	// config.CAData = []byte{}
	// config.CertFile = ""
	// config.CertData = []byte{}
	//
	// config.Transport = &loghttp.Transport{
	// 	LogResponse: func(resp *http.Response) {
	// 		dump, err := httputil.DumpResponse(resp, true)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// log.Debugf("RESPONSE: %q", dump)
	// log.Debugf("[%p] %d %s", resp.Request, resp.StatusCode, resp.Request.URL)
	// },
	// Transport: &http.Transport{
	// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// },
	// }

	restcli, err = rest.RESTClientFor(config)
	if err != nil {
		return &ThirdPartyResource{}
	}

	return &ThirdPartyResource{
		Interface: restcli,
		Namespace: c.Namespace,
		Type:      kind}
}

func listOptions() v1.ListOptions {
	return v1.ListOptions{
		LabelSelector: "creator=pipeline",
	}
}
