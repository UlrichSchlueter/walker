package collect

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

type IdParam int

const (
	GENERATED IdParam = iota + 1 // EnumIndex = 1
	NAME                         // EnumIndex = 2

)

type CollectType struct {
	namespace string
	resource  string
	version   string
	group     string
	IdType    IdParam
}

// https://stackoverflow.com/questions/67070548/how-to-get-schema-groupversionresource-for-given-k8s-resources-json
func resourcesToCollect() []CollectType {
	return []CollectType{
		{namespace: "hive", resource: "deployments", version: "v1", group: "apps", IdType: NAME},
		{namespace: "hive", resource: "pods", version: "v1", group: "", IdType: NAME},
		{namespace: "hive", resource: "services", version: "v1", group: "api", IdType: NAME},
		{namespace: "", resource: "nodes", version: "v1", group: "", IdType: GENERATED},
		{namespace: "", resource: "namespaces", version: "v1", group: "", IdType: NAME},
		{namespace: "", resource: "clusteroperators", version: "v1", group: "config.openshift.io", IdType: NAME},
		{namespace: "", resource: "ingresses", version: "v1", group: "networking.k8s.io", IdType: NAME},
	}
}

func GetFolderNameDate(file string) string {
	const layout = "02-01-2006-15-04"
	t := time.Now()
	return file + "-" + t.Format(layout)
}

func WriteFile(filename string, content string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func GetResourcesDynamically(dynamic dynamic.Interface, ctx context.Context,
	group string, version string, resource string, namespace string) (
	[]unstructured.Unstructured, error) {

	resourceId := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}
	list, err := dynamic.Resource(resourceId).Namespace(namespace).
		List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return list.Items, nil
}

func Collect(folder string) {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	dynamic := dynamic.NewForConfigOrDie(config)

	foldername := strings.Split(config.Host, ".")[1]

	na := GetFolderNameDate(foldername)
	err := os.Mkdir(na, 0755)
	if err != nil {
		panic(err)
	}

	for _, c := range resourcesToCollect() {

		items, err := GetResourcesDynamically(dynamic, ctx,
			c.group, c.version, c.resource, c.namespace)
		if err != nil {
			fmt.Println(err)
		} else {
			for i, item := range items {
				name := item.GetName()
				if name == "" {
					name = fmt.Sprintf("index-%d", i)
				}
				b, err := item.MarshalJSON()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(b))
					ns := c.namespace
					if c.namespace == "" {
						ns = "cluster"
					}
					ff := fmt.Sprintf("%s%s%s.%s.%s", na, string(os.PathSeparator), ns, c.resource, name)
					WriteFile(ff, string(b))
					WriteFile(ff+".meta", "{ \"name\": \""+ff+
						"\",\"type\": \""+c.resource+
						"\",\"prefix\": \""+c.namespace+
						"\",\"index\":\""+name+
						"\" }")
					//prettyPrint(b)
				}
			}

		}
	}

}
