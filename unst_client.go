import (
	//"context"
	//"encoding/json"
	"fmt"
	"io"
	"time"

	//"k8s.io/apimachinery/pkg/api/meta"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	//"k8s.io/apimachinery/pkg/types"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	//"k8s.io/client-go/discovery"
	//"k8s.io/client-go/discovery/cached/memory"
	//"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/restmapper"
)

const (
	applyTO = 30 * time.Second
)

var decUnstructured = yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)

// YAMLScanner scans a yaml stream and gets one doc at a time
type YAMLScanner struct {
	decoder io.ReadCloser
}

func NewYAMLScanner(r io.ReadCloser) *YAMLScanner {
	return &YAMLScanner{
		decoder: yamlutil.NewDocumentDecoder(r),
	}
}

func (y *YAMLScanner) Scan() ([]byte, error) {
	var doc []byte

	for {
		b := make([]byte, 16*1024)
		n, err := y.decoder.Read(b)
		switch err {
		case nil:
			doc = append(doc, b[:n]...)
			return doc, nil

		case io.ErrShortBuffer:
			doc = append(doc, b...)

		default:
			return nil, err
		}
	}

	return doc, nil
}

func applyYAML(r io.ReadCloser) error {
	scanner := NewYAMLScanner(r)
	for {
		doc, err := scanner.Scan()
		if err == nil {
			fmt.Printf("OBJ:\n %s\n", doc)
			continue
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}

// applyOneObject given a kubeconfig and unstructured object, apply it
// to the cluster using Server Side Apply
func applyOneObject(cfg *rest.Config, objYAML []byte) error {
/*
	// Discovery Client/RESTMapper to find GVR
	dc, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return err
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))

	// Dynamic client
	dyn, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}

	// Decode YAML unstructured
	obj := &unstructured.Unstructured{}
	_, gvk, err := decUnstructured.Decode(objYAML, nil, obj)
	if err != nil {
		return err
	}

	// Find GVR
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return err
	}

	// Get REST interface for the GVR
	var dr dynamic.ResourceInterface
	if mapping.Scope.Name() == meta.RESTScopeNameNamespace {
		// namespaced resources should specify the namespace
		dr = dyn.Resource(mapping.Resource).Namespace(obj.GetNamespace())
	} else {
		// for cluster-wide resources
		dr = dyn.Resource(mapping.Resource)
	}

	// Marshal object into JSON
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	// Create or Update the object with SSA
	// Force update on a conflict
	ctx, cancel := context.WithTimeout(context.Background(), applyTO)
	defer cancel()

	forceOnConf := true

	_, err = dr.Patch(ctx, obj.GetName(), types.ApplyPatchType, data, metav1.PatchOptions{
		FieldManager: "applier",
		Force:        &forceOnConf,
	})
*/

	return nil
}
